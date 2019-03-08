package concurrency

import (
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkWebsiteChecker(b *testing.B) {
	websites := []string{}
	for i := 0; i < 100; i++ {
		websites = append(websites, "a url")
	}
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, websites)
	}
}
