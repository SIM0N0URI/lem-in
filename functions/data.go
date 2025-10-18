package functions

import (
	"errors"
	"strings"

	"toolKit/units"
)

func ParseData(lines []string) (units.Room, units.Room, []units.Room, []units.Tunnel, error) {
	var rooms []units.Room
	var tunnels []units.Tunnel
	var startRoom, endRoom units.Room
	var foundStart, foundEnd bool

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if strings.HasPrefix(line, "#") {
			if line == "##start" && i+1 < len(lines) {
				room, err := ParseRoom(lines[i+1])
				if err != nil {
					return units.Room{}, units.Room{}, nil, nil, err
				}
				startRoom = room
				foundStart = true
				i++
			} else if line == "##end" && i+1 < len(lines) {
				room, err := ParseRoom(lines[i+1])
				if err != nil {
					return units.Room{}, units.Room{}, nil, nil, err
				}
				endRoom = room
				foundEnd = true
				i++
			}
			continue
		}

		if strings.Contains(line, "-") {
			tunnel, err := ParseTunnel(line)
			if err != nil {
				return units.Room{}, units.Room{}, nil, nil, err
			}
			tunnels = append(tunnels, tunnel)
			continue
		}

		room, err := ParseRoom(line)
		if err != nil {
			return units.Room{}, units.Room{}, nil, nil, err
		}
		rooms = append(rooms, room)
	}

	if !foundStart || !foundEnd {
		return units.Room{}, units.Room{}, nil, nil, errors.New("missing start or end room")
	}

	if err := ValidateRooms(rooms); err != nil {
		return units.Room{}, units.Room{}, nil, nil, err
	}

	return startRoom, endRoom, rooms, tunnels, nil
}
