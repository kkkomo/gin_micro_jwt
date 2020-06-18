package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type user struct {
	Name string `json:"name"`
	Age	int `json:"age"`
}
func MyGetFunc(ctx *gin.Context) {
	name := ctx.Param("name")
	action := ctx.Param("action")
	function := ctx.Param("function")
	query := ctx.DefaultQuery("page", "1")

	fmt.Println("name :" + name + "; action :" + action + "; func:" + function + "; query :"+query)
	ctx.JSON(200, "get success")
}

func MyPostFunc(ctx *gin.Context) {
	//form-data
	name := ctx.PostForm("name")
	age := ctx.DefaultPostForm("age", "1")
	fmt.Println("post form res: name. "+ name + " & age." + age)
	user := user{}
	//application/json
	_ = ctx.ShouldBindJSON(&user)
	fmt.Printf("json binding res : %v", user)
}

func MyPutFunc(ctx *gin.Context) {

}

func MyDeleteFunc(ctx *gin.Context) {

}