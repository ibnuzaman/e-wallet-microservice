package interfaces

import "github.com/gin-gonic/gin"

type IHealthcheckServices interface {
	HealthcheckServices() (string, error)
}
type IHealthcheckHandler interface {
	HealcheckHandlerHTTP(c *gin.Context)
}

type IHealthcheckRepo interface {
}
