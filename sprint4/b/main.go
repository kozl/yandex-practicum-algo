package main

import (
	"fmt"
	"math/rand"
)

var (
	hashes       = map[int]string{}
	lowerLetters = []rune("abcdefghijklmnopqrstuvwxyz")
)

func polyHash(s string, a, mod int) int {
	hash := 0
	b := 1
	for i := 0; i < len(s); i++ {
		hash += int(s[len(s)-1-i]) * b
		b = (b * a) % mod
		hash = hash % mod
	}
	return hash
}

func getRandString(l int) string {
	b := make([]rune, l)
	for i := range b {
		b[i] = lowerLetters[rand.Intn(len(lowerLetters))]
	}
	return string(b)
}

func main() {
	ch := make(chan string)
	go func(ch chan string) {
		for {
			ch <- getRandString(6)
		}
	}(ch)

	for s := range ch {
		h := polyHash(s, 1000, 123_987_123)
		if _, ok := hashes[h]; !ok {
			hashes[h] = s
		} else {
			fmt.Println(s)
			fmt.Println(hashes[h])
			return
		}
	}
}
