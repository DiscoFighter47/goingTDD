package concurrency

// WebsiteChecker checks a website
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites checks urls
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := map[string]bool{}
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
