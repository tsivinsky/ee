package main

import (
	"fmt"
	"log"

	"github.com/tsivinsky/ee"
)

func main() {
	e := ee.New()

	e.Once("message", func(data ...any) {
		msg := data[0].(string)
		fmt.Println(msg)
	})

	err := e.Emit("message", "Hello, World!")
	if err != nil {
		log.Fatal(err)
	}
	err = e.Emit("message", "Hello, World!")
	if err != nil {
		log.Fatal(err)
	}
}
