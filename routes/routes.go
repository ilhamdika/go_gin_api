package routes

import (
	"github.com/gin-gonic/gin"
	"nayaka/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/palindrome", controllers.CheckPalindrome)
	r.GET("/struct", controllers.GetStruct)

	r.GET("/", controllers.HomeHandler)
	r.GET("/language", controllers.GetProgrammingLanguage)
}
