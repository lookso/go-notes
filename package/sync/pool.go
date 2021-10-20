package main

import (
	"bytes"
	"sync"
)

func main() {
}

var buffers = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func GetBuffer() *bytes.Buffer {
	return buffers.Get().(*bytes.Buffer)
}
func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	buffers.Put(buf)
}
