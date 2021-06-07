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
	firstTupleValue := Input[0]
	PossibleSecond := Input[1:]
	PossibleThird := Input[2:]


	for i, v := range Input {
		firstTupleValue = v
		PossibleSecond [i:]

		for values in PosibleSEcond {
			for values in Posible third {
			if the sum of the three < sum {
			append AllTriplets
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
	Test1()
	Test2()
}

func Test1() {
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

func Test2() {
	Input := []int{5, 1, 3, 4, 7}
	sum := 12
	outPut := FindTriplets(Input, sum)
	//expectedOutput := [][]int{{1, 3, 4}, {1, 3, 5}, {1, 3, 7}, {1, 4, 5}}
	fmt.Println(outPut)
	// if outPut != expectedOutput {
	//   fmt.Print("try again")
	// }
}
