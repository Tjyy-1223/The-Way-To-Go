package main

import "fmt"

func mainMakeMaps() {
	var mapList map[string]int
	var mapAssigned map[string]int

	mapList = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapList

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapList["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapAssigned["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapList["ten"])

}
