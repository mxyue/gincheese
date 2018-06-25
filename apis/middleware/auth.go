package middleware

import (
	"errors"
	"fmt"
	"gincheese/app/db"
	"gincheese/app/util"
	"gopkg.in/gin-gonic/gin.v1"
)

func Auth(c *gin.Context) {
	token, _ := c.Request.Header["Token"]
	fmt.Println("token:", token)
	if len(token) > 0 {
		claims, err := util.Decrypt(token[0])
		if err != nil || claims["user_id"] == nil {
			fmt.Println("userId nil")
			noAuth(c)
		} else {
			userId := fmt.Sprintf("%s", claims["user_id"])
			fmt.Println("user id", userId)
			_, err := db.FindUserById(userId)
			if err != nil {
				fmt.Println(err)
				noAuth(c)
			} else {
				fmt.Println("====验证通过= userid: ", userId)
				c.Set("userId", userId)
				c.Next()
			}
		}
	} else {
		fmt.Println("token lenth: ", len(token))
		noAuth(c)
	}
}

func noAuth(c *gin.Context) {
	auth_error := errors.New("验证失败")
	c.JSON(401, gin.H{"errors": "验证失败"})
	c.AbortWithError(401, auth_error)
}
