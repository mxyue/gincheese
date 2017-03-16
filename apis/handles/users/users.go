package users

import (
	"fmt"
	"gincheese/app/config"
	"gincheese/app/db"
	"gincheese/app/util"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func Create(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	code := c.PostForm("code")
	cacheCode, found := util.CacheGet(config.RegistSPACE, email)
	if found && code == cacheCode {
		user := db.User{Email: email}
		util.CacheDelete(config.RegistSPACE, email)
		user_id, err := db.CreateUser(user, password)
		if err == nil {
			jwtToken := jwt.MapClaims{"user_id": user_id}
			c.JSON(http.StatusOK, gin.H{"success": "登陆成功", "token": util.Encrypt(jwtToken)})
			return
		} else {
			fmt.Println("error: ", err)
			c.JSON(601, gin.H{"error": "创建用户失败"})
		}
	} else {
		c.JSON(601, gin.H{"error": "验证码不正确"})
	}
}

func SendRegistCodeToEmail(c *gin.Context) {
	email, ok := c.GetQuery("email")
	if !ok {
		c.JSON(601, gin.H{"error": "请传入邮箱参数"})
	} else if !util.EmailRegex(email) {
		c.JSON(601, gin.H{"error": "不符合规则的邮箱"})
	} else if _, found := db.FindUser(bson.M{"email": email}); found {
		c.JSON(601, gin.H{"error": "该邮箱已经注册"})
	} else {
		err := util.SendCode(email, config.RegistSPACE)
		if err != nil {
			c.JSON(601, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"success": "发送成功"})
		}
	}

}
