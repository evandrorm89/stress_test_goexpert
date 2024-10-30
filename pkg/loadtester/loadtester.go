package loadtester

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type LoadTester struct {
	TotalTime          time.Duration
	TotalRequests      int
	SuccessfulRequests int
	OtherStatusCodes   map[int]int
}

func RunLoadTest(url string, totalRequests int, concurrency int) *LoadTester {
	var wg sync.WaitGroup
	startTime := time.Now()
	load := &LoadTester{
		OtherStatusCodes: make(map[int]int),
	}

	fmt.Println("Initializing test, please wait...")

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < totalRequests/concurrency; j++ {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Printf("Error making request: %v\n", err)
					continue
				}
				defer resp.Body.Close()

				load.TotalRequests++
				if resp.StatusCode == http.StatusOK {
					load.SuccessfulRequests++
				} else {
					load.OtherStatusCodes[resp.StatusCode]++
				}
			}

		}()
	}

	wg.Wait()
	load.TotalTime = time.Since(startTime)
	return load
}
