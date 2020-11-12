package main

import (
	"fmt"

	"github.com/Henry19910227/gympair-chat-service/internal/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello!!!")
	router := gin.New()
	controller.NewChatController(router)
	router.Run(":9091")
}
