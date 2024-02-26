package demo1

import (
	"reflect"
	"testing"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 11:47

func TestRange(t *testing.T) {
	type args struct {
		head *Link
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Range(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Range() = %v, want %v", got, tt.want)
			}
		})
	}
}
