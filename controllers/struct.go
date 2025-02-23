package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Relation struct {
	InfluencedBy []string `json:"influenced-by"`
	Influences   []string `json:"influences"`
}

type ProgrammingLanguage struct {
	Language       string   `json:"language"`
	Appeared       int      `json:"appeared"`
	Created        []string `json:"created"`
	Functional     bool     `json:"functional"`
	ObjectOriented bool     `json:"object-oriented"`
	Relation       Relation `json:"relation"`
}

func GetProgrammingLanguage(c *gin.Context) {
	lang := ProgrammingLanguage{
		Language:       "C",
		Appeared:       1972,
		Created:        []string{"Dennis Ritchie"},
		Functional:     true,
		ObjectOriented: false,
		Relation: Relation{
			InfluencedBy: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"},
			Influences:   []string{"C++", "Objective-C", "C#", "Java", "JavaScript", "PHP", "Go"},
		},
	}

	c.JSON(http.StatusOK, lang)
}
