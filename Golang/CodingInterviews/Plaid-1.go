package main

import "fmt"

/*
ach --> process in large batches. 1ce a day

netting--> take long series of transaction broken down into smaller transactions.



A-Z --> banks
Central bank algorithm
   2 step process
		calc net balance for each bank --> send or recieve
		output list of transfers


			AB1
			BA2
			BC1
			=
			BA2
			AC1


			A=central

change algorithm so that Output is always shorter than input

 */

var banks = []string{"A", "B", "C", "D"}
//A=CentralBank
func CalcBalanceForBanks(Input [][]int) []int {
	// [0, 0, 0]
	netBalance := make([]int, len(banks))
	for _, transaction := range Input {
		if ok := validateTransaction(transaction); !ok {
			//logerror
			fmt.Println("error")
		}
		netBalance[transaction[0]] -= transaction[2]
		netBalance[transaction[1]] += transaction[2]
	}
	return netBalance



	// b2b --> store in memory [][]int bankSender->bandRecievers->netamount. keep track of net balance between any/all pairs of banks

}

func validateTransaction(transaction []int) bool {
	//lengthOfTransaction := len(transaction)

	if len(transaction) != 3 {
		return false
	}
	if transaction[0] > len(transaction) {
		return false
	}
	return true
}

func PrintOutput(output []int) {
	for i, amount := range output {
		if i == 0 {
			continue
		}
		if amount < 0 {
			fmt.Printf("%sA%d\n", banks[i], amount*-1)
		} else {
			fmt.Printf("A%s%d\n", banks[i], amount)
		}

	}

}
func main() {
	TestNettingAlgorithm()
	//TestNettingAlgorithm2Inputs()
	TestNettingAlgorithmNoCentral()

}

//
//func TestNettingAlgorithm() {
//	Input := [][]int{
//	{0, 1, 1},
//	{1, 0, 2},
//	{1, 2, 1},
//	}
//	output := CalcBalanceForBanks(Input)
//	PrintOutput(output)
//	//BA2
//	//AC1
//}

func TestNettingAlgorithm() {
	Input := [][]int{
		{1, 2, 1},
	}
	output := CalcBalanceForBanks(Input)
	PrintOutput(output)

}
//BC1


func TestNettingAlgorithm2Inputs() {
	Input := [][]int{
		{1, 2, 1},
		{2, 3, 1},
	}
	output := CalcBalanceForBanks(Input)
	PrintOutput(output)

}
func TestNettingAlgorithmNoCentral() {
	Input := [][]int{
		{1, 2, 1},
		{1, 3, 1},
		{2, 3, 1},
	}
	//BC1
	//BD1
	//CD1
	output := CalcBalanceForBanks(Input)
	PrintOutput(output)
	//BA2
	//AD2
}

//BC1
//BD1
//output
//BA1
//AD1


func TestInvalidInput() {

}

func TestOutput(){

}