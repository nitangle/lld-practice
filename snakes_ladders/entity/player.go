package entity

type Player struct {
	name       string
	currentPos int
}

func NewPlayer(name string) *Player {
	return &Player{name: name, currentPos: 0}
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) SetPos(pos int) {
	p.currentPos = pos
}

func (p *Player) GetCurrentPos() int {
	return p.currentPos
}
func (p *Player) Roll(d *Dice) int {
	return d.Roll()
}
