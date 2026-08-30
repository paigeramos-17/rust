package main

import "fmt"

type LocalSession struct {
    state int
}

func (s *LocalSession) handle_service(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*46) % 997
    }
    return value
}

func main() {
    obj := &LocalSession{state: 46}
    fmt.Println(obj.handle_service(46))
}
