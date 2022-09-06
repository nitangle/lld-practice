package Publisher

import (
	"encoding/json"
	"fmt"
	"lld-practice/pub-sub/Queue"
)

type SimplePublisher struct{}

func NewSimplePublisher() SimplePublisher {
	return SimplePublisher{}
}
func (s SimplePublisher) Publish(topic string, q Queue.Queue) {
	actualData := struct {
		Message string
	}{
		Message: fmt.Sprintf("publisher 1 published to topic %s", topic),
	}
	dataToSend, err := json.Marshal(actualData)
	if err != nil {
		fmt.Println("could not encode message in publisher 1")
		return
	}
	q.SendToTopic(topic, dataToSend)
}
