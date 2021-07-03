package strings

import "testing"

func TestStringBuilderConcat(t *testing.T) {
	impId := "110"
	requestId := "100"
	res := StringBuilderConcat(impId, requestId)
	t.Logf("res:%+v", res)
}
