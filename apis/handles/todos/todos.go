package todos

import (
	"fmt"
	"gincheese/app/db"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"
)

//var db = database.Con
//
type NewTodo struct {
	Content string `json:"content" form:"content" binding:"required"`
}

func Create(c *gin.Context) {
	var content NewTodo
	err := c.Bind(&content)
	checkErr(err)

	if err != nil {
		c.JSON(601, gin.H{"error": "参数错误"})
	} else {
		todo := db.Todo{Content: content.Content, UserId: userId(c)}
		id, err := db.CreateTodo(todo)
		if err != nil {
			c.JSON(601, gin.H{"error": "保存失败"})
		} else {
			c.JSON(http.StatusOK, gin.H{"id": id.Hex()})
		}
	}
}

func List(c *gin.Context) {
	todos := db.GetUserTodos(userId(c))
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func Delete(c *gin.Context) {
	todoId := c.Params.ByName("id")
	if bson.IsObjectIdHex(todoId) {
		value, _ := c.Get("userId")
		userId := bson.ObjectIdHex(fmt.Sprintf("%s", value))
		err := db.DeleteUserTodoById(userId, todoId)
		if err != nil {
			fmt.Println("删除失败：", err)
			c.JSON(601, gin.H{"error": "删除失败"})
		} else {
			c.JSON(http.StatusOK, gin.H{"success": "删除成功"})
		}
	} else {
		c.JSON(601, gin.H{"error": "非法id"})
	}
}

func Show(c *gin.Context) {
	todoId := c.Params.ByName("id")
	if bson.IsObjectIdHex(todoId) {
		todo, err := db.FindTodoById(todoId)
		if err != nil {
			fmt.Println("查询失败：", err)
			c.JSON(601, gin.H{"error": "查询失败"})
		} else {
			c.JSON(http.StatusOK, todo)
		}
	} else {
		c.JSON(601, gin.H{"error": "非法id"})
	}
}

func CreateDone(c *gin.Context) {
	did_at_string := c.PostForm("did_at")
	todoId := c.Params.ByName("id")
	var todo db.Todo
	if bson.IsObjectIdHex(todoId) {
		todo = db.FindTodo(bson.M{"_id": bson.ObjectIdHex(todoId), "user_id": userId(c)})
	} else {
		fmt.Println("err:", "非法id")
		c.JSON(601, gin.H{"error": "非法id"})
	}
	if len(todo.Id.Hex()) < 10 {
		fmt.Println("err todo 不存在:", todo.Id.Hex())
		c.JSON(601, gin.H{"error": "todo 不存在"})
	} else {
		var did_at time.Time
		if did_at_string == "" {
			did_at = time.Now()
		} else {
			i, err := strconv.ParseInt(did_at_string, 10, 64)
			if err != nil {
				fmt.Println("err:", err)
				c.JSON(601, gin.H{"error": "时间戳不标准"})
			}
			did_at = time.Unix(i, 0)
		}
		done_id, err := db.CreateDone(todo, did_at)
		if err != nil {
			fmt.Println("todo err:", err)
			c.JSON(601, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"id": done_id.Hex()})
		}
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("err:", err)
	}
}

func userId(c *gin.Context) bson.ObjectId {
	value, _ := c.Get("userId")
	fmt.Println("user id: ", fmt.Sprintf("%s", value))
	return bson.ObjectIdHex(fmt.Sprintf("%s", value))
}
