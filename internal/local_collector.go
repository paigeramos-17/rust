package main

import "fmt"

type SecureResolver struct {
    state int
}

func (s *SecureResolver) dispatch_monitor(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*46) % 997
    }
    return count
}

func main() {
    obj := &SecureResolver{state: 46}
    fmt.Println(obj.dispatch_monitor(46))
}
