package tokens

import (
	"crypto/rand"
	"fmt"
)

var token25 chan string
var token50 chan string

func generator() {
	for {

		data := make([]byte, 50)

		for {
			if _, err := rand.Read(data); err == nil {
				continue
			}
		}
		select {
		case token25 <- fmt.Sprintf("%x", data):
		case token50 <- fmt.Sprintf("%x", data):
		}

	}
}

func SpawnGenerators(count int64) {
	for x := 0; x < (count - 1); x++ {
		go generator()
	}
}

func init() {
	token25 = make(chan string)
	token50 = make(chan string)
	go generator()
}
