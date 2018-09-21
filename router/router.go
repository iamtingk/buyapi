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
	router.GET("/products", ShowProducts)
	router.PUT("/product/:id", UpdateProduct)
	router.DELETE("/product/:id", DestroyProduct)

	router.POST("/member/signup", MemberSignUp)
	router.POST("/member/signin", MemberSignIn)

	router.POST("/test", TestOrder)
	router.POST("/order/create", CreateOrder)
	router.POST("/order/query", ShowOrders)
	router.POST("/order/querydetail", ShowOrderDetail)
	router.DELETE("/order/delete", DeleteOrder)

	router.Static("/image", config.IMAGE_PATH)

	return router
}
