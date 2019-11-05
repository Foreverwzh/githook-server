package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	_"os/exec"
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
		c.JSON(200, gin.H{"name":json.Repository.Name,"branch": json.Repository.Default_branch})
		// switch json.Repository.Default_branch {
		// case "master":
		// 	cmd := exec.Command("git", "pull", "origin", "master")
		// 	cmd.Dir = "/data/wzhlovelyw/"
		// 	err := cmd.Run()
		// 	if err!=nil {
		// 		fmt.Println(err)
		// 		c.JSON(200, gin.H{"error": err.Error()})
		// 	}
		// 	c.JSON(200, gin.H{"status": "Success"})
		// default:
		// 	fmt.Printf("unknow branch %v", json.Repository.Default_branch)
		// }
	})
	r.Run(":8080")
}