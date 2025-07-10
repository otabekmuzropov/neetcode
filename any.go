package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	fmt.Println(strings.Fields("task_mangement"))
	escapedQuery := url.QueryEscape("task")
	fmt.Println(escapedQuery)
	icon, err := searchIcons(escapedQuery)
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println("First icon found:", icon)
}

type IconSearchResponse struct {
	Icons []string `json:"icons"`
}

func searchIcons(query string) (string, error) {
	url := fmt.Sprintf("https://api.iconify.design/search?query=%s", query)

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result IconSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Icons) == 0 {
		return "", fmt.Errorf("no icons found for query: %s", query)
	}

	return result.Icons[0], nil
}
