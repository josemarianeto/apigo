package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	userContrutor()

	r.POST("/form_post", func(c *gin.Context) {
		UserCreating := User{
			Name:     c.PostForm("Name"),
			Password: c.PostForm("Password"),
		}

		userCreate(UserCreating.Name, UserCreating.Password)

		c.JSON(200, gin.H{
			"status": "Success",
			"Name":   UserCreating.Name,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
