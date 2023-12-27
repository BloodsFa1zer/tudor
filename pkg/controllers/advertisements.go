package controllers

import (
	"net/http"
	"strconv"
	reqm "study_marketplace/pkg/domen/mappers/reqresp_mappers"
	reqmodels "study_marketplace/pkg/domen/models/request_models"
	respmodels "study_marketplace/pkg/domen/models/response_models"
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
// @Param					advertisement-create	body	models.AdvertisementInput	true	"advertisement information"
// @Produce					json
// @Success					200	{object}	map[string]interface{}
// @Router					/protected/advertisement-create [post]
func (c *advertisementsController) AdvCreate(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	if userID == 0 {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed("user id error"))
		return
	}
	var inputModel reqmodels.CreateUpdateAdvertisementRequest
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed(err.Error()))
		return
	}
	advertisement, err := c.advertisementService.AdvCreate(ctx,
		reqm.CreateUpdateAdvRequestToAdvertisement(&inputModel, userID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, respmodels.NewResponseSuccess(
		reqm.AdvertisementToCreateUpdateAdvertisementResponse(advertisement)))
}

// @Advertisement-patch		godoc
// @Summary					PATCH request to update advertisement
// @Description				endpoint for advertisement update
// @Tags					advertisement-patch
// @Security				JWT
// @Param					Authorization		header	string						true	"Insert your access token"
// @Param					advertisement-patch	body	reqmodels.CreateUpdateAdvertisementRequest	true	"advertisement information"
// @Produce					json
// @Success					200	{object}	map[string]interface{}
// @Router					/protected/advertisement-patch [patch]
func (c *advertisementsController) AdvPatch(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	var inputModel reqmodels.CreateUpdateAdvertisementRequest
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed(err.Error()))
		return
	}
	advertisement, err := c.advertisementService.AdvPatch(ctx,
		reqm.CreateUpdateAdvRequestToAdvertisement(&inputModel, userID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, respmodels.NewResponseSuccess(
		reqm.AdvertisementToCreateUpdateAdvertisementResponse(advertisement)))
}

// @Advertisement-delete	godoc
// @Summary					DELETE request to delete advertisement
// @Description				endpoint for advertisement deletion by id
// @Tags					advertisement-delete
// @Security				JWT
// @Param					Authorization			header	string		true	"Insert your access token"
// @Param					advertisement-delete	body	reqmodels.DeleteAdvertisementRequest	true	"advertisement id"
// @Produce					json
// @Success					200	{object}	map[string]interface{}
// @Router					/protected/advertisement-delete [delete]
func (c *advertisementsController) AdvDelete(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	if userID == 0 {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed("user id error"))
		return
	}
	var inputModel reqmodels.DeleteAdvertisementRequest
	err := ctx.ShouldBindJSON(&inputModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed(err.Error()))
		return
	}
	err = c.advertisementService.AdvDelete(ctx, inputModel.ID, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, respmodels.NewResponseSuccess("advertisement is deleted"))
}

// @Summary			GET request to get 10 items sorted by creation date in desc order
// @Description		endpoint for getting all advertisements
// @Tags			advertisements-getall
// @Produce			json
// @Success			200	{object}	map[string]interface{}
// @Router			/open/advertisements/getall [get]
func (t *advertisementsController) AdvGetAll(ctx *gin.Context) {
	advertisements, err := t.advertisementService.AdvGetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, respmodels.NewResponseSuccess(
		reqm.AdvertisementsToAdvertisementResponses(advertisements)))
}

// @Summary		GET request to get advertisement by id
// @Description	endpoint to get advertisement based on it's id
// @Tags			open/advertisements/getbyid/{id}
// @Security		JWT
// @Param			id	path	int	true	"advertisement ID"
// @Produce		json
// @Success		200	{object}	map[string]interface{}
// @Router			/open/advertisements/getbyid/{id} [get]
func (c *advertisementsController) AdvGetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed("Faield to get advertisement ID"))
		return
	}
	advertisement, err := c.advertisementService.AdvGetByID(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, respmodels.NewResponseSuccess(
		reqm.AdvertisementToCreateUpdateAdvertisementResponse(advertisement)))
}

// @Advertisement-filter	godoc
// @Summary				POST request to get advertisement based on params in filter
// @Description			endpoint for getting specific advertisements
// @Tags					advertisement-filter
// @Security				JWT
// @Param					Authorization			header	string						true	"Insert your access token"
// @Param					advertisement-filter	body	models.AdvertisementFilter	true	"advertisement filter"
// @Produce				json
// @Success				200	{object}	map[string]interface{}
// @Router					/protected/advertisement-filter [post]
func (c *advertisementsController) AdvGetFiltered(ctx *gin.Context) {
	var filter reqmodels.AdvertisementFilterRequest
	err := ctx.ShouldBindJSON(&filter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed("Filter params not provided."))
		return
	}
	advertisements, err := c.advertisementService.AdvGetFiltered(ctx, &filter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, respmodels.NewResponseSuccess(
		reqm.AdvertisementPaginationToAdvertisementPaginationResponse(advertisements)))
}

// @Summary		GET request to get user created advertisements
// @Description	endpoint for getting user advertisements
// @Security		JWT
// @Param			Authorization	header	string	true	"Insert your access token"
// @Tags			advertisements-getmy
// @Produce		json
// @Success		200	{object}	map[string]interface{}
// @Router			/protected/advertisement-getmy [get]
func (c *advertisementsController) AdvGetMy(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed("Unauthorized."))
		return
	}
	advertisements, err := c.advertisementService.AdvGetMy(ctx, userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, respmodels.NewResponseSuccess(
		reqm.AdvertisementsToAdvertisementResponses(advertisements)))
}
