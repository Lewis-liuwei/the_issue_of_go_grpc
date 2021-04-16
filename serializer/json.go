package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// 将protobuf信息数据转换成json数据
func ProtobufToJson(message proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false, // 是否将枚举值用int/string类型数值代替
		EmitDefaults: true,  // 是否写入具有默认值的字段
		Indent:       "  ",  // 用两个空格进行缩进
		OrigName:     true,  // 是否要使用原始名称
	}

	// 将proto message 转换成json字符串返回
	return marshaler.MarshalToString(message)
}
