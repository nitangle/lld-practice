package Queue

import "pub-sub/Consumer"

type Queue interface {
	SendToTopic(topic string, data []byte)
	Subscribe(topic string, consumer Consumer.Consumer)
	Notify()
	Run()
}
