package controllers

import (
	"net/http"
	reqm "study_marketplace/pkg/domain/mappers/reqresp_mappers"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary		Show the status of server.
// @Description	get the status of server.
// @Tags			root
// @Accept			*/*
// @Produce		json
// @Success		200	{object}	respmodels.StringResponse
// @Router			/api/ [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, reqm.StrResponse("Server up and running."))
}
