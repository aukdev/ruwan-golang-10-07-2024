package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouteProduct(route *gin.RouterGroup) {
	route.GET("/", getAllProducts)
}

// get all products
func getAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "All product",
	})
}
