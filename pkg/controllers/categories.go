package controllers

import (
	"net/http"

	"study_marketplace/pkg/domen/models"
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

// @Summary		GET all categories parents with children in array
// @Description	endpoint for getting all categories
// @Tags			categories/getall
// @Produce		json
// @Success		200	{object}	map[string]interface{}
// @Router			/open/categories/getall [get]
func (t *categoriesController) CatGetAll(ctx *gin.Context) {
	categories, err := t.categoriesService.CatGetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResponseSuccess(categories))
}
