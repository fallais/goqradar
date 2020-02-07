package goqradar

import (
	"fmt"
	"strconv"
	"strings"
)

func parseContentRange(cr string) (int, int, int, error) {
	// Trim
	trimed := strings.Trim(cr, "items ")

	// Split min-max and total
	split := strings.Split(trimed, "/")
	if len(split) != 2 {
		return 0, 0, 0, fmt.Errorf("error with content-range")
	}

	// Split min and max
	minAndMax := strings.Split(split[0], "-")
	if len(minAndMax) != 2 {
		return 0, 0, 0, fmt.Errorf("error with content-range")
	}

	// Convert min
	min, err := strconv.Atoi(minAndMax[0])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error while converting the min into int")
	}

	// Convert max
	max, err := strconv.Atoi(minAndMax[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error while converting the max into int")
	}

	// Convert total
	total, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error while converting the total into int")
	}

	return min, max, total, nil
}
