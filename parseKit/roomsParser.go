package parsekit

import (
	"fmt"
	"strings"
)

func ParseRooms(lines []string) error {
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") || strings.Contains(line, "-") {
			continue
		}

		var name string
		var x, y int
		n, err := fmt.Sscanf(line, "%s %d %d", &name, &x, &y)

		if err != nil || n != 3 {
			return fmt.Errorf("invalid line found: %q", line)
		}

		if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
			return fmt.Errorf("invalid room name: %s", name)
		}

		for _, r := range Rooms {
			if r.X == x && r.Y == y {
				return fmt.Errorf("duplicate coordinates for room %s: (%d,%d)", name, x, y)
			}
		}

		Rooms[name] = &Room{Name: name, X: x, Y: y}
	}
	return nil
}
