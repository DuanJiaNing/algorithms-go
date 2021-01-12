package sort

func Bubble(arr Sortable) {
	for i := 0; i < arr.Len()-1; i++ {
		for j := arr.Len() - 1; j > i; j-- {
			if arr.Less(j, j-1) {
				arr.Swap(j, j-1)
			}
		}
	}
}
