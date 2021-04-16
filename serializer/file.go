package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

// 将protobf信息数据转换成json数据写入到json文件中
func WriteProtobufToJsonFile(message proto.Message, filename string) error {
	data, err := ProtobufToJson(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to json: %w", err)
	}
	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write json data to file: %w", err)
	}
	return nil
}

// 将protobuf信息转换成二进制写入到文件中
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	// 首先将protobuf信息转换成二进制
	data, err := proto.Marshal(message) // 此处导致溢栈
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

// 从二进制文件中读取数据，传入一个二进制文件和接收protobuf message数据的类型
func ReadProtobufFromBinary(filename string, message proto.Message) error {
	// 从二进制文件中读取数据
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary from file: %w", err)
	}
	// 反序列化
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot Unmashal binary to proto message: %w", err)
	}
	return nil
}
