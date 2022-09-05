package entity

import "fmt"

type Game struct {
	board           *Board
	players         map[int]*Player
	dice            *Dice
	lastMovedPlayer int
	status          int
	winner          *Player
}

func NewGame() *Game {
	g := &Game{}
	g.winner = nil
	g.board = NewBoard()
	for i := 1; i <= 4; i++ {
		var name string
		fmt.Println("Enter name of player i")
		fmt.Scanln(&name)
		g.players[i] = NewPlayer(name)
	}
	g.dice = NewDice(6)
	g.lastMovedPlayer = -1
	return g
}

func (g *Game) Start() {
	g.status = 1
	for true {
		playerNum := g.TakeTurn()
		currentPlayer := g.players[playerNum]
		steps := g.dice.Roll()
		possibleNewPos := currentPlayer.GetCurrentPos() + steps
		g.MovePlayer(currentPlayer, possibleNewPos)
		if g.IsWinner(currentPlayer) {
			g.winner = currentPlayer
			fmt.Printf("winner is %+v\n", currentPlayer.GetName())
			g.EndGame()
			break
		}
	}

}
func (g *Game) EndGame() {
	g.status = 2
}
func (g *Game) SetPlayerPos(p *Player, pos int) {
	p.SetPos(pos)
}

func (g *Game) TakeTurn() int {
	if g.lastMovedPlayer == -1 {
		return 0
	}
	return (g.lastMovedPlayer + 1) % len(g.players)
}

func (g *Game) IsWinner(p *Player) bool {
	if p.GetCurrentPos() == g.board.GetWinningPos() {
		return true
	}
	return false
}

func (g *Game) MovePlayer(p *Player, possibleNewPos int) {
	currentPos := p.GetCurrentPos()
	//rules
	//going outside the board
	if g.board.IsOutside(possibleNewPos) {
		p.SetPos(currentPos)
	}
	if g.board.CheckSnake(possibleNewPos) {
		p.SetPos(g.board.GetSnakeTail(possibleNewPos))
	}
	if g.board.CheckLadder(possibleNewPos) {
		p.SetPos(g.board.GetLadderTop(possibleNewPos))
	}

}
