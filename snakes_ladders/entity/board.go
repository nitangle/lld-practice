package entity

import (
	"fmt"
)

type Board struct {
	size    int
	snakes  map[int]*Snake
	ladders map[int]*Ladder
}

func NewBoard() *Board {
	b := &Board{}
	b.Setup()
	return b
}

func (b *Board) GetWinningPos() int {
	return b.size * b.size
}
func (b *Board) IsOutside(pos int) bool {
	return pos > (b.size * b.size)
}
func (b *Board) GetSnakeTail(head int) int {
	return b.snakes[head].GetTail()
}

func (b *Board) GetLadderTop(bottom int) int {
	return b.ladders[bottom].GetTop()
}
func (b *Board) Setup() {
	var (
		sz      int
		snakes  int
		ladders int
	)
	fmt.Scanln(&sz)
	if sz < 0 {
		fmt.Println("please enter positive size of board")
		fmt.Scanln(&sz)
	}
	b.size = sz
	fmt.Println("please enter number of snakes in board")
	fmt.Scanln(&snakes)
	fmt.Println("please enter HEAD and TAIL of snake in order")
	for i := 0; i < snakes; i++ {
		var (
			tail int
			head int
		)
		fmt.Scanln(&tail, &head)
		b.snakes[head] = NewSnake(head, tail)
	}
	fmt.Println("please enter number of ladders in board")
	fmt.Scanln(&ladders)
	fmt.Println("please enter BOTTOM and TOP of ladders in order")
	for i := 0; i < ladders; i++ {
		var (
			bottom int
			top    int
		)
		fmt.Scanln(&bottom, &top)
		b.ladders[bottom] = NewLadder(bottom, top)
	}

}

func (b *Board) CheckSnake(pos int) bool {
	if _, ok := b.snakes[pos]; ok {
		return true
	}
	return false
}

func (b *Board) CheckLadder(pos int) bool {
	if _, ok := b.ladders[pos]; ok {
		return true
	}
	return false
}
