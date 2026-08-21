// main package has examples shown
// in Go Data Structures and algorithms book
package main

// object is a simulated heap object for the garbage collector examples
type object struct {
	id         int
	generation int
	references []*object
}

// newObject creates a simulated heap object and registers it as live
func newObject(id int, generation int, references ...*object) *object {
	return &object{
		id:         id,
		generation: generation,
		references: references,
	}
}

// markBits simulates the mark bitmap of the collector
var markBits = map[*object]bool{}

// liveObjects simulates the heap the collector sweeps over
var liveObjects []*object

// registerObject adds an object to the simulated heap
func registerObject(o *object) {
	liveObjects = append(liveObjects, o)
}

// GetObjects returns every object currently on the simulated heap
func GetObjects() []*object {
	return liveObjects
}

// GetObjectsFromOldGeneration returns objects of the given generation
func GetObjectsFromOldGeneration(generation int) []*object {
	var old []*object
	for _, o := range liveObjects {
		if o.generation == generation {
			old = append(old, o)
		}
	}
	return old
}

// IfMarked reports whether the object is marked in the mark bitmap
func IfMarked(o *object) bool {
	return markBits[o]
}

// SetMarked marks the object in the mark bitmap
func SetMarked(o *object) {
	markBits[o] = true
}

// Release frees the object: it is removed from the heap and unmarked
func Release(o *object) {
	delete(markBits, o)
	for i, live := range liveObjects {
		if live == o {
			liveObjects = append(liveObjects[:i], liveObjects[i+1:]...)
			break
		}
	}
}

// Mark marks root and every object reachable from it
func Mark(root *object) {

	markedAlready := IfMarked(root)
	if !markedAlready {

		SetMarked(root)

	}

	references := GetReferences(root)

	for _, reference := range references {

		Mark(reference)
	}

}

// GetReferences returns the objects referenced by root
func GetReferences(root *object) []*object {
	return root.references
}

// MarkMain demonstrates marking a small object graph
func MarkMain() {

	a := newObject(1, 0)
	b := newObject(2, 0, a)
	registerObject(b)
	registerObject(a)

	Mark(b)

	println("marked a:", IfMarked(a))
	println("marked b:", IfMarked(b))
}
