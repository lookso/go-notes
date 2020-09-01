package main

import (
	"fmt"
	"github.com/google/uuid"
)

func getUuid() {
	uuidObj := uuid.New()
	fmt.Println(uuidObj.String())
}

func main() {
	getUuid()
}
