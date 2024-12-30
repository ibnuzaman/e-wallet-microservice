package cmd

import (
	"ewallet-framework/helpers"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MiddlewareValidateAuth(c *gin.Context, dependency Dependency) {

	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("authorization empty")
		helpers.SendResponse(c, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	_, err := dependency.UserRepository.GetUserSessionByToken(c.Request.Context(), auth)
	if err != nil {
		log.Println("failed to get user session on DB: ", err)
		helpers.SendResponse(c, http.StatusUnauthorized, "unauthorized", nil)
	}

	claim, err := helpers.ValidateToken(c.Request.Context(), auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponse(c, http.StatusUnauthorized, "unauthorized", nil)
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("jwt token is expired: ", claim.ExpiresAt)
		helpers.SendResponse(c, http.StatusUnauthorized, "unauthorized", nil)
	}
	c.Set("username", claim)

	c.Next()

}
