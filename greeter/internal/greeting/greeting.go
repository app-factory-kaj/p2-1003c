// Package greeting holds the domain logic for producing a Greeting.
package greeting

import "fmt"

// Greeting is the JSON response body for GET /hello.
type Greeting struct {
	Message string `json:"message"`
}

const defaultName = "World"

// New builds a Greeting addressed to name, or to the default name when name
// is empty.
func New(name string) Greeting {
	if name == "" {
		name = defaultName
	}
	return Greeting{Message: fmt.Sprintf("Hello, %s!", name)}
}
