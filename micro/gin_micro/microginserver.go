package gin_micro

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func RunServer() {

	engine := gin.Default()
	service := web.NewService(
		web.Name("gin.server"),
		web.Address(":8999"),
	)
	_ = service.Init()
	service.Handle("/", engine)
	if err := service.Run(); err != nil {
		fmt.Println("gin 服务启动失败")
		panic(err)
	}
}
