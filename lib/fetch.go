package lib

import (
	"io"
	"net/http"
)

func FetchPage(key string) string {

	resp, err := http.Get(key)
	
	if err != nil {
		return "[ERROR] Failed to fetch page"
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "[ERROR] Failed to fetch page"
	}

	return string(body)
}
