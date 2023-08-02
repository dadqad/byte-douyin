package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func PostHandler(c *gin.Context) {
	var user user
	// 解析POST请求的JSON数据到user结构体中
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 处理POST请求的逻辑
	// ... git push -u origin main git push -u origin master
	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"message": "Post request received",
		"name":    user.Name,
		"age":     user.Age,
	})
}
