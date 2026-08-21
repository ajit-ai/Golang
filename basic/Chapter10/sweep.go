// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// Sweep releases every unmarked object on the simulated heap
func Sweep() {

	var doomed []*object

	for _, o := range GetObjects() {

		if !IfMarked(o) {

			doomed = append(doomed, o)

		}
	}

	for _, o := range doomed {

		Release(o)

	}

}

// SweepMain demonstrates sweeping unmarked objects
func SweepMain() {

	before := len(GetObjects())

	kept := newObject(1, 0)
	freed := newObject(2, 0)
	registerObject(kept)
	registerObject(freed)

	SetMarked(kept)

	Sweep()

	println("only the unmarked object was released:", len(GetObjects()) == before+1)
}
