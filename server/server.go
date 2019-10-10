package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"log"
	hello "micro-me/protos"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.HelloRequest, rsp *hello.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// 启动的 grpc 服务的名字
	service := grpc.NewService(
		micro.Name("greeter"),
	)

	service.Init()

	err := hello.RegisterGreeterHandler(service.Server(), new(Say))
	if err != nil {
		log.Fatal("register server failed, err:%v \n", err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
