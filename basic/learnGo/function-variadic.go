// package main

// import "fmt"

// func sumAll(numbers ...int) int {
// 	fmt.Println("=========================================================")
// 	fmt.Println("==================== FUNCTION VARIADIC ==================")
// 	fmt.Println("=========================================================")

// 	total := 0
// 	for _, value := range numbers {
// 		total += value
// 	}
// 	return total

// }

// func main() {
// 	total := sumAll(10, 20, 20)
// 	fmt.Println(total)

// 	slice := []int{10, 10, 10}
// 	total = sumAll(slice...)
// 	fmt.Println(total)
// }
