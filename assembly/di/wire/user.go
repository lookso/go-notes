package wire

import (
	"errors"
	"github.com/google/wire"
)

// User 用户实体类
type User struct {
	Name string
	ID   int
}

// UserService User对象服务（实现类）
type UserService struct {
	userRepo IUserRepository
}

// IUserRepository 存放User对象的数据仓库接口
type IUserRepository interface {
	// GetUserByID 根据ID获取User, 如果找不到User返回对应错误信息
	GetUserByID(id int) (*User, error)
}

// NewUserService *UserService构造函数
func NewUserService(userRepo IUserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// UserRepoImpl 模拟一个 IUserRepository 实现
type UserRepoImpl struct {
	User
}

// GetUserByID UserRepository接口实现
// 意思意思
func (u *UserRepoImpl) GetUserByID(id int) (*User, error) {
	if id == 1 {
		return &u.User, nil
	}
	return nil, errors.New("未找到相关数据")
}

// NewMockUserRepo *mockUserRepo构造函数
func NewMockUserRepo() *UserRepoImpl {
	return &UserRepoImpl{
		User{Name: "Tom", ID: 1},
	}
}

// MockUserRepoSet 将 *mockUserRepo与UserRepository绑定
// Provider Set 是将一些 provider 组织起来使用，可以看下本包的 wire.go 文件
var MockUserRepoSet = wire.NewSet(
	NewMockUserRepo,
	wire.Bind(new(IUserRepository),
		new(*UserRepoImpl)))
