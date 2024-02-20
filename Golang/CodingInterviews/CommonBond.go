package main
/*
Input 1:
[-2, 0, 1, 3]
sum = 2

Output 1:
(-2, 0, 1) and (-2, 0, 3)

Input 2:
[5, 1, 3, 4, 7]
sum = 12

Ouput 2:
 (1, 3, 4), (1, 3, 5), (1, 3, 7) and (1, 4, 5)
 */



import (
	"fmt"
	"sort"
)

func FindTriplets(Input []int, sum int) [][]int {
	if len(Input) < 3 {
		return [][]int{}
	}
	sort.Ints(Input)

	var AllTriplets [][]int

	for i, v := range Input[:len(Input)-2] {
		PossibleSecond := Input[i+1:]
		for j, w := range PossibleSecond {
			PossibleThird := PossibleSecond[j+1:]
			for _, x := range PossibleThird {
				if v+w+x < sum {
					AllTriplets = append(AllTriplets, []int{v,w,x})
				}
			}
		}
	}

	//var runningTotal int
	// for _, v := range Input {
	//   fmt.Println(runningTuple, v)
	//   if runningTotal + v > sum {continue}
	//   if len(runningTuple)+1==3 {
	//     //fmt.Println("appendAllTriplets")
	//     AllTriplets = append(AllTriplets, append(runningTuple, v))
	//   } else {
	//     //fmt.Println("appendrunningTuple")
	//     runningTuple = append(runningTuple, v)
	//     runningTotal += v
	//   }
	// }

	return AllTriplets
}

func main() {
	TestCommonbond1()
	TestCommonbond2()
}

func TestCommonbond1() {
	Input := []int{-2, 0, 1, 3}
	sum := 2
	//call my func
	outPut := FindTriplets(Input, sum)
	fmt.Println(outPut)
	//expectedOutput := [][]int{[]int{-2, 0, 1}, []int{-2, 0, 3}}

	// if outPut != expectedOutput {
	//   fmt.Print("try again")
	// }
}

//[1, 3, 4, 5, 7]

func TestCommonbond2() {
	Input := []int{5, 1, 3, 4, 7}
	sum := 12
	outPut := FindTriplets(Input, sum)
	//expectedOutput := [][]int{{1, 3, 4}, {1, 3, 5}, {1, 3, 7}, {1, 4, 5}}
	fmt.Println(outPut)
	// if outPut != expectedOutput {
	//   fmt.Print("try again")
	// }
}
