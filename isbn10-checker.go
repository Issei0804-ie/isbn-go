package isbngo

import (
	"fmt"
	"strconv"
)

func IsValidIsbn10(maybeIsbn10 string) (bool, error) {
	if len(maybeIsbn10) != 10 {
		return false, nil
	}

	nums := make([]int, len(maybeIsbn10))

	for i, char := range maybeIsbn10 {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return false, fmt.Errorf("invalid args")
		}
		nums[i] = num
	}

	sum := 0

	for i := 0; i < 9; i++ {
		sum += nums[i] * (i + 1)
	}

	checkDigit := sum % 11

	return nums[9] == checkDigit, nil
}
