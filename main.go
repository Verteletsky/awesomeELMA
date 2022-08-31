package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

func main() {
	listUrls := []string{
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
		"https://go.dev/",
	}

	const numJobs = 5
	jobs := make(chan string, numJobs)
	results := make(chan int, numJobs)

	countWorker := len(listUrls)
	if countWorker > 5 {
		countWorker = 5
	} else {
		countWorker = len(listUrls)
	}

	for w := 0; w < countWorker; w++ {
		go worker(jobs, results)
	}

	for j := 0; j < len(listUrls); j++ {
		jobs <- listUrls[j]
	}
	close(jobs)

	total := 0
	for r := 0; r < len(listUrls); r++ {
		total += <-results
	}

	fmt.Println("Total: ", total)
}

func worker(jobs <-chan string, result chan<- int) {
	for j := range jobs {
		response, err := http.Get(j)
		if err != nil {
			log.Fatalln(err)
		}
		readAll, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		response.Body.Close()
		go func() {
			res := findGo(string(readAll))
			fmt.Println("Count for:", j, res)
			result <- res
		}()
	}
}

func findGo(str string) int {
	r, err := regexp.Compile("Go{1}")
	if err != nil {
		log.Fatalln(err)
	}
	strings := r.FindAllStringIndex(str, -1)
	return len(strings)
}
