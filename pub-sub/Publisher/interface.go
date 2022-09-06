package Publisher

import "lld-practice/pub-sub/Queue"

type Publisher interface {
	Publish(topic string, q Queue.Queue)
}
