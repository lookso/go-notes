package main

import (
	"fmt"
)

func main() {
	a := 30
	if a < 20 {
		fmt.Println(2222)
		goto end
	}
	//goto nb

	fmt.Println("a is less than 20")

end:
	fmt.Println("end")
	return
	// nb:
	//
	//	fmt.Println("nb")
	//	return
}

//func test(){
//	// 4. 从任务池获取任务
//	// 2. 权限校验、前置逻辑处理
//	task, err = t.beforeGrabTask(ctx, talId, producerId, pool, userPermission)
//	// 前置处理错误或者已经拿到任务，不再从任务池获取
//	if err != nil {
//		return
//	}
//	// 3. 从前置逻辑获取到任务(去重沙子、已领取任务等)
//	if task != nil {
//		goto grabed
//	}
//
//grabed:
//	// 7. 更新任务, 标识该任务的领取
//	fmt.Println("11111")
//	return
//}
