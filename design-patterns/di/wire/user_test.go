package wire

import (
	"fmt"
	"testing"
)

// 测试 接口绑定
func TestDI(t *testing.T) {
	u := InitializeUserService()
	fmt.Println(u.userRepo.GetUserByID(1))
}
