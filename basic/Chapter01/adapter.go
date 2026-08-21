package main

import (
	"fmt"
)

// IProces interface
type AdapterIProcess interface {
	process()
}

// Adapter struct
type Adapter struct {
	adaptee Adaptee
}

// Adapter class method process
func (adapter Adapter) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

// Adaptee Struct
type Adaptee struct {
	adapterType int
}

// Adaptee class method convert
func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

// AdapterMain method
func AdapterMain() {

	var processor AdapterIProcess = Adapter{}

	processor.process()

}

func main() {
	AdapterMain()
	BacktrackingMain()
	BridgeMain()
	BruteforceMain()
	ComplexityMain()
	CompositeMain()
	CubicComplexityMain()
	DecoratorMain()
	DivideMain()
	FacadeMain()
	FlyweightMain()
	HeapMain()
	LinearComplexityMain()
	ListMain()
	NumOperationsMain()
	PrivateClassMain()
	QuadraticComplexityMain()
	TreeMain()
	TuplesMain()
	VirtualProxyMain()
}
