package parsekit

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseAnts(lines []string) error {
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if IsEmpty(line) || IsComment(line) {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("invalid ant number: %v", err)
		}
		if num <= 0 {
			return fmt.Errorf("ant number must be > 0")
		}
		AntNum = num
		return nil
	}
	return fmt.Errorf("no valid ant number found")
}
