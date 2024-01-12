package controllers

import (
	"net/http"

	respmodels "study_marketplace/pkg/domain/models/response_models"
	"study_marketplace/pkg/services"

	"github.com/gin-gonic/gin"
)

type CategoriesController interface {
	CatGetAll(ctx *gin.Context)
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
func (t *categoriesController) CatGetAll(ctx *gin.Context) {
	categories, err := t.categoriesService.CatGetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.FailedResponse{Data: err.Error(), Status: "failed"})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}
