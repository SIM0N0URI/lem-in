package solverkit

import (
	"container/list"
	parsekit "toolKit/parseKit"
)

func BFS() [][]string {
	if parsekit.StartRoom == "" || parsekit.EndRoom == "" {
		return nil
	}

	graph := make(map[string][]string)
	for _, t := range parsekit.Tunnels {
		graph[t.From] = append(graph[t.From], t.To)
		graph[t.To] = append(graph[t.To], t.From)
	}

	var paths [][]string
	queue := list.New()
	queue.PushBack([]string{parsekit.StartRoom})

	for queue.Len() > 0 {
		elem := queue.Front()
		queue.Remove(elem)
		path := elem.Value.([]string)
		last := path[len(path)-1]

		if last == parsekit.EndRoom {
			paths = append(paths, path)
			continue
		}

		for _, neighbor := range graph[last] {
			if !contains(path, neighbor) {
				newPath := append(path, neighbor)
				queue.PushBack(newPath)
			}
		}
	}

	return paths
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
