package interfaces

import (
	"context"
	"ewallet-framework/internal/models"

	"github.com/gin-gonic/gin"
)

type IRegisService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}
type IRegisHandler interface {
	Register(c *gin.Context)
}
