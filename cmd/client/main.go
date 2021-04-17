package main

import (
	"context"
	"flag"
	"go_grpc/pb"
	"go_grpc/sample"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	// 从命令行中获取参数
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	// 打印正在与该服务器通信
	log.Printf("Dialing server %s", *serverAddress)

	// 连接服务器地址，使用不安全通信
	conne, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("connect server Failed ", err)
	}
	// 创建笔记本电脑服务客户端
	laptopClient := pb.NewLaptopServiceClient(conne)

	// 新建一个新的笔记本电脑示例
	laptop := sample.NewLaptop()
	// 创建一个新的笔记本电脑服务请求对象
	req := &pb.CreatLaptopRequest{
		Laptop: laptop,
	}
	// 使用laptopClient对象调用CreateLaptop()函数 // 此处导致溢栈
	res, err := laptopClient.CreatLaptop(context.Background(), req)
	if err != nil {
		// 将错误信息转换成状态对象，获取状态码便于校验
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			// 笔记本电脑id已存在，无须处理
			log.Printf("laptop already exsits")
		} else {
			//
			log.Fatal("cannot create laptop ", err)
		}
		return
	}
	// 打印创建的笔记本电脑id
	log.Printf("created laptop with id: %s", res.Id)

}
