package sample

import (
	"go_grpc/pb"
	"math/rand"

	"github.com/google/uuid"
)

// 随机一个键盘样式
func randomKeyboardLayout() pb.Keyboard_Layout {
	// 随机0～3（不含）
	r := rand.Intn(4)
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

//  随机CPU品牌
func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

// 随机一个CPU名称
func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	}

	return randomStringFromSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
		"RYzen 3 PRO 3200GE",
	)
}

// 随机GPU品牌
func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

// 随机GPU名称
func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		)
	}

	return randomStringFromSet(
		"RX 590",
		"Rx 580",
		"RX 5700-XT",
		"RX Vega-56",
	)
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	return &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}
}

func randomPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomLaptopBrand() string {
	return randomStringFromSet("Dell", "Lenovel", "Apple")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Lenovel":
		return randomStringFromSet("Thinkpad X1", "Thinpad P1", "Thinkpad P53")
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	default:
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	}
}

// 从集合中随机一个元素
func randomStringFromSet(a ...string) string {
	lenNum := len(a)
	if lenNum == 0 {
		return ""
	}
	return a[rand.Intn(lenNum)]
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

//随机系统核数
func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// 随机是否背光
func randomBool() bool {
	// 随机0～2（不含2）
	return rand.Intn(2) == 1
}

func randomID() string {
	return uuid.New().String()
}
