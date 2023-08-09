package gpic

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// Test good URLs over the Internet
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

// Test a bad favicon URL that returns just text instead of an image
func TestGetFaviconNotImage(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<html><body><p>I'll always return HTTP 200.</p></body></html>")
	}))
	defer ts.Close()

	inputURL, err := url.Parse(ts.URL)
	if err != nil {
		t.Errorf("The input URL failed to parse: %s", ts.URL)
		return
	}

	URL, err := GetFavicon(inputURL)
	fmt.Errorf("This is what is returned: %s", URL) //DEBUG
	if err != nil && errors.Is(err, ErrNotFound) {
		// good, it should have failed
		return
	} else if err != nil {
		t.Errorf("GetFavicon failed. Input URL: %s Error: %s", ts.URL, err)
		return
	}

	t.Errorf("GetFavicon returned successfully when it shouldn't have. Input URL: %s", ts.URL)

}
