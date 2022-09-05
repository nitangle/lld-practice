package Consumer

import (
	pub_sub "pub-sub/Message"
)

type Consumer interface {
	Consume(msg pub_sub.Message)
}
