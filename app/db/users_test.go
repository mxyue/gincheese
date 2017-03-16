package db

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func init() {
	db_name = "gocheese_test"
	UserColl().RemoveAll(nil)
	log.SetLevel(log.DebugLevel)
}

func TestCreateUser(t *testing.T) {
	user := User{Email: "basicbox@126.com", Mobile: "18280196887"}
	doc, err := CreateUser(user, "123456")
	log.Debug("create user doc:", doc)
	users := GetAllUsers()
	if len(users) == 1 && err == nil {
		switch {
		case users[0].Email != "basicbox@126.com":
			t.Error("邮箱不对")
		case users[0].Mobile != "18280196887":
			t.Error("手机号不对")
		case !users[0].ValidPassword("123456"):
			t.Error("密码不对")
		}
	} else {
		t.Error("用户创建不成功")
	}
}

func TestFindUser(t *testing.T) {
	user := User{Email: "test@126.com", Mobile: "18280199999"}
	_, err := CreateUser(user, "123456")
	if err != nil {
		t.Error("创建用户数据出错")
	}
	findUser := FindUser(bson.M{"mobile": "18280199999"})
	if findUser.Mobile != "18280199999" {
		t.Error("查询用户失败")
	}
}

func TestReCreateUser(t *testing.T) {
	user := User{Email: "basicbox@126.com", Mobile: "18280196887"}
	doc, err := CreateUser(user, "123456")
	users := GetAllUsers()
	log.Debug("create user doc:", doc)
	if len(users) != 1 && err == nil {
		t.Error("重复创建用户, 用户数量:", len(users))
	}
}
