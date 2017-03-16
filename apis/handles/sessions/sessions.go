package sessions

import (
	"fmt"
	"gincheese/app/db"
	 "gincheese/app/util"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

func Create(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	fmt.Println("email: ", email)
	fmt.Println("password: ", password)
	user, found := db.FindUser(bson.M{"email": email})
	if found {
		if user.ValidPassword(password) {
			 jwtToken := jwt.MapClaims{"user_id": user.Id.Hex()}
			c.JSON(http.StatusOK, gin.H{"token": util.Encrypt(jwtToken), "email": email})
		} else {
			fmt.Println("密码错误")
			c.JSON(601, gin.H{"error": "密码错误"})
		}
	} else {
		fmt.Println("账号不存在")
		c.JSON(601, gin.H{"error": "账号不存在"})
	}
}
