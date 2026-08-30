package main

import "fmt"

type RemoteProcessor struct {
    state int
}

func (s *RemoteProcessor) collect_processor(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*23) % 997
    }
    return total
}

func main() {
    obj := &RemoteProcessor{state: 23}
    fmt.Println(obj.collect_processor(23))
}
