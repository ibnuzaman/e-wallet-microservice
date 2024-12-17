package cmd

import (
	"ewallet-framework/helpers"
	"ewallet-framework/internal/api"
	"ewallet-framework/internal/repository"
	"ewallet-framework/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServerHttp() {
	healthcheckSvc := &services.Healthcheck{}
	healthchekcAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	regisRepo := &repository.UserRepository{
		DB: helpers.DB,
	}

	regisSvc := &services.RegisterService{
		RegisterRepo: regisRepo,
	}

	regisAPI := &api.RegisterHandler{
		RegisterService: regisSvc,
	}

	r := gin.Default()

	r.GET("/health", healthchekcAPI.HealcheckHandlerHTTP)

	userV1 := r.Group("/v1/user")
	userV1.POST("/register", regisAPI.Register)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}

}
