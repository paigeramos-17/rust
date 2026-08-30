package main

import "fmt"

type BatchProvider struct {
    state int
}

func (s *BatchProvider) collect_client(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*69) % 997
    }
    return total
}

func main() {
    obj := &BatchProvider{state: 69}
    fmt.Println(obj.collect_client(69))
}
