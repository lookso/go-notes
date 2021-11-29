package util

import (
	"fmt"
	"testing"
)

func TestBytes2Bits(t *testing.T)  {
	fmt.Println("Hello, ", Bytes2Bits([]byte{64, 0}))
}
