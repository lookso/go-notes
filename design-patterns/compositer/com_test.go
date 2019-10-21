package composite

import (
	"testing"
)

func TestComposite_Swimmer(t *testing.T) {
	swimmer := CompositeSwimmer{
		ITrainer: &Athlete{},
		ISwimmer: &Swimmer{},
	}

	swimmer.Train()
	swimmer.Swim()
}

func TestTree(t *testing.T) {
	root := Tree{
		LeafValue: 0,
		Left: &Tree{
			LeafValue: 20,
			Left:      nil,
			Right:     nil,
		},
		Right: &Tree{
			LeafValue: 20,
			Left:      nil,
			Right: &Tree{
				LeafValue: 8,
				Left:      nil,
				Right:     nil,
			},
		},
	}

	println(root.Right.Right.LeafValue)
}
