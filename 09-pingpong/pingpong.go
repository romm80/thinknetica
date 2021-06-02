package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	res := make(map[string]int)
	p1 := "Player 1"
	p2 := "Player 2"
	for res[p1] < 10 && res[p2] < 10 {
		wg.Add(2)
		go player(p1, ch, &wg)
		go player(p2, ch, &wg)
		ch <- "begin"
		wg.Wait()
		res[<-ch]++
		fmt.Printf("%s (%v) - %s (%v)\n", p1, res[p1], p2, res[p2])
	}
}

func player(name string, ch chan string, wg *sync.WaitGroup) {
	for {
		msg := <-ch
		if msg == "stop" {
			wg.Done()
			ch <- name
			return
		}
		rand.Seed(time.Now().UnixNano())
		if msg != "begin" && rand.Intn(100) < 20 {
			ch <- "stop"
			wg.Done()
			return
		}
		if msg == "ping" {
			msg = "pong"
		} else {
			msg = "ping"
		}
		fmt.Printf("%s: %v\n", name, msg)
		time.Sleep(time.Millisecond * 100)
		ch <- msg
	}
}
