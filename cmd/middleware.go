package cmd

import (
	"ewallet-framework/helpers"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateAuth(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("authorization empty")
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	_, err := d.UserRepository.GetUserSessionByToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println("failed to get user session on DB: ", err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	claim, err := helpers.ValidateToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("jwt token is expired: ", claim.ExpiresAt)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	ctx.Set("token", claim)

	ctx.Next()
}
