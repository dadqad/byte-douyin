package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("Hello, World!")
	router := gin.Default()
	// POST请求处理
	router.POST("/post", postHandler)
	router.Run(":8080")
}

func postHandler(c *gin.Context) {
	var user User
	// 解析POST请求的JSON数据到user结构体中
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 处理POST请求的逻辑
	// ... git push -u origin main
	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"message": "Post request received",
		"name":    user.Name,
		"age":     user.Age,
	})
}
