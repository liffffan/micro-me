package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	proto "micro-me/protos"
)

func main() {

	// etcd
	//etcdRegistry := etcdv3.NewRegistry(
	//	func(options *registry.Options) {
	//		// 这里可以改成起 etcd 的ip地址
	//		options.Addrs = []string{"127.0.0.1:2379"}
	//		// 如果有账号和密码
	//		//etcdv3.Auth("root", "root")(options)
	//	})
	//
	//// Create a new service. Optionally include some options here.
	//service := micro.NewService(
	//	micro.Name("greeter.client"),
	//	micro.Registry(etcdRegistry),
	//	)
	//service.Init()
	//
	//// Create new greeter client
	//greeter := proto.NewGreeterService("greeter.service", service.Client())
	//
	//// Call the greeter
	//rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//// Print response
	//fmt.Println(rsp.Greeting)

	etcdRegistry := etcdv3.NewRegistry(
		func(options *registry.Options) {
			// 这里可以改成起 etcd 的ip地址
			options.Addrs = []string{"127.0.0.1:2379"}
			// 如果有账号和密码
			//etcdv3.Auth("root", "root")(options)
		})

	service := micro.NewService(
		micro.Name("greeter.client"),
		micro.Registry(etcdRegistry),
	)
	service.Init()

	// 创建客户端
	greeter := proto.NewGreeterService("greeter.service", service.Client())

	// 调用greeter服务
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Micro中国"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印响应结果
	fmt.Println(rsp.Greeting)
}
