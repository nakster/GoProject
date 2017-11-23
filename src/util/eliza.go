// adapted from https://github.com/kennysong/goeliza
package util

import (
	"math/rand"
	"strings"
	"regexp"
	"fmt"
)

//this func returns a random string from the array
func choice(ls []string) string {
	randnum := rand.Intn(len(ls))
	return ls[randnum]
}

//ReplyQuestion returns the answer to the question
func ReplyQuestion(input string) string {

	//trip and lowercase it 
	input = preprocess(input)


	if IsQuitStatement(input) {
		return choice(Goodbyes)
	}

	for pattern, res := range Responses {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(input)

		// If the statement matched any recognizable statements.
		if len(matches) > 0 {
			// If we matched a regex group in parentheses, get the first match.
			var fragment string
			if len(matches) > 1 {
				fragment = reflect(matches[1])
			}
			// Choose a random appropriate response, and format it with the
			// fragment, if needed.
			response := choice(res)
			if strings.Contains(response, "%s") {
				response = fmt.Sprintf(response, fragment)
			}
			return response
		}
	}

	return choice(defaults)
}

//this check if the user wants to end the conversation
func IsQuitStatement(input string) bool {
	input = preprocess(input)
	for _, quitStatement := range byeCheck {
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

// reflect flips words i to you
func reflect(fragment string) string {
	words := strings.Split(fragment, " ")
	for i, word := range words {
		if reflectedWord, ok := ReflectedWords[word]; ok {
			words[i] = reflectedWord
		}
	}
	return strings.Join(words, " ")
}
