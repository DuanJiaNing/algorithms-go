package sort

import (
	"reflect"
	"testing"
)

type sortFunc func(Sortable)

// Insertion sort.
func TestInsertion(t *testing.T) {
	testSort(t, Insertion)
}

func BenchmarkInsertion(b *testing.B) {
	// BenchmarkInsertion-8   	   65006	     17960 ns/op
	benchmarkSort(b, Insertion)
}

// Selection sort.
func TestSelection(t *testing.T) {
	testSort(t, Selection)
}

func BenchmarkSelection(b *testing.B) {
	// BenchmarkSelection-8   	   41934	     28265 ns/op
	benchmarkSort(b, Selection)
}

// Bubble sort.
func TestBubble(t *testing.T) {
	testSort(t, Bubble)
}

func BenchmarkBubble(b *testing.B) {
	// BenchmarkBubble-8   	   39454	     29224 ns/op
	benchmarkSort(b, Bubble)
}

func benchmarkSort(b *testing.B, sort sortFunc) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort(NewIntArr(81, 87, 47, 59, 81, 18, 25, 40, 56, 0, 94, 11, 62, 89, 28, 74, 11, 45, 37, 6, 95, 66, 28, 58, 47, 47, 87, 88, 90, 15, 41, 8, 87, 31, 29, 56, 37, 31, 85, 26, 13, 90, 94, 63, 33, 47, 78, 24, 59, 53, 57, 21, 89, 99, 0, 5, 88, 38, 3, 55, 51, 10, 5, 56, 66, 28, 61, 2, 83, 46, 63, 76, 2, 18, 47, 94, 77, 63, 96, 20, 23, 53, 37, 33, 41, 59, 33, 43, 91, 2, 78, 36, 46, 7, 40, 3, 52, 43, 5, 98))
	}
}

func testSort(t *testing.T, sort sortFunc) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "0",
			args: args{arr: []int{}},
			want: []int{},
		},
		{
			name: "1",
			args: args{arr: []int{3}},
			want: []int{3},
		},
		{
			name: "2",
			args: args{arr: []int{13, 13}},
			want: []int{13, 13},
		},
		{
			name: "3",
			args: args{arr: []int{3, 13}},
			want: []int{3, 13},
		},
		{
			name: "4",
			args: args{arr: []int{13, 3}},
			want: []int{3, 13},
		},
		{
			name: "5",
			args: args{arr: []int{23, 4, 13}},
			want: []int{4, 13, 23},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewIntArr(tt.args.arr...)
			sort(arr)
			if !reflect.DeepEqual(arr.arr, tt.want) {
				t.Fatalf("got: %v, want: %v", arr.arr, tt.want)
			}
		})
	}
}
