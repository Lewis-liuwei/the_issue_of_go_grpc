package service

import (
	"context"
	"errors"
	"go_grpc/pb"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 声明一个LaptopServer结构体, 提供笔记本电脑服务的服务器
type LaptopServer struct {
	Store LaptopStore // 继承笔记本电脑存储接口
}

// 使用NewLaptopServer函数返回新的实例
func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{Store: store}
}

/**
实现一个CreateLaptop函数，LaptopServiceServer接口需要它来实现
传入一个创建笔记本电脑请求结构体，返回一个创建笔记本电脑响应结构体
还可以传入一个自定义的函数
*/
func (server *LaptopServer) CreatLaptop(ctx context.Context, req *pb.CreatLaptopRequest) (*pb.CreateLaptopReaponse, error) {
	// 首先调用GetLaptop函数，从请求中获取笔记本电脑对象
	laptop := req.GetLaptop()
	// 打印日志说明已经收到带有此ID的create-laptop请求
	log.Printf("receive a create-laptop request with id: %s", laptop.Id)
	// 如果客户端已经生成了uuid,则验证该uuid是否有效
	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			// 表示uuid无效，给客户端返回nil以及错误状态码和错误信息（使用grpc状态码和代码子包，传入InvalidArgument代码）
			return nil, status.Errorf(codes.InvalidArgument, "laptop id is not a valid uuid: %v", err)
		}
	} else {
		// 如果客户端没有传入笔记本电脑uuid，则重新生成一个
		id, err := uuid.NewRandom()
		if err != nil {
			// 生成失败
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}
		// uuid生成成功，则将该id赋值给laptop中的id
		laptop.Id = id.String() // 将类型为uuid.UUID转换成字符串类型
	}

	// 将笔记本电脑信息存入到数据库中，此处为了演示将存入到内存中
	err := server.Store.Save(laptop)
	if err != nil {
		// 为了客户端更清晰的处理错误信息，此处判断错误信息是否是记录已存在导致的
		code := codes.Internal
		if errors.Is(err, ErrorAlreadyExists) {
			code = codes.AlreadyExists
		}
		// 如果保存失败，则用grpc中状态码返回错误信息
		return nil, status.Errorf(code, "cannot save laotop to the store: %v", err)
	}
	// 如果保存成功，打印成功日志
	log.Printf("saved laptop with id: %s", laptop.Id)
	// 使用该id创建一个新的笔记本电脑对象
	res := &pb.CreateLaptopReaponse{
		Id: laptop.Id,
	}
	// 返回新的对象
	return res, nil
}

func (server *LaptopServer) SearchLaptop(req *pb.SearchLaptopRequest, stream pb.LaptopService_SearchLaptopServer) error {
	// 获取请求的过滤器
	filter := req.Filter
	log.Printf("receive a search-laptop request with  filter: %v", filter)
	// 根据过滤器搜索不笔记本电脑数据
	err := server.Store.Search(filter, func(laptop *pb.Laptop) error {
		// 获取搜索响应数据
		res := &pb.SearchLaptopResponse{Laptop: laptop}

		// 发送流数据（rpc提供）
		err := stream.Send(res)
		if err != nil {
			return err
		}
		return nil
	})

	// 如果搜索失败
	if err != nil {
		return status.Errorf(codes.Internal, "unexpected erro: %v", err)
	}
	return nil
}
