package routes

import "github.com/gin-gonic/gin"

func RootRoute(apiRoute *gin.RouterGroup) {
	var userRoute = apiRoute.Group("user")
	var productRoute = apiRoute.Group("product")
	RouteUser(userRoute)
	RouteProduct(productRoute)
}
