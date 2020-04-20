package main

import (
	"fmt"
	"sync"
)

func main() {
	sum, err := Sum()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("sum", sum)
}
func Sum() (sum int, err error) {
	var nums = []int{0,5, 1, 7, 8}
	var errCh = make(chan error, len(nums))
	var wg sync.WaitGroup
	for k, v := range nums {
		wg.Add(1)
		go func(k, v int) {
			num, err := dire(k, v)
			if err != nil {
				errCh <- err
				return
			}
			sum += num
			defer wg.Done()
		}(k, v)
	}
	wg.Wait()
	close(errCh)
	for ec := range errCh {
		if ec != nil {
			err = ec
			break
		}
	}
	return sum, err
}
func dire(k, v int) (num int, err error) {
	num = k / v
	return num, err
}
