package entity

type Snake struct {
	start int
	tail  int
}

func NewSnake(start int, tail int) *Snake {
	return &Snake{start: start, tail: tail}
}

func (s *Snake) GetTail() int {
	return s.tail
}
