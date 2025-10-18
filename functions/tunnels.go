package functions

import (
	"errors"
	"strings"

	"toolKit/units"
)

func ParseTunnel(line string) (units.Tunnel, error) {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return units.Tunnel{}, errors.New("invalid tunnel format")
	}
	return units.Tunnel{From: parts[0], To: parts[1]}, nil
}
