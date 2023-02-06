package main

import "fmt"

func main() {

	//map : key val pair
	//all the keys must of same type and all the values must be of same type

	//By default Golang prints the map with sorted keys but while iterating over a map, it follows the order of the keys appearing as it is.

	//Method 1:
	// var colors map[string]string -> declaring map with var keyword initializes it with zero val

	//Method 2:
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// delete(colors, "white")

	//Method 3:
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#67728",
	}
	fmt.Println(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, "is", hex)
	}
}

//********Diff b/w map and struct
/*
MAP:
1. All keys must be of same type and all values must be of same type
2. Used to represent a collection of related properties
3. No need to know all the keys at compile time
4. We can iterate over the keys
5. Reference type

Struct: (used more than maps)
1. Values can be of different type
2. Need to know all fiels during compile time
3. Key's don't support indexing
4. Used to represent a thing with diff properties
5. Value Type

*/
