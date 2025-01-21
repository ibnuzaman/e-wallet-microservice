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
	RegisterService interfaces.IRegisService
}

func (api *RegisterHandler) Register(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	req := models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Failed to parse request", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadParseRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadParseRequest, nil)
		return
	}

	resp, err := api.RegisterService.Register(c.Request.Context(), req)
	if err != nil {
		log.Error("Failed to register user", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrEmailorUsernameAlreadyExist, nil)
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, resp)
}
