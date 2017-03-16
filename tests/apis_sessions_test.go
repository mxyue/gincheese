package main

import (
	"encoding/json"
	"fmt"
	"gocheese/util"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestCreateSession(t *testing.T) {
	params := url.Values{"email": {firstUser.Email}, "password": {"123456"}}
	res, err := http.PostForm(fmt.Sprintf("%s/sessions", server.URL), params)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var data map[string]string
	err = json.Unmarshal(body, &data)
	fmt.Println("body:", data)
	if res.StatusCode == 200 && err == nil {
		tokenString := data["token"]
		claims, err := util.Decrypt(fmt.Sprintf("%s", tokenString))
		if err != nil {
			t.Error("decrypt失败:", err)
		} else if claims["user_id"] == firstUser.Id.Hex() {
			t.Log("通过")
		} else {
			t.Log("claims:", claims["user_id"])
			t.Log("first user id: ", firstUser.Id.Hex())
			t.Error("user id不正确")
		}
	} else {
		t.Log(res.StatusCode)
		t.Error(err)
	}
}
