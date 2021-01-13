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

func Selection(arr Sortable) {
	for i := 0; i < arr.Len()-1; i++ {
		minOneIndex := i
		for j := i + 1; j < arr.Len(); j++ {
			if arr.Less(j, minOneIndex) {
				minOneIndex = j
			}
		}
		if minOneIndex != i {
			arr.Swap(minOneIndex, i)
		}
	}
}

func Insertion(arr Sortable) {
	for i := 1; i < arr.Len(); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr.Less(j+1, j) {
				arr.Swap(j+1, j)
			} else {
				break
			}
		}
	}
}
