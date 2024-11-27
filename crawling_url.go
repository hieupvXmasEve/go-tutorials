package main

import (
	"fmt"
	"sync"
)

// Fetcher interface
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, visited *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done() // Báo rằng goroutine hiện tại đã hoàn thành

	// Điều kiện dừng nếu depth <= 0
	if depth <= 0 {
		return
	}

	// Kiểm tra nếu URL đã được duyệt
	if _, ok := visited.LoadOrStore(url, true); ok {
		// Nếu URL đã tồn tại trong visited, dừng lại
		return
	}

	// Fetch nội dung và các URL liên kết
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// In kết quả
	fmt.Printf("found: %s %q\n", url, body)

	// Tạo goroutine mới cho mỗi URL trong danh sách
	for _, u := range urls {
		wg.Add(1) // Thêm một công việc mới
		go Crawl(u, depth-1, fetcher, visited, wg)
	}
}

func main() {
	// Đối tượng sync.Map để lưu trữ các URL đã duyệt
	visited := &sync.Map{}

	// WaitGroup để chờ tất cả các goroutines hoàn thành
	var wg sync.WaitGroup

	// Bắt đầu từ URL gốc với độ sâu tối đa là 4
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, visited, &wg)

	// Chờ tất cả các goroutines hoàn thành
	wg.Wait()

	fmt.Println("Crawl complete")
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
