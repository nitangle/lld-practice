package main

import (
	"fmt"
	"lld-practice/pub-sub/Consumer"
	"lld-practice/pub-sub/Publisher"
	"lld-practice/pub-sub/Queue"
	"time"
)

func main() {
	q := Queue.NewSimpleQueue()
	ticker := time.NewTicker(4 * time.Second)
	done := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-done:
				{
					return
				}
			case <-ticker.C:
				{
					fmt.Println("calling notify")
					q.Notify()
				}

			}
		}
	}()
	p1 := Publisher.SimplePublisher{}
	p1.Publish("t1", &q)
	c1 := Consumer.NewSimpleConsumer()
	q.Subscribe("t1", c1)
	q.Subscribe("t2", c1)
	p1.Publish("t2", &q)
	for _ = range done {
		//
	}
}
