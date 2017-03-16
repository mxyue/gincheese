package main

import (
	"encoding/json"
	"fmt"
	"gincheese/app/db"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestGetTodos(t *testing.T) {
	var client = &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/todos", server.URL), nil)
	req.Header.Add("Token", firstUserToken)
	checkErr(err)
	res, err := client.Do(req)
	checkErr(err)
	fmt.Println("response code: ", res.StatusCode)
	defer res.Body.Close()
	type Data struct {
		Todos []db.Todo
	}
	var data Data
	err = json.NewDecoder(res.Body).Decode(&data)
	checkErr(err)
	fmt.Println("response body: ", data.Todos)
	if len(data.Todos) == 1 && res.StatusCode == 200 && err == nil {
		fmt.Printf("first Content: %v\n", data.Todos[0].Content)
		t.Log("通过")
	} else if err != nil {
		t.Error(err)
	} else {
		t.Error("帖子数量为", len(data.Todos))
		fmt.Println("code: ", res.StatusCode)
		fmt.Println("todos len: ", len(data.Todos))
	}
}

func TestCreateTodoDone(t *testing.T) {
	var todo db.Todo
	db.TodoColl().Find(bson.M{}).One(&todo)
	var client = &http.Client{}
	todo_id := todo.Id.Hex()
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/todos/%s/dones", server.URL, todo_id), nil)
	req.Header.Add("token", firstUserToken)
	res, err := client.Do(req)
	if res.StatusCode == 200 && err == nil {
		db.TodoColl().Find(bson.M{}).One(&todo)
		if len(todo.Dones) == 1 {
			t.Log("通过")
		} else {
			t.Error("失败")
		}
	} else {
		t.Log(res.StatusCode)
		t.Error(err)
	}
}

func TestDeleteTodo(t *testing.T) {
	var todo db.Todo
	db.TodoColl().Find(bson.M{}).One(&todo)
	var client = &http.Client{}
	id := todo.Id.Hex()
	fmt.Println("get id==>", id)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/todos/%s", server.URL, id), nil)
	req.Header.Add("token", firstUserToken)
	res, err := client.Do(req)
	if res.StatusCode == 200 && err == nil {
		t.Log("通过")
	} else {
		t.Log(res.StatusCode)
		t.Error(err)
	}
}

func TestCreateTodo(t *testing.T) {
	var client = &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/todos", server.URL), strings.NewReader("content=新任务"))
	req.Header.Add("token", firstUserToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var data map[string]string
	err = json.Unmarshal(body, &data)
	fmt.Println("body:", data)
	id := data["id"]
	todo, err := db.FindTodoById(fmt.Sprint(id))
	if res.StatusCode == 200 && err == nil && todo.Content == "新任务" {
		if todo.UserId.Hex() == firstUser.Id.Hex() {
			t.Log("通过")
		} else {
			t.Error("todo user id =", todo.UserId.Hex())
		}
	} else {
		t.Log(res.StatusCode)
		t.Error(err)
	}
}

