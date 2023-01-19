// package main

// import "fmt"

// func main() {
// 	fmt.Println("=======================================================")
// 	fmt.Println("======================== SLICE ========================")
// 	fmt.Println("=======================================================")

// 	var months = [...]string{
// 		"January",
// 		"February",
// 		"March",
// 		"April",
// 		"May",
// 		"June",
// 		"July",
// 		"August",
// 		"September",
// 		"October",
// 		"November",
// 		"December",
// 	}
// 	var slice1 = months[4:7]
// 	fmt.Println(slice1)
// 	fmt.Println(len(slice1))
// 	fmt.Println(cap(slice1))

// 	// months[5] = "Diubah"
// 	// fmt.Println(slice1)

// 	// slice1[0] = "More May"
// 	// fmt.Println(months)

// 	var slice2 = months[10:]
// 	fmt.Println(slice2)

// 	var slice3 = append(slice2, "Asep")
// 	fmt.Println(slice3)

// 	newSlice := make([]string, 2, 5)
// 	newSlice[0] = "Asep"
// 	newSlice[1] = "Asep2"

// 	fmt.Println(newSlice)
// 	fmt.Println(len(newSlice))
// 	fmt.Println(cap(newSlice))

// 	copySlice := make([]string, len(newSlice), cap(newSlice))
// 	copy(copySlice, newSlice)
// 	fmt.Println(copySlice)
// }
