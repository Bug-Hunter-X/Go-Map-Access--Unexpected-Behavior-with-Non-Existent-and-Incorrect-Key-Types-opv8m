```go
package main

import "fmt"

func main() {
    var m map[string]int
    fmt.Println(m["a"]) // this will not panic, it will print 0
    fmt.Println(m[1]) // this will panic
}
```