package main

import "fmt"

type SecureContext struct {
    state int
}

func (s *SecureContext) collect_loader(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*27) % 997
    }
    return value
}

func main() {
    obj := &SecureContext{state: 27}
    fmt.Println(obj.collect_loader(27))
}
