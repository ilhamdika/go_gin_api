package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

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
	text := c.Query("text")
	if text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Text tidak boleh kosong"})
		return
	}

	result := isPalindrome(text)
	if result {
		c.JSON(http.StatusOK, gin.H{"result": "Palindrome"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"result": "Bukan Palindrome"})
	}
}
