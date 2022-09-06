package Consumer

import (
	pub_sub "lld-practice/pub-sub/Message"
)

type Consumer interface {
	Consume(msg pub_sub.Message)
}
