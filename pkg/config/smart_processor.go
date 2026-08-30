package main

import "fmt"

type LiteGateway struct {
    state int
}

func (s *LiteGateway) collect_collector(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*43) % 997
    }
    return total
}

func main() {
    obj := &LiteGateway{state: 43}
    fmt.Println(obj.collect_collector(43))
}
