package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return len(layers) * 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	countLayer := 0
	countSauce := 0.0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			countLayer += 1
		}
		if layers[i] == "sauce" {
			countSauce += 1
		}
	}
	return countLayer * 50, countSauce * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, num int) []float64 {
	scaled := make([]float64, len(quantities))
	for i := range len(quantities) {
		scaled[i] = quantities[i] * float64(num) / 2
	}
	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
