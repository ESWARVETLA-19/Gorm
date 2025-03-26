package middleweares

import (
	"fmt"
	"strings"
	"myproject/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token:=context.Request.Header.Get("Authorization")
	if token==""{
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"unauthorized","error":"token not found"})
		return 
		}
		token = strings.TrimPrefix(token, "Bearer ")

// Debug: Print the token to check its value after removing "Bearer "
fmt.Println("Token to verify:", token)
	userId,err:=utils.VerifyToken(token)
	if err!=nil{
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"unauthorized","error":err.Error()})
		return
	}
	context.Set("userId",userId)
	context.Next()
}