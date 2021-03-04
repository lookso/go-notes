package listNode

import "testing"

func TestDoubleList(t *testing.T) {
	var dl DoubleList
	dl.Init()
	dl.Insert(1, &DoubleNode{
		Data: DoubleObject(5),
		Prev: &DoubleNode{
			Data: DoubleObject(6),
			Prev: nil,
			Next: &DoubleNode{
				Data: DoubleObject(8),
				Prev: nil,
			},
		},
		Next: &DoubleNode{
			Data: DoubleObject(7),
			Prev: &DoubleNode{
				Data: DoubleObject(10),
				Prev: &DoubleNode{
					Data: DoubleObject(6),
					Prev: nil,
				},
			},
			Next: nil,
		},
	})
	dl.Display()

}
