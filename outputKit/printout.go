package outputkit

import (
	"fmt"
	parsekit "toolKit/parseKit"
	solverkit "toolKit/solverKit"
)

func PrintAntMovements(paths solverkit.Paths, antCount int) {
	type Ant struct {
		ID      int
		PathIdx int
		Pos     int
		Done    bool
	}

	ants := make([]*Ant, antCount)
	for i := 0; i < antCount; i++ {
		ants[i] = &Ant{
			ID:      i + 1,
			PathIdx: i % len(paths),
			Pos:     0,
			Done:    false,
		}
	}

	occupied := make(map[string]int)
	occupied[parsekit.StartRoom] = -1
	occupied[parsekit.EndRoom] = -1

	moving := true
	for moving {
		moving = false
		line := ""

		for _, ant := range ants {
			if ant.Done {
				continue
			}

			path := paths[ant.PathIdx]
			var nextRoom string
			if ant.Pos < len(path) {
				nextRoom = path[ant.Pos]
			} else {
				nextRoom = parsekit.EndRoom
			}

			if nextRoom == parsekit.EndRoom || occupied[nextRoom] == 0 {
				if ant.Pos > 0 {
					currRoom := path[ant.Pos-1]
					if currRoom != parsekit.StartRoom {
						occupied[currRoom] = 0
					}
				}

				if nextRoom != parsekit.EndRoom {
					occupied[nextRoom] = ant.ID
				}

				line += fmt.Sprintf("L%d-%s ", ant.ID, nextRoom)
				ant.Pos++
				moving = true

				if nextRoom == parsekit.EndRoom {
					ant.Done = true
				}
			}
		}

		if line != "" {
			fmt.Println(line)
		}
	}
}
