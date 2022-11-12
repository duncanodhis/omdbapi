package readFile

import (
	"reflect"
	"testing"
)

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
		name string
		args args
		want []string
	}{
		{name: "1", args: args{lines: []string{"tt0000005\tshort\tBlacksmith Scene\tBlacksmith Scene\t0\t1893\t\\N\t1\tComedy,Short"},
			title_Type:      "short",
			primary_Title:   "Blacksmith Scene",
			original_Title:  "Blacksmith Scene",
			start_Year:      "1893",
			end_Year:        "\\N",
			runtime_Minutes: "1",
			genres_:         "Short"},
			want: []string{"tt0000005"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTconst(tt.args.lines, tt.args.title_Type, tt.args.primary_Title, tt.args.original_Title, tt.args.start_Year, tt.args.end_Year, tt.args.runtime_Minutes, tt.args.genres_); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findTconst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLine(t *testing.T) {
	type args struct {
		filename string
		line     chan string
		readerr  chan error
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getLine(tt.args.filename, tt.args.line, tt.args.readerr)
		})
	}
}
