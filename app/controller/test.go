package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestIndex(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}

func TestShow(c *gin.Context) {
	c.String(http.StatusOK, "show hello world")
}
