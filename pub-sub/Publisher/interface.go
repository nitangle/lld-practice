package Publisher

import "pub-sub/Queue"

type Publisher interface {
	Publish(topic string, q Queue.Queue)
}
