package statuschecker

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
)

func New() StatusChecker {
	return &checker{}
}

func (c *checker) Check(concurrency int, src string, urls ...string) (<-chan *Result, <-chan struct{}, error) {
	resChan := make(chan *Result)
	done := make(chan struct{}, 1)

	parsedURLS := make([]string, 0)

	parsedURLS = append(parsedURLS, urls...)
	parsedURLS = c.parseURLsFromFIle(parsedURLS, src)

	go func() {
		var wg sync.WaitGroup

		sem := make(chan struct{}, limiter)

		for _, url := range parsedURLS {
			sem <- struct{}{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				res, err := c.checkStatus(url)
				if err != nil {
					os.Stderr.WriteString(err.Error())
				} else {
					resChan <- &Result{
						Up:  res,
						Url: url,
					}
				}

				<-sem
			}()
		}

		wg.Wait()
		close(done)
	}()

	return resChan, done, nil
}

func (c *checker) parseURLsFromFIle(parsedURLS []string, src string) []string {
	f, err := os.Open(src)
	if err != nil {
		fmt.Println("could not open " + src)

		return parsedURLS
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		url := s.Text()
		if url != "" {
			parsedURLS = append(parsedURLS, url)
		}
	}

	return parsedURLS
}

func (c *checker) checkStatus(url string) (bool, error) {
	resp, err := http.DefaultClient.Head(url)
	if err != nil {
		return false, fmt.Errorf("failed to call perform HTTP HEAD request for %s: %w", url, err)
	}

	if resp.StatusCode >= 400 {
		retry, errR := http.DefaultClient.Get(url)
		if errR != nil {
			return false, fmt.Errorf("failed to call perform HTTP HEAD & GET requests for %s: %w", url, err)
		}

		if retry.StatusCode >= 400 {
			return false, nil
		}
	}

	return true, nil
}
