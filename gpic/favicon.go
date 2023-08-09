package gpic

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"slices"
)

var faviconFileTypes = []string{"svg", "png", "ico"}
var faviconContentTypes = []string{
	"image/svg+xml",
	"image/x-png",
	"image/png",
	"image/x-icon",
	"image/vnd.microsoft.icon",
}
var ErrNotFound = errors.New("A favicon was not found.")

/*
 * Finds and returns a URL to a favicon. The favicon is searched for on the
 * hostname provided. Currently will return only the very first favicon found
 * regardless of size, filetype, etc.
 *
 * For now this fuction only checks for favicons at the root of the website.
 */
func GetFavicon(URL *url.URL) (string, error) {

	// Prep the URL. Drop any path from the URL as we're only concerned about
	// hostnames
	inputURL := fmt.Sprintf("%s://%s/favicon.", URL.Scheme, URL.Host)

	client := &http.Client{
		Timeout: time.Duration(time.Second * 10),
	}

	// try for each favicon filetype at the root of the hostname
	for _, fileType := range faviconFileTypes {

		resp, err := client.Get(inputURL + fileType)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 && slices.Contains(faviconContentTypes, resp.Header.Get("Content-Type")) {
			return inputURL + fileType, nil
		}
	}

	return "", ErrNotFound
}
