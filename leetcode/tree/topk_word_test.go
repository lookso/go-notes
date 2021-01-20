package tree

import (
	"fmt"
	"testing"
)

func TestTopKWord(t *testing.T) {
	var s = []string{"i", "love", "leetcode", "i", "love", "coding"}
	fmt.Println(topKFrequent(s, 2))
}
