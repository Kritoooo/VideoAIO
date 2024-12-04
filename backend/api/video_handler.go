package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
}

func NewVideoHandler() *VideoHandler {
	return &VideoHandler{}
}

func (h *VideoHandler) Upload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "upload endpoint",
	})
}

func (h *VideoHandler) Process(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "process endpoint",
	})
}
