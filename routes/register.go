package routes

import (
	"net/http"
	"myproject/models"
	"strconv"

	"github.com/gin-gonic/gin"
)
func RegisterForEvent(context *gin.Context){
	userID:=context.GetUint("userId")
	eventID,err:=strconv.ParseInt(context.Param("id"),10,64)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Invalid event id"})
		return
	}
	event,err:=models.GetEvent(uint(eventID))
	if err!=nil{
		context.JSON(http.StatusNotFound,gin.H{"message":"Event not found"})
		return
	}
	err=event.Register(userID)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not register for event"})
		return
	}
	context.JSON(http.StatusCreated,gin.H{"message":"Registered for event"})
}


func CancelRegistration(context *gin.Context){
	userID:=context.GetUint("userId")
	eventID,err:=strconv.ParseInt(context.Param("id"),10,64)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Invalid event id"})
		return
	}
	var event models.Event
	event.ID=uint(eventID)
	err=event.Unregister(userID)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not unregister for event"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message":"Unregistered for event"})

}