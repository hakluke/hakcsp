package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	concurrencyPtr := flag.Int("t", 8, "Number of threads to utilise. Default is 8.")
	flag.Parse()

	work := make(chan string)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			work <- s.Text()
		}
		close(work)
	}()

	wg := &sync.WaitGroup{}

	for i := 0; i < *concurrencyPtr; i++ {
		wg.Add(1)
		go doWork(work, wg)
	}
	wg.Wait()
}

func doWork(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range work {
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Error fetching url:", err)
			continue
		}
		value := resp.Header.Get("Content-Security-Policy")
		for _, s := range strings.Split(value, " ") {
			if strings.Contains(s, ".") { // weird way to check if it's a domain, unsure if this will work in all situations
				fmt.Println(s)
			}
		}
	}
}
