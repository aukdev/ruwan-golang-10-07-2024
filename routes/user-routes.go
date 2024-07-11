package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginInfo struct {
	Username string `json:"username"`
	Password string `json:"pass"`
}

type registerInfo struct {
	Username string `json:"username"`
	Password string `json:"pass"`
	Email    string `json:"email"`
}

func RouteUser(route *gin.RouterGroup) {
	route.POST("/login", userLogin)
	route.POST("/register", userRegister)
}

func userLogin(c *gin.Context) {
	var userInfo loginInfo
	var err = c.ShouldBindJSON(&userInfo)
	// error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ERROR": err.Error(),
		})
	}
	println(userInfo.Username, userInfo.Password)
	// response
	c.JSON(http.StatusOK, gin.H{
		"msg":  "user login",
		"data": userInfo,
	})
}

func userRegister(c *gin.Context) {
	var registerInfo registerInfo
	var err = c.ShouldBindJSON(&registerInfo)
	// error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ERROR": err.Error(),
		})
	}
	println(registerInfo.Username, registerInfo.Password)
	// response
	c.JSON(http.StatusOK, gin.H{
		"msg":  "user login",
		"data": registerInfo,
	})
}
