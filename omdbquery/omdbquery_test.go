package omdbquery

import (
	"reflect"
	"testing"
)

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
			args: args{tconst: "tt0000005",
				data:       "{\"Title\":\"Blacksmith Scene\",\"Year\":\"1893\",\"Rated\":\"Unrated\",\"Released\":\"09 May 1893\",\"Runtime\":\"1 min\",\"Genre\":\"Short, Comedy\",\"Director\":\"William K.L. Dickson\",\"Writer\":\"N/A\",\"Actors\":\"Charles Kayser, John Ott\",\"Plot\":\"Three men hammer on an anvil and pass a bottle of beer around.\",\"Language\":\"None\",\"Country\":\"United States\",\"Awards\":\"1 win\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BNDg0ZDg0YWYtYzMwYi00ZjVlLWI5YzUtNzBkNjlhZWM5ODk5XkEyXkFqcGdeQXVyNDk0MDg4NDk@._V1_SX300.jpg\",\"Ratings\":[{\"Source\":\"Internet Movie Database\",\"Value\":\"6.2/10\"}],\"Metascore\":\"N/A\",\"imdbRating\":\"6.2\",\"imdbVotes\":\"2,541\",\"imdbID\":\"tt0000005\",\"Type\":\"movie\",\"DVD\":\"N/A\",\"BoxOffice\":\"N/A\",\"Production\":\"N/A\",\"Website\":\"N/A\",\"Response\":\"True\"}",
				plotFilter: "Three men"},
			want: []string{"tt0000005", "Blacksmith Scene", "Three men hammer on an anvil and pass a bottle of beer around."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := responseData(tt.args.tconst, tt.args.data, tt.args.plotFilter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("responseData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOmdapiquery(t *testing.T) {
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
			if got := Omdapiquery(tt.args.tconst, tt.args.plotFilter, tt.args.maxRequest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Omdapiquery() = %v, want %v", got, tt.want)
			}
		})
	}
}
