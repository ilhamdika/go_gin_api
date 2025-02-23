package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nayaka/models"
)

func HomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello Go developers")
}

func GetProgrammingLanguage(c *gin.Context) {
	lang := models.ProgrammingLanguage{
		Language:       "C",
		Appeared:       1972,
		Created:        []string{"Dennis Ritchie"},
		Functional:     true,
		ObjectOriented: false,
		Relation: models.Relation{
			InfluencedBy: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"},
			Influences:   []string{"C++", "Objective-C", "C#", "Java", "JavaScript", "PHP", "Go"},
		},
	}

	c.JSON(http.StatusOK, lang)
}
