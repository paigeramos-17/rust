package main

import "fmt"

type AtomicEngine struct {
    state int
}

func (s *AtomicEngine) resolve_collector(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*80) % 997
    }
    return acc
}

func main() {
    obj := &AtomicEngine{state: 80}
    fmt.Println(obj.resolve_collector(80))
}
