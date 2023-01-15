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
	if !db.Migrator().HasTable(&data.Todo{}) {
		fmt.Println("Creating todo table")
		// Create table for `User`
		db.Migrator().CreateTable(&data.Todo{})
		fmt.Println("CreateTodo table is ready")
	} else {
		fmt.Println("CreateTodo table already exists")
	}
}
