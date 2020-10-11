package main

import "math/rand"

type Session struct {
}

func (s *Session) GenerateUuid() string {
	chars := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	rand.Shuffle(len(chars), func(i, j int) {
		chars[i], chars[j] = chars[j], chars[i]
	})

	return string(chars[3:9])
}
