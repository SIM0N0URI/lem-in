package solverkit

type Path []string

type Paths []Path

func (p Paths) FindShortest() Paths {
	minLen := -1
	var shortest Paths

	for _, path := range p {
		if minLen == -1 || len(path) < minLen {
			minLen = len(path)
			shortest = Paths{path}
		} else if len(path) == minLen {
			shortest = append(shortest, path)
		}
	}

	return shortest
}
