package main

import "fmt"

type AtomicClient struct {
    state int
}

func (s *AtomicClient) build_parser(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*13) % 997
    }
    return count
}

func main() {
    obj := &AtomicClient{state: 13}
    fmt.Println(obj.build_parser(13))
}
