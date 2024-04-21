package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Test_deleteDuplicate_2(t *testing.T) {
	type args struct {
		numbers string
	}
	x := args{numbers: "0, 0, 0, 0"}
	fmt.Println(x)
	type test_md struct {
		name string
		args args
		want string
	}
	tests := []test_md{
		{
			name: "all same",
			args: args{
				numbers: "0, 0, 0",
			},
			want: "0",
		},
		{
			name: "all double",
			args: args{
				numbers: "0, 0, 1, 1, 2, 2, 3, 3",
			},
			want: "0, 1, 2, 3",
		},
		{
			name: "all double",
			args: args{
				numbers: "0, 0, 1, 2, 1, 5, 6, 71, 9, 6, 6, 0, 1",
			},
			want: "0, 1, 2, 5, 6, 71, 9",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicate(strings.Split(tt.args.numbers, ", ")); !reflect.DeepEqual(got, strings.Split(tt.want, ", ")) {
				t.Errorf("deleteDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
