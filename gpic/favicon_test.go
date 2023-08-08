package gpic

import (
	"net/url"
	"testing"
)

func TestGetFavicon(t *testing.T) {

	testCases := []struct {
		inputURL   string
		faviconURL string
	}{
		{"http://neverssl.com", "http://neverssl.com/favicon.ico"},
		{"https://www.espn.com", "https://www.espn.com/favicon.png"},
	}

	for idx, testCase := range testCases {

		inputURL, err := url.Parse(testCase.inputURL)
		if err != nil {
			t.Errorf("The input URL failed to parse: %s", testCase.inputURL)
			continue
		}

		URL, err := GetFavicon(inputURL)
		if err != nil {
			t.Errorf("Failed to get any favicon with the input URL: %s", testCase.inputURL)
			continue
		}

		if URL != testCase.faviconURL {
			t.Errorf("The favicon URL doesn't match. id: %d got: %s want: %s", idx+1, URL, testCase.faviconURL)
			continue
		}
	}
}
