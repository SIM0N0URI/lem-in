package functions

import (
	"errors"
	"strconv"
)

func ValidateAntNumber(line string) (int, error) {
	num, err := strconv.Atoi(line)
	if err != nil || num <= 0 {
		return 0, errors.New("invalid ant number on first line")
	}
	return num, nil
}
