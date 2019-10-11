package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	hello "micro-me/protos"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.HelloRequest, rsp *hello.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// etcd 注册
	//etcdRegistry := etcdv3.NewRegistry(
	//	func(options *registry.Options) {
	//		// 这里可以改成起 etcd 的ip地址
	//		options.Addrs = []string{"127.0.0.1:2379"}
	//		// 如果有账号和密码
	//		//etcdv3.Auth("root", "root")(options)
	//	})
	//
	//
	//// 启动的 grpc 服务的名字
	//// 创建服务，除了服务名，其它选项可加可不加，比如Version版本号、Metadata元数据等
	//service := grpc.NewService(
	//	micro.Name("greeter.service"),
	//	// 默认使用的是 127.0.0.1:2379
	//	micro.Registry(etcdRegistry),
	//)
	//
	//service.Init()
	//
	//err := hello.RegisterGreeterHandler(service.Server(), new(Say))
	//if err != nil {
	//	log.Fatal("register server failed, err:%v \n", err)
	//}
	//
	//if err := service.Run(); err != nil {
	//	log.Fatal(err)
	//}

	etcdRegistry := etcdv3.NewRegistry(
		func(options *registry.Options) {
			// 这里可以改成起 etcd 的ip地址
			options.Addrs = []string{"127.0.0.1:2379"}
			// 如果有账号和密码
			//etcdv3.Auth("root", "root")(options)
		})

	service := micro.NewService(
		micro.Name("greeter.service"),
		// 默认使用的是 127.0.0.1:2379
		micro.Registry(etcdRegistry),
	)
	service.Init()

	// 注册服务
	hello.RegisterGreeterHandler(service.Server(), new(Say))

	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
