package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	userContrutor()

	r.POST("/user", func(c *gin.Context) {

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

	r.GET("/user", func(c *gin.Context) {
		var users []User = userAll()
		c.JSON(200, users)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
