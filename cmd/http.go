package cmd

import (
	"ewallet-framework/helpers"
	"ewallet-framework/internal/api"
	"ewallet-framework/internal/interfaces"
	"ewallet-framework/internal/repository"
	"ewallet-framework/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServerHttp() {

	dependency := dependencyInject()

	r := gin.Default()

	r.GET("/health", dependency.HealthchekAPI.HealcheckHandlerHTTP)

	userV1 := r.Group("/v1/user")
	userV1.POST("/register", dependency.RegisterAPI.Register)
	userV1.POST("/login", dependency.LoginAPI.Login)
	userV1.DELETE("/logout", dependency.MiddlewareValidateAuth, dependency.LogoutAPI.Logout)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}

}

type Dependency struct {
	UserRepository interfaces.IUserRepository

	HealthchekAPI interfaces.IHealthcheckHandler
	RegisterAPI   interfaces.IRegisHandler
	LoginAPI      interfaces.ILoginHandler
	LogoutAPI     interfaces.ILogoutHandler
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

	logoutSvc := &services.LogoutService{
		UserRepo: userRepo,
	}

	logoutAPI := &api.LogoutHandler{
		LogoutService: logoutSvc,
	}

	return Dependency{
		UserRepository: userRepo,
		HealthchekAPI:  healthchekcAPI,
		RegisterAPI:    regisAPI,
		LoginAPI:       loginAPI,
		LogoutAPI:      logoutAPI,
	}

}
