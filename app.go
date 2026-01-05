package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
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
		itemType = "1"
	}

	address := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.DialTimeout("tcp", address, 10*time.Second)
	if err != nil {
		return GopherResponse{Err: fmt.Sprintf("Connection failed: %v", err)}
	}
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(10 * time.Second))

	_, err = conn.Write([]byte(selector + "\r\n"))
	if err != nil {
		return GopherResponse{Err: fmt.Sprintf("Failed to write to connection: %v", err)}
	}

	data, err := io.ReadAll(conn)
	if err != nil {
		return GopherResponse{Err: fmt.Sprintf("Failed to read from connection: %v", err)}
	}

	if itemType == "g" || itemType == "I" {
		encoded := base64.StdEncoding.EncodeToString(data)
		return GopherResponse{
			ContentType: "image",
			Raw:         encoded,
		}
	}

	rawText := string(data)
	scanner := bufio.NewScanner(strings.NewReader(rawText))
	var items []GopherItem
	for scanner.Scan() {
		line := scanner.Text()

		if line == "." {
			break
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
	if len(line) < 1 {
		return GopherItem{
			Type:     "i",
			Display:  " ",
			Selector: "",
			Host:     "error.host",
			Port:     "1",
		}
	}

	if line[0] == '\t' || !strings.Contains(line, "\t") {
		return GopherItem{
			Type:     "i",
			Display:  line,
			Selector: "",
			Host:     "error.host",
			Port:     "1",
		}
	}

	itemType := string(line[0])
	parts := strings.Split(line[1:], "\t")
	item := GopherItem{Type: itemType, Display: ""}

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
