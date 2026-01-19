package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"strings"

	gopher "codeberg.org/bryzcolson/net-gopher"
)

type App struct {
	ctx context.Context
}

type GopherItem struct {
	Type     string `json:"type"`
	Display  string `json:"display"`
	Selector string `json:"selector"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type GopherResponse struct {
	Items       []GopherItem `json:"items"`
	Raw         string       `json:"raw"`
	Err         string       `json:"err"`
	ContentType string       `json:"contentType"`
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Fetch(host string, port string, selector string, itemType string) GopherResponse {
	if port == "" {
		port = "70"
	}
	if selector == "" {
		selector = "/"
	}
	if itemType == "" {
		itemType = string(gopher.TypeDirectory)
	}

	// Check if selector contains a search query (tab-separated)
	query := ""
	if idx := strings.Index(selector, "\t"); idx != -1 {
		query = selector[idx+1:]
		selector = selector[:idx]
	}

	gopherURL := fmt.Sprintf("gopher://%s:%s/%s%s", host, port, itemType, selector)
	req, err := gopher.NewRequest(gopherURL)
	if err != nil {
		return GopherResponse{Err: fmt.Sprintf("Failed to create request: %v", err)}
	}
	req.Query = query

	resp, err := gopher.DefaultClient.Do(req)
	if err != nil {
		return GopherResponse{Err: fmt.Sprintf("Failed to fetch: %v", err)}
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return GopherResponse{Err: fmt.Sprintf("Failed to read response: %v", err)}
	}

	if itemType == string(gopher.TypeGIF) || itemType == string(gopher.TypeImage) {
		encoded := base64.StdEncoding.EncodeToString(data)
		return GopherResponse{ContentType: "image", Raw: encoded}
	}

	if itemType == string(gopher.TypeText) {
		rawText := string(data)
		rawText = strings.TrimSuffix(rawText, ".\r\n")
		rawText = strings.TrimSuffix(rawText, ".\n")
		return GopherResponse{ContentType: "text", Raw: rawText}
	}

	rawText := string(data)
	scanner := bufio.NewScanner(strings.NewReader(rawText))
	var items []GopherItem
	for scanner.Scan() {
		line := scanner.Text()

		if line == "." {
			break
		}

		if len(line) == 0 {
			continue
		}

		if !strings.Contains(line, "\t") {
			if len(line) > 0 && line[0] == byte(gopher.TypeInfo) {
				items = append(items, GopherItem{Type: string(gopher.TypeInfo), Display: line[1:]})
			} else {
				items = append(items, GopherItem{Type: string(gopher.TypeInfo), Display: line})
			}
			continue
		}

		item := parseGopherLine(line)
		items = append(items, item)
	}

	if err := scanner.Err(); err != nil {
		return GopherResponse{Err: fmt.Sprintf("Failed to read from connection: %v", err)}
	}

	return GopherResponse{Items: items, Raw: rawText}
}

func parseGopherLine(line string) GopherItem {
	parts := strings.Split(line[1:], "\t")
	item := GopherItem{Type: string(line[0])}

	if len(parts) > 0 {
		item.Display = parts[0]
	}
	if len(parts) > 1 {
		item.Selector = parts[1]
	}
	if len(parts) > 2 {
		item.Host = parts[2]
	}
	if len(parts) > 3 {
		item.Port = parts[3]
	}
	return item
}
