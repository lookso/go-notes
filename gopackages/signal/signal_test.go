package signal

import (
	"testing"
)

// 监听全部信号
func TestListenAll(t *testing.T) {
	ListenAll()
	t.Log("ListenAllTest success")
}

func TestListenSigUsr(t *testing.T) {

	ListenSigUsr()
	t.Log("TestListenSigUsr success")
}

func TestGoExit(t *testing.T) {
	GoExit()
	t.Log("TestGoExit success")
}
