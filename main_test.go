package main

import (
	"reflect"
	"testing"
	"time"
)

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

func Test_requestLimiter(t *testing.T) {
	type args struct {
		url        string
		maxRequest int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{url: "http://www.omdbapi.com/?apikey=5226c193&i=tt0000005", maxRequest: 100},
			want: "{\"Title\":\"Blacksmith Scene\",\"Year\":\"1893\",\"Rated\":\"Unrated\",\"Released\":\"09 May 1893\",\"Runtime\":\"1 min\",\"Genre\":\"Short, Comedy\",\"Director\":\"William K.L. Dickson\",\"Writer\":\"N/A\",\"Actors\":\"Charles Kayser, John Ott\",\"Plot\":\"Three men hammer on an anvil and pass a bottle of beer around.\",\"Language\":\"None\",\"Country\":\"United States\",\"Awards\":\"1 win\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BNDg0ZDg0YWYtYzMwYi00ZjVlLWI5YzUtNzBkNjlhZWM5ODk5XkEyXkFqcGdeQXVyNDk0MDg4NDk@._V1_SX300.jpg\",\"Ratings\":[{\"Source\":\"Internet Movie Database\",\"Value\":\"6.2/10\"}],\"Metascore\":\"N/A\",\"imdbRating\":\"6.2\",\"imdbVotes\":\"2,541\",\"imdbID\":\"tt0000005\",\"Type\":\"movie\",\"DVD\":\"N/A\",\"BoxOffice\":\"N/A\",\"Production\":\"N/A\",\"Website\":\"N/A\",\"Response\":\"True\"}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := requestLimiter(tt.args.url, tt.args.maxRequest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestLimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_omdapiquery(t *testing.T) {
	type args struct {
		tconst     string
		plotFilter string
		maxRequest int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1", args: args{
			tconst:     "tt0000005",
			plotFilter: "Three men",
			maxRequest: 10,
		}, want: []string{"tt0000005", "Blacksmith Scene", "Three men hammer on an anvil and pass a bottle of beer around."}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := omdapiquery(tt.args.tconst, tt.args.plotFilter, tt.args.maxRequest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("omdapiquery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_responseData(t *testing.T) {
	type args struct {
		tconst     string
		data       string
		plotFilter string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{tconst: "tt0000005", data: "{\"Title\":\"Blacksmith Scene\",\"Year\":\"1893\",\"Rated\":\"Unrated\",\"Released\":\"09 May 1893\",\"Runtime\":\"1 min\",\"Genre\":\"Short, Comedy\",\"Director\":\"William K.L. Dickson\",\"Writer\":\"N/A\",\"Actors\":\"Charles Kayser, John Ott\",\"Plot\":\"Three men hammer on an anvil and pass a bottle of beer around.\",\"Language\":\"None\",\"Country\":\"United States\",\"Awards\":\"1 win\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BNDg0ZDg0YWYtYzMwYi00ZjVlLWI5YzUtNzBkNjlhZWM5ODk5XkEyXkFqcGdeQXVyNDk0MDg4NDk@._V1_SX300.jpg\",\"Ratings\":[{\"Source\":\"Internet Movie Database\",\"Value\":\"6.2/10\"}],\"Metascore\":\"N/A\",\"imdbRating\":\"6.2\",\"imdbVotes\":\"2,541\",\"imdbID\":\"tt0000005\",\"Type\":\"movie\",\"DVD\":\"N/A\",\"BoxOffice\":\"N/A\",\"Production\":\"N/A\",\"Website\":\"N/A\",\"Response\":\"True\"}", plotFilter: "Three men"},
			want: []string{"tt0000005", "Blacksmith Scene", "Three men hammer on an anvil and pass a bottle of beer around."}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := responseData(tt.args.tconst, tt.args.data, tt.args.plotFilter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("responseData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gracefull(t *testing.T) {
	type args struct {
		input      []string
		maxRuntime time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{input: []string{"tt0000005", "Blacksmith Scene", "Three men hammer on an anvil and pass a bottle of beer around."}, maxRuntime: 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gracefull(tt.args.input, tt.args.maxRuntime)
		})
	}
}
