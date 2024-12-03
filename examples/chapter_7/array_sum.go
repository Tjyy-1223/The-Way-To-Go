package main

import "fmt"

func mainArraySum() {
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator

	fmt.Printf("The sum of the array is : %f\n", x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a { // derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}
