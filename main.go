package main

import (
	"fmt"
	"go_grpc/sample"
)

/**
protoc --proto_path=proto proto/*.proto --gofast_out=plugins=grpc:pb
--proto_path 协议缓冲区代码文件路径
proto/*.proto 该文件夹内的文件
--gofast_out=plugins=grpc:路径 将协议缓冲区代码转译后存储的路径设置
*/
func main() {
	laptop := sample.NewLaptop()
	fmt.Printf("%v\n", laptop)
}
