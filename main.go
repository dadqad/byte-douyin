package main

import (
	"awesomeProject/api"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	router := gin.Default()
	// POST请求处理
	router.POST("/post", api.PostHandler)
	router.Run(":8080")
}
