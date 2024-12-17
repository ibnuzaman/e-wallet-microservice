package api

import (
	"ewallet-framework/helpers"
	"ewallet-framework/internal/constants"
	"ewallet-framework/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
}

func (api *LoginHandler) Login(c *gin.Context) {
	var (
		log  = helpers.Logger
		req  = models.LoginRequest{}
		resp = models.LoginResponse{}
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Failed to parse request", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadParseRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadParseRequest, nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, resp)

}
