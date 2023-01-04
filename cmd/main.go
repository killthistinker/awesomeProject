package main

import (
	"awesomeProject/internal/tutor"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	r := router.Group("/tutors")
	{
		r.GET("/getAll", tutor.GetAllTutorHandler)
		r.GET(":id", tutor.GetTutorById)
	}

	router.Run(":8080")
}
