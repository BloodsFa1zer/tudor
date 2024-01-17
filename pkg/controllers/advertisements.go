package controllers

import (
	"net/http"
	"strconv"
	reqm "study_marketplace/pkg/domain/mappers/reqresp_mappers"
	reqmodels "study_marketplace/pkg/domain/models/request_models"
	v "study_marketplace/pkg/infrastructure/validator"
	"study_marketplace/pkg/services"

	"github.com/gin-gonic/gin"
)

type AdvertisementsController interface {
	AdvCreate(ctx *gin.Context)
	AdvPatch(ctx *gin.Context)
	AdvDelete(ctx *gin.Context)
	AdvGetAll(ctx *gin.Context)
	AdvGetByID(ctx *gin.Context)
	AdvGetFiltered(ctx *gin.Context)
	AdvGetMy(ctx *gin.Context)
}

type advertisementsController struct {
	advertisementService services.AdvertisementService
}

func NewAdvertisementsController(sa services.AdvertisementService) AdvertisementsController {
	return &advertisementsController{sa}
}

// @Advertisement-create	godoc
// @Summary					POST request to create advertisement
// @Description				endpoint for advertisement creation
// @Tags					advertisement-create
// @Security				JWT
// @Param					Authorization			header	string						true	"Insert your access token"
// @Param					advertisement-create	body	reqmodels.CreateAdvertisementRequest	true	"advertisement information"
// @Produce					json
// @Success					200	{object}	respmodels.AdvertisementResponse
// @Failure					400	{object}	respmodels.FailedResponse
// @Router					/protected/advertisement-create [post]
func (c *advertisementsController) AdvCreate(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	var inputModel reqmodels.CreateAdvertisementRequest
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}

	if err := v.Validate(inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}

	advertisement, err := c.advertisementService.AdvCreate(ctx,
		reqm.CreateAdvRequestToAdvertisement(&inputModel, userID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, reqm.AdvertisementToCreateUpdateAdvertisementResponse(advertisement))
}

// @Advertisement-patch		godoc
// @Summary					PATCH request to update advertisement
// @Description				endpoint for advertisement update
// @Tags					advertisement-patch
// @Security				JWT
// @Param					Authorization		header	string						true	"Insert your access token"
// @Param					advertisement-patch	body	reqmodels.UpdateAdvertisementRequest	true	"advertisement information"
// @Produce					json
// @Success					200	{object}	respmodels.AdvertisementResponse
// @Failure					400	{object}	respmodels.FailedResponse
// @Router					/protected/advertisement-patch [patch]
func (c *advertisementsController) AdvPatch(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	var inputModel reqmodels.UpdateAdvertisementRequest
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	if err := v.Validate(inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	advertisement, err := c.advertisementService.AdvPatch(ctx,
		reqm.UpdateAdvRequestToAdvertisement(&inputModel, userID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, reqm.AdvertisementToCreateUpdateAdvertisementResponse(advertisement))
}

// @Advertisement-delete	godoc
// @Summary					DELETE request to delete advertisement
// @Description				endpoint for advertisement deletion by id
// @Tags					advertisement-delete
// @Security				JWT
// @Param					Authorization			header	string		true	"Insert your access token"
// @Param					advertisement-delete	body	reqmodels.DeleteAdvertisementRequest	true	"advertisement id"
// @Produce					json
// @Success					200	{object} respmodels.StringResponse
// @Failure					400	{object} respmodels.FailedResponse
// @Router					/protected/advertisement-delete [delete]
func (c *advertisementsController) AdvDelete(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	if userID == 0 {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse("user id error"))
		return
	}
	var inputModel reqmodels.DeleteAdvertisementRequest
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	if err := c.advertisementService.AdvDelete(ctx, inputModel.ID, userID); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, reqm.StrResponse("Advertisement deleted"))
}

// @Summary			GET request to get 10 items sorted by creation date in desc order
// @Description		endpoint for getting all advertisements
// @Tags			advertisements-getall
// @Produce			json
// @Success			200	{object}	respmodels.AdvertisementsResponse
// @Failure			400	{object}	respmodels.FailedResponse
// @Router			/open/advertisements/getall [get]
func (t *advertisementsController) AdvGetAll(ctx *gin.Context) {
	advertisements, err := t.advertisementService.AdvGetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, reqm.AdvertisementsToAdvertisementsResponses(advertisements))
}

// @Summary			GET request to get advertisement by id
// @Description		endpoint to get advertisement based on it's id
// @Tags			open/advertisements/getbyid/{id}
// @Security		JWT
// @Param			id	path	int	true	"advertisement ID"
// @Produce			json
// @Success			200	{object}	respmodels.AdvertisementResponse
// @Failure			400	{object}	respmodels.FailedResponse
// @Router			/open/advertisements/getbyid/{id} [get]
func (c *advertisementsController) AdvGetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	advertisement, err := c.advertisementService.AdvGetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, reqm.AdvertisementToCreateUpdateAdvertisementResponse(advertisement))
}

// @Advertisement-filter	godoc
// @Summary					POST request to get advertisement based on params in filter
// @Description				endpoint for getting specific advertisements
// @Tags					advertisement-filter
// deprecated(@Security				JWT)
// deprecated(@Param					Authorization			header	string						true	"Insert your access token"))
// @Param					advertisement-filter	body	reqmodels.AdvertisementFilterRequest	true	"advertisement filter"
// @Produce					json
// @Success					200	{object}	respmodels.AdvertisementPaginationResponse
// @Failure					400	{object}	respmodels.FailedResponse
// @Router					/open/advertisements/adv-filter [post]
func (c *advertisementsController) AdvGetFiltered(ctx *gin.Context) {
	var filter reqmodels.AdvertisementFilterRequest
	err := ctx.ShouldBindJSON(&filter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	if err := v.Validate(filter); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	advertisements, err := c.advertisementService.AdvGetFiltered(ctx, &filter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, reqm.AdvertisementPaginationToAdvertisementPaginationResponse(advertisements))
}

// @Summary			GET request to get user created advertisements
// @Description		endpoint for getting user advertisements
// @Security		JWT
// @Param			Authorization	header	string	true	"Insert your access token"
// @Tags			advertisements-getmy
// @Produce			json
// @Success			200	{object}	respmodels.AdvertisementsResponse
// @Failure			400	{object}	respmodels.FailedResponse
// @Router			/protected/advertisement-getmy [get]
func (c *advertisementsController) AdvGetMy(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse("Unauthorized."))
		return
	}
	advertisements, err := c.advertisementService.AdvGetMy(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, reqm.AdvertisementsToAdvertisementsResponses(advertisements))
}
