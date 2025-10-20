package parsekit

import (
	"fmt"
	"strings"
)

func ParseStartEnd(lines []string) error {
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "##start" && i+1 < len(lines) {
			next := strings.TrimSpace(lines[i+1])
			name, _, _, err := parseRoomDataStrict(next)
			if err != nil {
				return fmt.Errorf("invalid room under ##start: %v", err)
			}
			StartRoom = name
		} else if line == "##end" && i+1 < len(lines) {
			next := strings.TrimSpace(lines[i+1])
			name, _, _, err := parseRoomDataStrict(next)
			if err != nil {
				return fmt.Errorf("invalid room under ##end: %v", err)
			}
			EndRoom = name
		}
	}

	if StartRoom == "" || EndRoom == "" {
		return fmt.Errorf("missing start or end room")
	}

	return nil
}

func parseRoomDataStrict(line string) (string, int, int, error) {
	var name string
	var x, y int
	n, err := fmt.Sscanf(line, "%s %d %d", &name, &x, &y)
	if err != nil || n != 3 {
		return "", 0, 0, fmt.Errorf("line is not a valid room: %q", line)
	}
	if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
		return "", 0, 0, fmt.Errorf("invalid room name: %s", name)
	}
	return name, x, y, nil
}
