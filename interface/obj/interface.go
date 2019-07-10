package obj

import "fmt"

type Football interface {
	buy() float32
}

type features struct {
	name    string
	country string
}

type toBuy struct {
	site  string
	price float32
	nums  int
}

func (b *toBuy) buy() float32 {
	return b.price * float32(b.nums)
}

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{}

func main() {
	var features = new(features)
	features.name = "football"
	features.country = "baxi"

	var ball Football = &toBuy{"taobao", 12.32, 10}
	fmt.Printf(features.name+"来自 %s,需要花费%f元\n", features.country, ball.buy())

	var val Any
	val = 5
	fmt.Printf("val has the value: %v\n", val)
	val = str
	fmt.Printf("val has the value: %v\n", val)
	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1
	fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}

}
