package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/montanaflynn/stats"
)

/*
A 30×30 grid of squares contains 900 fleas, initially one flea per square.
When a bell is rung, each flea jumps to an adjacent square at random
(usually 4 possibilities, except for fleas on the edge of the grid or at the corners).

What is the expected number of unoccupied squares after 50 rings of the bell?
Give your answer rounded to six decimal places.
 */
const (
	RowLength = 30
	NumOfColumns = 30
	GridSize = RowLength*NumOfColumns
	print = false
	numOfJumps = 61
	NintyFivePercentConfidentZScore = 1.96
)

type Square struct {
	newFleas int
	Fleas int
}


func FleasJump(grid []Square){
	for i, square := range grid {
		// jump all the fleas to new squares
		for j := 0; j < square.Fleas; j++ {
			grid[jump(i)].newFleas += 1
		}
	}
	for i, _ := range grid {
		grid[i].Fleas = grid[i].newFleas
		grid[i].newFleas = 0
	}
}

func GenerateGrid() []Square {
	newGrid := make([]Square, GridSize)
	for i, _ := range newGrid {
		newGrid[i].Fleas = 1
	}
	return newGrid
}


func jump(position int) int{
	//generate random number between 1-4, including 1 and 4
	var newPosition int
	rand.Seed(time.Now().UnixNano())
	jumpdirection := (rand.Int() % 4) + 1
	switch jumpdirection {
	case 1:
		//up
		newPosition = position - 30
	case 2:
		//right
		newPosition = position + 1
	case 3:
		//down
		newPosition = position + 30
	case 4:
		//left
		newPosition = position - 1
	}
	if isOffEdge(position, newPosition) {
		return jump(position)
	}
	return newPosition
}
func isOffEdge(position, newPosition int) bool {
	if newPosition < 0 {
		// off the top
		return true
	}
	if newPosition > GridSize - 1 {
		// off the bottom
		return true
	}
	if position % RowLength == 0 && position + 1 == newPosition {
		// right edge of board
		return true
	}
	if position % RowLength == 1 && position - 1 == newPosition {
		// left edge of board
		return true
	}
	return false
}

func PrintGrid(grid []Square) {
	if print {
		for i := 1; i <= GridSize; i++ {
			fmt.Printf("[%d]", grid[i-1].Fleas)
			if i%30 == 0 {
				fmt.Print("\n")
			}
		}
		fmt.Println("\n\n")
	}
}

func CountFleas(grid []Square) int {
	totalFleas := 0
	for i := 0; i < GridSize; i ++ {
		totalFleas += grid[i].Fleas
	}
	return totalFleas
}

func CountUnoccupiedSpaces(grid []Square) int {
	totalUnoccupied := 0
	for i := 0; i < GridSize; i ++ {
		if grid[i].Fleas == 0 {
			totalUnoccupied += 1
		}
	}
	return totalUnoccupied
}
func  calcTValue(degreesOfFreedom int) float64 {
	df := float64(degreesOfFreedom)
	var tTable95ConfidenceLevel []float64

	tTable := func (x float64) float64 {
		return math.Exp(lgamma((df+1)/2)-lgamma(df/2)) /
		math.Sqrt(df*math.Pi) * math.Pow(1+(x*x)/df, -(df+1)/2)
	}(2)
	fmt.Println(tTable)

	for i := 0.0; i < df; i++ {
		tTable95ConfidenceLevel = append(tTable95ConfidenceLevel, tTable)
	}
	return tTable95ConfidenceLevel[degreesOfFreedom]
}

func lgamma(x float64) float64 {
	y, _ := math.Lgamma(x)
	return y
}

func factorial(n float64) float64 {
	var factVal float64 = 1
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	}else{
		for i:=1.0; i<=n; i++ {
			factVal *= i  // mismatched types int64 and int
		}

	}
	return factVal  /* return from function*/
}

func calculateStats(Population []int) (float64, float64){
	d := stats.LoadRawData(Population)
	sd, _ := d.StandardDeviation()
	m, _ := d.Mean()
	// for population
	nintyfivePercentConfidenceLevel := NintyFivePercentConfidentZScore * (sd/math.Sqrt(float64(len(Population))))
	// for sample
	df := len(Population)-1
	tvalue := calcTValue(df)
	nintyfivePercentConfidenceLevel = tvalue * (sd/math.Sqrt(float64(len(Population))))
	return m-nintyfivePercentConfidenceLevel, m+nintyfivePercentConfidenceLevel
}

func main() {
	grid := GenerateGrid()
	var unoccupiedPopulation []int

	for i := 0; i < numOfJumps; i++ {
		PrintGrid(grid)

		FleasJump(grid)
		if CountFleas(grid) < GridSize{
			fmt.Print("WrongNumberOfFleas")
		}
		unoccupiedPopulation = append(unoccupiedPopulation, CountUnoccupiedSpaces(grid))
	}
	PrintGrid(grid)

	lowerCL, upperCL := calculateStats(unoccupiedPopulation)
	fmt.Printf("We can Be 95 Percent sure that the expected number of unoccupied squares after %d rings of the bell is between %f, and %f", numOfJumps, lowerCL, upperCL)

}
