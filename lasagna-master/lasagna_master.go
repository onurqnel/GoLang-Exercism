package lasagna

// PreparationTime() calculates the total preparation time based on the number of layers.
// If avgTimePerLayer is 0, it defaults to 2 minutes per layer.
func PreparationTime(layerNames []string, prepTime int) int {
	if prepTime > 0 {
		return len(layerNames) * prepTime
	} else {
		return len(layerNames) * 2
	}
}

// Quantities calculates the amount of noodles and sauce needed for lasagna based on the given layers.
// It returns the quantity of noodles in grams (int) and sauce in liters (float64).
func Quantities(layerNames []string) (int, float64) {
	var sauceQuantitie int
	var noodleQuantitite float64

	for i := 0; i < len(layerNames); i++ {
		if layerNames[i] == "noodles" {
			sauceQuantitie += 50
		} else if layerNames[i] == "sauce" {
			noodleQuantitite += 0.2
		}
	}
	return sauceQuantitie, noodleQuantitite
}

/**
OR
func Quantities(layerNames []string) (int, float64) {
	var noodleQuantity int
	var sauceQuantity float64

	for _, layer := range layerNames {
		if layer == "noodles" {
			noodleQuantity += 50
		} else if layer == "sauce" {
			sauceQuantity += 0.2
		}
	}
	return noodleQuantity, sauceQuantity
}
**/

// AddSecretIngredient replaces the last element "?" in ownIngredients with the last element from friendsIngredients.
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

/**
OR
func AddSecretIngredient(friendsList, myList []string) {
	for i := range myList {
		if i == len(myList)-1 {
			myList[i] = friendsList[len(friendsList)-1]
		}
	}
}
**/

// ScaleRecipe scales the recipe amounts for the desired number of portions.
// It takes a slice of float64 amounts for 2 portions (quantities) and an integer (portionCount).
// Returns a new slice of float64 with scaled amounts for the desired number of portions.
func ScaleRecipe(quantities []float64, portionCount int) []float64 {
	scaledQuantities := []float64{}

	for i := 0; i < len(quantities); i++ {
		scaledAmount := (float64(portionCount) / 2) * quantities[i]
		scaledQuantities = append(scaledQuantities, scaledAmount)
	}
	return scaledQuantities
}

/**
OR
func ScaleRecipe(quantities []float64, portionCount int) []float64 {
	scaledQuantities := make([]float64, len(quantities))

	for i, quantity := range quantities {
		scaledQuantities[i] = (float64(portionCount) / 2) * quantity
	}

	return scaledQuantities
}
**/
