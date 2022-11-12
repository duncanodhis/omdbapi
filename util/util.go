package util

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//performs Gracefull exit when runtime is achievd or sigterm is called
func Gracefull(input []string, maxRuntime time.Duration) {
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

// RemoveDuplicateStr Removes duplicates from a slice
func RemoveDuplicateStr(strSlice []string) []string {
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
