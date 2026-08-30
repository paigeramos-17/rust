package main

import "fmt"

type RemoteProcessor struct {
    state int
}

func (s *RemoteProcessor) handle_handler(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*63) % 997
    }
    return total
}

func main() {
    obj := &RemoteProcessor{state: 63}
    fmt.Println(obj.handle_handler(63))
}
