package main

import (
	"fmt"
	"math"
	"time"
)

type CostComputation interface{
	Price(startTime time.Time, ps ParkingSpot) float64
}

type defaultCostComputation struct {}

func NewDefaultCostComputation() CostComputation {
	return defaultCostComputation{}
}

func (c defaultCostComputation) Price(startTime time.Time, ps ParkingSpot) float64 {
	return 20
}


type twoWheelerCostComputation struct {}


func NewTwoWheelerCostComputation() CostComputation {
	return twoWheelerCostComputation{}
}

func (c twoWheelerCostComputation) Price(startTime time.Time, ps ParkingSpot) float64 {
	// cost computation logic for two wheeler

	currTime := time.Now()

	duration := currTime.Sub(startTime)

	fmt.Println(duration.Minutes())

	if duration.Minutes() < 5 {
		return float64(ps.Price())
	}

	// Calculate the price based on the duration (in minutes)
	price := duration.Minutes() * float64(ps.Price())

	// Round to 2 decimal places
	roundedPrice := math.Round(price*100) / 100

	return roundedPrice
}


type fourWheelerCostComputation struct {}

func NewFourWheelerCostComputation() CostComputation {
	return fourWheelerCostComputation{}
}

func (c fourWheelerCostComputation) Price(startTime time.Time, ps ParkingSpot) float64 {
	// cost computation logic for four wheeler

	currTime := time.Now()

	duration := currTime.Sub(startTime)

	if duration.Hours() < 1 {
		return float64(ps.Price())
	}

	return duration.Hours() * float64(ps.Price())
}
