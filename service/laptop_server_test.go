package service_test

import (
	"context"
	"go_grpc/pb"
	"go_grpc/sample"
	"go_grpc/service"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel() // 保证多个测试可以并行处理

	// 声明一个笔记本电脑对象
	laptopNoId := sample.NewLaptop()
	laptopNoId.Id = ""
	laptopInvalidId := sample.NewLaptop()
	laptopInvalidId.Id = "Invalid-uuid"
	// 重复id
	laptopDuplicateId := sample.NewLaptop()
	// 声明一个内存存储对象
	storeDuplicateId := service.NewInMemoryLaptopStore()
	// 保存已存在的id，返回错误信息
	err := storeDuplicateId.Save(laptopDuplicateId)
	require.Nil(t, err)

	// 声明多个测试用例
	testCase := []struct {
		name   string              // 测试用例名称
		laptop *pb.Laptop          // 笔记本垫对象
		store  service.LaptopStore // 电脑存储对象
		code   codes.Code          // 状态码
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: laptopNoId, // 传入一个id为空的laptop
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_invalid_id",
			laptop: laptopInvalidId, // 传入一个id为无效的laptop
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument, // 无效参数
		},
		{
			name:   "failure_duplicate_id", // 因id重复而失败
			laptop: laptopDuplicateId,
			store:  storeDuplicateId,
			code:   codes.AlreadyExists, // 无效参数
		},
	}

	// 遍历测试用例进行测试
	for _, i := range testCase {
		t.Run(i.name, func(t *testing.T) {
			t.Parallel()
			// 创建请求实例
			req := &pb.CreatLaptopRequest{
				Laptop: i.laptop,
			}

			// 创建服务
			server := service.NewLaptopServer(i.store)
			// 调用服务中CreateLaptop方法
			res, err := server.CreatLaptop(context.Background(), req)
			// 判断测试用例状态是否ok
			if i.code == codes.OK {
				// ok
				// 要求没有错误信息
				require.NoError(t, err)
				// 要爱u返回值不能为nil
				require.NotNil(t, res)
				// 要求返回的uuid不能为空
				require.NotEmpty(t, res.Id)
				// 如果原始laptopId不为空
				if len(i.laptop.Id) > 0 {
					// 要求与返回的id相等
					require.Equal(t, i.laptop.Id, res.Id)
				}
			} else {
				// 不ok
				// 要求有错误信息
				require.Error(t, err)
				// 要求返回值为nil
				require.Nil(t, res)
				// 要求成功获取失败的状态码
				st, ok := status.FromError(err)
				require.True(t, ok)
				// 要求测试失败状态码和创建服务失败状态码保持一致
				require.Equal(t, i.code, st.Code())
			}
		})
	}
}
