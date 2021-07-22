package strings

import (
	"bytes"
	"fmt"
	"strings"
)

// https://blog.csdn.net/wangzhezhilu001/article/details/91421451

// buffer.WriteString():可以当成可变字符使用，对内存的增长也有优化
// 如果能预估字符串的长度，还可以用 buffer.Grow() 接口来设置 capacity
func builderConcat(impId, requestId string) string {
	var build strings.Builder
	build.WriteString("http://www.your_dsp.com/win?sid=")
	build.WriteString(requestId)
	build.WriteString("&impid=")
	build.WriteString(impId)
	build.WriteString("&price={")
	build.WriteString(fmt.Sprintf("%d", 10))
	build.WriteString("}")
	build.WriteString(strings.Join([]string{"lcc", "lyh", "ry"}, "-"))
	return build.String()
}
func bufferConcat(impId, requestId string) string {
	buf := new(bytes.Buffer)
	buf.WriteString("http://www.your_dsp.com/win?sid=")
	buf.WriteString(requestId)
	buf.WriteString("&impid=")
	buf.WriteString(impId)
	buf.WriteString("&price={")
	buf.WriteString(fmt.Sprintf("%d", 10))
	buf.WriteString("}")
	buf.WriteString(strings.Join([]string{"lcc", "lyh", "ry"}, "-"))
	return buf.String()
}
func byteConcat(n int, str string) string {
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}