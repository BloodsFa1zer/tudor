package controllers

import (
	"net/http"

	reqm "study_marketplace/pkg/domain/mappers/reqresp_mappers"
	"study_marketplace/pkg/services"

	"github.com/gin-gonic/gin"
)

type CategoriesControllerInterface interface {
	CategoriesGetAll(ctx *gin.Context)
}

type categoriesController struct {
	categoriesService services.CategoriesService
}

func NewCatController(sc services.CategoriesService) *categoriesController {
	return &categoriesController{sc}
}

// @Summary			GET all categories parents with children in array
// @Description		endpoint for getting all categories
// @Tags			open/allcategories
// @Produce			json
// @Success			200	{object}	[]queries.GetCategoriesWithChildrenRow
// @Router			/open/allcategories [get]
func (t *categoriesController) CategoriesGetAll(ctx *gin.Context) {
	categories, err := t.categoriesService.CatGetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, categories)
}
