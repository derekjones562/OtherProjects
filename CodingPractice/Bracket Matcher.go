package main

import (
	"fmt"
	"source.vivint.com/propm/testly/asserts"
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
		fmt.Println(BracketStack)
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
	TestCorrectInputParens()
	TestIncorrectInputParens()
	TestCorrectInputBracketAndParens()
	TestIncorrectInputBracketAndParens()
	//input := "(coder)(byte))"
	//fmt.Println(BracketMatcher(input))

}

//Big O(n), where n is the length of the input string



func TestCorrectInputParens() {
	asserts.Assert.IsEqualTo(BracketMatcher("(coder)(byte))"), Matches)
}
func TestIncorrectInputParens() {
	asserts.Assert.IsEqualTo(BracketMatcher("(c(oder)) b(yte)"), NotMatches)
	asserts.Assert.IsEqualTo(BracketMatcher(")(c(oder)) b(yte)"), NotMatches)
}
func TestCorrectInputBracketAndParens() {
	asserts.Assert.IsEqualTo(BracketMatcher("{(c(oder))} b({y}te)"), Matches)
}
func TestIncorrectInputBracketAndParens() {
	asserts.Assert.IsEqualTo(BracketMatcher("}(c(ode{r)) b}(yte)"), NotMatches)
	asserts.Assert.IsEqualTo(BracketMatcher("(c(ode{r)) b}(yte)"), NotMatches)
}