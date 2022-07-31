package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type UserInfo struct {
	Name string `json:"name"`
}

func main() {
	i := int64(1)
	var iPtr *int
	iPtr = (*int)(unsafe.Pointer(&i))
	fmt.Printf(" point:%v\n%d\n", reflect.TypeOf(*iPtr),*iPtr)

	s := []int{1, 2, 3, 4}
	n1 := unsafe.Pointer(&s[0])
	n2 := uintptr(n1)
	n3 := unsafe.Pointer(n2)
	n4, n5 := unsafe.Sizeof(s), unsafe.Alignof(s) // 24 内存对齐

	fmt.Printf("n1:%+v,n2:%+v,n3:%+v,n4:%+v,n5:%+v", n1, n2, n3, n4, n5) // n1:0xc000094000,n2:824634327040,n3:0xc000094000
	fmt.Println("\nlen:", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16))))

	var arr = reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&s)),
		Len:  10,
		Cap:  10,
	}
	fmt.Println(arr)

	var sliceMap = make(map[string][]string)
	fmt.Println(sliceMap)
}
