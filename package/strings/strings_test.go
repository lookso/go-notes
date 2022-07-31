package strings

import "testing"

func TestStringBuilderConcat(t *testing.T) {
	impId := "110"
	requestId := "100"
	buildConcat := builderConcat(impId, requestId)
	t.Logf("buildConcat:%+v", buildConcat)

	bufferConcat := bufferConcat(impId, requestId)
	t.Logf("bufferConcat:%+v", bufferConcat)

	byteConcat := byteConcat(10, requestId)
	t.Logf("byteConcat:%+v", byteConcat)
}
