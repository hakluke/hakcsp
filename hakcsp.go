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

	numWorkers := *concurrencyPtr
	work := make(chan string)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			work <- s.Text()
		}
		close(work)
	}()

	wg := &sync.WaitGroup{}

	for i := 0; i < numWorkers; i++ {
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
		}
		value := resp.Header.Get("Content-Security-Policy")
		for _, s := range strings.Split(value, " ") {
			if strings.HasPrefix(s, "http") {
				fmt.Println(s)
			}
		}
	}
}
