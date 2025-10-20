package main

import (
	"fmt"

	parsekit "toolKit/parseKit"
)

func main() {
	lines, err := parsekit.ReadFileLines()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rest, err := parsekit.ParseAnts(lines)
	if err != nil {
		fmt.Println("Ant parsing error:", err)
		return
	}

	if err := parsekit.ParseRooms(rest); err != nil {
		fmt.Println("Room parsing error:", err)
		return
	}

	if err := parsekit.ParseStartEnd(rest); err != nil {
		fmt.Println("Start/End parsing error:", err)
		return
	}

	if err := parsekit.ParseTunnels(rest); err != nil {
		fmt.Println("Tunnel parsing error:", err)
		return
	}

	if err := parsekit.WriteParsedOutput("parsed_output.txt"); err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}

	// fmt.Println("✅ Results written to parsed_output.txt successfully!")
	fmt.Println("Parsing completed successfully!")
}
