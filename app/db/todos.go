package db

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Todo struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UserId    bson.ObjectId `bson:"user_id,omitempty" json:"user_id"`
	Content   string        `bson:"content" json:"content"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	Dones     []Done        `json:"dones"`
}

type Done struct {
	Id    bson.ObjectId `bson:"_id" json:"id"`
	DidAt time.Time     `bson:"did_at" json:"did_at"`
}

func (t *Todo) AddDone(done Done) int {
	t.Dones = append(t.Dones, done)
	return len(t.Dones)
}

func GetAllTodos() []Todo {
	var todos []Todo
	TodoColl().Find(nil).Sort("-created_at").All(&todos)
	return todos
}

func GetUserTodos(user_id bson.ObjectId) []Todo {
	var todos []Todo
	TodoColl().Find(bson.M{"user_id": user_id}).Sort("-created_at").All(&todos)
	return todos
}

func FindTodo(query bson.M) Todo {
	var todo Todo
	TodoColl().Find(query).One(&todo)
	return todo
}

func FindTodoById(id string) (todo Todo, err error) {
	if bson.IsObjectIdHex(id) {
		bsonObjectID := bson.ObjectIdHex(id)
		TodoColl().FindId(bsonObjectID).One(&todo)
	} else {
		err = errors.New("非法id")
	}
	return todo, err
}

func CreateTodo(todo Todo) (bson.ObjectId, interface{}) {
	id := bson.NewObjectId()
	todo.Id = id
	todo.CreatedAt = time.Now()
	doc := TodoColl().Insert(todo)
	return id, doc
}

func DeleteTodoById(id string) (err error) {
	if bson.IsObjectIdHex(id) {
		fmt.Println("token count: ", len(GetAllTodos()))

		objId := bson.ObjectIdHex(id)
		err = TodoColl().RemoveId(objId)
	} else {
		err = errors.New("非法id")
	}
	if err != nil {
		fmt.Print(err)
	}
	return err
}


func DeleteUserTodoById(userId bson.ObjectId, id string) (err error) {
	if bson.IsObjectIdHex(id) {
		objId := bson.ObjectIdHex(id)
		todo := FindTodo(bson.M{"id": objId, "user_id": userId})
		fmt.Println(todo.Content)
		err = TodoColl().Remove(bson.M{"_id": objId, "user_id": userId})
	} else {
		err = errors.New("非法 todo id")
	}
	if err != nil {
		fmt.Print(err)
	}
	return err
}

func CreateDone(todo Todo, did_at time.Time) (bson.ObjectId, error) {
	id := bson.NewObjectId()
	done := Done{Id: id, DidAt: did_at}
	dones := append(todo.Dones, done)
	doc := bson.M{"$set": bson.M{
		"dones": dones,
	}}
	err := TodoColl().UpdateId(todo.Id, doc)
	return id, err
}
