package routes

import (
	"github.com/gin-gonic/gin"
	"scroll2top.com/golang-rest-api/middlewares"
)

func RegisterRoutes(ginServer *gin.Engine){
    
	ginServer.GET("/events/:id", getEvent)
	ginServer.GET("/events", getEvents)

    authenticated := ginServer.Group("/")

    authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	ginServer.POST("/signup", signUp)
	ginServer.POST("/login", login)
}