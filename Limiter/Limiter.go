package Limiter

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//RLHTTPClient Rate Limited HTTP Client
type RLHTTPClient struct {
	client      *http.Client
	Ratelimiter *rate.Limiter
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

//limits the number of requests sent then returns http.get() as a response in string format
func RequestLimiter(url string, maxRequest int) string {
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
