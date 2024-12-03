package main

import "fmt"

func mainSliceForrange2() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}

	for ix := range seasons {
		seasons[ix] = seasons[ix] + "-2"
	}

	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}
}
