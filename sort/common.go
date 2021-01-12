package sort

import (
	"fmt"
	"strings"
)

type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type IntArr struct {
	arr []int
}

func NewIntArr(arr ...int) *IntArr {
	return &IntArr{arr}
}

func (ia *IntArr) String() string {
	ss := make([]string, 0, len(ia.arr))
	for _, a := range ia.arr {
		ss = append(ss, fmt.Sprintf("%d", a))
	}

	return fmt.Sprintf("[%s]", strings.Join(ss, ","))
}

func (ia *IntArr) Swap(i, j int) {
	old := ia.arr[i]
	ia.arr[i] = ia.arr[j]
	ia.arr[j] = old
}

func (ia *IntArr) Less(i, j int) bool {
	return ia.arr[i] < ia.arr[j]
}

func (ia *IntArr) Len() int {
	return len(ia.arr)
}
