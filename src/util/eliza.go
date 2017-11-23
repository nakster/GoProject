package util

import (
	"math/rand"
)

//this func
func choice(ls []string) string {
	randnum := rand.Intn(len(ls))
	return ls[randnum]
}

//ReplyQuestion returns the answer to the question
func ReplyQuestion(input string) string {

	return choice(defaults)
}
