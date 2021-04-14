package sample

import (
	"go_grpc/pb"
	"math/rand"
)

// 随机一个键盘样式
func randoKeyboardLayout() pb.Keyboard_Layout {
	r := rand.Intn(3)
	switch r {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	case 3:
		return pb.Keyboard_AZERTY
	default:
		return pb.Keyboard_UNKNOWN
	}
}

// 随机是否背光
func randomBool() bool {
	return rand.Intn(2) == 1
}
