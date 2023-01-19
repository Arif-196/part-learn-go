package main

import (
	"fmt"
	dsa "go-dsa/dsa"
	search "go-dsa/dsa/a_search"
	sorts "go-dsa/dsa/b_sort"
	"os"
	"sort"
)

func main() {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	fmt.Println("========================================= LINEAR SEARCH =========================================\n")
	fmt.Println("LINEAR SEARCH :", search.Linearsearch(items, 58), "\n") // True

	fmt.Println("========================================= BINARY SEARCH =========================================\n")
	fmt.Println("BINARY SEARCH :", search.BinarySearch(78, items), "\n") // False

	fmt.Println("====================================== INTERPOLATION SEARCH =====================================\n")
	fmt.Println("INTERPOLATION SEARCH :", search.InterpolationSearch(items, 251), "\n") // FINDING INDEX AT 251 IS 7

	fmt.Println("========================================== BUBBLE SORT ==========================================\n") // SORTING SOME NUMBER RANDOM
	slice := sorts.GeneratSlice(20)
	fmt.Println("UNSORTED :", slice)
	sorts.BubbleSort(slice)
	fmt.Println("SORTED   :", slice, "\n")

	fmt.Println("========================================== QUICK SORT ==========================================\n") // SORTING SOME NUMBER RANDOM
	slice2 := sorts.GeneratSlice(20)
	fmt.Println("UNSORTED :", slice2)
	sorts.QuickSort(slice2)
	fmt.Println("SORTED   :", slice2, "\n")

	fmt.Println("========================================== SELECTION SORT ==========================================\n") // SORTING SOME NUMBER RANDOM
	slice3 := sorts.GeneratSlice(20)
	fmt.Println("UNSORTED :", slice3)
	sorts.SelectionSort(slice3)
	fmt.Println("SORTED   :", slice3, "\n")

	fmt.Println("========================================== SHELL SORT ==========================================\n") // SORTING SOME NUMBER RANDOM
	slice4 := sorts.GeneratSlice(20)
	fmt.Println("UNSORTED :", slice4)
	sorts.ShellSort(slice4)
	fmt.Println("SORTED   :", slice4, "\n")

	fmt.Println("========================================== INSERTION SORT ==========================================\n") // SORTING SOME NUMBER RANDOM
	slice5 := sorts.GeneratSlice(20)
	fmt.Println("UNSORTED :", slice5)
	sorts.InsertionSort(slice5)
	fmt.Println("SORTED   :", slice5, "\n")

	fmt.Println("=========================================== COMB SORT ===========================================\n") // SORTING SOME NUMBER RANDOM
	slice6 := sorts.GeneratSlice(20)
	fmt.Println("UNSORTED :", slice6)
	sorts.CombSort(slice6)
	fmt.Println("SORTED   :", slice6, "\n")

	fmt.Println("=========================================== MERGE SORT ===========================================\n") // SORTING SOME NUMBER RANDOM
	slice7 := sorts.GeneratSlice(20)
	fmt.Println("UNSORTED :", slice7)
	// sorts.MergeSort(slice7)
	fmt.Println("SORTED   :", sorts.MergeSort(slice7), "\n")

	fmt.Println("=========================================== RADIX SORT ===========================================\n") // SORTING SOME NUMBER RANDOM
	var data = []int32{421, 15, -175, 90, -2, 214, -52, -166}
	fmt.Println("UNSORTED :", data)
	sorts.RadixSort(data)
	fmt.Println("SORTED   :", data, "\n")

	fmt.Println("=========================================== PANCAKE SORT ===========================================\n") // SORTING SOME NUMBER RANDOM
	var list = sorts.Data{28, 11, 59, -26, 503, 158, 997, 193, -23, 44}
	fmt.Println("UNSORTED :", list)
	list.Pancakesort()
	fmt.Println("SORTED   :", list, "\n")

	fmt.Println("=========================================== BINARY TREE ===========================================\n") // SORTING SOME NUMBER RANDOM
	var tree *dsa.BinaryTree = &dsa.BinaryTree{}
	tree.Insert(100).
		Insert(-20).
		Insert(-50).
		Insert(-15).
		Insert(-60).
		Insert(50).
		Insert(60).
		Insert(55).
		Insert(85).
		Insert(15).
		Insert(5).
		Insert(-10)
	dsa.Print(os.Stdout, tree.Root, 0, 'M')

	fmt.Println("=========================================== LINKED LIST ===========================================\n") // SORTING SOME NUMBER RANDOM
	var link *dsa.List = &dsa.List{}
	link.Insert(5)
	link.Insert(9)
	link.Insert(13)
	link.Insert(22)
	link.Insert(28)
	link.Insert(36)

	fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.Head.Key)
	fmt.Printf("Tail: %v\n", link.Tail.Key)
	link.Display()
	fmt.Println("\n==============================\n")
	fmt.Printf("head: %v\n", link.Head.Key)
	fmt.Printf("tail: %v\n", link.Tail.Key)
	link.Reverse()
	fmt.Println("\n==============================\n")

	fmt.Println("=========================================== RABIN KARP ===========================================\n") // SORTING SOME NUMBER RANDOM
	txt := "Australia is a country and continent surrounded by the Indian and Pacific oceans."
	patterns := []string{"and", "the", "surround", "Pacific", "Germany"}
	matches := dsa.Search(txt, patterns)
	fmt.Println(matches)

	fmt.Println("=========================================== RABIN KARP ===========================================\n") // SORTING SOME NUMBER RANDOM
	intSlice := []int{5, 9, 77, 62, 71, 11, 22, 46, 36, 18, 19, 33, 75, 17, 39, 41, 73, 50, 217, 79, 120}
	sort.Ints(intSlice)
	for _, j := range []int{5, 10, 15, 20} {
		i := dsa.MedianOfMedians(intSlice, j, 5)
		fmt.Println(j, "Smallest Element = ", i)
		v := intSlice[j-1]
		fmt.Println("arr[", j-1, "] = ", v)
		if i != v {
			fmt.Println("Oops! Algorithm is wrong")
		}
	}

}
