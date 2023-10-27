package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckWebsites(t *testing.T) {
	mockChecker := func(url string) bool {
		if url == "not_a_url" {
			return false
		}
		return true
	}
	t.Run("test that it returns the right result", func(t *testing.T) {
		urls := []string{
			"https://www.google.com",
			"https://www.duckduckgo.com",
			"https://www.facebook.com",
			"not_a_url",
		}

		results := CheckWebsite(mockChecker, urls)
		want := map[string]bool{
			"https://www.google.com":     true,
			"https://www.duckduckgo.com": true,
			"https://www.facebook.com":   true,
			"not_a_url":                  false,
		}

		if !reflect.DeepEqual(results, want) {
			t.Errorf("got %#v, but want %#v", results, want)
		}
	})
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsite(slowStubWebsiteChecker, urls)
	}
}
