package api

import (
	"ewallet-framework/helpers"
	"ewallet-framework/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Healthcheck struct {
	HealthcheckServices interfaces.IHealthcheckServices
}

func (api *Healthcheck) HealcheckHandlerHTTP(c *gin.Context) {
	msg, err := api.HealthcheckServices.HealthcheckServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": msg})
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, msg, nil)
}
