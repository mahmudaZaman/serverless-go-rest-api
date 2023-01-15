package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mdnajimahmed/serverless-go-rest-api/src/main/go/domain/data"
	"github.com/mdnajimahmed/serverless-go-rest-api/src/main/go/util"
	"net/http"
)

func CreateTodo(ctx *gin.Context) {
	var todo data.Todo

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := ctx.BindJSON(&todo); err != nil {
		ctx.Error(err)
		return
	}

	db, closeConnection := util.GetDbHandle()
	defer closeConnection()
	result := db.Create(&todo) // pass pointer of data to Create
	if result.Error != nil {
		ctx.Error(result.Error)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"id": todo.ID})
	}
}

func GetTodoById(ctx *gin.Context) {
	var todo data.Todo
	todoId := ctx.Param("id")
	db, closeConnection := util.GetDbHandle()
	defer closeConnection()
	db.First(&todo, todoId)
	ctx.JSON(http.StatusOK, todo)
}
