package Queue

import "lld-practice/pub-sub/Consumer"

type Queue interface {
	// SendToTopic add message to topic.
	SendToTopic(topic string, data []byte)

	// Subscribe add consumer to topic.
	Subscribe(topic string, consumer Consumer.Consumer)

	// Notify will trigger consumer.consume method for the messages in topic queue
	// TODO: Remove from interface
	Notify()

	// Run just runs
	Run()
}
