// package main

// import "fmt"

// type Address struct {
// 	City, Province, Country string
// }

// func main() {
// 	address1 := Address{"Subang,", "Jawa Barat,", "Indonesia"}
// 	address2 := &address1
// 	address3 := &address1

// 	address2.City = "Bandung"

// 	*address2 = Address{"Malang,", "Jawa Timur,", "Indonesia"}

// 	fmt.Println(address1)
// 	fmt.Println(address2)
// 	fmt.Println(address3)

// 	address4 := new(Address)
// 	address4.City = "Bandung"
// 	address4.Province = "Jabar"
// 	address4.Country = "Indonesia"

// 	fmt.Println(address4)

// }
