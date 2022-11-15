package rand

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T)  {
	randNew()
	fmt.Println("---------")
	randSeed()
	t.Log("TestRand success")
}
