package pattern

import (
	"fmt"
	"software-design/builder"
	"software-design/factory"
	"software-design/observer"
)

const Factory = "factory"
const Observer = "observer"
const ObserverChannels = "channels"
const Builder = "builder"

type Pattern interface {
	// Shows the example of the pattern
	ProcessPattern() bool
}

// CreatePattern creates a pattern based on the name
func CreatePattern(name string) (Pattern, error) {
	switch name {
	case Factory:
		return factory.FactoryPattern{}, nil
	case Observer:
		return observer.ObserverPattern{}, nil
	case ObserverChannels:
		return observer.ChannelsObserverPattern{}, nil
	case Builder:
		return builder.BuilderPattern{}, nil
	default:
		return nil, fmt.Errorf("pattern not supported: %s", name)
	}
}
