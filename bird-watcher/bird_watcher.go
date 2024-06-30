package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sumBird int
	for i := 0; i < len(birdsPerDay); i++ {
		sumBird += birdsPerDay[i]
	}
	return sumBird
}

/**
OR
func TotalBirdCount(birdsPerDay []int) int{
	var sumBird int
	for_, count := range birdsPerdDay {
		sumBird += count
	}
	return sumBird
}
**/

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var sumBird int
	startWeek := (week - 1) * 7
	endWeek := startWeek + 7

	if endWeek > len(birdsPerDay) {
		endWeek = len(birdsPerDay)
	}

	for i := startWeek; i < endWeek; i++ {
		sumBird += birdsPerDay[i]
	}
	return sumBird
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i == 0 || i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}

/**
OR
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i+= 2{
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}
**/
