package lasagna

// PreparationTime estimates how long it will take to prepare a lasagna.
func PreparationTime(layers []string, avgLayerPrepMins int) int {
	if avgLayerPrepMins == 0 {
		avgLayerPrepMins = 2
	}

	return len(layers) * avgLayerPrepMins
}

// Quantities computes the amount of noodles and sauce needed.
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50 // grams
		}
		if layer == "sauce" {
			sauce += 0.2 // liters
		}
	}
	return
}

// AddSecretIngredient replaces the last ingredient from myList with the last ingredient from friendsList.
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
	return
}

// ScaleRecipe calculates the amounts needed for different numbers of portions.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var totalQuantities []float64 = make([]float64, len(quantities))

	for i, quantity := range quantities {
		totalQuantities[i] = (quantity / 2) * float64(portions)
	}
	return totalQuantities
}
