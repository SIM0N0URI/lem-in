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
			name, _, _ := parseRoomData(next)
			StartRoom = name
		} else if line == "##end" && i+1 < len(lines) {
			next := strings.TrimSpace(lines[i+1])
			name, _, _ := parseRoomData(next)
			EndRoom = name
		}
	}
	if StartRoom == "" || EndRoom == "" {
		return fmt.Errorf("missing start or end room")
	}
	return nil
}

func parseRoomData(line string) (string, int, int) {
	var name string
	var x, y int
	fmt.Sscanf(line, "%s %d %d", &name, &x, &y)
	return name, x, y
}
