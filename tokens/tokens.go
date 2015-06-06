package tokens

import (
	"crypto/rand"
	"fmt"
)

var Token4 chan string
var Token10 chan string
var Token24 chan string
var Token50 chan string


func generator(response chan string, length int64) {
	
	for {
		data := make([]byte, length / 2)
		if _, err := rand.Read(data); err != nil {
			fmt.Println(err)
			continue
		}
		response <- fmt.Sprintf("%x", data)
	}
}


func SpawnGenerators(count int64) {
	var x int64
	for x = 0; x < (count); x++ {
		go generator(Token4, 4)
		go generator(Token10, 10)
		go generator(Token24, 24)
		go generator(Token50, 50)
	}
}

func init() {
	Token4 = make(chan string)
	Token10 = make(chan string)
	Token24 = make(chan string)
	Token50 = make(chan string)
	SpawnGenerators(1)
}
