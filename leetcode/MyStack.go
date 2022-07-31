package main

func main() {
	obj := ConstructorQueue()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	//fmt.Println(obj.queue)
	//param1 := obj.Pop()
	//
	//param2 := obj.Top()
	//param3 := obj.Empty()
	//	fmt.Println(param1, param2, param3)
}
type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func ConstructorQueue() (s MyStack) {
	return
}

/** Push element x onto stack. */
func (s *MyStack) Push(x int) {
	n := len(s.queue)
	s.queue = append(s.queue, x)
	//fmt.Println(s.queue)
	for ; n > 0; n-- {
		s.queue = append(s.queue, s.queue[0])
		s.queue = s.queue[1:]
	}
}

/** Removes the element on top of the stack and returns that element. */
func (s *MyStack) Pop() int {
	v := s.queue[0]
	s.queue = s.queue[1:]
	return v
}

/** Get the top element. */
func (s *MyStack) Top() int {
	return s.queue[0]
}

/** Returns whether the stack is empty. */
func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}

//链接：https://leetcode-cn.com/problems/implement-stack-using-queues/solution/yong-dui-lie-shi-xian-zhan-by-leetcode-solution/
