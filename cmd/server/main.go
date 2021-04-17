package main

import (
	"flag"
	"fmt"
	"go_grpc/pb"
	"go_grpc/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 从命令行中获取参数
	port := flag.Int("port", 0, "the server port")
	// 解析参数
	flag.Parse()
	// 日志打印
	log.Printf("server start on port: %d", *port)

	// 创建服务器
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	// 注册服务器
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	// 服务器地址
	address := fmt.Sprintf("000:%d", *port)
	// 创建tcp连接
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	// 开启grpc监听服务
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
