// main package demonstrates Go interfaces: implicit satisfaction,
// polymorphism, interface embedding and fmt.Stringer
package main

import (
	"fmt"
	"strings"
)

// Speaker is satisfied implicitly by any type with a Speak method
type Speaker interface {
	Speak() string
}

// Runner adds movement behaviour
type Runner interface {
	Run() string
}

// Swimmer adds swimming behaviour
type Swimmer interface {
	Swim() string
}

// Triathlete composes two interfaces into a larger one
type Triathlete interface {
	Runner
	Swimmer
	Speak() string
}

// Dog implements Speaker without declaring it anywhere
type Dog struct{ Name string }

func (d Dog) Speak() string { return d.Name + " says woof" }

// Robot implements Speaker
type Robot struct{ Model string }

func (r Robot) Speak() string { return r.Model + " beeps" }

// Human implements Speaker, Runner and Swimmer, therefore also Triathlete
type Human struct{ Name string }

func (h Human) Speak() string { return h.Name + " says hello" }
func (h Human) Run() string   { return h.Name + " runs 10km" }
func (h Human) Swim() string  { return h.Name + " swims 1km" }

// GreetAll works on any Speaker: polymorphism in action
func GreetAll(speakers []Speaker) []string {
	lines := make([]string, 0, len(speakers))
	for _, s := range speakers {
		lines = append(lines, s.Speak())
	}
	return lines
}

// TrainAthlete accepts the composed interface
func TrainAthlete(t Triathlete) string {
	return strings.Join([]string{t.Run(), t.Swim(), t.Speak()}, "; ")
}

// Person satisfies fmt.Stringer so fmt prints it nicely
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

// InterfacesMain demonstrates polymorphism and Stringer
func InterfacesMain() {
	speakers := []Speaker{Dog{Name: "Rex"}, Robot{Model: "RX-7"}, Human{Name: "Ajit"}}
	for _, line := range GreetAll(speakers) {
		fmt.Println(line)
	}

	var p fmt.Stringer = Person{Name: "Ajit", Age: 30}
	fmt.Println("Stringer:", p)
}

// EmbeddingMain demonstrates composed interfaces
func EmbeddingMain() {
	var t Triathlete = Human{Name: "Priya"}
	fmt.Println(TrainAthlete(t))
}

// main runs the demo entry points of this package
func main() {
	InterfacesMain()
	EmbeddingMain()
}
