package Consumer

import (
	"encoding/json"
	"fmt"
	pub_sub "pub-sub/Message"
)

type SimpleConsumer struct{}

type SimpleMessage struct {
	Message string
}

func NewSimpleConsumer() SimpleConsumer {
	return SimpleConsumer{}
}
func (c SimpleConsumer) Consume(msg pub_sub.Message) {
	var decodedData SimpleMessage
	err := json.Unmarshal(msg.Data, &decodedData)
	if err != nil {
		fmt.Println("err in decoding message in simple consumer")
		return
	}
	fmt.Printf("consumed message : %+v\n", decodedData)
}
