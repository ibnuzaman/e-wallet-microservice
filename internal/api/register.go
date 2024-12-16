package api

import (
	"ewallet-framework/helpers"
	"ewallet-framework/internal/constants"
	"ewallet-framework/internal/interfaces"
	"ewallet-framework/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	RegisterService interfaces.IRegisterService
}

func (api *RegisterHandler) Register(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	req := models.User{}

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

	resp, err := api.RegisterService.Register(c.Request.Context(), req)
	if err != nil {
		log.Error("Failed to register user", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, resp)
}
