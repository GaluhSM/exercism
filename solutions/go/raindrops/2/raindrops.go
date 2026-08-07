package raindrops

import "strconv"

var drops = []struct {
	factor int
	sound  string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(number int) string {
	var result string

	for _, drop := range drops {
		if number%drop.factor == 0 {
			result += drop.sound
		}
	}

	if result == "" {
		return strconv.Itoa(number)
	}

	return result
	panic("Please implement the Convert function")
}
