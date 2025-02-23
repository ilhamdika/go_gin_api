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

	missingFields := make(map[string]string)
	if newLanguage.Language == "" {
		missingFields["language"] = "Language harus diinput"
	}
	if newLanguage.Appeared == 0 {
		missingFields["appeared"] = "Appeared harus diinput"
	}
	if len(newLanguage.Created) == 0 {
		missingFields["created"] = "Created harus diinput"
	}
	if len(newLanguage.Relation.InfluencedBy) == 0 {
		missingFields["influenced-by"] = "Influenced-by harus diinput"
	}
	if len(newLanguage.Relation.Influences) == 0 {
		missingFields["influences"] = "Influences harus diinput"
	}

	if len(missingFields) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": missingFields})
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

	missingFields := make(map[string]string)
	if updatedLanguage.Language == "" {
		missingFields["language"] = "Language harus diinput"
	}
	if updatedLanguage.Appeared == 0 {
		missingFields["appeared"] = "Appeared harus diinput"
	}
	if len(updatedLanguage.Created) == 0 {
		missingFields["created"] = "Created harus diinput"
	}
	if len(updatedLanguage.Relation.InfluencedBy) == 0 {
		missingFields["influenced-by"] = "Influenced-by harus diinput"
	}
	if len(updatedLanguage.Relation.Influences) == 0 {
		missingFields["influences"] = "Influences harus diinput"
	}

	if len(missingFields) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": missingFields})
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