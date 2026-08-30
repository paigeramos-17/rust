package main

import "fmt"

type HybridParser struct {
    state int
}

func (s *HybridParser) dispatch_client(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*31) % 997
    }
    return result
}

func main() {
    obj := &HybridParser{state: 31}
    fmt.Println(obj.dispatch_client(31))
}
