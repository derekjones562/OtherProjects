package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// The cost of a stock on each day is given in an array, find the max profit that you can
// make by buying and selling in those days. For example, if the given array is:
// {100, 180, 260, 310, 40, 535, 695}, the maximum profit can earned by buying on day 0,
// selling on day 3. Again buy on day 4 and sell on day 6. If the given array of prices
// is sorted in decreasing order, then profit cannot be earned at all.

func maxProfit (prices []int, start, end int) int {
	if ok := validatePrices(prices); !ok {return 0}
	if start >= end {
		return 0
	}

	buyIndex := -1
	var profit int
	for start < end {
		if buyIndex == -1 && prices[start] < prices[start + 1] {
			buyIndex = start
		}
		if buyIndex != -1 && prices[start] > prices[start + 1]{
			profit += prices[start] - prices[buyIndex]
			buyIndex = -1
		}
		start ++
	}
	if buyIndex != -1 {
		profit += prices[end] - prices[buyIndex]
	}
	return profit
}

func validatePrices(prices []int) bool {
	lengthOfPrices := len(prices)
	if lengthOfPrices <= 1 {
		return false
	}
	for _, price := range prices {
		if price < 0 {
			return false
		}
	}
	return true
}

func main() {
	prices := []int {100, 180, 260, 310, 40, 535, 695}
	mp := maxProfit(prices, 0, len(prices)-1)
	fmt.Print(mp)
}

func TestEmptyPrices(t *testing.T) {
	fmt.Println("TestEmptyPrices")
	var prices []int
	mp := maxProfit(prices, 0, len(prices)-1)
	assert.Equal(t, 0, mp, "Error EmptyPrices")
}

func TestDescending(t *testing.T) {
	prices := []int {500, 400, 300, 200, 100}
	mp := maxProfit(prices, 0, len(prices)-1)
	assert.Equal(t, 0, mp, "Error Descending")
}

func TestFirstPriceGreatestDifferentStarting(t *testing.T) {
	prices := []int {500, 400, 300, 200, 100}
	mp := maxProfit(prices, 1, len(prices)-1)
	assert.Equal(t, 0, mp, "Error DescendingDifferentStarting")
}

func TestAscending(t *testing.T) {
	prices := []int {100, 200, 300, 400, 500}
	mp := maxProfit(prices, 0, len(prices)-1)
	assert.Equal(t, 400, mp, "Error Ascending")
}

func TestMixedPrices(t *testing.T) {
	prices := []int {100, 180, 260, 310, 40, 535, 695}
	mp := maxProfit(prices, 0, len(prices)-1)
	assert.Equal(t, 865, mp,"Error MixedPrices")
}

func TestMixedPricesDifferentStart(t *testing.T) {
	prices := []int {100, 180, 260, 310, 40, 535, 695}
	mp := maxProfit(prices, 1, len(prices)-1)
	assert.Equal(t, 785, mp,"Error MixedPricesDifferentStart")
}

func TestMixedPricesUpsAndDowns (t *testing.T) {
	prices := []int {100, 180, 260, 310, 40, 535, 695, 235, 108, 78, 85, 88, 88, 285, 316, 374}
	mp := maxProfit(prices, 0, len(prices)-1)
	assert.Equal(t, 1161, mp,"Error MixedPricesUpsAndDowns")
}

func TestNegativePricesDescending(t *testing.T) {
	prices := []int {-100, -200, -300, -400, -500}
	mp := maxProfit(prices, 0, len(prices)-1)
	assert.Equal(t, 0, mp, "Error NegativePricesDescending")
}

func TestNegativePricesAscending(t *testing.T) {
	prices := []int {-500, -400, -300, -200, -100}
	mp := maxProfit(prices, 0, len(prices)-1)
	assert.Equal(t, 0, mp, "Error NegativePricesAscending")
}

func TestFirstBuyInMiddle (t *testing.T){
	prices := []int {12, 11, 13, 16}
	mp := maxProfit(prices, 0, len(prices)-1)
	assert.Equal(t, 5, mp, "ErrorFirstBuyInMiddle")
}

func Test (t *testing.T){
	prices := []int {12, 11, 13, 16, 12, 11}
	mp := maxProfit(prices, 0, len(prices)-1)
	assert.Equal(t, 5, mp, "Error")
}

func TestSellOnBeforeEnd (t *testing.T){
	prices := []int {12, 11, 13, 16, 12}
	mp := maxProfit(prices, 0, len(prices)-1)
	assert.Equal(t, 5, mp, "Error SellOneBeforeEnd")
}

func TestShortenWeek(t *testing.T) {
	prices := []int {12, 11, 13, 16, 12}
	mp := maxProfit(prices, 0, 3)
	assert.Equal(t, 5, mp, "Error ShortenWeek")
}

func TestTwoDayWeek(t *testing.T){
	prices := []int {11, 12}
	mp := maxProfit(prices, 0, len(prices)-1)
	assert.Equal(t, 1, mp, "Error TwoDayWeek")
}

func TestOneDayWeek (t *testing.T) {
	prices := []int {11}
	mp := maxProfit(prices, 0, len(prices)-1)
	assert.Equal(t, 0, mp, "Error OneDayWeek")
}

func TestWorstCaseBigORuntime(t *testing.T) {
	prices := []int {11, 12, 10, 15, 13, 19, 15, 20}
	mp := maxProfit(prices, 0, len(prices)-1)
	assert.Equal(t, 17, mp, "Error WorstCaseBigORuntime")
}