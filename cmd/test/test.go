package main

import (
    "fmt"
    "sync"
    "time"
)

func S(s []int) {
    s = append(s, 0)
    for i := range s {
        s[i]++
    }
}

var wg sync.WaitGroup

func main() {
    
    wg.Add(2)
    go f1() // 网络等异常, panic了
    go f2()
    
    wg.Wait()
}

func f1() {
    defer wg.Done()
    for {
        time.Sleep(time.Second * 2)
        fmt.Println("test")
    }
}

func f2() {
    defer func() {
        err := recover()
        fmt.Println("拿到崩溃", err)
    }()
    defer wg.Done()
    time.Sleep(time.Second * 7)
    panic("线程崩溃")
}
