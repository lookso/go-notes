//+build wireinject

// https://gitee.com/llh-gitee/wire-demo/tree/master
package wire

import (
	"github.com/google/wire"
)

// InitGreeteEvent1 使用 wire 生成初始化代码
func InitGreeteEvent1() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{} // 这里的返回值不重要。
}

// InitGreeteEvent2 使用 wire 生成初始化代码，并返回产生的错误
func InitGreeteEvent2() (Event, error) {
	wire.Build(NewEvent2, NewGreeter2, NewMessage)
	return Event{}, nil
}

// InitGreeteEvent3 使用 wire 生成初始化代码，带参数
func InitGreeteEvent3(phrase string) (Event, error) {
	wire.Build(NewEvent2, NewGreeter2, NewMessage2)
	return Event{}, nil
}

// ------ 下面两方法展示了生成代码与输入参数的顺序无关 ------

func InitGreeteEvent4() (Event, error) {
	wire.Build(NewEvent2, NewMessage, NewGreeter2)
	return Event{}, nil
}

func InitGreeteEvent5() (Event, error) {
	wire.Build(NewMessage, NewEvent2, NewGreeter2)
	return Event{}, nil
}


// ---user--

// user InitializeUserService 初始化
func InitializeUserService() *UserService {
	wire.Build(NewUserService, MockUserRepoSet) // 使用MockUserRepoSet
	return nil
}
// --setting

func InitializeFooBar() Person {
	wire.Build(Set)

	return Person{}
}


// error
func InitService() (Service, error) {
	wire.Build(NewService)
	return Service{}, nil
}

func InitServiceClean() (Service, func(), error) {
	wire.Build(NewServiceClean)
	return Service{}, nil, nil
}

// player
func InitMission(name string) Mission {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}
}