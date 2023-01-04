package tutor

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllTutorHandler(c *gin.Context) {
	var actions Actions = &Tutor{}
	test := actions.GetAll()
	c.JSONP(http.StatusOK, test)
}

func GetTutorById(c *gin.Context) {
	var tutorActions = &Tutor{}
	tutor := tutorActions.GetById(5)
	c.JSONP(http.StatusOK, tutor)
}
