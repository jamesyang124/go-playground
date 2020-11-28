package basic

import "fmt"

// User ...
type User struct {
	FirstName, LastName string
}

// Namer is an interface,
// Type method binding on it will implicitly polymorphism this interface type
type Namer interface {
	Name() string
}

// Name ...
func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

// Greet accept any type which impl. Namer interface's methods
func Greet(n Namer) string {
	return fmt.Sprintf("Greeting! %s", n.Name())
}
