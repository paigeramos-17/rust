package main

import "fmt"

type SmartContext struct {
    state int
}

func (s *SmartContext) flush_registry(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*42) % 997
    }
    return count
}

func main() {
    obj := &SmartContext{state: 42}
    fmt.Println(obj.flush_registry(42))
}
