package DbConn

import (
	"Models"
	"fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"net/url"
	"sync"
)

var wg sync.WaitGroup


//-------------------------------------------------------------- Migrating tables ---------------------------------------------------------------------


var Db *gorm.DB

func InitDB()  {
	var err error

	Db, err = gorm.Open("mysql", "root:nikitesh@/url_health?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	//Migrate the schema

	fmt.Println("Creating URL table ")
	Db.AutoMigrate(&Models.Url{})

	// fmt.Println("Creating UrlHits table")
	Db.AutoMigrate(&Models.UrlHits{})

}










// --------------------------------------------- get hit log for given url id ------------------------------------------------------





















// 	//transforms the todos for building a good response
// 	for _, item := range todos {
// 		completed := false
// 		if item.Completed == 1 {
// 			completed = true
// 		} else {
// 			completed = false
// 		}
// 		_todos = append(_todos, transformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
// 	}
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
// }

// // fetchSingleTodo fetch a single todo
// func fetchSingleTodo(c *gin.Context) {
// 	var todo todoModel
// 	todoID := c.Param("id")

// 	db.First(&todo, todoID)

// 	if todo.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
// 		return
// 	}

// 	completed := false
// 	if todo.Completed == 1 {
// 		completed = true
// 	} else {
// 		completed = false
// 	}

// 	_todo := transformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
// }

// // updateTodo update a todo
// func updateTodo(c *gin.Context) {
// 	var todo todoModel
// 	todoID := c.Param("id")

// 	db.First(&todo, todoID)

// 	if todo.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
// 		return
// 	}

// 	db.Model(&todo).Update("title", c.PostForm("title"))
// 	completed, _ := strconv.Atoi(c.PostForm("completed"))
// 	db.Model(&todo).Update("completed", completed)
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
// }

// // deleteTodo remove a todo
// func deleteTodo(c *gin.Context) {
// 	var todo todoModel
// 	todoID := c.Param("id")

// 	db.First(&todo, todoID)

// 	if todo.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
// 		return
// 	}

// 	db.Delete(&todo)
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
// }
