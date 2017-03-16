package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var db_name = "gocheese"
var DbSession *mgo.Session
var Database *mgo.Database

const (
	TODOS = "todos"
	USERS = "users"
)

func init() {
	url := "mongodb://localhost"
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Println(err)
	}
	DbSession = session
	Database = DbSession.DB(db_name)
}

func SetDBName(new_name string) {
	db_name = new_name
}

var TodoColl = func() *mgo.Collection {
	return Database.C(TODOS)
}

var UserColl = func() *mgo.Collection {
	return Database.C(USERS)
}
