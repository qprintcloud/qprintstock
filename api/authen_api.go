package api

import (
	"github.com/gin-gonic/gin"
)

//SetupAuthenAPI for qpritnstock
func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{
		authenAPI.POST("/login", Login)
		authenAPI.POST("/register", Register)
	}

}

//Login route for login
func Login(c *gin.Context) {
	c.JSON(200, gin.H{"result": "Login !!"})
}

//Register route for register
func Register(c *gin.Context) {
	c.JSON(200, gin.H{"result": "Register !!"})
}
