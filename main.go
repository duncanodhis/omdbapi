package main

import (
	"flag"
	"fmt"
	"omdbapi/readFile"
	"time"
)

//main function

func main() {

	var filePath, titleType, primaryTitle,
		originalTitle, genres_, startYear,
		endYear, runtimeMinutes, plotFilter string
	var maxRequests int
	var maxRunTime time.Duration

	flag.StringVar(&filePath, "f", "./az.csv", "File path")
	flag.StringVar(&titleType, "t", "short", "filter on TitleType column ")
	flag.StringVar(&primaryTitle, "p", "Blacksmith Scene", "filter on primaryType column")
	flag.StringVar(&originalTitle, "o", "Blacksmith Scene", "filter on originalTitle column")
	flag.StringVar(&genres_, "g", "Short", "filter on genres column")
	flag.StringVar(&startYear, "s", "1893", "filter on startYear column")
	flag.StringVar(&endYear, "e", "\\N", "filter endYear column")
	flag.StringVar(&runtimeMinutes, "r", "1", "filter on runtimeMinutes column")
	flag.IntVar(&maxRequests, "mR", 5, "maximum number of requests to send to omdbapi")
	flag.DurationVar(&maxRunTime, "mT", 10, " maximum run time of the application")
	flag.StringVar(&plotFilter, "pF", "Three men", "regex pattern  filter")
	flag.Parse()
	tconst := make([]string, 0, 0)
	fmt.Println("Input values")
	fmt.Println(filePath, titleType, primaryTitle, originalTitle, genres_, startYear, endYear, plotFilter)
	tconst = readFile.filter(
		filePath,
		titleType,
		originalTitle,
		primaryTitle,
		startYear,
		endYear,
		runtimeMinutes,
		genres_,
	)
	//tconst = filter("./az.csv", "short", "Blacksmith Scene",
	//	"Blacksmith Scene", "1893", "\\N",
	//	"1", "Short")
	//output := make([]string, 0)

	//fmt.Println(tconst)
	for _, key := range removeDuplicateStr(tconst) {
		//fmt.Println(key)
		i := omdapiquery(key, plotFilter, 10)
		gracefull(i, maxRunTime)
	}

}
