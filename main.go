package main

import "fmt"


func netherCoords(coord1 float64, coord3 float64) (float64, float64) {
	coord1 *= 8.0
	coord3 *= 8.0
	return coord1, coord3
}


func overworldCoords(coord1 float64, coord3 float64) (float64, float64) {
	coord1 /= 8.0
	coord3 /= 8.0
	return coord1, coord3
}


func main() {
	var location int8

	fmt.Print("Enter 1 to convert Nether Coords, Enter 2 to convert Overworld Coords: ")
	fmt.Scanln(&location)  // Cannot use Scanf as it will pull the 1 from the print string as the location.

	var x, y, z float64
	fmt.Print("Enter your x Value: ")
	fmt.Scanf("%f\n", &x)
	fmt.Print("Enter your y Value: ")
	fmt.Scanf("%f\n", &y)
	fmt.Print("Enter your z Value: ")
	fmt.Scanf("%f\n", &z)

	switch location {
	case 1:
		x, z = netherCoords(x, z)
		fmt.Println("Your converted coords are:", "x:", x, "y:", y, "z:", z)
		main()
	case 2:
		x, z = overworldCoords(x, z)
		fmt.Println("Your converted coords are:", "x:", x, "y:", y, "z:", z)
		main()
	default:
		fmt.Println("Please choose 1 or 2, no spaces or other characters allowed...")
		main()
	}
}