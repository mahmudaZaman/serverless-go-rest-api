package main

import (
	"fmt"
	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"github.com/mdnajimahmed/serverless-go-rest-api/src/main/go/domain/data"
	"github.com/mdnajimahmed/serverless-go-rest-api/src/main/go/handler"
	"github.com/mdnajimahmed/serverless-go-rest-api/src/main/go/util"
	"log"
	"net/http"
)

var port = ":8287"

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", handler.Ping)
	r.POST("/todo", handler.CreateTodo)
	r.GET("/todo/:id", handler.GetTodoById)
	r.PATCH("/todo/:id", handler.UpdateTodo)
	r.PATCH("/todo/batch", handler.UpdateTodoBatch)
	r.DELETE("/todo/:id", handler.DeleteTodo)
	r.GET("/todos", handler.GetAllTodos)
	r.POST("/todo/batch", handler.CreateBatchTodo)
	r.POST("/mssp/batch", handler.CreateBatchMSSP)
	r.GET("/orMapping", handler.GetAllOrMapping)
	return r
}

func main() {
	prepareDataBase()
	if util.InLambda() {
		fmt.Println("running aws lambda in aws")
		log.Fatal(gateway.ListenAndServe(port, setupRouter()))
	} else {
		fmt.Println("running aws lambda in local")
		log.Fatal(http.ListenAndServe(port, setupRouter()))
	}
}

func prepareDataBase() {
	db, closeConnection := util.GetDbHandle()
	defer closeConnection()
	if !db.Migrator().HasTable(&data.OrMapping{}) {
		fmt.Println("Creating mssp table")
		// Create table for `User`
		//db.Migrator().CreateTable(&data.Todo{})
		db.Migrator().CreateTable(&data.OrMapping{})
		fmt.Println("CreateMSSP table is ready")
	} else {
		fmt.Println("CreateMSSP table already exists")
	}
}
