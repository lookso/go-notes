package tree

import (
	"testing"
)

func TestTrie(t *testing.T) {
	obj := Constructor()
	obj.Insert("apple")
	obj.Insert("aps")
	p1 := obj.Search("apple")
	p2 := obj.StartsWith("ap")
	t.Logf("p1:%v,p2:%v", p1, p2)
}
