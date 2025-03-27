package utils

import (
	"errors"
	"time"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

const secretKey="supersecreat"


func GenerateToken(email string,userId uint)(string,error){
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
		"userId":userId,
		"exp":time.Now().Add(time.Hour*2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

// func VarifyToken(token string)(uint,error){
// 	parsedToken,err:=jwt.Parse(token,func(token *jwt.Token)(interface{},error){
// 		_,status:=token.Method.(*jwt.SigningMethodHMAC)
// 		if !status{
// 			return nil,errors.New("unexpected signing method")
// 		}
// 		return []byte(secretKey),nil
// 	})
// 	if err!=nil{
// 		return 0,errors.New("could not parseToken")
// 	}
// 	tokenisValid:=parsedToken.Valid
// 	if !tokenisValid{
// 		return 0,errors.New("Invalid Token")
// 	}
// 	claims,ok:=parsedToken.Claims.(jwt.MapClaims)//type checking 
// 	if !ok{
// 		return 0,errors.New("invalid token claims")
// 	}
// 	// email:=claims["email"].(string)
// 	userId:=claims["userId"].(float64)
// 	return uint(userId),nil
// }

func VerifyToken(token string) (uint, error) {
    fmt.Println("Token to verify:", token)

    parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
        // Ensure the signing method is HMAC (HS256)
        _, status := token.Method.(*jwt.SigningMethodHMAC)
        if !status {
            fmt.Println("Unexpected signing method:", token.Method)
            return nil, errors.New("unexpected signing method")
        }
        return []byte(secretKey), nil
    })

    if err != nil {
        fmt.Println("Error parsing token:", err)
        return 0, errors.New("could not parse token")
    }

    if !parsedToken.Valid {
        fmt.Println("Invalid token")
        return 0, errors.New("invalid token")
    }

    claims, ok := parsedToken.Claims.(jwt.MapClaims)
    if !ok {
        fmt.Println("Invalid claims")
        return 0, errors.New("invalid token claims")
    }

    userId, ok := claims["userId"].(float64)
    if !ok {
        fmt.Println("Invalid userId")
        return 0, errors.New("invalid or missing userId claim")
    }

    fmt.Println("Parsed userId:", userId)
    return uint(userId), nil
}

