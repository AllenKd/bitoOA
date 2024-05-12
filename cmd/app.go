package main

import (
	"bitoOA/internal/config"
	"bitoOA/internal/route/handler"
	"bitoOA/internal/route/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	route          *gin.Engine
	UserController *handler.UserController
	Config         *config.Config
}

func (a *App) SetupRoutes() *gin.Engine {
	a.route = gin.Default()
	a.route.Use(middleware.Log)

	apiV1 := a.route.Group("/api/v1")

	user := apiV1.Group("/user")
	user.POST("", a.UserController.AddSinglePersonAndMatch)

	users := apiV1.Group("/users")
	users.DELETE("/:userId", a.UserController.RemoveSinglePerson)

	matches := apiV1.Group("/matches")
	matches.GET("/popular-users", a.UserController.QuerySinglePeople)
	matches.POST("/like", a.UserController.Like)

	if a.Config.Env == "local" {
		a.route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return a.route
}

func NewApp(userController *handler.UserController, config *config.Config) *App {
	return &App{
		UserController: userController,
		Config:         config,
	}
}
