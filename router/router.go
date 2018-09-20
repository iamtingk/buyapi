package router

import (
	. "buyapi/apis"
	"buyapi/config"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", Users)

	router.POST("/user", Storee)

	router.PUT("/user/:id", Update)

	router.DELETE("/user/:id", Destroy)

	router.POST("/product", CreateProduct)
	router.GET("/products", Products)
	router.PUT("/product/:id", UpdateProduct)

	router.Static("/image", config.IMAGE_PATH)

	return router
}
