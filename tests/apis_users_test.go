package main

import (
	"encoding/json"
	"fmt"
	"gincheese/app/config"
	"gincheese/app/db"
	"gincheese/app/util"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"gocheese/apis"
)

func TestCreateUser(t *testing.T) {
	var beforeUsers []db.User

	email := "626690641@qq.com"
	res1, err := http.Get(fmt.Sprintf("%s/valid_email?email=%s", server.URL, email))
	code, found := util.CacheGet(config.RegistSPACE, email)
	if res1.StatusCode == 200 && found {
		t.Log("通过")
	} else {
		t.Log(res1.StatusCode)
		bodyMsg(res1.Body)
		t.Error(err)
	}

	//创建用户
	err = db.UserColl().Find(nil).All(&beforeUsers)
	params := url.Values{"email": {email}, "password": {"123456"}, "code": {code.(string)}}
	res, err := http.PostForm(fmt.Sprintf("%s/users", server.URL), params)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var data apis.ResponseBody
	err = json.Unmarshal(body, &data)
	fmt.Println("body:", data)
	var users []db.User
	err = db.UserColl().Find(nil).All(&users)
	if res.StatusCode == 200 && err == nil && len(users) == (len(beforeUsers)+1) {
		t.Log("通过")
	} else {
		t.Log(res.StatusCode)
		t.Error(err)
	}
}
