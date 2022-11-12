package util

import (
	"reflect"
	"testing"
	"time"
)

func TestGracefull(t *testing.T) {
	type args struct {
		input      []string
		maxRuntime time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Gracefull(tt.args.input, tt.args.maxRuntime)
		})
	}
}

func TestRemoveDuplicateStr(t *testing.T) {
	type args struct {
		strSlice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{name: "1", args: args{strSlice: []string{"tt0000005", "tt0000005", "tt0000005"}},
			want: []string{"tt0000005"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicateStr(tt.args.strSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicateStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
