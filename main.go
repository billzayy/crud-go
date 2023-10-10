package main

import (
	"os"

	"github.com/billzayy/crud-go/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/billzayy/crud-go/docs"
)

func main() {
	godotenv.Load()
	// gin.SetMode(gin.ReleaseMode)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example GoLang CRUD API"
	docs.SwaggerInfo.Description = "This is a sample Golang server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + os.Getenv("PORT")
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Link"},
		AllowAllOrigins:  true,
		AllowCredentials: false,
		MaxAge:           300,
	}))

	api := app.Group("/api")
	{
		api.GET("/get-all-user", controller.GetUserController)
		api.GET("/get-user/:id", controller.GetUserById)
		api.POST("/create-user", controller.CreateUserController)
		api.PUT("/update-user", controller.UpdateUserController)
		api.DELETE("/delete-user", controller.DeleteUserController)
	}

	// use ginSwagger middleware to serve the API docs
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Run(":" + os.Getenv("PORT"))
}
