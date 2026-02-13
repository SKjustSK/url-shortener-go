package helpers

import (
	"net/url"
	"os"
	"strings"
)

func EnforceHTTP(url string) string {
	if len(url) < 4 || url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func RemoveDomainError(rawUrl string) bool {
	if rawUrl == os.Getenv("DOMAIN") {
		return false
	}

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return true
	}

	// parsedUrl.Host automatically handles http/https and paths
	// We just need to strip 'www.' if it exists
	cleanHost := strings.TrimPrefix(parsedUrl.Host, "www.")

	if cleanHost == os.Getenv("DOMAIN") {
		return false
	}
	return true
}
