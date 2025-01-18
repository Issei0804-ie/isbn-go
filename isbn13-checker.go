package isbngo

import (
	"fmt"
	"math"
	"strconv"
)

func IsValidIsbn13(maybeIsbn13 string) (bool, error) {
	if len(maybeIsbn13) != 13 {
		return false, nil
	}

	nums := make([]int, len(maybeIsbn13))

	for i, char := range maybeIsbn13 {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return false, fmt.Errorf("invalid args")
		}
		nums[i] = num
	}

	sum := 0

	for i := 0; i < 12; i++ {
		if (i % 2) == 0 {
			sum += nums[i] * 1
			continue
		} else {
			sum += nums[i] * 3
			continue
		}
	}

	checkDigit := int(math.Abs(float64(((sum % 10) - 10) % 10)))

	return nums[12] == checkDigit, nil
}
