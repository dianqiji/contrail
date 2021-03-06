package cluster

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/Juniper/asf/pkg/keystone"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/Juniper/contrail/pkg/deploy/base"
	"github.com/pkg/errors"
)

const (
	defaultTemplateRoot = "./pkg/cluster/configs"
)

type deployCluster struct {
	base.Deploy
	cluster     *Cluster
	clusterID   string
	action      string
	clusterData *base.Data
}

func newDeployCluster(c *Cluster, cData *base.Data, moduleName string) *deployCluster {
	return &deployCluster{
		cluster:     c,
		clusterID:   c.config.ClusterID,
		action:      c.config.Action,
		clusterData: cData,
		Deploy: base.Deploy{
			Reporter: newReporter(c),
			Log:      logutil.NewFileLogger(moduleName, c.config.LogFile),
		},
	}
}

func newReporter(cluster *Cluster) *report.Reporter {
	return report.NewReporter(
		cluster.APIServer,
		fmt.Sprintf("%s/%s", DefaultResourcePath, cluster.config.ClusterID),
		logutil.NewFileLogger("reporter", cluster.config.LogFile),
	)
}

func (p *deployCluster) isCreated() bool {
	state := p.clusterData.ClusterInfo.ProvisioningState
	if p.action == "create" && (state == StatusNoState || state == "") {
		return false
	}
	p.Log.Infof("Cluster %s already deployed, STATE: %s", p.clusterID, state)
	return true
}

func (p *deployCluster) getTemplateRoot() string {
	templateRoot := p.cluster.config.TemplateRoot
	if templateRoot == "" {
		templateRoot = defaultTemplateRoot
	}
	return templateRoot
}

func (p *deployCluster) getWorkRoot() string {
	workRoot := p.cluster.config.WorkRoot
	if workRoot == "" {
		workRoot = DefaultWorkRoot
	}
	return workRoot
}

func (p *deployCluster) getClusterHomeDir() string {
	return filepath.Join(p.getWorkRoot(), p.clusterID)
}

func (p *deployCluster) getWorkingDir() string {
	return filepath.Join(p.getClusterHomeDir())
}

func (p *deployCluster) createWorkingDir() error {
	return os.MkdirAll(p.getWorkingDir(), os.ModePerm)
}

func (p *deployCluster) deleteWorkingDir() error {
	return os.RemoveAll(p.getClusterHomeDir())
}

func (p *deployCluster) keystoneProxyURL() (string, error) {
	auth, err := url.Parse(p.cluster.APIServer.Keystone.URL)
	if err != nil {
		return "", errors.Wrap(err, "parsing AuthURL from cluster data")
	}
	auth.Path = fmt.Sprintf("/proxy/%s/keystone/v3", p.clusterID)

	return auth.String(), nil
}

func (p *deployCluster) ensureServiceUserCreated() error {
	if p.clusterData.ClusterInfo.Orchestrator != "openstack" {
		return nil
	}
	ctx := context.Background()
	name, pass := p.clusterData.KeystoneAdminCredential()

	keystoneURL, err := p.keystoneProxyURL()
	if err != nil {
		return err
	}

	keystoneClient := &keystone.Client{
		URL:      keystoneURL,
		HTTPDoer: p.cluster.APIServer.Keystone.HTTPDoer,
	}

	token, err := keystoneClient.ObtainToken(
		ctx, name, pass, keystone.NewScope("default", "", "", keystone.AdminRoleName),
	)
	if err != nil {
		return errors.Wrap(err, "failed to obtain scoped token")
	}
	ctx = keystone.WithXAuthToken(ctx, token)

	if _, err = keystoneClient.EnsureServiceUserCreated(ctx, keystone.User{
		Name:     p.cluster.config.ServiceUserID,
		Password: p.cluster.config.ServiceUserPassword,
	}); err != nil {
		return errors.Wrapf(err, "cannot ensure that service user %q was created", p.cluster.config.ServiceUserID)
	}
	return nil
}

func (p *deployCluster) createEndpoints() error {
	e := &base.EndpointData{
		ClusterID:   p.clusterID,
		ResManager:  base.NewResourceManager(p.cluster.APIServer, p.cluster.config.LogFile),
		ClusterData: p.clusterData,
		Log:         p.Log,
	}

	return e.Create()
}

func (p *deployCluster) updateEndpoints() error {
	e := &base.EndpointData{
		ClusterID:   p.clusterID,
		ResManager:  base.NewResourceManager(p.cluster.APIServer, p.cluster.config.LogFile),
		ClusterData: p.clusterData,
		Log:         p.Log,
	}

	return e.Update()
}

func (p *deployCluster) deleteEndpoints() error {
	e := &base.EndpointData{
		ClusterID:  p.clusterID,
		ResManager: base.NewResourceManager(p.cluster.APIServer, p.cluster.config.LogFile),
		Log:        p.Log,
	}

	return e.Remove()
}
