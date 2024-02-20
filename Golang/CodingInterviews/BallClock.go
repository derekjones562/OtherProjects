package main

import (
	"fmt"
	"os"
	"os/exec"

	//"github.com/inancgumus/screen"
	"time"
)

var (
	ElapsedTwelveHourPeriods = 0
	Start = time.Now()
)

type BallClock struct {
	WaitingBalls []int
	OneMinSlot   []int
	FiveMinSlot  []int
	HourSlot     []int
	On bool
}

func (b *BallClock) Start() {
	b.On = true
	b.AddBallsToClock()

	for b.On == true {
		//time.Sleep(1*time.Minute)
		b.DropBall()
		//b.PrintCurrentTime()
	}
}

func (b *BallClock)Stop() {
	b.On = false
}

func (b *BallClock) AddBallsToClock() {
	b.WaitingBalls = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
}
func (b *BallClock) DropBall() {
	b.AddOneMinute()
}

func (b *BallClock)AddOneMinute() {
	if len(b.OneMinSlot) < 4 {
		b.OneMinSlot = append(b.OneMinSlot, b.WaitingBalls[0])
	} else {
		b.EmptySlot(b.OneMinSlot)
		b.OneMinSlot = []int{}
		b.AddFiveMinutes(b.WaitingBalls[0])
	}
	b.WaitingBalls = b.WaitingBalls[1:]
}

func (b *BallClock) AddFiveMinutes(Ball int) {
	if len(b.FiveMinSlot) < 11 {
		b.FiveMinSlot = append(b.FiveMinSlot, Ball)
	} else {
		b.EmptySlot(b.FiveMinSlot)
		b.FiveMinSlot = []int{}
		b.AddHour(Ball)
	}
}

func (b *BallClock) AddHour(Ball int) {
	if len(b.HourSlot) < 12 {
		b.HourSlot = append(b.HourSlot, Ball)
	} else {
		b.HourSlot = append(b.HourSlot, Ball)
		b.HourSlot = b.HourSlot[1:]
		b.EmptySlot(b.HourSlot)
		b.HourSlot = []int{-1}
		//b.HourSlot = append(b.HourSlot, Ball)
		ElapsedTwelveHourPeriods += 1
		if b.CheckForOriginalState() {
			fmt.Println(fmt.Sprintf("%d balls takes %d days. Computed in %d minutes", len(b.WaitingBalls), ElapsedTwelveHourPeriods/2, Start.Sub(time.Now()).Minutes()))
		} else {
			//fmt.Println(float64(ElapsedTwelveHourPeriods)/2)
		}
	}

}

func (b *BallClock) EmptySlot(Slot []int) {
	//ReverseTheSlot
	for i, j := 0, len(Slot)-1; i < j; i, j = i+1, j-1 {
		Slot[i], Slot[j] = Slot[j], Slot[i]
	}
	b.WaitingBalls = append(b.WaitingBalls, Slot...)
}

func (b *BallClock) PrintCurrentTime () {
	//screen.Clear()
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	_ = cmd.Run() //Only on Linux
	minutes := len(b.OneMinSlot) + (len(b.FiveMinSlot) * 5)
	fmt.Println(fmt.Sprintf("%d:%02d", len(b.HourSlot), minutes))
}

func (b *BallClock)CheckForOriginalState () bool {
	for i, v := range b.WaitingBalls {
		if v != i {
			return false
		}
	}
	return true
}

func main() {
	ballClock := BallClock{}
	go ballClock.Start()
	time.Sleep(90*time.Minute)
	ballClock.Stop()
}
