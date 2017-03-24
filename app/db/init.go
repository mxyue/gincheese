package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

var db_name = "gocheese"
var DbSession *mgo.Session
var url = "mongodb://localhost"
var trying = false

// var Database *mgo.Database

const (
	TODOS = "todos"
	USERS = "users"
)

func init() {
	setSession()
}

func setSession() {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	DbSession = session
}

func SetDBName(new_name string) {
	db_name = new_name
}

var Database = func() *mgo.Database {
	err := DbSession.Ping()
	if err != nil && !trying {
		Reconnect()
	}
	return DbSession.DB(db_name)
}

var TodoColl = func() *mgo.Collection {
	return Database().C(TODOS)
}

var UserColl = func() *mgo.Collection {
	return Database().C(USERS)
}

func Reconnect() {
	trying = true
	fmt.Println("refresh session")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			time.Sleep(5 * time.Second)
			Reconnect()
		}
	}()
	setSession()
	trying = false
}
