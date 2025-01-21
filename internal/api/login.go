package api

import (
	"ewallet-framework/helpers"
	"ewallet-framework/internal/constants"
	"ewallet-framework/internal/interfaces"
	"ewallet-framework/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginService interfaces.ILoginService
}

func (api *LoginHandler) Login(c *gin.Context) {
	var (
		log  = helpers.Logger
		req  models.LoginRequest
		resp models.LoginResponse
	)

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

	user, err := api.LoginService.Login(c.Request.Context(), req)
	if err != nil {
		log.Error("Internal Server Error ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	resp = models.LoginResponse{
		UserID:       user.UserID,
		Username:     user.Username,
		Email:        user.Email,
		FullName:     user.FullName,
		Token:        user.Token,
		RefreshToken: user.RefreshToken,
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, resp)

}
