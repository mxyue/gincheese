package db

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	Email     string        `bson:"email"`
	Mobile    string        `bson:"mobile"`
	Password  []byte        `bson:"password"`
	CreatedAt time.Time     `bson:"created_at"`
}

func CreateUser(user User, password string) (string, error) {
	if _, found := FindUser(bson.M{"email": user.Email}); found {
		return "", errors.New("该邮箱已经注册")
	}
	// if _, found := FindUser(bson.M{"mobile": user.Mobile}); found {
	// 	return "", errors.New("该手机已经注册")
	// }
	id := bson.NewObjectId()
	user.Id = id
	user.CreatedAt = time.Now()
	bt_password := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bt_password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword
	err = UserColl().Insert(user)
	return id.Hex(), err
}

func GetAllUsers() []User {
	var users []User
	UserColl().Find(nil).All(&users)
	return users
}

func FindUser(query bson.M) (user User, found bool) {
	UserColl().Find(query).One(&user)
	if user.Email == "" {
		found = false
	} else {
		found = true
	}
	return user, found
}

func FindUserById(id string) (user User, err error) {
	if bson.IsObjectIdHex(id) {
		bsonObjectID := bson.ObjectIdHex(id)
		UserColl().FindId(bsonObjectID).One(&user)
		if user.Email == "" {
			err = errors.New("用户不存在")
		}
		return user, err
	} else {
		return user, errors.New("不正确的id")
	}

}

func (user *User) ValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
