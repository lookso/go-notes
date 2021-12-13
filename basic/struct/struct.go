package main

import (
	"fmt"
	"reflect"
)

type T1 struct {
	Name  string
	Age   int
	Arr   [2]bool
	ptr   *int
	slice []int
	map1  map[string]string
}

func main() {
	t1 := T1{
		Name:  "yxc",
		Age:   1,
		Arr:   [2]bool{true, false},
		ptr:   new(int),
		slice: []int{1, 2, 3},
		map1:  make(map[string]string, 0),
	}
	t2 := T1{
		Name:  "yxc",
		Age:   1,
		Arr:   [2]bool{true, false},
		ptr:   new(int),
		slice: []int{1, 2, 3},
		map1:  make(map[string]string, 0),
	}

	// 报错 实例不能比较 Invalid operation: t1 == t2 (operator == not defined on T1)
	// fmt.Println(t1 == t2)
	// 指针可以比较
	fmt.Println(&t1 == &t2) // false

	t3 := &T1{
		Name:  "yxc",
		Age:   1,
		Arr:   [2]bool{true, false},
		ptr:   new(int),
		slice: []int{1, 2, 3},
		map1:  make(map[string]string, 0),
	}

	t4 := &T1{
		Name:  "yxc",
		Age:   1,
		Arr:   [2]bool{true, false},
		ptr:   new(int),
		slice: []int{1, 2, 3},
		map1:  make(map[string]string, 0),
	}

	fmt.Println(t3 == t4)                               // false
	fmt.Println("DeepEqual", reflect.DeepEqual(t1, t2)) // true
	fmt.Printf("%p, %p \n", t3, t4)                     // 0xc000046050, 0xc0000460a0
	fmt.Printf("%p, %p \n", &t3, &t4)                   // 0xc000006030, 0xc000006038

	// 前面加*，表示指针指向的值，即结构体实例，不能用==
	// Invalid operation: *t3 == *t4 (operator == not defined on T1)
	// fmt.Println(*t3 == *t4)

	t5 := t3
	fmt.Println(t3 == t5)                  // true
	fmt.Println(reflect.DeepEqual(t3, t5)) // true
	fmt.Printf("%p, %p \n", t3, t5)        // 0xc000046050, 0xc000046050
	fmt.Printf("%p, %p \n", &t3, &t5)      // 0xc000006030, 0xc000006040
	MyUser()
}

type User struct {
	name string `json:"name"`
}

// 我们可以看到，它的两个成员变量都是非大写字母开头，只能在包内使用，
//现在我们为其中的 firstName 来定义 setter 与 getter
func (u *User) Getter() string {
	return u.name
}
func (u *User) Setter(newName string) {
	u.name = newName
}
func MyUser() {
	var user = User{name: "jack"}
	fmt.Println(user.name)
}
