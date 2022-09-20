package main

import (
	"digitalmoney/api/controllers"
	"digitalmoney/api/middlewares"
	"digitalmoney/api/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	repository.Connect()
	repository.Migrate()
	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", controllers.Ping)
	api := router.Group("/api")
	{
		api.POST("/login", controllers.GenerateToken)

		apiUser := api.Group("/users") 
		{
			apiUser.POST("", controllers.RegisterUser)
			apiUser.PUT("/:id", controllers.Ping).Use(middlewares.Auth())
		}

		accountApi := api.Group("/account")
		{
			accountApi.GET("", controllers.FindAccount).Use(middlewares.Auth())
		}
		
		cardsApi := accountApi.Group("/:account_id/cards").Use(middlewares.Auth())
		{
			cardsApi.GET("", controllers.GetCardsByAccountID)
			cardsApi.POST("", controllers.CreateCard)
			cardsApi.GET("/:id", controllers.Ping)
		}

		transfApi := accountApi.Group("/:account_id/transferences").Use(middlewares.Auth())
		{
			transfApi.GET("", controllers.GetTransferencesByAccountID)
			transfApi.POST("", controllers.CreateTransference)
			transfApi.GET("/:id", controllers.Ping)
		}
	}
	return router
}
