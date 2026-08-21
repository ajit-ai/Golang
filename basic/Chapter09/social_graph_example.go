// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// Name type
type Name string

// Social Graph class
type ExampleSocialGraph struct {
	GraphNodes map[Name]struct{}
	Links      map[Name]map[Name]struct{}
}

// NewExampleSocialGraph method
func NewExampleSocialGraph() *ExampleSocialGraph {
	return &ExampleSocialGraph{
		GraphNodes: make(map[Name]struct{}),
		Links:      make(map[Name]map[Name]struct{}),
	}
}

// AddEntity method
func (ExampleSocialGraph *ExampleSocialGraph) AddEntity(name Name) bool {

	var exists bool
	if _, exists = ExampleSocialGraph.GraphNodes[name]; exists {
		return true
	}
	ExampleSocialGraph.GraphNodes[name] = struct{}{}
	return true
}

// Add Link
func (ExampleSocialGraph *ExampleSocialGraph) AddLink(name1 Name, name2 Name) {
	var exists bool
	if _, exists = ExampleSocialGraph.GraphNodes[name1]; !exists {
		ExampleSocialGraph.AddEntity(name1)
	}
	if _, exists = ExampleSocialGraph.GraphNodes[name2]; !exists {
		ExampleSocialGraph.AddEntity(name2)
	}

	if _, exists = ExampleSocialGraph.Links[name1]; !exists {
		ExampleSocialGraph.Links[name1] = make(map[Name]struct{})
	}
	ExampleSocialGraph.Links[name1][name2] = struct{}{}

}

func (ExampleSocialGraph *ExampleSocialGraph) PrintLinks() {
	var root Name
	root = Name("Root")

	fmt.Printf("Printing all links adjacent to %s\n", root)

	var node Name
	for node = range ExampleSocialGraph.Links[root] {
		fmt.Printf("Link: %s -> %s\n", root, node)
	}

	var m map[Name]struct{}
	fmt.Println("Printing all links.")
	for root, m = range ExampleSocialGraph.Links {
		var vertex Name
		for vertex = range m {
			fmt.Printf("Link: %s -> %s\n", root, vertex)
		}
	}
}

// main method
func SocialGraphExampleMain() {

	var ExampleSocialGraph *ExampleSocialGraph

	ExampleSocialGraph = NewExampleSocialGraph()

	var root Name = Name("Root")
	var john Name = Name("John Smith")
	var per Name = Name("Per Jambeck")
	var cynthia Name = Name("Cynthia Gibas")

	ExampleSocialGraph.AddEntity(root)
	ExampleSocialGraph.AddEntity(john)
	ExampleSocialGraph.AddEntity(per)
	ExampleSocialGraph.AddEntity(cynthia)

	ExampleSocialGraph.AddLink(root, john)
	ExampleSocialGraph.AddLink(root, per)
	ExampleSocialGraph.AddLink(root, cynthia)

	var mayo Name = Name("Mayo Smith")
	var lorrie Name = Name("Lorrie Jambeck")
	var ellie Name = Name("Ellie Vlocksen")

	ExampleSocialGraph.AddLink(john, mayo)
	ExampleSocialGraph.AddLink(john, lorrie)
	ExampleSocialGraph.AddLink(per, ellie)

	ExampleSocialGraph.PrintLinks()
}
