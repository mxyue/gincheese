package main

import (
	"gincheese/apis/routes"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Token, token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET,POST,PUT,DELETE")
		if c.Request.Method == "OPTIONS" {
			c.JSON(200, gin.H{"cross": "allow"})
			return
		}
		c.Next()
	})
	routes.RouteEngine(r).Run(":5100")

}
