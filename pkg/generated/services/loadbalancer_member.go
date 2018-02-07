package services

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"

	log "github.com/sirupsen/logrus"
)

//RESTLoadbalancerMemberUpdateRequest for update request for REST.
type RESTLoadbalancerMemberUpdateRequest struct {
	Data map[string]interface{} `json:"loadbalancer-member"`
}

//RESTCreateLoadbalancerMember handle a Create REST service.
func (service *ContrailService) RESTCreateLoadbalancerMember(c echo.Context) error {
	requestData := &models.CreateLoadbalancerMemberRequest{
		LoadbalancerMember: models.MakeLoadbalancerMember(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_member",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateLoadbalancerMember(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateLoadbalancerMember handle a Create API
func (service *ContrailService) CreateLoadbalancerMember(
	ctx context.Context,
	request *models.CreateLoadbalancerMemberRequest) (*models.CreateLoadbalancerMemberResponse, error) {
	model := request.LoadbalancerMember
	if model.UUID == "" {
		model.UUID = uuid.NewV4().String()
	}
	auth := common.GetAuthCTX(ctx)
	if auth == nil {
		return nil, common.ErrorUnauthenticated
	}

	if model.FQName == nil {
		if model.DisplayName == "" {
			return nil, common.ErrorBadRequest("Both of FQName and Display Name is empty")
		}
		model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.DisplayName}
	}
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.CreateLoadbalancerMember(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_member",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateLoadbalancerMemberResponse{
		LoadbalancerMember: request.LoadbalancerMember,
	}, nil
}

//RESTUpdateLoadbalancerMember handles a REST Update request.
func (service *ContrailService) RESTUpdateLoadbalancerMember(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateLoadbalancerMemberRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_member",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateLoadbalancerMember(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateLoadbalancerMember handles a Update request.
func (service *ContrailService) UpdateLoadbalancerMember(
	ctx context.Context,
	request *models.UpdateLoadbalancerMemberRequest) (*models.UpdateLoadbalancerMemberResponse, error) {
	model := request.LoadbalancerMember
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateLoadbalancerMember(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_member",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateLoadbalancerMemberResponse{
		LoadbalancerMember: model,
	}, nil
}

//RESTDeleteLoadbalancerMember delete a resource using REST service.
func (service *ContrailService) RESTDeleteLoadbalancerMember(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteLoadbalancerMemberRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteLoadbalancerMember(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteLoadbalancerMember delete a resource.
func (service *ContrailService) DeleteLoadbalancerMember(ctx context.Context, request *models.DeleteLoadbalancerMemberRequest) (*models.DeleteLoadbalancerMemberResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteLoadbalancerMember(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteLoadbalancerMemberResponse{
		ID: request.ID,
	}, nil
}

//RESTGetLoadbalancerMember a REST Get request.
func (service *ContrailService) RESTGetLoadbalancerMember(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetLoadbalancerMemberRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetLoadbalancerMember(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetLoadbalancerMember a Get request.
func (service *ContrailService) GetLoadbalancerMember(ctx context.Context, request *models.GetLoadbalancerMemberRequest) (response *models.GetLoadbalancerMemberResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListLoadbalancerMemberRequest{
		Spec: spec,
	}
	var result *models.ListLoadbalancerMemberResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListLoadbalancerMember(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.LoadbalancerMembers) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetLoadbalancerMemberResponse{
		LoadbalancerMember: result.LoadbalancerMembers[0],
	}
	return response, nil
}

//RESTListLoadbalancerMember handles a List REST service Request.
func (service *ContrailService) RESTListLoadbalancerMember(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListLoadbalancerMemberRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListLoadbalancerMember(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListLoadbalancerMember handles a List service Request.
func (service *ContrailService) ListLoadbalancerMember(
	ctx context.Context,
	request *models.ListLoadbalancerMemberRequest) (response *models.ListLoadbalancerMemberResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListLoadbalancerMember(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}