package main

import (
	"github.com/gin-gonic/gin"
	"web-service-gin/controller"
	"web-service-gin/database"
	_ "web-service-gin/database"
	"web-service-gin/middleware"
	"web-service-gin/model"
	_ "web-service-gin/model"
)

func loadDatabase() {
	database.Connect()
	err := database.Database.AutoMigrate(&model.Company{})
	err = database.Database.AutoMigrate(&model.User{})
	if err != nil {
		return
	}
}
func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())

	protectedRoutes.POST("/comp", controller.AddRecordController)

	protectedRoutes.GET("/comp", controller.GetRecordController)

	protectedRoutes.GET("/comp/:id", controller.GetRecordByIdController)

	protectedRoutes.PUT("/comp/:id", controller.UpdateRecordController)

	protectedRoutes.DELETE("/comp/:id", controller.DeleteRecordController)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func main() {
	loadDatabase()
	serveApplication()
}
