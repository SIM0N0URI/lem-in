package main

import (
	"fmt"

	"toolKit/functions"
)

func main() {
	lines, err := functions.ParseFile("map.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	antCount, err := functions.ValidateAntNumber(lines[0])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	startRoom, endRoom, rooms, tunnels, err := functions.ParseData(lines[1:])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Ants:", antCount)
	fmt.Println("Start Room:", startRoom)
	fmt.Println("End Room:", endRoom)
	fmt.Println("Rooms:", rooms)
	fmt.Println("Tunnels:", tunnels)
}
