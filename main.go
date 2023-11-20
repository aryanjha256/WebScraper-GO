package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: webscraper <url>")
		os.Exit(1)
	}

	url := os.Args[1]
	content, err := fetch(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		os.Exit(1)
	}

	links, err := extractLinks(content)
	if err != nil {
		fmt.Println("Error extracting links:", err)
		os.Exit(1)
	}

	fmt.Println("Links on", url)
	for _, link := range links {
		fmt.Println(link)
	}
}

func fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	// Read the response body
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func extractLinks(content string) ([]string, error) {
	var links []string
	tokenizer := html.NewTokenizer(strings.NewReader(content))

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			if tokenizer.Err() == io.EOF {
				return links, nil
			}
			return links, tokenizer.Err()
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}
