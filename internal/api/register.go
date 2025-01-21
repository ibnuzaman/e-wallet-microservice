package api

import (
	"ewallet-framework/helpers"
	"ewallet-framework/internal/constants"
	"ewallet-framework/internal/interfaces"
	"ewallet-framework/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterHandler struct {
	RegisterService interfaces.IRegisService
}

func (api *RegisterHandler) Register(c *gin.Context) {
	validate := validator.New()
	var (
		log = helpers.Logger
	)
	req := models.User{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Failed to bind JSON", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadParseRequest, nil)
		return
	}
	if err := validate.Struct(&req); err != nil {

		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {

			validationErrors = append(validationErrors, err.Field()+" is "+err.Tag())
		}

		log.Error("Validation failed", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadParseRequest, validationErrors)
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
