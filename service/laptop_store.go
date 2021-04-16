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
	// 声明一个笔记本电脑对象
	otherLaptop := &pb.Laptop{}
	// 调用copier中的copy方法，将laptop复制到otherLaptop中
	err := copier.Copy(otherLaptop, laptop)
	if err != nil {
		// 如果复制失败，则返回包装后的错误信息
		return fmt.Errorf("cannot copy laptop to orther_laptop: %w", err)
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
	otherLapto := &pb.Laptop{}
	err := copier.Copy(otherLapto, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	// 返回复制的数据
	return otherLapto, nil
}
