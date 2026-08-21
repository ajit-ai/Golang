// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// GenerationCollect marks every object of the given old generation
func GenerationCollect() {

	currentGeneration := 3

	objects := GetObjectsFromOldGeneration(currentGeneration)

	for _, o := range objects {

		markedAlready := IfMarked(o)
		if !markedAlready {

			SetMarked(o)

		}
	}

}

// GenerationCollectMain demonstrates generational marking
func GenerationCollectMain() {

	old := newObject(1, 3)
	young := newObject(2, 1)
	registerObject(old)
	registerObject(young)

	GenerationCollect()

	println("old marked:", IfMarked(old))
	println("young marked:", IfMarked(young))
}
