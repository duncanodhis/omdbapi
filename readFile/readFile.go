package readFile

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const maxWorkers = 10

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
func Filter(filepath string, title_Type string, original_Title string,
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
