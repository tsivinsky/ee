package main

import (
	"fmt"

	"github.com/tsivinsky/ee"
)

func main() {
	e := ee.New()

	first := e.On("message", func(data ...any) {
		msg := data[0].(string)
		fmt.Printf("first: %s\n", msg)
	})

	second := e.On("message", func(data ...any) {
		msg := data[0].(string)
		fmt.Printf("second: %s\n", msg)
	})

	third := e.On("message", func(data ...any) {
		msg := data[0].(string)
		fmt.Printf("third: %s\n", msg)
	})

	_ = first
	_ = second
	_ = third

	e.Emit("message", "Hello, World!")

	e.Off("message", second)

	e.Emit("message", "Hello, World!")
}
