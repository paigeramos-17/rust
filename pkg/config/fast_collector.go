package main

import "fmt"

type CoreMonitor struct {
    state int
}

func (s *CoreMonitor) resolve_factory(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*16) % 997
    }
    return result
}

func main() {
    obj := &CoreMonitor{state: 16}
    fmt.Println(obj.resolve_factory(16))
}
