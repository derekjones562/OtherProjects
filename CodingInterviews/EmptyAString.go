/*
You are given a string S with only 0 and 1. You can delete the string 100 from any position of S an infinite number of times and obtain a new S after concatenation. Is it possible to make the string empty?

As for example, if S=101000 then 101000->100->empty

If S=1010001 then 1010001->1001->1->not empty

Input

Input starts with an integer T (≤ 100), denoting the number of test cases.

Each case contains a string S. The size of string is at most 120000.

Output

For each test case, print the case number and “yes” if it is possible to make the string S empty, print “no” otherwise.

Sample Input

Output for Sample Input

2
101000
1010001

Case 1: yes
Case 2: no
*/

package main
import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	Pattern = "100"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	inputData := ""
	for scanner.Scan() {
		inputData += scanner.Text() + "\n"
	}
	// Do not change: Log the output
	output := code_here(inputData)
	fmt.Println(output)
}

func code_here(inputData string) string {
	// Use this function to return the solution.
	var outputData string
	inputStrings := strings.Split(inputData, "\n")
	d, err := strconv.Atoi(inputStrings[0])
	if err != nil {
		//log Error
		return ""
	}
	if d > 100 {
		//log Error. only 100 or less input cases are allowed
		return ""
	}
	for i, binaryString := range inputStrings {
		if i > 0  && binaryString != ""{
			if IsEmptyString(binaryString) {
				outputData = outputData + "yes\n"
			} else {
				outputData = outputData + "no\n"
			}
		}
	}
	return outputData
}
func IsEmptyString( input string)bool{

	if  input == ""{
		return true
	}

	re := regexp.MustCompile(Pattern)
	remainder := re.Split(input, -1)
	if len(remainder) == len(input) || len(input) < len(Pattern){
		return false
	}
	return IsEmptyString(strings.Join(remainder, ""))
}