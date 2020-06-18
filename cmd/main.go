package main

import "gin_micro_jwt/gin"

func main() {
	/*
	//jwt
	token, err := jwt.SignJwtToken("uuu")
	if err == nil {
		fmt.Println(token)
		jwtToken, b := jwt.ValidJwtToken(token)
		fmt.Printf("parse token result : %v, %s", b, jwtToken)
	} else {
		fmt.Println("sign token error")
	}*/

	//gin
	gin.InitRouter()
}
