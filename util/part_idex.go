package util

func BeginIdxPart(partstring string) int {
	// Create a map that maps parts to values
	strBytes := []byte(partstring)
	parts := rune(strBytes[len(strBytes)-1])
	partsString := string(parts)
	partToValue := map[string]int{
		"1": 0,
		"2": 6,
		"3": 31,
		"4": 70,
		"5": 100,
		"6": 130,
		"7": 146,
		"8": 200,
	}

	// Check if the part exists in the map, and return the corresponding value
	if value, exists := partToValue[partsString]; exists {
		return value
	}

	// Default value for unknown parts
	return -1
}
func LimitPerPart(partstring string) int {
	// Create a map that maps parts to values
	strBytes := []byte(partstring)
	parts := rune(strBytes[len(strBytes)-1])
	partsString := string(parts)
	partToValue := map[string]int{
		"1": 6,
		"2": 25,
		"3": 39,
		"4": 30,
		"5": 30,
		"6": 16,
		"7": 54,
	}

	// Check if the part exists in the map, and return the corresponding value
	if value, exists := partToValue[partsString]; exists {
		return value * 4
	}

	// Default value for unknown parts
	return -1
}
