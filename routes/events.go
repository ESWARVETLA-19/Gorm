package routes

import (
	"net/http"
	

	"fmt"
	"myproject/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func GetEventByID(context *gin.Context){
	id,err:=strconv.ParseInt(context.Param("id"),10,64)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Invalid event id"})
		return
	}
	event,err:=models.GetEvent(uint(id))
	if err!=nil{
		context.JSON(http.StatusNotFound,gin.H{"message":"Event not found"})
		return
	}
	context.JSON(http.StatusOK,event)

}

func CreateEvent(context *gin.Context) {
	
	var event models.Event
	err:= context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	userId:=context.GetUint("userId")
	event.UserID = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func UpdateEvent(context *gin.Context){

	id,err:=strconv.ParseInt(context.Param("id"),10,64)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Invalid event id"})
		return
	}
	userID:=context.GetUint("userId")
	event,err:=models.GetEvent(uint(id))
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch data"})
		return
	}

	if event.UserID!=userID{
		context.JSON(http.StatusUnauthorized,gin.H{"message":"Unauthorized"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	updatedEvent.ID=uint(id)
	err=updatedEvent.Update()
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not Update event"})
		return 
	}
	context.JSON(http.StatusOK,gin.H{"message":"Event Updated Successfully"})
}
	


func DeleteEvent(context *gin.Context){
	id,err:=strconv.ParseInt(context.Param("id"),10,64)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Invalid event id"})
		return
	}
	event,err:=models.GetEvent(uint(id))
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch data"})
		return
	}
	userID:=context.GetUint("userId")
	if event.UserID!=userID{
		context.JSON(http.StatusUnauthorized,gin.H{"message":"Unauthorized"})
		return
	}
	err=event.Delete()
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not delete data"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message": fmt.Sprintf("Event with id (%d) deleted", event.ID)})
}