package util

import (
	"math/rand"
	"strings"
)

//this func
func choice(ls []string) string {
	randnum := rand.Intn(len(ls))
	return ls[randnum]
}

//ReplyQuestion returns the answer to the question
func ReplyQuestion(input string) string {

	if IsQuitStatement(input) {
		return choice(Goodbyes)
	}

	return choice(defaults)
}

func IsQuitStatement(input string) bool {
	input = preprocess(input)
	for _, quitStatement := range QuitStatements {
		if input == quitStatement {
			return true
		}
	}
	return false
}

//this makes the string better to match with the regex
func preprocess(input string) string {
	input = strings.TrimRight(input, "\n.!")
	input = strings.ToLower(input)
	return input
}
