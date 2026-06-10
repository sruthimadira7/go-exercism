package lasagna

// OvenTime represents the number of minutes lasagna takes to cook in oven.
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
    return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    return numberOfLayers * 2
}

// ElapsedTime calculates the time elapsed cooking the lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    return PreparationTime(numberOfLayers) + actualMinutesInOven
}
