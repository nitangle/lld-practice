package Message

type Message struct {
	Id   int64
	Data []byte
	Ch   chan int64
}

func (m Message) Finish() {
	m.Ch <- m.Id
}
