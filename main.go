package main

import (
	"github.com/car-api/controller"
	"github.com/car-api/db"
	docs "github.com/car-api/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Car Api
// @version		1.0
// @description	This is a sample page
// @termsOfService	http://swagger.io/terms/
func main() {
	db.InitDb()
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1/users"

	r.POST("/signup", controller.SignUp)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8081")

}
