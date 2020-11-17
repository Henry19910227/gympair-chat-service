package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ChatController ...
type ChatController struct {
}

// NewChatController ...
func NewChatController(router *gin.Engine) {
	chatController := &ChatController{}
	router.POST("/publish", chatController.Publish)
	router.POST("/connect", chatController.Connect)
}

// Publish ...
func (chatVC *ChatController) Publish(c *gin.Context) {
	var body map[string]interface{}
	c.ShouldBindJSON(&body)
	fmt.Println(body)
	c.JSON(http.StatusOK, gin.H{"result": gin.H{}})
}

// Connect ...
func (chatVC *ChatController) Connect(c *gin.Context) {
	fmt.Println("Connect!!!!!!")
	var body map[string]interface{}
	c.ShouldBindJSON(&body)
	fmt.Println(body)
	c.JSON(http.StatusOK, gin.H{"result": gin.H{}})
}
