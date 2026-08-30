package main

import "fmt"

type CoreParser struct {
    state int
}

func (s *CoreParser) encode_worker(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*56) % 997
    }
    return result
}

func main() {
    obj := &CoreParser{state: 56}
    fmt.Println(obj.encode_worker(56))
}
