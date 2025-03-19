package prototype

import (
	"fmt"
)

// Address represents a nested struct
type Address struct {
	City  string
	State string
}

// This is defined in basic.go file
// Prototype interface for cloning
// type Prototype interface {
// 	Clone() Prototype
// 	GetInfo() string
// }

// User represents a complex object with a nested struct and a map
type User struct {
	Name        string
	Age         int
	Address     *Address
	Preferences map[string]string
}

// Clone method to perform a deep copy
func (u *User) Clone() Prototype {
	// Deep copy of Address (creating a new instance)
	newAddress := &Address{
		City:  u.Address.City,
		State: u.Address.State,
	}

	// Deep copy of Preferences map (creating a new instance)
	newPreferences := make(map[string]string)
	for key, value := range u.Preferences {
		newPreferences[key] = value
	}

	// Return a new User instance with independent data
	return &User{
		Name:        u.Name,
		Age:         u.Age,
		Address:     newAddress,
		Preferences: newPreferences,
	}
}

// GetInfo returns user details
func (u *User) GetInfo() string {
	return fmt.Sprintf("User: %s, Age: %d, Location: %s, %s, Preferences: %v",
		u.Name, u.Age, u.Address.City, u.Address.State, u.Preferences)
}

// This is a struct to be invoked by main.go
type DeepPrototype struct{}

func (d DeepPrototype) ProcessPattern() bool {
	// Create an original User
	originalUser := &User{
		Name:    "Alice",
		Age:     30,
		Address: &Address{City: "New York", State: "NY"},
		Preferences: map[string]string{
			"theme": "dark",
			"lang":  "EN",
		},
	}

	// Clone the User
	clonedUser := originalUser.Clone().(*User)

	// Modify cloned data to ensure deep copy works
	clonedUser.Name = "Bob"
	clonedUser.Age = 25
	clonedUser.Address.City = "Los Angeles"
	clonedUser.Address.State = "CA"
	clonedUser.Preferences["theme"] = "light"

	// Print original and cloned user details
	fmt.Println("Original and Cloned Users:")
	fmt.Println(originalUser.GetInfo()) // User: Alice, Age: 30, Location: New York, NY, Preferences: map[theme:dark lang:EN]
	fmt.Println(clonedUser.GetInfo())   // User: Bob, Age: 25, Location: Los Angeles, NY, Preferences: map[theme:light lang:EN]
	return true
}
