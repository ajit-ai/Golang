// main package demonstrates struct embedding, promotion,
// receiver semantics and struct tags
package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Engine is embedded by value into Car
type Engine struct {
	Horsepower int
	Type       string
}

// Start is promoted to Car through embedding
func (e Engine) Start() string {
	return e.Type + " engine started"
}

// Wheels is embedded as a pointer
type Wheels struct {
	Count int
}

// Car embeds Engine and Wheels: their fields and methods are promoted
type Car struct {
	Engine // embedded struct: Car has Horsepower, Type, Start()
	*Wheels
	Brand string
}

// Describe uses both its own field and promoted fields
func (c Car) Describe() string {
	return fmt.Sprintf("%s with %d HP on %d wheels", c.Brand, c.Horsepower, c.Count)
}

// Counter demonstrates value vs pointer receivers
type Counter struct {
	n int
}

// IncrementValue works on a copy: the caller never sees the change
func (c Counter) IncrementValue() Counter {
	c.n++
	return c
}

// IncrementPointer mutates the caller's value
func (c *Counter) IncrementPointer() {
	c.n++
}

// Config shows struct tags consumed via reflection
type Config struct {
	Host   string `json:"host" validate:"required"`
	Port   int    `json:"port" validate:"min=1,max=65535"`
	Debug  bool   `json:"debug,omitempty" validate:"optional"`
	APIKey string `json:"-"` // "-" means skip entirely
}

// TagLookup reads the validate tag of every exported field via reflection
func TagLookup(v interface{}) map[string]string {
	tags := make(map[string]string)
	t := reflect.TypeOf(v)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("validate")
		if tag != "" {
			tags[field.Name] = tag
		}
	}
	return tags
}

// EmbeddingMain demonstrates embedding and promotion
func EmbeddingMain() {
	car := Car{
		Engine: Engine{Horsepower: 300, Type: "V6"},
		Wheels: &Wheels{Count: 4},
		Brand:  "GoMobile",
	}
	fmt.Println(car.Start())    // promoted method from Engine
	fmt.Println(car.Describe()) // mixes own + promoted fields
}

// ReceiversMain contrasts value and pointer receivers
func ReceiversMain() {
	c := Counter{}
	_ = c.IncrementValue()
	fmt.Println("after value receiver:", c.n) // still 0

	c.IncrementPointer()
	c.IncrementPointer()
	fmt.Println("after pointer receiver:", c.n) // 2
}

// TagsMain demonstrates struct tags read back with reflection
func TagsMain() {
	for name, tag := range TagLookup(Config{}) {
		fmt.Printf("%s -> %s\n", name, tag)
	}
	var _ = strings.TrimSpace // keep strings import for future helpers
}

// main runs the demo entry points of this package
func main() {
	EmbeddingMain()
	ReceiversMain()
	TagsMain()
}
