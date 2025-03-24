package main

import (
	"authservice/handlers"
	"authservice/middleware"
	"authservice/models"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DbInit() *gorm.DB {
	db, err := models.Setup()
	if err != nil {
		log.Println("Connection error")
	}
	return db
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db := DbInit()

	server := handlers.NewServer(db)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router := r.Group("/api")
	r.Static("/uploads", "./uploads")

	router.POST("/register", server.Register)
	router.POST("/login", server.Login)

	auth := r.Group("/auth")
	auth.Use(middleware.JWTMiddleware())
	auth.GET("/profile", server.GetUserInfo)
	auth.PATCH("/update-profile", server.UpdateUserInfo)

	return r

}

func main() {
	r := SetupRouter()

	r.Run(":8080")
}
