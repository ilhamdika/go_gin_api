package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type RequestBody struct {
	Text string `json:"text" binding:"required"`
}

func isPalindrome(text string) bool {
	cleanText := strings.ToLower(strings.ReplaceAll(text, " ", ""))
	length := len(cleanText)
	for i := 0; i < length/2; i++ {
		if cleanText[i] != cleanText[length-1-i] {
			return false
		}
	}
	return true
}

func CheckPalindrome(c *gin.Context) {
	var reqBody RequestBody

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Text tidak boleh kosong"})
		return
	}

	result := isPalindrome(reqBody.Text)
	c.JSON(http.StatusOK, gin.H{
		"text":      reqBody.Text,
		"palindrome": result,
	})
}
