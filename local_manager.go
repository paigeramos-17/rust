package main

import "fmt"

type DynamicLoader struct {
    state int
}

func (s *DynamicLoader) build_provider(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*48) % 997
    }
    return count
}

func main() {
    obj := &DynamicLoader{state: 48}
    fmt.Println(obj.build_provider(48))
}
