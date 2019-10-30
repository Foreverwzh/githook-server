package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

type repository struct {
	Name string `json:"name" binding:"required"`
	Default_branch string `json:"default_branch" binding:"required"`
}

type push struct {
	Repository repository `json:"repository" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/push", func(c *gin.Context) {
		var json push
		err := c.ShouldBindJSON(&json)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(json.Repository.Name)
		fmt.Println(json.Repository.Default_branch)
	})
	r.Run("localhost:8081")
}