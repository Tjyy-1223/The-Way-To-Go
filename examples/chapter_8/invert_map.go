package main

import "fmt"

var (
	barVal2 = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func mainInvertMap() {
	invMap := make(map[int]string, len(barVal2))
	for k, v := range barVal2 {
		invMap[v] = k
	}

	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
}
