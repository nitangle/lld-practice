package entity

import (
	"math/rand"
	"time"
)

type Dice struct {
	side int
}

func NewDice(side int) *Dice {
	return &Dice{side: side}
}

func (d *Dice) Roll() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(d.side) + 1
}
