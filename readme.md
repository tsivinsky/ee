# ee

Event Emitter in Go

## Install

```bash
go get -u github.com/tsivinsky/ee
```

## Example

```go
package main

import (
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
```
