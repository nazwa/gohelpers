package tokens

import (
	"crypto/rand"
	"fmt"
)

var Token25 chan string
var Token50 chan string

func generator() {
	for {

		data := make([]byte, 50)

		if _, err := rand.Read(data); err != nil {
			fmt.Println(err)
			continue
		}
		select {
		case Token25 <- fmt.Sprintf("%x", data):
		case Token50 <- fmt.Sprintf("%x", data):
		}

	}
}

func SpawnGenerators(count int64) {
	var x int64
	for x = 0; x < (count); x++ {
		go generator()
	}
}

func init() {
	Token25 = make(chan string)
	Token50 = make(chan string)
	SpawnGenerators(1)
}
