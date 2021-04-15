package serializer_test

import (
	"go_grpc/sample"
	"go_grpc/serializer"
	"testing"

	"github.com/stretchr/testify/require"
)

// 单元测试，必须以Test前缀开头，并且指向tests.T对象的指针作为输入
// 所有单元测试都要调用t.Parallel()函数，便于他们可以并行运行，而为任何高并发状态都可以检测到

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	// 存储二进制信息文件路径
	binaryFile := "../tmp/laptop.bin"
	// 创建一个笔记本电脑样本
	laptop1 := sample.NewLaptop()
	// 将笔记本电脑信息写入到二进制文件中
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	/**
	`require`软件包具有与`assert`软件包相同的全局功能，但是它们没有返回布尔结果，而是调用了t.FailNow（）。
	每个断言函数还将一个可选的字符串消息作为最后一个参数，从而允许将自定义错误消息附加到断言方法输出的消息中
	NoError断言一个函数没有返回错误（即`nil`）
	*/
	require.NoError(t, err)
}
