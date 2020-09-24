package api

import (
	"github.com/gin-gonic/gin"
)

//SetupAuthenAPI for qpritnstock
func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}

}

func login(c *gin.Context) {
	c.JSON(200, gin.H{"result": "Login !!"})
}

func register(c *gin.Context) {
	c.JSON(200, gin.H{"result": "Register !!"})
}
