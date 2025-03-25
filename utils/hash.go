package utils

import(
	"golang.org/x/crypto/bcrypt"
)


func HashGenerator(password string)(string,error){
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password),14)

		return string(bytes),err

}

func CheckHashPassword(password,hashedpasswd string)bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hashedpasswd),[]byte(password))
	return err==nil
}