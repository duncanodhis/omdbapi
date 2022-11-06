package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"golang.org/x/time/rate"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"
)

const maxWorkers = 10

//RLHTTPClient Rate Limited HTTP Client
type RLHTTPClient struct {
	client      *http.Client
	Ratelimiter *rate.Limiter
}

type OMDAPI struct {
	Title string `json:"Title"`
	Plot  string `json:"Plot"`
}

//Do dispatches the HTTP request to the network
func (c *RLHTTPClient) Do(req *http.Request) (*http.Response, error) {
	// Comment out the below 5 lines to turn off ratelimiting
	ctx := context.Background()
	err := c.Ratelimiter.Wait(ctx) // This is a blocking call. Honors the rate limit
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//NewClient return http client with a ratelimiter
func NewClient(rl *rate.Limiter) *RLHTTPClient {
	c := &RLHTTPClient{
		client:      http.DefaultClient,
		Ratelimiter: rl,
	}
	return c
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

//limits the number of requests sent then returns http.get() as a response in string format
func requestLimiter(url string, maxRequest int) string {
	rl := rate.NewLimiter(rate.Every(time.Second), maxRequest) // max request every 10 seconds
	c := NewClient(rl)
	reqURL := url
	var results []byte
	req, err := http.NewRequest("GET", reqURL, nil)
	//fmt.Println(req)
	if err != nil {

		fmt.Print(err.Error())
		os.Exit(1)
	}

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < maxRequest; i++ {
		resp, err := c.Do(req)
		response, err := ioutil.ReadAll(resp.Body)
		results = response
		//fmt.Println(string(response))
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(resp.StatusCode)

		}
		if resp.StatusCode == 429 {

			fmt.Printf("Rate limit reached after %d requests", i)
			return string(results)
		}
	}
	return string(results)
}

//omdapi query  returns the data sent to responseData for processing and cleaning
func omdapiquery(tconst string, plotFilter string, maxRequest int) []string {

	temp := make([]string, 0)
	var str = "http://www.omdbapi.com/?apikey=5226c193&i=" + tconst
	respData := requestLimiter(str, maxRequest)
	//fmt.Println("we are in query" + respData)
	temp = responseData(tconst, respData, plotFilter)
	return temp
}

//performs gracefull exit when runtime is achievd or sigterm is called
func gracefull(input []string, maxRuntime time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), maxRuntime*time.Second)
	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		cancel()
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		select {

		case <-ctx.Done():
			fmt.Println(input[0] + "|" + input[1] + "|" + input[2])
			break
		case <-time.After(1 * time.Second):
			fmt.Println(input[0] + "|" + input[1] + "|" + input[2])
			break

		}
	}()
	wg.Wait()
	fmt.Println("Done")
}

//Removes duplicates from a slice
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

//finds the unique identifier tconst
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

//gets each line of the file
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

//filters the data according to flags given ,it reads the file using limited goroutine
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
	
	ch := make(chan struct{}, maxWorkers)
	ch <- struct{}{}
	go func() {
		getLine(filepath, line, readerr)
		<-ch
	}()
	fmt.Println("Processing file Complete")
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

//main function

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
	flag.IntVar(&maxRequests, "mR", 5, "maximum number of requests to send to omdbapi")
	flag.DurationVar(&maxRunTime, "mT", 10, " maximum run time of the application")
	flag.StringVar(&plotFilter, "pF", "Three men", "regex pattern  filter")
	flag.Parse()
	tconst := make([]string, 0, 0)
	fmt.Println("Input values")
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
	//output := make([]string, 0)

	//fmt.Println(tconst)
	for _, key := range removeDuplicateStr(tconst) {
		//fmt.Println(key)
		i := omdapiquery(key, plotFilter, 10)
		gracefull(i, maxRunTime)
	}

}
