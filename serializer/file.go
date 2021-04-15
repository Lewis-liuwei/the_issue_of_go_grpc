package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

// 将protobuf信息转换成二进制写入到文件中
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	// 首先将protobuf信息转换成二进制
	data, err := proto.Marshal(message)

	if err != nil {
		// 返回错误信息
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}
	// 将二进制数据写入到文件中
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		// 如果写入失败，返回信息
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}
	return nil
}
