package omdbquery

import (
	"encoding/json"
	"fmt"
	"log"
	"omdbapi/Limiter"
	"regexp"
)

type OMDAPI struct {
	Title string `json:"Title"`
	Plot  string `json:"Plot"`
}

//takes the tconst,the string data ,and the plotfilter
//then formats the data{originally in json} returns a slice
func responseData(tconst string, data string, plotFilter string) []string {
	results := make([]string, 0)
	omdapi := OMDAPI{}
	//fmt.Println("date =" + data)
	e := json.Unmarshal([]byte(data), &omdapi)
	if e != nil {
		fmt.Println(e)
	}
	//fmt.Println("hello?How are you ?just see me if i exist :-(" + omdapi.Plot)
	res1, er := regexp.MatchString(plotFilter, omdapi.Plot)
	if er != nil {
		log.Fatal(er)
	}
	if res1 {

		//fmt.Println(tconst + "|" + omdapi.Title + "|" + omdapi.Plot)
		results = append(results, tconst)
		results = append(results, omdapi.Title)
		results = append(results, omdapi.Plot)
	} else {
		fmt.Println(" The plot does not match regex provided")
	}

	return results

}

//omdapi query  returns the data sent to responseData for processing and cleaning
func Omdapiquery(tconst string, plotFilter string, maxRequest int) []string {

	temp := make([]string, 0)
	var str = "http://www.omdbapi.com/?apikey=5226c193&i=" + tconst
	respData := Limiter.RequestLimiter(str, maxRequest)
	//fmt.Println("we are in query" + respData)
	temp = responseData(tconst, respData, plotFilter)
	return temp
}
