package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_deleteDuplicate(t *testing.T) {
	type args struct {
		numbers []int
	}
	x := args{numbers: []int{0, 0, 0, 0}}
	fmt.Println(x)
	type test_md struct {
		name string
		args args
		want []int
	}
	tests := []test_md{
		{
			name: "all same",
			args: args{
				numbers: []int{0, 0, 0},
			},
			want: []int{0},
		},
		{
			name: "all double",
			args: args{
				numbers: []int{0, 0, 1, 1, 2, 2, 3, 3},
			},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "all double",
			args: args{
				numbers: []int{0, 0, 1, 2, 1, 5, 6, 71, 9, 6, 6, 0, 1},
			},
			want: []int{0, 1, 2, 5, 6, 71, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicate(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
