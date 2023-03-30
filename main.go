package main

import (
	"github.com/Noush-012/JWT/controllers"
	"github.com/Noush-012/JWT/initializer"
	"github.com/Noush-012/JWT/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnToDB()
	initializer.SyncDatabase()
}

func main() {

	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.UserAuth, controllers.Validate)
	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
