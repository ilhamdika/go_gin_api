package routes

import (
	"github.com/gin-gonic/gin"
	"nayaka/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/palindrome", controllers.CheckPalindrome)
}
