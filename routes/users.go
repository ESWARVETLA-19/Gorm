package routes

import (
	"net/http"
	"myproject/models"
	"myproject/utils"
	"log"
	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		log.Println("Error parsing request data:", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	err=user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created!"})
}

func Login(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	err=user.ValidateCreds()
	if err!=nil{
		context.JSON(http.StatusUnauthorized,err) 
		return
	}
	token,err:=utils.GenerateToken(user.Email,user.ID)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"user was Not Authenticated"})
	}
	context.JSON(http.StatusOK,gin.H{"token":token})
}

