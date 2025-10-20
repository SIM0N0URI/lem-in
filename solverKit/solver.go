package solverkit

import (
	"fmt"
	parsekit "toolKit/parseKit"
)

func Solve() Paths {
	if parsekit.StartRoom == "" || parsekit.EndRoom == "" {
		fmt.Println("Error: StartRoom or EndRoom not defined")
		return nil
	}

	rawPaths := BFS()

	allPaths := make(Paths, len(rawPaths))
	for i, p := range rawPaths {
		allPaths[i] = p
	}

	shortestPaths := allPaths.FindShortest()

	fmt.Println("Found paths:")
	for _, p := range shortestPaths {
		fmt.Println(p)
	}

	return shortestPaths
}
