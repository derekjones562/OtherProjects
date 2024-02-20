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
	numOfJumps = 50
	NintyFivePercentConfidentZScore = 1.96
	NintyFivePercent = 0.95
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
func betacf(x, a, b float64) (h float64){
	var fpmin = 1e-30
	var qab = a + b
	var qap = a + 1
	var qam = a - 1
	var c = 1.0
	var d = 1 - qab * x / qap
	var m2, aa, del float64

	// These q's will be used in factors that occur in the coefficients
	if math.Abs(d) < fpmin {
		d = fpmin
	}
	d = 1 / d
	h = d

	for m := 1.0; m <= 100; m++ {
		m2 = 2 * m
		aa = m * (b - m) * x / ((qam + m2) * (a + m2))
		// One step (the even one) of the recurrence
		d = 1 + aa * d
		if math.Abs(d) < fpmin {
			d = fpmin
		}
		c = 1.0 + aa / c
		if math.Abs(c) < fpmin {
			c = fpmin
		}
		d = 1 / d
		h *= d * c
		aa = -(a + m) * (qab + m) * x / ((a + m2) * (qap + m2))
		// Next step of the recurrence (the odd one)
		d = 1 + aa * d
		if math.Abs(d) < fpmin {
			d = fpmin
		}
		c = 1 + aa / c
		if math.Abs(c) < fpmin {
			c = fpmin
		}
		d = 1 / d
		del = d * c
		h *= del
		if math.Abs(del - 1.0) < 3e-7 {
			break
		}
	}
	return h
}


func gammaln(x float64) float64 {
	var cof = []float64{
		76.18009172947146, -86.50532032941677, 24.01409824083091,
		-1.231739572450155, 0.1208650973866179e-2, -0.5395239384953e-5,
	}
	var ser = 1.000000000190015
	xx := x
	y := x
	tmp := x + 5.5
	tmp -= (xx + 0.5) * math.Log(tmp)
	for j:=0; j < 6; j++ {
		y++
		ser += cof[j] / y
	}
	return math.Log(2.5066282746310005 * ser / xx) - tmp
}

func ibeta(x, a, b float64) float64 {
	// Factors in front of the continued fraction.
	var bt = 0.0
	if !(x == 0 || x == 1) {
		bt = math.Exp(gammaln(a + b) - gammaln(a) - gammaln(b) + a * math.Log(x) + b * math.Log(1 - x))
	}
	if x < 0 || x > 1 {
		return 0
	}
	if x < (a + 1) / (a + b + 2) {
		// Use continued fraction directly.
		return bt * betacf(x, a, b) / a
	}
	// else use continued fraction after making the symmetry transformation.
	return 1.0 - bt * betacf(1 - x, b, a) / b
}

func ibetainv(p, a, b float64) float64 {
	var EPS = 1e-8
	var a1 = a - 1
	var b1 = b - 1
	var lna, lnb, pp, t, u, err, x, al, h, w, afac float64
	if p <= 0 {
		return 0
	}
	if p >= 1{
		return 1
	}
    if a >= 1 && b >= 1 {
    	pp = 1 - p
    	if p < 0.5 {
    		pp = p
    	}
    	t = math.Sqrt(-2 * math.Log(pp))
    	x = (2.30753 + t * 0.27061) / (1 + t* (0.99229 + t * 0.04481)) - t
    	if p < 0.5 {
			x = -x
		}
		al = (x * x - 3) / 6
		h = 2 / (1 / (2 * a - 1)  + 1 / (2 * b - 1))
		w = (x * math.Sqrt(al + h) / h) - (1 / (2 * b - 1) - 1 / (2 * a - 1)) * (al + 5 / 6 - 2 / (3 * h))
		x = a / (a + b * math.Exp(2 * w))
    } else {
    	lna = math.Log(a / (a + b))
    	lnb = math.Log(b / (a + b))
    	t = math.Exp(a * lna) / a
    	u = math.Exp(b * lnb) / b
    	w = t + u
    	if p < t / w {
    		x = math.Pow(a * w * p, 1 / a)
		} else {
			x = 1 - math.Pow(b*w*(1-p), 1/b)
		}
    }
    afac = -gammaln(a) - gammaln(b) + gammaln(a + b)
    for j := 0; j < 10; j++ {
    	if x == 0 || x == 1 {
			return x
		}
		err = ibeta(x, a, b) - p
		t = math.Exp(a1 * math.Log(x) + b1 * math.Log(1 - x) + afac)
		u = err / t
		t = u / (1 - 0.5 * math.Min(1, u * (a1 / x - b1 / (1 - x))))
		x -= t
		if x <= 0 {
			x = 0.5 * (x + t)
		}
		if x >= 1 {
			x = 0.5 * (x + t + 1)
		}
		if math.Abs(t) < EPS * x && j > 0 {
			break
		}
    }
    return x
}

func tinv(p, dof float64) float64 {
	var x = ibetainv(2.0 * math.Min(p, 1 - p), 0.5 * dof, 0.5)
	x = math.Sqrt(dof * (1.0 - x) / x)
	if p > 0 {
		return x
	}
	return -x
}

func calculate(df, conf float64) float64 {
	if df>0  {
		conf = conf + (1-conf)/2
		//alert(tinv(p, 250))
		var cv = tinv(conf, df)
		cv = math.Round(cv*1000)
		cv=cv/1000
		return cv
	}
	return 0
}


func  calcTValue(degreesOfFreedom int) float64 {
	df := float64(degreesOfFreedom)
	var tTable95ConfidenceLevel []float64

	//tTable := func (x float64) float64 {
	//	return math.Exp(lgamma((df+1)/2)-lgamma(df/2)) /
	//	math.Sqrt(df*math.Pi) * math.Pow(1+(x*x)/df, -(df+1)/2)
	//}(2)

	for i := 0.0; i <= df; i++ {
		tTable95ConfidenceLevel = append(tTable95ConfidenceLevel, calculate(i, NintyFivePercent))
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
