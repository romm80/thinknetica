package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	r := bufio.NewReader(conn)
	for scanner.Scan() {
		if scanner.Err() != nil {
			log.Fatal(err)
		}
		b := append(scanner.Bytes(), '\r', '\n')
		_, err = conn.Write(b)
		if err != nil {
			log.Fatal(err)
		}
		for {
			b, _, err = r.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			if len(b) == 0 {
				break
			}
			fmt.Println(string(b))
		}
	}

}
