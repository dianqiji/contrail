package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestBaremetalNode(t *testing.T) {
	t.Parallel()
	db := testDB
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	common.UseTable(db, "metadata")
	common.UseTable(db, "baremetal_node")
	defer func() {
		common.ClearTable(db, "baremetal_node")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeBaremetalNode()
	model.UUID = "baremetal_node_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "baremetal_node_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "baremetal_node_admin_project_uuid"
	projectModel.FQName = []string{"default-domain-test", "admin-test"}
	projectModel.Perms2.Owner = "admin"
	var createShare []*models.ShareType
	createShare = append(createShare, &models.ShareType{Tenant: "default-domain-test:admin-test", TenantAccess: 7})
	model.Perms2.Share = createShare
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateProject(ctx, tx, &models.CreateProjectRequest{
			Project: projectModel,
		})
	})
	if err != nil {
		t.Fatal("project create failed", err)
	}

	//    //populate update map
	//    updateMap := map[string]interface{}{}
	//
	//
	//    common.SetValueByPath(updateMap, ".UUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".UpdatedAt", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".TargetProvisionState", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".TargetPowerState", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ProvisionState", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PowerState", ".", "test")
	//
	//
	//
	//    if ".Perms2.Share" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Perms2.Share", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Perms2.Share", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.OwnerAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.Owner", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.GlobalAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ParentUUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ParentType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Name", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".MaintenanceReason", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Maintenance", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".LastError", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".InstanceUUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".InstanceInfo.Vcpus", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".InstanceInfo.SwapMB", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".InstanceInfo.RootGB", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".InstanceInfo.NovaHostID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".InstanceInfo.MemoryMB", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".InstanceInfo.LocalGB", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".InstanceInfo.ImageSource", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".InstanceInfo.DisplayName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".InstanceInfo.Capabilities", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.UserVisible", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.OwnerAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.Owner", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.OtherAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.GroupAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.Group", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.LastModified", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Enable", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Description", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Creator", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Created", ".", "test")
	//
	//
	//
	//    if ".FQName" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".FQName", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".FQName", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DriverInfo.IpmiUsername", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DriverInfo.IpmiPassword", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DriverInfo.IpmiAddress", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DriverInfo.DeployRamdisk", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DriverInfo.DeployKernel", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DisplayName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".CreatedAt", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ConsoleEnabled", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".BMProperties.MemoryMB", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".BMProperties.DiskGB", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".BMProperties.CPUCount", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".BMProperties.CPUArch", ".", "test")
	//
	//
	//
	//    if ".Annotations.KeyValuePair" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test": "test"}`)
	//    }
	//
	//
	//    common.SetValueByPath(updateMap, "uuid", ".", "baremetal_node_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateBaremetalNode(ctx, tx,
			&models.CreateBaremetalNodeRequest{
				BaremetalNode: model,
			})
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateBaremetalNode(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	//Delete the project created for sharing
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteProject(ctx, tx, &models.DeleteProjectRequest{
			ID: projectModel.UUID})
	})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListBaremetalNode(ctx, tx, &models.ListBaremetalNodeRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.BaremetalNodes) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteBaremetalNode(ctxDemo, tx,
			&models.DeleteBaremetalNodeRequest{
				ID: model.UUID},
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteBaremetalNode(ctx, tx,
			&models.DeleteBaremetalNodeRequest{
				ID: model.UUID})
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateBaremetalNode(ctx, tx,
			&models.CreateBaremetalNodeRequest{
				BaremetalNode: model})
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListBaremetalNode(ctx, tx, &models.ListBaremetalNodeRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.BaremetalNodes) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}