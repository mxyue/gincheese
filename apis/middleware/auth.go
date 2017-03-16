package middleware

import (
	"errors"
	"gopkg.in/gin-gonic/gin.v1"
	"fmt"
	"gincheese/app/util"
	"gincheese/app/db"
)

func Auth(c *gin.Context){
	token, _ := c.Request.Header["Token"]
	fmt.Println("header:",c.Request.Header["Token"][0])
	if len(token) > 0 {
		claims, err := util.Decrypt(token[0])
		if err != nil || claims["user_id"] == nil {
			fmt.Println("userId nil")
			noAuth(c)
		} else {
			userId := fmt.Sprintf("%s", claims["user_id"])
			_, err := db.FindUserById(userId)
			if err != nil {
				noAuth(c)
			} else {
				fmt.Println("====验证通过= userid: ",userId)
				c.Set("userId", userId)
				c.Next()
			}
		}
	}else{
		fmt.Println("token lenth: ", len(token))
		noAuth(c)
	}
}

func noAuth(c *gin.Context){
	auth_error := errors.New("验证失败")
	c.JSON(401, gin.H{"errors": "验证失败"})
	c.AbortWithError(401, auth_error)
}
