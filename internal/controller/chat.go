package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ChatController ...
type ChatController struct {
}

// NewChatController ...
func NewChatController(router *gin.Engine) {
	chatController := &ChatController{}
	router.GET("/test", chatController.Test)
}

// Test ...
func (chatVC *ChatController) Test(c *gin.Context) {
	fmt.Println("Test!!!!!")
}
