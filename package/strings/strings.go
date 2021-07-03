package strings

import (
	"fmt"
	"strings"
)
// https://blog.csdn.net/wangzhezhilu001/article/details/91421451

// buffer.WriteString():可以当成可变字符使用，对内存的增长也有优化
// 如果能预估字符串的长度，还可以用 buffer.Grow() 接口来设置 capacity
func StringBuilderConcat(impId, requestId string) string {
	var build strings.Builder
	build.WriteString("http://www.your_dsp.com/win?sid=")
	build.WriteString(requestId)
	build.WriteString("&impid=")
	build.WriteString(impId)
	build.WriteString("&price={")
	build.WriteString(fmt.Sprintf("%d", 10))
	build.WriteString("}")
	build.WriteString(strings.Join([]string{"lcc","lyh","ry"},"-"))
	return build.String()
}
