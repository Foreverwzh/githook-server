package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

type repository struct {
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
	Default_branch string `form:"default_branch" json:"default_branch" xml:"default_branch"  binding:"required"`
}

type push struct {
	Repository repository `form:"user" json:"user" xml:"user"  binding:"required"`
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