package functions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"toolKit/units"
)

func ParseRoom(line string) (units.Room, error) {
	if strings.HasPrefix(line, "L") || strings.HasPrefix(line, "#") {
		return units.Room{}, errors.New("room name cannot start with L or #")
	}

	parts := strings.Fields(line)
	if len(parts) != 3 {
		return units.Room{}, errors.New("invalid room format")
	}

	name := parts[0]
	x, err1 := strconv.Atoi(parts[1])
	y, err2 := strconv.Atoi(parts[2])
	if err1 != nil || err2 != nil {
		return units.Room{}, errors.New("invalid coordinates")
	}

	return units.Room{Name: name, X: x, Y: y}, nil
}

func ValidateRooms(rooms []units.Room) error {
	for i := 0; i < len(rooms); i++ {
		for j := i + 1; j < len(rooms); j++ {
			if rooms[i].X == rooms[j].X && rooms[i].Y == rooms[j].Y {
				return fmt.Errorf("duplicate coordinates between %s and %s", rooms[i].Name, rooms[j].Name)
			}
		}
	}
	return nil
}
