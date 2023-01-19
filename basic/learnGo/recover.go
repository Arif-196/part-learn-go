// package main

// import "fmt"

// func endApp() {
// 	message := recover()
// 	if message != nil {
// 		fmt.Println("Terjadi Error", message)
// 	}
// 	fmt.Println("Aplikasi Selesai ")

// }

// func runApp(error bool) {
// 	defer endApp()
// 	if error {
// 		panic("ERROR")
// 	}
// 	fmt.Println("Aplikasi Berjalan")
// }
// func main() {
// 	runApp(true)
// 	fmt.Println("ASEP")

// }
