package main

import (
	"fmt"
	"gincheese/apis/routes"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin == "null" {
			origin = "*"
		}
		fmt.Println("request origin 2 :", origin)
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Token, token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET,POST,PUT,DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.JSON(200, gin.H{"cross": "allow"})
			return
		}
		c.Next()
	})
	routes.RouteEngine(r).Run(":5100")

}
