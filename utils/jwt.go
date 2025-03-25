package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secreatKey="supersecreat"


func GenerateToken(email string,userId uint)(string,error){
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
		"userId":userId,
		"exp":time.Now().Add(time.Hour*2).Unix(),
	})
	return token.SignedString([]byte(secreatKey))
}

func VarifyToken(token string)(uint,error){
	parsedToken,err:=jwt.Parse(token,func(token *jwt.Token)(interface{},error){
		_,status:=token.Method.(*jwt.SigningMethodHMAC)
		if !status{
			return nil,errors.New("unexpected signing method")
		}
		return []byte(secreatKey),nil
	})
	if err!=nil{
		return 0,errors.New("could not parseToken")
	}
	tokenisValid:=parsedToken.Valid
	if !tokenisValid{
		return 0,errors.New("Invalid Token")
	}
	claims,ok:=parsedToken.Claims.(jwt.MapClaims)//type checking 
	if !ok{
		return 0,errors.New("invalid token claims")
	}
	// email:=claims["email"].(string)
	userId:=uint(claims["userId"].(float64))
	return userId,nil
}