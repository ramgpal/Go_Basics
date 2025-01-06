package main

import "fmt"

func main() {
	mapExample := make(map[string]string)

	mapExample["JS"] = "JavaScript"
	mapExample["RB"] = "Ruby"
	mapExample["Py"] = "Python"

	fmt.Println(mapExample)

	//access the value from map based on key
	fmt.Println("The value for JS is: ", mapExample["JS"])

	// removing key from map
	delete(mapExample, "RB")

	fmt.Println(mapExample)

	// iterating the map-> Map
	for key, value := range mapExample {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
