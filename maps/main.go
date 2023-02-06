package main

import "fmt"

func main() {

	//map : key val pair
	//all the keys must of same type and all the values must be of same type

	//By default Golang prints the map with sorted keys but while iterating over a map, it follows the order of the keys appearing as it is. 


	//Method 1:
	// var colors map[string]string -> declaring map with var keyword initializes it with zero val

	//Method 2:
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#67728",
		"app":"wee",

	}
	fmt.Println(colors)
}
