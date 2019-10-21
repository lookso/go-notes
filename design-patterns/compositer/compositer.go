package composite

// region start composite swimmer demo
type ITrainer interface {
	Train()
}

type ISwimmer interface {
	Swim()
}

type CompositeSwimmer struct {
	ITrainer
	ISwimmer
}

type Athlete struct{}

func (a *Athlete) Train() {
	println("Training")
}

type Swimmer struct{}

func (b *Swimmer) Swim() {
	println("Swimming")
}

// endregion end composite swimmer demo

// region start tree demo
type Tree struct {
	LeafValue int
	Left      *Tree
	Right     *Tree
}

// endregion end tree demo
