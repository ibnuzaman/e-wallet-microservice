package api

import (
	"ewallet-framework/helpers"
	"ewallet-framework/internal/constants"
	"ewallet-framework/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogoutHandler struct {
	LogoutService interfaces.ILogoutService
}

func (api *LogoutHandler) Logout(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	token := c.GetHeader("Authorization")

	err := api.LogoutService.Logout(c.Request.Context(), token)
	if err != nil {
		log.Error("Internal Server Error ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, nil)
}
