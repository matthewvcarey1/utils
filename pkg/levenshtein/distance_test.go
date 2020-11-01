package levenshtein

import (
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "apple1",
			args: args{
				s: "apple",
				t: "crabapple",
			},
			want: 4,
		},
		{
			name: "apple2",
			args: args{
				s: "apple",
				t: "apples",
			},
			want: 1,
		},
		{
			name: "blocks1",
			args: args{
				s: "blocks",
				t: "boxes",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevenshteinDistance(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("LevenshteinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimum(t *testing.T) {
	type args struct {
		a int
		v []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				a: 1,
				v: []int{2, 3},
			},
			want: 1,
		},
		{
			name: "b",
			args: args{
				a: 3,
				v: []int{2, 1},
			},
			want: 1,
		},
		{
			name: "c",
			args: args{
				a: 2,
				v: []int{1, 3},
			},
			want: 1,
		},
		{
			name: "d",
			args: args{
				a: 1,
				v: []int{},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimum(tt.args.a,tt.args.v...); got != tt.want {
				t.Errorf("minimum() = %v, want %v", got, tt.want)
			}
		})
	}
}
