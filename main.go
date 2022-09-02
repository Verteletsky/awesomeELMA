package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

func main() {
	startParseGo(urls())
}

func startParseGo(listUrls []string) {
	// создаем два канала
	jobs := make(chan string, 5)
	results := make(chan int, 5)

	// максимум до 5 воркеров
	countWorker := len(listUrls)
	if countWorker > 5 {
		countWorker = 5
	} else {
		countWorker = len(listUrls)
	}

	// анализируем регулярку
	r, err := regexp.Compile("Go{1}")
	if err != nil {
		log.Println(err)
	}
	worker(countWorker, jobs, results, r) // запускаем воркер

	// заполняет джобы данными
	go func() {
		for j := 0; j < len(listUrls); j++ {
			jobs <- listUrls[j]
		}
		close(jobs) // закрываем каналы
	}()

	printTotal(listUrls, results) // выводим результат
}

func printTotal(listUrls []string, results chan int) {
	total := 0
	for re := 0; re < len(listUrls); re++ {
		total += <-results
	}

	fmt.Println("Total: ", total)
}

func worker(countWorker int, jobs chan string, results chan int, r *regexp.Regexp) {
	for w := 0; w < countWorker; w++ {
		go func(jobs <-chan string, results chan<- int, r *regexp.Regexp) {
			for j := range jobs {
				res := findGo(r, j)
				results <- res
				fmt.Println("Count for:", j, res)
			}
		}(jobs, results, r)
	}
}

func findGo(r *regexp.Regexp, j string) int {
	response, err := http.Get(j) // todo refactor to NewRequestWithContext
	if err != nil {
		log.Println(err)
	}
	readAll, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	response.Body.Close()
	strings := r.FindAllStringIndex(string(readAll), -1)
	return len(strings)
}

func urls() []string {
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
	}
	return listUrls
}
