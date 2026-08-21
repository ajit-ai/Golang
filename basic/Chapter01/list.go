// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and container list packages
import (
	"container/list"
	"fmt"
)

func BuildIntList(values ...int) *list.List {
	var intList list.List

	for _, value := range values {
		intList.PushBack(value)
	}

	return &intList
}

// ListMain method
func ListMain() {
	var intList = BuildIntList(11, 23, 34)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
