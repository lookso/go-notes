package rand

import (
	"fmt"
	"math/rand"
	"time"
)

func randNew() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= 20; i++ {
		fmt.Println(rand.Intn(i)) // 0到i,i不能为0
	}
}

func randSeed() {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 20; i++ {
		fmt.Println(rand.Intn(i)) // 0到i,i不能为0
	}
}
