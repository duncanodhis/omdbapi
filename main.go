package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"
)

const maxWorkers = 10

type OMDAPI struct {
	Title string `json:"Title"`
	Plot  string `json:"Plot"`
}

func omdapiquery(tconst string, plotFilter string) []string {
	results := make([]string, 0)

	var str = "http://www.omdbapi.com/?apikey=5226c193&i=" + tconst
	fmt.Println(str)
	go func() {
		response, err := http.Get(str)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(string(responseData))

		jsonData := []byte(responseData)
		var omdapi OMDAPI
		e := json.Unmarshal(jsonData, &omdapi)
		if err != nil {
			log.Println(e)
		}
		res1, er := regexp.MatchString(plotFilter, omdapi.Plot)
		if er != nil {
			log.Fatal(er)
		}
		if res1 {

			fmt.Println(tconst + "|" + omdapi.Title + "|" + omdapi.Plot)
		} else {
			fmt.Println(" The plot does not match regex provided")
		}
		results = append(results, tconst)
		results = append(results, omdapi.Title)
		results = append(results, omdapi.Plot)
	}()
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	for {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			fmt.Println("sigint")
		case syscall.SIGTERM:
			fmt.Println("sigterm")
			return nil
		}
	}
	return results
}
func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
func findTconst(lines []string, title_Type string, primary_Title string, original_Title string, start_Year string, end_Year string, runtime_Minutes string, genres_ string) []string {
	tconst := make([]string, 0, 0)
	for _, text := range lines {

		split := strings.SplitN(text, "	", 9)
		titleType := strings.TrimSpace(split[1])
		primaryTitle := strings.TrimSpace(split[2])
		originalTitle := strings.TrimSpace(split[3])
		startYear := strings.TrimSpace(split[5])
		endYear := strings.TrimSpace(split[6])
		runtimeMinutes := strings.TrimSpace(split[7])
		genres := strings.TrimSpace(split[8])

		if strings.Contains(titleType, title_Type) && strings.Contains(primaryTitle, primary_Title) && strings.Contains(originalTitle, original_Title) && strings.Contains(startYear, start_Year) && strings.Contains(endYear, end_Year) && strings.Contains(runtimeMinutes, runtime_Minutes) && strings.Contains(genres, genres_) {
			tconst = append(tconst, split[0])

		}

	}
	// fmt.Println(tconst)
	return tconst
}
func getLine(filename string, line chan string, readerr chan error) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line <- scanner.Text()
	}
	close(line)
	readerr <- scanner.Err()
}
func filter(filepath string, title_Type string, original_Title string,
	primary_Title string, start_Year string,
	end_Year string,
	runtime_Minutes string, genres_ string) []string {

	line := make(chan string)
	rows := make([]string, 0, 0)
	tconst := make([]string, 0, 0)
	//var workers int = 10000
	readerr := make(chan error)
	// fmt.Println("Getting file")

	go func() {
		getLine(filepath, line, readerr)
	}()
	fmt.Println("Processing Complete")
	for l := range line {
		//fmt.Println(l)
		rows = append(rows, l)
		tconst = findTconst(rows, title_Type, primary_Title, original_Title, start_Year, end_Year, runtime_Minutes, genres_)

	}
	if err := <-readerr; err != nil {
		log.Fatal(err)
	}

	return tconst
}
func main() {

	var filePath string
	var title_Type string
	var primary_Title string
	var original_Title string
	var genres_ string
	var start_Year string
	var end_Year string
	var runtime_Minutes string
	var maxRequests int
	var maxRunTime time.Duration
	var plotFilter string

	flag.StringVar(&filePath, "f", "./az.csv", "File path")
	flag.StringVar(&title_Type, "t", "short", "filter on TitleType column ")
	flag.StringVar(&primary_Title, "p", "Blacksmith Scene", "filter on primaryType column")
	flag.StringVar(&original_Title, "o", "Blacksmith Scene", "filter on originalTitle column")
	flag.StringVar(&genres_, "g", "Short", "filter on genres column")
	flag.StringVar(&start_Year, "s", "1893", "filter on startYear column")
	flag.StringVar(&end_Year, "e", "\\N", "filter endYear column")
	flag.StringVar(&runtime_Minutes, "r", "1", "filter on runtimeMinutes column")
	flag.IntVar(&maxRequests, "mR", 0, "maximum number of requests to send to omdbapi")
	flag.DurationVar(&maxRunTime, "mT", 0, " maximum run time of the application")
	flag.StringVar(&plotFilter, "pF", "Three men", "regex pattern  filter")
	flag.Parse()
	tconst := make([]string, 0, 0)
	fmt.Println(filePath, title_Type, primary_Title, original_Title, genres_, start_Year, end_Year, plotFilter)
	tconst = filter(
		filePath,
		title_Type,
		original_Title,
		primary_Title,
		start_Year,
		end_Year,
		runtime_Minutes,
		genres_,
	)
	//tconst = filter("./az.csv", "short", "Blacksmith Scene",
	//	"Blacksmith Scene", "1893", "\\N",
	//	"1", "Short")

	fmt.Println(tconst)
	for _, key := range removeDuplicateStr(tconst) {

		omdapiquery(key, plotFilter)
	}

}
