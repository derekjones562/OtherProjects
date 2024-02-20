package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Bracket Matcher
Have the function BracketMatcher(str) take the str parameter being passed and
return 1 if the brackets are correctly matched and each one is accounted for.
Otherwise return 0. For example: if str is "(hello (world))", then the output
should be 1, but if str is "((hello (world))" the the output should be 0
because the brackets do not correctly match up. Only "(" and ")" will be used
as brackets. If str contains no brackets return 1.

Examples

Input: "(coder)(byte))"
Output: 0

Input: "(c(oder)) b(yte)"
Output: 1
*/

var (
	Matches = "1"
	NotMatches = "0"
)

func BracketMatcher(str string) string {

	var BracketStack []rune
	var ok bool

	for _, v  := range str {
		switch v {
		case '(', '{', '[':
			BracketStack, ok = addToStack(BracketStack, v)
		case ')', '}', ']':
			BracketStack, ok = removeFromStack(BracketStack, v)
		default :
			continue
		}
		if !ok {
			return NotMatches
		}
		//fmt.Println(BracketStack)
	}
	if len(BracketStack) > 0 {
		return NotMatches
	}
	return Matches

}

func addToStack(stack []rune, r rune) ([]rune, bool) {
	stack = append(stack, r)
	return stack, true
}

func removeFromStack(stack []rune, r rune) ([]rune, bool) {
	if len(stack) == 0 {
		return nil, false
	}
	endRune := '.'
	switch r {
	case ')':
		endRune = '('
	case '}':
		endRune = '{'
	case ']':
		endRune = '['
	}
	n := len(stack) - 1
	if stack[n] != endRune {
		return nil, false
	}
	stack = stack[:n]
	return stack, true
}



func main() {
	t := &testing.T{}
	TestCorrectInputParens(t)
	TestIncorrectInputParens(t)
	TestCorrectInputBracketAndParens(t)
	TestIncorrectInputBracketAndParens(t)
	//input := "(coder)(byte))"
	//fmt.Println(BracketMatcher(input))

}

//Big O(n), where n is the length of the input string



func TestCorrectInputParens(t *testing.T) {
	assert.Equal(t, BracketMatcher("(coder)(byte))"), Matches)
}
func TestIncorrectInputParens(t *testing.T) {
	assert.Equal(t, BracketMatcher("(c(oder)) b(yte)"), NotMatches)
	assert.Equal(t, BracketMatcher(")(c(oder)) b(yte)"), NotMatches)
}
func TestCorrectInputBracketAndParens(t *testing.T) {
	assert.Equal(t, BracketMatcher("{(c(oder))} b({y}te)"), Matches)
}
func TestIncorrectInputBracketAndParens(t *testing.T) {
	assert.Equal(t, BracketMatcher("}(c(ode{r)) b}(yte)"), NotMatches)
	assert.Equal(t, BracketMatcher("(c(ode{r)) b}(yte)"), NotMatches)
}