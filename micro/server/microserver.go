package server

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func RunServer() {
	//注册中心
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})

	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(consulReg),
		micro.Address(":9877"),
	)
	service.Init()

	if err := service.Run(); err != nil {
		fmt.Println("服务启动失败")
		panic(err)
	}
}
