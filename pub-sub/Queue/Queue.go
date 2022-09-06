package Queue

import (
	"fmt"
	"lld-practice/pub-sub/Consumer"
	"lld-practice/pub-sub/Message"
	"sync"
	"time"
)

type MessageList []Message.Message

type TopicInfo struct {
	messages MessageList
	len      int
}

type ConsumerInfo struct {
	consumer Consumer.Consumer
	offset   int
}

// TODO: Can use interface as type for Subscribers which will implement get and put to accommodate variable data types

type SimpleQueue struct {
	wg          sync.WaitGroup
	Subscribers map[string][]ConsumerInfo
	TopicQueues map[string]TopicInfo
}

func NewSimpleQueue() SimpleQueue {
	subscribers := make(map[string][]ConsumerInfo)
	topicQueues := make(map[string]TopicInfo)
	return SimpleQueue{
		wg:          sync.WaitGroup{},
		Subscribers: subscribers,
		TopicQueues: topicQueues,
	}
}

func (s *SimpleQueue) Subscribe(topic string, consumer Consumer.Consumer) {
	if _, ok := s.Subscribers[topic]; !ok {
		s.Subscribers[topic] = make([]ConsumerInfo, 0)
	}
	s.Subscribers[topic] = append(s.Subscribers[topic], ConsumerInfo{
		consumer: consumer,
		offset:   0,
	})
	fmt.Println("subscribed to topic ", topic)
}

func (s *SimpleQueue) SendToTopic(topic string, data []byte) {
	if _, ok := s.TopicQueues[topic]; !ok {
		messageList := make(MessageList, 0)
		ti := TopicInfo{
			messages: messageList,
			len:      0,
		}
		s.TopicQueues[topic] = ti
	}
	id := time.Now().UnixNano()
	ch := make(chan int64)
	msg := Message.Message{Data: data, Id: id, Ch: ch}
	ti := s.TopicQueues[topic]
	ti.messages = append(ti.messages, msg)
	//TODO: remove
	ti.len++
	s.TopicQueues[topic] = ti

	fmt.Println("added message to topic ", topic)
}

// Tightly coupled with internal ds implementation
func (s *SimpleQueue) Notify() {
	fmt.Println("running notify")
	for topic, info := range s.TopicQueues {
		if s.TopicQueues[topic].len > 0 && len(s.Subscribers[topic]) > 0 {
			for idx, consumer := range s.Subscribers[topic] {
				s.wg.Add(1)
				fmt.Printf("consumer offset %+v for topic %+v and message len:%+v\n", consumer.offset, topic, info.len)
				if consumer.offset < info.len {
					go func(idx int, topic string, info TopicInfo) {
						defer s.wg.Done()
						fmt.Println("called consumer for topic", topic)
						subscriberInfo := s.Subscribers[topic][idx]
						subscriberInfo.consumer.Consume(info.messages[subscriberInfo.offset])
						subscriberInfo.offset += 1
						s.Subscribers[topic][idx] = subscriberInfo
					}(idx, topic, info)
				}

			}
		}
	}

	s.wg.Wait()
	fmt.Println("finished waiting for consumers")
}

func (s *SimpleQueue) Run() {
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
					s.Notify()
				}

			}
		}
	}()
	done <- true
	ticker.Stop()
}
