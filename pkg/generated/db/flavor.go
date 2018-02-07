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

const insertFlavorQuery = "insert into `flavor` (`vcpus`,`uuid`,`swap`,`rxtx_factor`,`ram`,`property`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`name`,`type`,`rel`,`href`,`is_public`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`id`,`fq_name`,`ephemeral`,`display_name`,`disk`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteFlavorQuery = "delete from `flavor` where uuid = ?"

// FlavorFields is db columns for Flavor
var FlavorFields = []string{
	"vcpus",
	"uuid",
	"swap",
	"rxtx_factor",
	"ram",
	"property",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"name",
	"type",
	"rel",
	"href",
	"is_public",
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
	"id",
	"fq_name",
	"ephemeral",
	"display_name",
	"disk",
	"key_value_pair",
}

// FlavorRefFields is db reference fields for Flavor
var FlavorRefFields = map[string][]string{}

// FlavorBackRefFields is db back reference fields for Flavor
var FlavorBackRefFields = map[string][]string{}

// FlavorParentTypes is possible parents for Flavor
var FlavorParents = []string{}

// CreateFlavor inserts Flavor to DB
func CreateFlavor(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreateFlavorRequest) error {
	model := request.Flavor
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFlavorQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertFlavorQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, int(model.Vcpus),
		string(model.UUID),
		int(model.Swap),
		int(model.RXTXFactor),
		int(model.RAM),
		string(model.Property),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		string(model.Name),
		string(model.Links.Type),
		string(model.Links.Rel),
		string(model.Links.Href),
		bool(model.IsPublic),
		bool(model.IDPerms.UserVisible),
		int(model.IDPerms.Permissions.OwnerAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OtherAccess),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Group),
		string(model.IDPerms.LastModified),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Creator),
		string(model.IDPerms.Created),
		string(model.ID),
		common.MustJSON(model.FQName),
		int(model.Ephemeral),
		string(model.DisplayName),
		int(model.Disk),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "flavor",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "flavor", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanFlavor(values map[string]interface{}) (*models.Flavor, error) {
	m := models.MakeFlavor()

	if value, ok := values["vcpus"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Vcpus = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["swap"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Swap = castedValue

	}

	if value, ok := values["rxtx_factor"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.RXTXFactor = castedValue

	}

	if value, ok := values["ram"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.RAM = castedValue

	}

	if value, ok := values["property"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Property = castedValue

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["parent_uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentUUID = castedValue

	}

	if value, ok := values["parent_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentType = castedValue

	}

	if value, ok := values["name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Name = castedValue

	}

	if value, ok := values["type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Links.Type = castedValue

	}

	if value, ok := values["rel"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Links.Rel = castedValue

	}

	if value, ok := values["href"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Links.Href = castedValue

	}

	if value, ok := values["is_public"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IsPublic = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["ephemeral"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Ephemeral = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["disk"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Disk = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListFlavor lists Flavor with list spec.
func ListFlavor(ctx context.Context, tx *sql.Tx, request *models.ListFlavorRequest) (response *models.ListFlavorResponse, err error) {
	var rows *sql.Rows
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "flavor"
	qb.Fields = FlavorFields
	qb.RefFields = FlavorRefFields
	qb.BackRefFields = FlavorBackRefFields
	result := models.MakeFlavorSlice()

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filter.AppendValues("parent_uuid", []string{parentMetaData.UUID})
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
		m, err := scanFlavor(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListFlavorResponse{
		Flavors: result,
	}
	return response, nil
}

// UpdateFlavor updates a resource
func UpdateFlavor(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdateFlavorRequest,
) error {
	//TODO
	return nil
}

// DeleteFlavor deletes a resource
func DeleteFlavor(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeleteFlavorRequest) error {
	deleteQuery := deleteFlavorQuery
	selectQuery := "select count(uuid) from flavor where uuid = ?"
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