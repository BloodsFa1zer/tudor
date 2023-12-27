package controllers

import (
	"net/http"
	respmodels "study_marketplace/pkg/domen/models/response_models"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary		Show the status of server.
// @Description	get the status of server.
// @Tags			root
// @Accept			*/*
// @Produce		json
// @Success		200	{object}	map[string]interface{}
// @Router			/api/ [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, respmodels.NewResponseSuccess("Server up and running."))
}
