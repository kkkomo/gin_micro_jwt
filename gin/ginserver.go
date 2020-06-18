package gin

import (
	"fmt"
	"gin_micro_jwt/gin/service"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	engine.GET("/api/:name/:action/*function", service.MyGetFunc)

	wxGroup := engine.Group("/me")
	wxGroup.POST("/info", service.MyPostFunc)
	wxGroup.PUT("/info", service.MyPutFunc)
	wxGroup.DELETE("/info", service.MyDeleteFunc)
	addr := ":7887"
	if err := engine.Run(addr); err != nil {
		fmt.Println("server run error ...")
		panic(err)
	}
	fmt.Println("Listen and serve " + addr)
}