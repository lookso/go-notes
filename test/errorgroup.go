package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
)

func main() {
	sumNum, err := getSum()
	if err != nil {
		fmt.Println(err)
		os.Exit(123)
	}
	fmt.Println("sum", sumNum)
}

func getSum() (sum int, err error) {
	group, _ := errgroup.WithContext(context.Background())
	for i := 0; i < 5; i++ {
		index := i
		group.Go(func() error {

			if index == 2 {
				return errors.New("fetch err")
			}
			fmt.Println("index", index)
			sum = index + sum
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		return sum, err
	}
	return sum, nil
}
