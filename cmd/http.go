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

	// healthchekcAPI := dependencyInject().HealthchekAPI
	// regisAPI := dependencyInject().RegisterAPI

	dependency := dependencyInject()

	r := gin.Default()

	r.GET("/health", dependency.HealthchekAPI.HealcheckHandlerHTTP)

	userV1 := r.Group("/v1/user")
	userV1.POST("/register", dependency.RegisterAPI.Register)
	userV1.POST("/login", dependency.LoginAPI.Login)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}

}

type Dependency struct {
	HealthchekAPI *api.Healthcheck
	RegisterAPI   *api.RegisterHandler
	LoginAPI      *api.LoginHandler
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthchekcAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}

	regisSvc := &services.RegisterService{
		RegisterRepo: userRepo,
	}

	regisAPI := &api.RegisterHandler{
		RegisterService: regisSvc,
	}

	loginSvc := &services.LoginService{
		UserRepo: userRepo,
	}

	loginAPI := &api.LoginHandler{
		LoginService: loginSvc,
	}

	return Dependency{
		HealthchekAPI: healthchekcAPI,
		RegisterAPI:   regisAPI,
		LoginAPI:      loginAPI,
	}

}
