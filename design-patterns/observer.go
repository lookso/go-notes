package main

import "fmt"

type Subject struct {
	observers []Observer
	context   string
}

type Observer interface {
	Update(*Subject)
}

type Reader struct {
	name string
}
func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}
func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}


func main() {
	subject := NewSubject()
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

	subject.UpdateContext("observer mode")
	// Output:
	// reader1 receive observer mode
	// reader2 receive observer mode
	// reader3 receive observer mode
}
