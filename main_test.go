package main

import (
	"reflect"
	"testing"
)

func Test_omdapiquery(t *testing.T) {
	type args struct {
		tconst     string
		plotFilter string
	}
	var tests = []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{"tt0000005", "Three men hammer"},
			want: []string{"tt0000005", "Blacksmith Scene", "Three men hammer on an anvil and pass a bottle of beer around."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := omdapiquery(tt.args.tconst, tt.args.plotFilter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("omdapiquery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicateStr(t *testing.T) {
	type args struct {
		strSlice []string
	}
	tests := []struct {
		name     string
		strSlice []string
		want     []string
	}{
		// TODO: Add test cases.
		{name: "1", strSlice: []string{"tt0000005", "tt0000005", "tt0000005"},
			want: []string{"tt0000005"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateStr(tt.strSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicateStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findTconst(t *testing.T) {
	type args struct {
		lines           []string
		title_Type      string
		primary_Title   string
		original_Title  string
		start_Year      string
		end_Year        string
		runtime_Minutes string
		genres_         string
	}
	tests := []struct {
		name            string
		lines           []string
		title_Type      string
		primary_Title   string
		original_Title  string
		start_Year      string
		end_Year        string
		runtime_Minutes string
		genres_         string
		want            []string
	}{
		{name: "1", lines: []string{"tt0000005\tshort\tBlacksmith Scene\tBlacksmith Scene\t0\t1893\t\\N\t1\tComedy,Short"},
			title_Type:      "short",
			primary_Title:   "Blacksmith Scene",
			original_Title:  "Blacksmith Scene",
			start_Year:      "1893",
			end_Year:        "\\N",
			runtime_Minutes: "1",
			genres_:         "Short",
			want:            []string{"tt0000005"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTconst(tt.lines, tt.title_Type, tt.primary_Title, tt.original_Title, tt.start_Year, tt.end_Year, tt.runtime_Minutes, tt.genres_); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findTconst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_getLine1(t *testing.T) {
	type args struct {
		filename string
		line     chan string
		readerr  chan error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getLine(tt.args.filename, tt.args.line, tt.args.readerr)
		})
	}
}
