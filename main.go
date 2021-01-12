package main

import (
	"algorithms-go/sort"
	"fmt"
)

func main() {
	iarr := sort.NewIntArr(2, 3, 6, 1, 0, 12)
	sort.Bubble(iarr)
	fmt.Println(iarr)
}
