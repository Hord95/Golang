package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{"https://metanit.com/go/tutorial/9.5.php", "https://metanit.com/go/tutorial/9.4.php"}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wtf := FetchUrls(urls, ctx)
	for u, x := range wtf {
		fmt.Println(u, x)
	}
	time.Sleep(100 * time.Millisecond)
}
func FetchUrls(urls []string, ctx context.Context) map[string]string {
	var wg sync.WaitGroup
	var mu sync.Mutex
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	res := make(map[string]string)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				mu.Lock()
				res[url] = "error"
				mu.Unlock()

			}
			resp, err := client.Do(req)
			if err != nil {
				mu.Lock()
				res[url] = "error"
				mu.Unlock()

			}

			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			if len(body) >= 100 {
				body = body[:100]
			}

			mu.Lock()
			res[url] = string(body)
			mu.Unlock()

		}(url)
		wg.Wait()

	}
	return res
}
