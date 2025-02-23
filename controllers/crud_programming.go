package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nayaka/models"
	"strconv"
)

var languages []models.ProgrammingLanguage

func GetProgrammingLanguageByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 || id >= len(languages) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	c.JSON(http.StatusOK, languages[id])
}

func GetAllProgrammingLanguages(c *gin.Context) {
	c.JSON(http.StatusOK, languages)
}

func CreateProgrammingLanguage(c *gin.Context) {
	var newLanguage models.ProgrammingLanguage
	if err := c.ShouldBindJSON(&newLanguage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	languages = append(languages, newLanguage)
	c.JSON(http.StatusCreated, newLanguage)
}

func UpdateProgrammingLanguage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 || id >= len(languages) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var updatedLanguage models.ProgrammingLanguage
	if err := c.ShouldBindJSON(&updatedLanguage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	languages[id] = updatedLanguage
	c.JSON(http.StatusOK, updatedLanguage)
}

func DeleteProgrammingLanguage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 || id >= len(languages) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	languages = append(languages[:id], languages[id+1:]...)
	c.JSON(http.StatusOK, gin.H{"message": "Language deleted"})
}