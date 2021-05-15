package main

import (
	"context"
	"flag"
	"go_grpc/pb"
	"go_grpc/sample"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreatLaptopClient(laptopClient pb.LaptopServiceClient) {
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

func searchLaptop(laptopClient pb.LaptopServiceClient, filter *pb.Filter) {
	log.Printf("search filter: %v", filter)

	/**
	此函数类似于 context.WithDeadline。不同之处在于它将持续时间作为参数输入而不是时间对象。
	此函数返回派生 context，如果调用取消函数或超出超时持续时间，则会取消该派生 context
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 实例搜索请求
	req := &pb.SearchLaptopRequest{Filter: filter}
	// 通过laptopClient中的SearchLaptop方法获取数据流
	stream, err := laptopClient.SearchLaptop(ctx, req)
	if err != nil {
		log.Fatal("cannot search laptop: ", err)
	}

	// 循环处理数据流
	for {
		// 获取SearchLaptopResponse指针
		res, err := stream.Recv()
		if err != io.EOF {
			// 如果数据流读取未结束
			return
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}
		// 调用SearchLaptopResponse指针的GetLaptop方法获取笔记本电脑信息
		laptop := res.GetLaptop()
		// 打印出笔记本电脑参数
		log.Print("- found: ", laptop.GetId())
	}
}

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
	// 创建10个client
	for i := 0; i < 10; i++ {
		CreatLaptopClient(laptopClient)
	}

	// 实例一个过滤器
	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
	}

	// 执行搜索函数
	searchLaptop(laptopClient, filter)
}
