package main

import "fmt"

func main() {
	fmt.Println("Maps pairs in go")

	idToCategoryMap := make(map[int]string)
	idToCategoryMap[1] = "Sports"
	idToCategoryMap[2] = "News"
	idToCategoryMap[3] = "Social Media"
	idToCategoryMap[4] = "Politics"
	idToCategoryMap[5] = "Technology"

	fmt.Printf("The type of idToCategoryMap is %T\n", idToCategoryMap)
	fmt.Println("The id to category map is:", idToCategoryMap)

	fmt.Println("Category '3' maps to:", idToCategoryMap[3])

	delete(idToCategoryMap, 3)
	fmt.Println("The id to category map after del is:", idToCategoryMap)

	fmt.Println("")
	// Looping over a map
	for key, value := range idToCategoryMap {
		fmt.Printf("For key:%v, the value is %v\n", key, value)
	}
}
