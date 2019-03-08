package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://zahid.al.tair"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"waat://zahid.al.tair",
	}
	want := map[string]bool{
		"https://www.google.com/":   true,
		"https://www.facebook.com/": true,
		"waat://zahid.al.tair":      false,
	}
	got := CheckWebsites(mockWebsiteChecker, websites)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
