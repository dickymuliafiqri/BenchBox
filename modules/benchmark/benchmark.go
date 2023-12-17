package benchmark

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

func StartBenchmark(listenPort uint) map[string]int {
	result := map[string]int{}

	proxyClient, _ := url.Parse(fmt.Sprintf("socks5://0.0.0.0:%d", listenPort))
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyClient),
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 3)

	for _, benchService := range BenchmarkList {
		wg.Add(1)

		go func(service BenchmarkListType) {
			defer wg.Done()

			semaphore <- struct{}{}
			defer func() {
				<-semaphore
			}()

			result[service.Name] = 0
			resp, err := httpClient.Get(service.Domain)
			if err != nil {
				fmt.Printf("Error while testing %s: %s\n", service.Name, err)
			}

			if resp != nil {
				defer resp.Body.Close()

				// fmt.Printf("%s: %d\n", service.Name, resp.StatusCode)
				if resp.StatusCode == service.ExpectedStatus {
					result[service.Name] = 1
				}
			}
		}(benchService)
	}

	wg.Wait()

	return result
}
