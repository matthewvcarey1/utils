package qsort

import (
	"reflect"
	"testing"
)

func TestQuickSortStrings(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Happy Day",
			args: args{
				slice: []string{"hedgehog", "ibis", "jackal", "kangeroo", "cat", "mongoose", "dog", "elephant", "fly", "giraffe", "lion", "narwhale", "otter", "porcupine", "ant", "badger"},
			},
			want: []string{"ant", "badger", "cat", "dog", "elephant", "fly", "giraffe", "hedgehog", "ibis", "jackal", "kangeroo", "lion", "mongoose", "narwhale", "otter", "porcupine"},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSortStrings(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSortStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSortInts(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Happy Day",
			args: args{
				slice: []int{6,3,2,7,1,0,9,8,5,4},
			},
			want: []int{0,1,2,3,4,5,6,7,8,9},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSortInts(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSortInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
