package service

import (
	"errors"
	"fmt"
	"go_grpc/pb"
	"sync"

	"github.com/jinzhu/copier"
)

// 声明一个导出的错误变量
var ErrorAlreadyExists = errors.New("record already exists")

// 定义一个笔记本电脑信息存储接口，具有保存功能
type LaptopStore interface {
	Save(laptop *pb.Laptop) error       // 保存电脑信息函数
	Find(id string) (*pb.Laptop, error) // 根据id查询电脑信息
	// 使用过滤器搜索电脑信息，逐个返回
	Search(filter *pb.Filter, found func(laptop *pb.Laptop) error) error
}

// 实现存储到内存中的LaptopStore接口
type InMemoryLaptopStore struct {
	// 会出现多个并发请求来保存笔记本电脑信息，所以需要一个读写互斥锁来处理并发
	mutex sync.RWMutex
	// 使用map来存储数据，键是笔记本电脑的ID，值就是笔记本电脑对象
	data map[string]*pb.Laptop
}

// 实现存储到数据库中的LaptopStore接口
type InDBLaptopStore struct{}

// 声明一个函数，返回一个初始化的InMemoryLaptopStore对象
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

// 实现保存笔记本电脑信息的函数
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	// 首先在保存新对象数据之前先获得写锁
	store.mutex.Lock()
	// 推迟解锁命令
	defer store.mutex.Unlock()
	// 下一步，检查笔记本电脑id是否已经存在于store.data中
	if store.data[laptop.Id] != nil {
		// 如果存在，则将错误直接返回给调用方
		return ErrorAlreadyExists
	}

	// 如果不存在，则保存在map中
	// 为了安全起见，需要对笔记本电脑对象进行深度复制
	otherLaptop, err := deepCopy(laptop)
	if err != nil {
		// 如果复制失败，则返回包装后的错误信息
		return err
	}
	// 复制成功之后，将otherLaptop保存到store.data中
	store.data[otherLaptop.Id] = otherLaptop
	return nil
}

// 通过uuid查询笔记本电脑信息
func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	// 获取读锁
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	// 为了数据安全性，复制查询的数据
	return deepCopy(laptop)
}

//
func (store *InMemoryLaptopStore) Search(filter *pb.Filter, found func(laptop *pb.Laptop) error) error {
	// 获取读锁
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	// 遍历store中所有数据
	for _, laptop := range store.data {
		// 根据过滤器，判断是否是有效数据
		if isQualified(filter, laptop) {
			// 复制对象
			otherLaptop, err := deepCopy(laptop)
			if err != nil {
				return nil
			}
			// 传入的函数found
			err = found(otherLaptop)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func isQualified(filter *pb.Filter, laptop *pb.Laptop) bool {

	if laptop.PriceUsd > filter.MaxPriceUsd {
		// 笔记本电脑价格不能超过过滤器中的最高价格
		return false
	}

	if laptop.GetCpu().GetNumberCores() < filter.GetMinCpuCores() {
		// 笔记本电脑cup核数小于过滤器中最小核数
		return false
	}

	if laptop.GetCpu().GetMinGhz() < filter.GetMinCpuGhz() {
		// 笔记本电脑cpu频率小于过滤其中最小频率
		return false
	}

	if toBit(laptop.GetRam()) < toBit(filter.GetMinRam()) {
		return false
	}

	return true
}

func toBit(memory *pb.Memory) uint64 {
	value := memory.GetValue()

	// 根据单位换算成位
	switch memory.Unit {
	case pb.Memory_BIT: // 位
		return value
	case pb.Memory_BYTE: // 1字节=8位
		return value << 3 // 8 = 2^3
	case pb.Memory_KILOBYTE: // kb
		return value << 13 // 1024 * 8 = 2^10 * 2^3 = 2^13
	case pb.Memory_MEGABYTE: // mb
		return value << 23 // 10240 * 8 =
	case pb.Memory_GIGABYTE: // gb
		return value << 33
	case pb.Memory_TERABYTE: // tb
		return value << 43
	default:
		return 0
	}
}

func deepCopy(laptop *pb.Laptop) (*pb.Laptop, error) {
	// 为了数据安全性，复制查询的数据
	otherLapto := &pb.Laptop{}
	err := copier.Copy(otherLapto, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	// 返回复制的数据
	return otherLapto, nil
}
