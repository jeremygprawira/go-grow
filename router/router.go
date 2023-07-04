package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	handler "go-grow/delivery/http"
)

var r *gin.Engine

func InitRouter(user *handler.UserHandler) {
	r = gin.Default()
	r.Static("/images", "./mock/images")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	api := r.Group("v1.0/api")

	api.POST("/user/register", user.RegisterUser)
	api.POST("/user/is-email-available", user.EmailAvailability)
	
	/*apiAuth := r.Group("v1/api")
	apiAuth.Use(middleware.AuthMiddleWare())*/
	
}

func Start() error {
	return r.Run(":8080")
}