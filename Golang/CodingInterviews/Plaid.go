package main
//aanderson@plaid.com
import (
	"fmt"
	"time"
)

/*
website. track visitors. implement simple api... contrived functions
log_hit() // everytime it gets hit
get_hits() // in the last five min
 */
// log -- timestamp
const (
	test = false
)

var numberOfHitsInLastMinute int
var last5MinutesOfHits []int // -->[300,555, 300,111, 222]

func log_hits(ts string) {
	fmt.Println(ts)
	numberOfHitsInLastMinute++
}

func get_hits() int {
	var TotalHits int
	for _, v := range last5MinutesOfHits{
		TotalHits += v
	}
	return TotalHits + numberOfHitsInLastMinute
}

func UpdateLogs() {
	for {
		if !test {
			for {
				if time.Now().Nanosecond() % 10000000 == 0  {
					break
				}
			}
		}
		last5MinutesOfHits = append(last5MinutesOfHits[1:], numberOfHitsInLastMinute)
		numberOfHitsInLastMinute = 0
		fmt.Printf("UpdatedLogs: %d", last5MinutesOfHits)
	}

}
func main () {
	last5MinutesOfHits = make([]int, 5*60)
	go UpdateLogs()
	Test1()
}

func Test1() {
	for i:=0; i< 1000; i++ {
		time.Sleep(100000)
		log_hits(time.Now().String())
	}
	fmt.Println(get_hits())
}