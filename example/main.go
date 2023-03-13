package main

import (
	"fmt"

	"github.com/tsivinsky/ee"
)

func main() {
	e := ee.New()

	e.On("message", func(data ...any) {
		msg := data[0].(string)

		fmt.Println(msg)
	})

	e.Emit("message", "Hello, World!")
}
