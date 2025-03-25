package routes

import (
	"myproject/middleweares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){
	server.GET("/events", GetEvents) //route to get all events  // GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", GetEventByID)//route to get event by single id
	authentication:=server.Group("/")
	authentication.Use(middleweares.Authenticate)
	authentication.POST("/events", CreateEvent)//route to create event
	authentication.PUT("/events/:id",UpdateEvent)//route to update event
	authentication.DELETE("/events/:id",DeleteEvent)//route to Delete Event
	authentication.POST("/events/:id/register",RegisterForEvent)//route to register for an event
	authentication.DELETE("/events/:id/register",CancelRegistration)//route to unregister for an event
	server.POST("/signup",SignUp)//User creation Route
	server.POST("/login",Login)//User Login route
}