package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTableHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "vsye ok")
}
