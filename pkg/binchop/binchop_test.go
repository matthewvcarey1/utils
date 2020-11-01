package binchop

import (
	"testing"
)

func TestBinChopStrings(t *testing.T) {
	type args struct {
		value string
		slice []string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
	}{
		{
			name: "Happy day",
			args: args{
				value: "jackal",
				slice: []string{"ant", "badger", "cat", "dog", "elephant", "fly", "giraffe", "hedgehog", "ibis", "jackal", "kangeroo", "lion", "mongoose", "narwhale", "otter", "porcupine"},
			},
			want:  true,
			want1: 9,
		},
		{
			name: "unhappy day",
			args: args{
				value: "alligator",
				slice: []string{"ant", "badger", "cat", "dog", "elephant", "fly", "giraffe", "hedgehog", "ibis", "jackal", "kangeroo", "lion", "mongoose", "narwhale", "otter", "porcupine"},
			},
			want:  false,
			want1: -1,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BinChopStrings(tt.args.value, tt.args.slice)
			if got != tt.want {
				t.Errorf("BinChopStrings() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BinChopStrings() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBinChopInts(t *testing.T) {
	type args struct {
		value int
		slice []int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
	}{
		{
			name: "Happy day",
			args: args{
				value: 3,
				slice: []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21},
			},
			want:  true,
			want1: 1,
		},
		{
			name: "unhappy day",
			args: args{
				value: 20,
				slice: []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21},
			},
			want:  false,
			want1: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BinChopInts(tt.args.value, tt.args.slice)
			if got != tt.want {
				t.Errorf("BinChopInts() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BinChopInts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
