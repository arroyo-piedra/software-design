package observer

import "fmt"

// Defines the interface for receiving updates from the subject.
type Observer interface {
	Update(string) // Method to update the observer, used by the subject.
}

// In a classic Oriented Object Language this would defines the Subject's interface
// Subject defines the subject interface.
// type Subject interface {
// 	RegisterObserver(o Observer)
// 	RemoveObserver(o Observer)
// 	NotifyObservers(message string)
// }

// WeatherStation defines the subject that observers observe.
type WeatherStation struct {
	observers []Observer
}

// RegisterObserver adds an observer to the list.
func (cs *WeatherStation) RegisterObserver(o Observer) {
	cs.observers = append(cs.observers, o)
}

// RemoveObserver removes an observer from the list.
func (cs *WeatherStation) RemoveObserver(o Observer) {
	for i, observer := range cs.observers {
		if observer == o {
			cs.observers = append(cs.observers[:i], cs.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers sends updates to all observers.
func (cs *WeatherStation) NotifyObservers(message string) {
	for _, observer := range cs.observers {
		observer.Update(message)
	}
}

// Specific observers (Phone case)
type PhoneDisplay struct {
	weather string
}

// Update is the method that is called by the subject to notify the observer.
func (cd *PhoneDisplay) Update(weather string) {
	cd.weather = weather
	fmt.Printf("Phone Display gets the weather: %s\n", cd.weather)
}

// Specific observers (Tv case)
type TvDisplay struct {
	weather string
}

// Update is the method that is called by the subject to notify the observer.
func (cd *TvDisplay) Update(weather string) {
	cd.weather = weather
	fmt.Printf("Tv Display gets the weather: %s\n", cd.weather)
}

// This is a struct to be invoked by main.go
type ObserverPattern struct{}

func (o ObserverPattern) ProcessPattern() bool {
	fmt.Println("------------ Observer Pattern ------------")

	fmt.Println("Creating Weather Station")
	weatherStation := &WeatherStation{}

	// Observers
	fmt.Println("Creating Observers")
	phoneDisplay := &PhoneDisplay{}
	tvDisplay := &TvDisplay{}

	// Register observers
	fmt.Println("Registering Observers")
	weatherStation.RegisterObserver(phoneDisplay)
	weatherStation.RegisterObserver(tvDisplay)

	// Notify observers
	fmt.Println("Notifying Observers")
	weatherStation.NotifyObservers("It's raining")

	// Remove observer
	fmt.Println("Removing Observer")
	weatherStation.RemoveObserver(phoneDisplay)

	// Notify observers
	fmt.Println("Notifying Observers")
	weatherStation.NotifyObservers("It's sunny")
	return true
}
