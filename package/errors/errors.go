package main

import (
	"errors"
	"fmt"
	"time"
)

// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}
func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}
func main() {
	err := errors.New("emit macho dwarf: elf header corrupted\n")
	if err != nil {
		fmt.Print(err)
	}
	const name, id = "bimmler", 17
	fmtErr := fmt.Errorf("user %q (id %d) not found", name, id)
	if fmtErr!=nil{
		fmt.Println(fmtErr)
	}
	if err := oops(); err != nil {
		fmt.Println(err)
	}
	// Output: 1989-03-15 22:30:00 +0000 UTC: the file system has gone away
}
