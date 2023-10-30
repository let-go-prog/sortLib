package main

import (
	"fmt"
	"github.com/let-go-prog/sortLib/pkg/customSort"
)

func main() {
	nums := []int{1, 3, 2, 5, 10}
	fmt.Println(customSort.BubbleSort(nums))
}
