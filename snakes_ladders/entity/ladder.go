package entity

type Ladder struct {
	bottom int
	top    int
}

func NewLadder(bottom int, top int) *Ladder {
	return &Ladder{
		bottom: bottom,
		top:    top,
	}
}
func (l *Ladder) GetTop() int {
	return l.top
}
