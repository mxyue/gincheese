package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gincheese/apis/routes"
	"gincheese/app/db"
	"gincheese/app/util"
	"gopkg.in/mgo.v2/bson"
	"net/http/httptest"
	"io"
	"io/ioutil"
	"encoding/json"
)

var server *httptest.Server
var firstUser db.User
var firstUserToken string

func init() {
	db.SetDBName("gocheese_test")

	db.UserColl().RemoveAll(nil)
	db.TodoColl().RemoveAll(nil)

	password := "123456"
	userData := db.User{Email: "basic@126.com", Mobile: "18280196887", Password: []byte(password)}
	_, err := db.CreateUser(userData, password)

	firstUser, _ = db.FindUser(bson.M{"email": userData.Email})
	fmt.Println("first user_id :", firstUser.Id.Hex())
	mapClaims := jwt.MapClaims{"user_id": firstUser.Id.Hex()}
	firstUserToken = util.Encrypt(mapClaims)

	dones := []db.Done{}
	todo := db.Todo{Id: bson.NewObjectId(),
		UserId:  firstUser.Id,
		Content: "第一个任务",
		Dones:   dones,
	}
	err = db.TodoColl().Insert(todo)
	if err != nil || todo.UserId.Hex() == "" {
		fmt.Println("数据存储不成功:", err)
	}else{
		fmt.Println("todo ===> ",todo.UserId.Hex())
	}
	server = httptest.NewServer(routes.RouteEngine())
}

func bodyMsg(resBody io.ReadCloser) {
	var data map[string]string
	body, err := ioutil.ReadAll(resBody)
	err = json.Unmarshal(body, &data)
	checkErr(err)
	fmt.Println("body: ", data)
}
func checkErr(err error) {
	if err != nil {
		fmt.Println("错误：", err)
	}
}
