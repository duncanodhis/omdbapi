package Limiter

import "testing"

func TestRequestLimiter(t *testing.T) {
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
			if got := RequestLimiter(tt.args.url, tt.args.maxRequest); got != tt.want {
				t.Errorf("RequestLimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}
