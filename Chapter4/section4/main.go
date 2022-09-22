package main

import "fmt"
import "gopl/Chapter4/section4/sort"

func main() {
	var values = []int{1, 5, 3, 8, 6, 4, 9}
	values = sort.Sort(values)
	fmt.Println(values)
}
