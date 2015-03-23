package gohelpers

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomAlphaNumeric(length int) string {

	var alphanum []byte = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	var dictLen = len(alphanum)
	var slice []byte = make([]byte, length)

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < length; i++ {
		slice[i] = alphanum[rand.Intn(dictLen)]
	}

	return fmt.Sprintf("%s", slice)
}
