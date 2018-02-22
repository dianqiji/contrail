package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertLocationQuery = "insert into `location` (`uuid`,`type`,`provisioning_state`,`provisioning_start_time`,`provisioning_progress_stage`,`provisioning_progress`,`provisioning_log`,`private_redhat_subscription_user`,`private_redhat_subscription_pasword`,`private_redhat_subscription_key`,`private_redhat_pool_id`,`private_ospd_vm_vcpus`,`private_ospd_vm_ram_mb`,`private_ospd_vm_name`,`private_ospd_vm_disk_gb`,`private_ospd_user_password`,`private_ospd_user_name`,`private_ospd_package_url`,`private_ntp_hosts`,`private_dns_servers`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`gcp_subnet`,`gcp_region`,`gcp_asn`,`gcp_account_info`,`fq_name`,`display_name`,`aws_subnet`,`aws_secret_key`,`aws_region`,`aws_access_key`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteLocationQuery = "delete from `location` where uuid = ?"

// LocationFields is db columns for Location
var LocationFields = []string{
	"uuid",
	"type",
	"provisioning_state",
	"provisioning_start_time",
	"provisioning_progress_stage",
	"provisioning_progress",
	"provisioning_log",
	"private_redhat_subscription_user",
	"private_redhat_subscription_pasword",
	"private_redhat_subscription_key",
	"private_redhat_pool_id",
	"private_ospd_vm_vcpus",
	"private_ospd_vm_ram_mb",
	"private_ospd_vm_name",
	"private_ospd_vm_disk_gb",
	"private_ospd_user_password",
	"private_ospd_user_name",
	"private_ospd_package_url",
	"private_ntp_hosts",
	"private_dns_servers",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"user_visible",
	"permissions_owner_access",
	"permissions_owner",
	"other_access",
	"group_access",
	"group",
	"last_modified",
	"enable",
	"description",
	"creator",
	"created",
	"gcp_subnet",
	"gcp_region",
	"gcp_asn",
	"gcp_account_info",
	"fq_name",
	"display_name",
	"aws_subnet",
	"aws_secret_key",
	"aws_region",
	"aws_access_key",
	"key_value_pair",
}

// LocationRefFields is db reference fields for Location
var LocationRefFields = map[string][]string{}

// LocationBackRefFields is db back reference fields for Location
var LocationBackRefFields = map[string][]string{

	"physical_router": []string{
		"uuid",
		"server_port",
		"server_ip",
		"resource",
		"physical_router_vnc_managed",
		"physical_router_vendor_name",
		"username",
		"password",
		"version",
		"v3_security_name",
		"v3_security_level",
		"v3_security_engine_id",
		"v3_privacy_protocol",
		"v3_privacy_password",
		"v3_engine_time",
		"v3_engine_id",
		"v3_engine_boots",
		"v3_context_engine_id",
		"v3_context",
		"v3_authentication_protocol",
		"v3_authentication_password",
		"v2_community",
		"timeout",
		"retries",
		"local_port",
		"physical_router_snmp",
		"physical_router_role",
		"physical_router_product_name",
		"physical_router_management_ip",
		"physical_router_loopback_ip",
		"physical_router_lldp",
		"service_port",
		"physical_router_image_uri",
		"physical_router_dataplane_ip",
		"share",
		"owner_access",
		"owner",
		"global_access",
		"parent_uuid",
		"parent_type",
		"user_visible",
		"permissions_owner_access",
		"permissions_owner",
		"other_access",
		"group_access",
		"group",
		"last_modified",
		"enable",
		"description",
		"creator",
		"created",
		"fq_name",
		"display_name",
		"key_value_pair",
	},
}

// LocationParentTypes is possible parents for Location
var LocationParents = []string{}

