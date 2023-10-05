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

### Remove registered event

```go
package main

func main() {
    e := ee.New()

    e.On("message", func(data ..any) {})

    e.Remove("message")
}
```

### Remove one handler from event

```go
package main

func main() {
    e := ee.New()

    index := e.On("message", func(data ..any) {})

    e.Emit("message", "hi") // events["message"] length is 1

    e.Off("message", index) // events["message"] length is 0
}
```
