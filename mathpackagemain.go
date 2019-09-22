package main

import (
	"fmt"

	m "./goPackages/mathPackage"
	//  "github.com/damilarelana/introToGo/gomathpackage"
)

func main() {
	fmt.Println("Enter the array elements: ")
	initialArray := make([]float64, 5)

	// Reads input variables (i.e. elements of the array) as numbers separated by spaces
	// initialArray := []float64{48, 96, 86, 68, 57}
	fmt.Scanf("%f %f %f %f %f", &initialArray[0], &initialArray[1], &initialArray[2], &initialArray[3], &initialArray[4])

	// Print the inputted array
	fmt.Println(" ")
	fmt.Println("==== ==== ====")
	fmt.Println(" ")
	fmt.Println("Given array: ", initialArray)

	// Call the Average function to get average value of the array elements
	fmt.Println(" ")
	fmt.Println("==== ==== ====")
	fmt.Println(" ")
	fmt.Println("average is: ", m.Average(initialArray))

	// Call the Max function to get largest array element
	fmt.Println(" ")
	fmt.Println("==== ==== ====")
	fmt.Println(" ")
	fmt.Println("largest element is: ", m.Max(initialArray))

	// Call the Min function to get smallest array element
	fmt.Println(" ")
	fmt.Println("==== ==== ====")
	fmt.Println(" ")
	fmt.Println("smallest element is: ", m.Min(initialArray))
}
