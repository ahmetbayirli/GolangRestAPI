package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"scroll2top.com/golang-rest-api/db"
	"scroll2top.com/golang-rest-api/routes"
)

func main() {
	fmt.Println("Hello world")
	db.InitDB()
	ginServer := gin.Default()

    routes.RegisterRoutes(ginServer)
	ginServer.Run(":8080")

}

