// package main

// import "fmt"

// type Filter func(string) string

// func wordsFilter(name string, filter Filter) {
// 	fmt.Println("=========================================================")
// 	fmt.Println("==================== FUNCTION VALUES ==================")
// 	fmt.Println("=========================================================")

// 	wordsFiltered := filter(name)
// 	fmt.Println("Hello ", wordsFiltered)
// }

// func reqWords(name string) string {
// 	if name == "Anjing" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }
// func main() {
// 	wordsFilter("Eko", reqWords)
// 	wordsFilter("Anjing", reqWords)
// }
