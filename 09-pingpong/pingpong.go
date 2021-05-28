package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var response map[string]string

const MatchPoint = 15

func main() {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	response = make(map[string]string)
	response["begin"] = "ping"
	response["ping"] = "pong"
	response["PING"] = "stop"
	response["pong"] = "ping"
	response["PONG"] = "stop"
	response["stop"] = "begin"

	wg.Add(2)
	go player("Player 1", ch, &wg)
	go player("Player 2", ch, &wg)
	ch <- "begin"
	wg.Wait()
	fmt.Printf("%s  -  %s", <-ch, <-ch)
}

func player(name string, ch chan string, wg *sync.WaitGroup) {
	scr := 0
	for {
		msg := <-ch
		if msg == "I win" {
			wg.Done()
			ch <- fmt.Sprintf("%s (%v)", name, scr)
			break
		}
		if msg == "stop" {
			scr++
			if scr == MatchPoint {
				ch <- "I win"
				wg.Done()
				ch <- fmt.Sprintf("%s (%v)", name, scr)
				break
			}
		}
		msg = randomize(response[msg])
		fmt.Printf("%s: %s\n", name, msg)
		ch <- msg
	}
}

func randomize(msg string) string {
	if msg == "ping" || msg == "pong" {
		rand.Seed(time.Now().UnixNano())
		if (rand.Intn(100)) < 20 {
			msg = strings.ToUpper(msg)
		}
	}
	return msg
}
