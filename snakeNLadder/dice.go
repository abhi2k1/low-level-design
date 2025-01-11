package main

import "math/rand"

type Dice struct {
	noOFDice int
	maxNum   int
}

func NewDice(count int, maxNum int) *Dice {
	return &Dice{
		noOFDice: count,
		maxNum:   maxNum,
	}
}

func (d *Dice) RollDice() int {
	sum := 0

	for i := 0; i < d.noOFDice; i++ {
		sum += rand.Intn(d.maxNum) + 1
	}

	return sum
}
