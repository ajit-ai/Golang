// main package has examples shown
// in Go Data Structures and algorithms book
package main

// Weighted Reference Counter
type WeightedReferenceCounter struct {
	weight int
}

// new Weighted Reference Counter method
func newWeightedReferenceCounter(weight int) WeightedReferenceCounter {
	return WeightedReferenceCounter{
		weight: weight,
	}
}

// GetWeightedReferences returns the weighted references to sum over
func GetWeightedReferences() []WeightedReferenceCounter {
	return []WeightedReferenceCounter{
		newWeightedReferenceCounter(10),
		newWeightedReferenceCounter(20),
		newWeightedReferenceCounter(12),
	}
}

// WeightedReference method
func WeightedReference() int {

	references := GetWeightedReferences()

	var reference WeightedReferenceCounter

	var sum int
	for _, reference = range references {

		sum = sum + reference.weight

	}

	return sum

}

// WeightedReferenceMain method
func WeightedReferenceMain() {
	println("total weight:", WeightedReference())
}
