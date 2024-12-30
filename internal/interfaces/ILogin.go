package interfaces

import (
	"context"
	"ewallet-framework/internal/models"

	"github.com/gin-gonic/gin"
)

type ILoginService interface {
	Login(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error)
}
type ILoginHandler interface {
	Login(c *gin.Context)
}
