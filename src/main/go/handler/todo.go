package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mdnajimahmed/serverless-go-rest-api/src/main/go/domain/data"
	"github.com/mdnajimahmed/serverless-go-rest-api/src/main/go/util"
	"net/http"
)

func CreateTodo(ctx *gin.Context) {
	var todo data.Todo
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

func CreateBatchTodo(ctx *gin.Context) {
	var todos []data.Todo
	if err := ctx.BindJSON(&todos); err != nil {
		ctx.Error(err)
		return
	}
	db, closeConnection := util.GetDbHandle()
	defer closeConnection()
	result := db.Create(&todos)
	if result.Error != nil {
		ctx.Error(result.Error)
	} else {
		ctx.JSON(http.StatusOK, nil)
	}
}

func CreateBatchMSSP(ctx *gin.Context) {
	var allOrMapping []data.OrMapping
	if err := ctx.BindJSON(&allOrMapping); err != nil {
		ctx.Error(err)
		return
	}
	db, closeConnection := util.GetDbHandle()
	defer closeConnection()
	result := db.Create(&allOrMapping)
	if result.Error != nil {
		ctx.Error(result.Error)
	} else {
		ctx.JSON(http.StatusOK, nil)
	}
}

func UpdateTodo(ctx *gin.Context) {
	var todo data.Todo
	if err := ctx.BindJSON(&todo); err != nil {
		ctx.Error(err)
		return
	}
	if ctx.Param("id") != fmt.Sprintf("%d", todo.ID) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id provided in the url does not match with id in the request body"})
	}

	// Validate input
	db, closeConnection := util.GetDbHandle()
	defer closeConnection()
	//if err := db.Where("id = ?", ctx.Param("id")).First(&todo).Error; err != nil {
	if err := db.First(&todo, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	result := db.Save(&todo)
	//result := db.Updates(&todo)
	if result.Error != nil {
		ctx.Error(result.Error)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"id": todo.ID})
	}
}

func UpdateTodoBatch(ctx *gin.Context) {
	var todos []data.Todo
	// Validate input
	if err := ctx.BindJSON(&todos); err != nil {
		ctx.Error(err)
		return
	}

	db, closeConnection := util.GetDbHandle()
	defer closeConnection()
	updateErrors := make([]error, 0)
	for _, todo := range todos {
		var todoTmp data.Todo
		if err := db.First(&todoTmp, todo.ID).Error; err != nil {
			updateErrors = append(updateErrors, err)
		} else {
			db.Save(&todo)
		}
	}
	ctx.JSON(http.StatusOK, updateErrors)
	//if err := db.Where("description = ?", ctx.Param("description")).First(&todo).Error; err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	//	return
	//}
	//
	//result := db.Where("description = ?", "same description updated by batch test").Updates(data.Todo{Description: "updated by batch test"})
	//if result.Error != nil {
	//	ctx.Error(result.Error)
	//} else {
	//	ctx.JSON(http.StatusOK, gin.H{"data": true})
	//}
}

func DeleteTodo(ctx *gin.Context) {
	var todo data.Todo
	db, closeConnection := util.GetDbHandle()
	defer closeConnection()
	if err := db.Where("id = ?", ctx.Param("id")).First(&todo).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	result := db.Delete(&todo)
	if result.Error != nil {
		ctx.Error(result.Error)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func GetAllTodos(ctx *gin.Context) {
	var todos []data.Todo
	db, closeConnection := util.GetDbHandle()
	defer closeConnection()
	result := db.Find(&todos)
	if result.Error != nil {
		ctx.Error(result.Error)
	} else {
		ctx.JSON(http.StatusOK, todos)
	}
}

func GetAllOrMapping(ctx *gin.Context) {
	var allOrMapping []data.OrMapping
	db, closeConnection := util.GetDbHandle()
	defer closeConnection()
	result := db.Find(&allOrMapping)
	if result.Error != nil {
		ctx.Error(result.Error)
	} else {
		ctx.JSON(http.StatusOK, allOrMapping)
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