// CreateLocation inserts Location to DB
func CreateLocation(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreateLocationRequest) error {
	model := request.Location
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLocationQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertLocationQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		string(model.GetType()),
		string(model.GetProvisioningState()),
		string(model.GetProvisioningStartTime()),
		string(model.GetProvisioningProgressStage()),
		int(model.GetProvisioningProgress()),
		string(model.GetProvisioningLog()),
		string(model.GetPrivateRedhatSubscriptionUser()),
		string(model.GetPrivateRedhatSubscriptionPasword()),
		string(model.GetPrivateRedhatSubscriptionKey()),
		string(model.GetPrivateRedhatPoolID()),
		string(model.GetPrivateOspdVMVcpus()),
		string(model.GetPrivateOspdVMRAMMB()),
		string(model.GetPrivateOspdVMName()),
		string(model.GetPrivateOspdVMDiskGB()),
		string(model.GetPrivateOspdUserPassword()),
		string(model.GetPrivateOspdUserName()),
		string(model.GetPrivateOspdPackageURL()),
		string(model.GetPrivateNTPHosts()),
		string(model.GetPrivateDNSServers()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		bool(model.GetIDPerms().GetUserVisible()),
		int(model.GetIDPerms().GetPermissions().GetOwnerAccess()),
		string(model.GetIDPerms().GetPermissions().GetOwner()),
		int(model.GetIDPerms().GetPermissions().GetOtherAccess()),
		int(model.GetIDPerms().GetPermissions().GetGroupAccess()),
		string(model.GetIDPerms().GetPermissions().GetGroup()),
		string(model.GetIDPerms().GetLastModified()),
		bool(model.GetIDPerms().GetEnable()),
		string(model.GetIDPerms().GetDescription()),
		string(model.GetIDPerms().GetCreator()),
		string(model.GetIDPerms().GetCreated()),
		string(model.GetGCPSubnet()),
		string(model.GetGCPRegion()),
		int(model.GetGCPAsn()),
		string(model.GetGCPAccountInfo()),
		common.MustJSON(model.GetFQName()),
		string(model.GetDisplayName()),
		string(model.GetAwsSubnet()),
		string(model.GetAwsSecretKey()),
		string(model.GetAwsRegion()),
		string(model.GetAwsAccessKey()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "location",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "location", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanLocation(values map[string]interface{}) (*models.Location, error) {
	m := models.MakeLocation()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["type"]; ok {

		m.Type = common.InterfaceToString(value)

	}

	if value, ok := values["provisioning_state"]; ok {

		m.ProvisioningState = common.InterfaceToString(value)

	}

	if value, ok := values["provisioning_start_time"]; ok {

		m.ProvisioningStartTime = common.InterfaceToString(value)

	}

	if value, ok := values["provisioning_progress_stage"]; ok {

		m.ProvisioningProgressStage = common.InterfaceToString(value)

	}

	if value, ok := values["provisioning_progress"]; ok {

		m.ProvisioningProgress = common.InterfaceToInt64(value)

	}

	if value, ok := values["provisioning_log"]; ok {

		m.ProvisioningLog = common.InterfaceToString(value)

	}

	if value, ok := values["private_redhat_subscription_user"]; ok {

		m.PrivateRedhatSubscriptionUser = common.InterfaceToString(value)

	}

	if value, ok := values["private_redhat_subscription_pasword"]; ok {

		m.PrivateRedhatSubscriptionPasword = common.InterfaceToString(value)

	}

	if value, ok := values["private_redhat_subscription_key"]; ok {

		m.PrivateRedhatSubscriptionKey = common.InterfaceToString(value)

	}

	if value, ok := values["private_redhat_pool_id"]; ok {

		m.PrivateRedhatPoolID = common.InterfaceToString(value)

	}

	if value, ok := values["private_ospd_vm_vcpus"]; ok {

		m.PrivateOspdVMVcpus = common.InterfaceToString(value)

	}

	if value, ok := values["private_ospd_vm_ram_mb"]; ok {

		m.PrivateOspdVMRAMMB = common.InterfaceToString(value)

	}

	if value, ok := values["private_ospd_vm_name"]; ok {

		m.PrivateOspdVMName = common.InterfaceToString(value)

	}

	if value, ok := values["private_ospd_vm_disk_gb"]; ok {

		m.PrivateOspdVMDiskGB = common.InterfaceToString(value)

	}

	if value, ok := values["private_ospd_user_password"]; ok {

		m.PrivateOspdUserPassword = common.InterfaceToString(value)

	}

	if value, ok := values["private_ospd_user_name"]; ok {

		m.PrivateOspdUserName = common.InterfaceToString(value)

	}

	if value, ok := values["private_ospd_package_url"]; ok {

		m.PrivateOspdPackageURL = common.InterfaceToString(value)

	}

	if value, ok := values["private_ntp_hosts"]; ok {

		m.PrivateNTPHosts = common.InterfaceToString(value)

	}

	if value, ok := values["private_dns_servers"]; ok {

		m.PrivateDNSServers = common.InterfaceToString(value)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		m.Perms2.OwnerAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["owner"]; ok {

		m.Perms2.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["global_access"]; ok {

		m.Perms2.GlobalAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = common.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = common.InterfaceToString(value)

	}

	if value, ok := values["user_visible"]; ok {

		m.IDPerms.UserVisible = common.InterfaceToBool(value)

	}

	if value, ok := values["permissions_owner_access"]; ok {

		m.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["permissions_owner"]; ok {

		m.IDPerms.Permissions.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["other_access"]; ok {

		m.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["group_access"]; ok {

		m.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["group"]; ok {

		m.IDPerms.Permissions.Group = common.InterfaceToString(value)

	}

	if value, ok := values["last_modified"]; ok {

		m.IDPerms.LastModified = common.InterfaceToString(value)

	}

	if value, ok := values["enable"]; ok {

		m.IDPerms.Enable = common.InterfaceToBool(value)

	}

	if value, ok := values["description"]; ok {

		m.IDPerms.Description = common.InterfaceToString(value)

	}

	if value, ok := values["creator"]; ok {

		m.IDPerms.Creator = common.InterfaceToString(value)

	}

	if value, ok := values["created"]; ok {

		m.IDPerms.Created = common.InterfaceToString(value)

	}

	if value, ok := values["gcp_subnet"]; ok {

		m.GCPSubnet = common.InterfaceToString(value)

	}

	if value, ok := values["gcp_region"]; ok {

		m.GCPRegion = common.InterfaceToString(value)

	}

	if value, ok := values["gcp_asn"]; ok {

		m.GCPAsn = common.InterfaceToInt64(value)

	}

	if value, ok := values["gcp_account_info"]; ok {

		m.GCPAccountInfo = common.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["aws_subnet"]; ok {

		m.AwsSubnet = common.InterfaceToString(value)

	}

	if value, ok := values["aws_secret_key"]; ok {

		m.AwsSecretKey = common.InterfaceToString(value)

	}

	if value, ok := values["aws_region"]; ok {

		m.AwsRegion = common.InterfaceToString(value)

	}

	if value, ok := values["aws_access_key"]; ok {

		m.AwsAccessKey = common.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["backref_physical_router"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakePhysicalRouter()
			m.PhysicalRouters = append(m.PhysicalRouters, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["server_port"]; ok && propertyValue != nil {

				childModel.TelemetryInfo.ServerPort = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["server_ip"]; ok && propertyValue != nil {

				childModel.TelemetryInfo.ServerIP = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["resource"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.TelemetryInfo.Resource)

			}

			if propertyValue, ok := childResourceMap["physical_router_vnc_managed"]; ok && propertyValue != nil {

				childModel.PhysicalRouterVNCManaged = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["physical_router_vendor_name"]; ok && propertyValue != nil {

				childModel.PhysicalRouterVendorName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["username"]; ok && propertyValue != nil {

				childModel.PhysicalRouterUserCredentials.Username = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["password"]; ok && propertyValue != nil {

				childModel.PhysicalRouterUserCredentials.Password = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["version"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.Version = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["v3_security_name"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.V3SecurityName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["v3_security_level"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.V3SecurityLevel = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["v3_security_engine_id"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.V3SecurityEngineID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["v3_privacy_protocol"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.V3PrivacyProtocol = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["v3_privacy_password"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.V3PrivacyPassword = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["v3_engine_time"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.V3EngineTime = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["v3_engine_id"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.V3EngineID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["v3_engine_boots"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.V3EngineBoots = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["v3_context_engine_id"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.V3ContextEngineID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["v3_context"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.V3Context = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["v3_authentication_protocol"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["v3_authentication_password"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.V3AuthenticationPassword = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["v2_community"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.V2Community = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["timeout"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.Timeout = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["retries"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.Retries = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["local_port"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMPCredentials.LocalPort = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["physical_router_snmp"]; ok && propertyValue != nil {

				childModel.PhysicalRouterSNMP = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["physical_router_role"]; ok && propertyValue != nil {

				childModel.PhysicalRouterRole = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["physical_router_product_name"]; ok && propertyValue != nil {

				childModel.PhysicalRouterProductName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["physical_router_management_ip"]; ok && propertyValue != nil {

				childModel.PhysicalRouterManagementIP = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["physical_router_loopback_ip"]; ok && propertyValue != nil {

				childModel.PhysicalRouterLoopbackIP = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["physical_router_lldp"]; ok && propertyValue != nil {

				childModel.PhysicalRouterLLDP = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["service_port"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.PhysicalRouterJunosServicePorts.ServicePort)

			}

			if propertyValue, ok := childResourceMap["physical_router_image_uri"]; ok && propertyValue != nil {

				childModel.PhysicalRouterImageURI = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["physical_router_dataplane_ip"]; ok && propertyValue != nil {

				childModel.PhysicalRouterDataplaneIP = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListLocation lists Location with list spec.
func ListLocation(ctx context.Context, tx *sql.Tx, request *models.ListLocationRequest) (response *models.ListLocationResponse, err error) {
	var rows *sql.Rows
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "location"
	qb.Fields = LocationFields
	qb.RefFields = LocationRefFields
	qb.BackRefFields = LocationBackRefFields
	result := []*models.Location{}

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = common.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
	}

	query := qb.BuildQuery()
	columns := qb.Columns
	values := qb.Values
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, errors.Wrap(err, "select query failed")
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "row error")
	}

	for rows.Next() {
		valuesMap := map[string]interface{}{}
		values := make([]interface{}, len(columns))
		valuesPointers := make([]interface{}, len(columns))
		for _, index := range columns {
			valuesPointers[index] = &values[index]
		}
		if err := rows.Scan(valuesPointers...); err != nil {
			return nil, errors.Wrap(err, "scan failed")
		}
		for column, index := range columns {
			val := valuesPointers[index].(*interface{})
			valuesMap[column] = *val
		}
		m, err := scanLocation(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListLocationResponse{
		Locations: result,
	}
	return response, nil
}

// UpdateLocation updates a resource
func UpdateLocation(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdateLocationRequest,
) error {
	//TODO
	return nil
}

// DeleteLocation deletes a resource
func DeleteLocation(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeleteLocationRequest) error {
	deleteQuery := deleteLocationQuery
	selectQuery := "select count(uuid) from location where uuid = ?"
	var err error
	var count int
	uuid := request.ID
	auth := common.GetAuthCTX(ctx)
	if auth.IsAdmin() {
		row := tx.QueryRowContext(ctx, selectQuery, uuid)
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid)
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRowContext(ctx, selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid, auth.ProjectID())
	}

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	err = common.DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}
