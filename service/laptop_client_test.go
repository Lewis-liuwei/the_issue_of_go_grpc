package service_test

import (
	"context"
	"go_grpc/pb"
	"go_grpc/sample"
	"go_grpc/serializer"
	"go_grpc/service"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	// 启动服务
	laptopServer, serverAddress := startTestLaptopServer(t)
	// 创建一个函数返回笔记本电脑服务客户端进行测试
	laptopClient := newTestLaptopClient(t, serverAddress)

	// 新建一个新的笔记本电脑示例
	laptop := sample.NewLaptop()
	// 将id保存在此，便后期做对比
	expectedId := laptop.Id
	// 创建一个新的笔记本电脑服务请求对象
	req := &pb.CreatLaptopRequest{
		Laptop: laptop,
	}
	// 使用laptopClient对象调用CreateLaptop()函数
	res, err := laptopClient.CreatLaptop(context.Background(), req)
	// 要求没有错误信息
	require.NoError(t, err)
	// 返回数据不能为nil
	require.NotNil(t, res)
	// 返回的id与原始id相等
	require.Equal(t, expectedId, res.Id)

	// 需要判断当前id确实存在服务器中
	otherLaptop, err := laptopServer.Store.Find(res.Id)
	// 要求没有错误信息
	require.NoError(t, err)
	require.NotNil(t, otherLaptop)

	// 验证保存的笔记本电脑信息是否与发送的信息一致
	reqiureSameLaptop(t, otherLaptop, laptop)
}

// 启动一个处理笔记本电脑信息的服务（grpc服务）以及服务器网络地址字符串
func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())

	// 通过调用grpc.NewServer()创建grpc服务
	grpcServer := grpc.NewServer()
	// 在该grpc服务器上注册笔记本电脑服务的服务器
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	// 创建一个监听器，监听TCP连接，0表示分配任何随机可用的端口号
	listener, err := net.Listen("tcp", ":0")
	// 要求不能报错
	require.NoError(t, err)
	// 调用grpcServer.serve()开始监听请求
	go grpcServer.Serve(listener) // 开启一个协程监听

	// 返回笔记本电脑服务服务器以及监听器的地址
	return laptopServer, listener.Addr().String()
}

// 创建一个函数返回笔记本电脑服务客户端进行测试
func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	// 使用grpc.Dial()函数连接服务器地址，此处只是测试，所以使用不安全地址模式
	conne, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	// 返回带有服务器连接的笔记本电脑服务客户端
	return pb.NewLaptopServiceClient(conne)
}

func reqiureSameLaptop(t *testing.T, otherLaptop *pb.Laptop, laptop *pb.Laptop) {
	// 先将protobuf数据转换成json数据再做比较，因为在笔记本电脑结构体中有grpc在内部用于实例化对象用的参数，所以要比较就得忽略这些特殊参数
	otherLaptopJson, err := serializer.ProtobufToJson(otherLaptop)
	require.NoError(t, err)

	laptopJson, err := serializer.ProtobufToJson(laptop)
	require.NoError(t, err)

	require.Equal(t, otherLaptopJson, laptopJson)
}
