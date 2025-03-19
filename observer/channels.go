package observer

import (
	"fmt"
	"sync"
	"time"
)

// This is an example using channels to improve the behavior of the observer pattern in Go.
// This implementation showcases how channels and goroutines can be used in Go to implement the
// Observer pattern, providing a way to manage and notify multiple observers efficiently.

// WeatherStation defines the subject that observers observe.
// (Subject)
type WeatherStationChannel struct {
	observers map[chan string]struct{}
	mu        sync.Mutex
}

func NewWeatherStationChannel() *WeatherStationChannel {
	return &WeatherStationChannel{
		observers: make(map[chan string]struct{}),
	}
}

// Subscribe allows an observer to subscribe to the subject.
func (s *WeatherStationChannel) Subscribe() chan string {
	s.mu.Lock()
	defer s.mu.Unlock()

	ch := make(chan string)
	s.observers[ch] = struct{}{}
	return ch
}

// Unsubscribe allows an observer to unsubscribe from the subject.
func (s *WeatherStationChannel) Unsubscribe(ch chan string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.observers[ch]; ok {
		delete(s.observers, ch)
		close(ch)
	}
}

// The Notify method sends updates to all registered observers by iterating over the map and
// sending the message to each channel.
func (s *WeatherStationChannel) Notify(msg string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for ch := range s.observers {
		ch <- msg
	}
	// The code executes too fast, so we need to wait for the observers to process the message
	time.Sleep(100 * time.Millisecond)
}

// The ObserverDevice function simulates an observer that receives updates. It runs in a separate
// goroutine and listens for messages on its channel, printing them out when received.
func ObserverDevice(name string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range ch {
		fmt.Printf("Observer %s received: %s\n", name, msg)
	}
}

// This is a struct to be invoked by main.go
type ChannelsObserverPattern struct{}

func (o ChannelsObserverPattern) ProcessPattern() bool {
	fmt.Println("------------ Observer Pattern with Channels ------------")

	fmt.Println("-> Creating Weather Station")
	subject := NewWeatherStationChannel()

	var wg sync.WaitGroup
	fmt.Println("-> Creating Observers")
	phone := subject.Subscribe()
	tv := subject.Subscribe()

	wg.Add(2)

	fmt.Println("-> Starting Observers")
	go ObserverDevice("Phone", phone, &wg)
	go ObserverDevice("TV", tv, &wg)

	fmt.Println("-> Notifying Observers")
	subject.Notify("It's getting cold")
	subject.Notify("RUN! The winter is comming")

	fmt.Println("-> Unsubscribing Observers")
	subject.Unsubscribe(phone)

	fmt.Println("-> Notifying Observers")
	subject.Notify("It's raining")

	fmt.Println("-> Unsubscribing Observers")
	subject.Unsubscribe(tv)

	// Wait for goroutines to finish
	wg.Wait()
	return true
}
