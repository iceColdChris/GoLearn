package main

import (
	"fmt"
)

func main() {
	// Initial creation
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	//creates an empty map
	//var colors map[string]string
	//colors := make(map[string]string)

	//Adds a value to the map
	colors["white"] = "#ffffff"

	//deletes from the map
	//delete(colors, "white")

	fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println("Hex code for", k, "is", v)
	}
}
