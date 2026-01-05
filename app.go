package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"strings"
	"time"
)

type App struct {
	ctx context.Context
}

type GopherItem struct {
	Type        string `json:"type"`
	Display     string `json:"display"`
	Selector    string `json:"selector"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Description string `json:"description"`
}

type GopherResponse struct {
	Items []GopherItem `json:"items"`
	Raw   string       `json:"raw"`
	Err   string       `json:"err"`
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Fetch(host string, port string, selector string) GopherResponse {
	if port == "" {
		port = "70"
	}
	if selector == "" {
		selector = "/"
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

	scanner := bufio.NewScanner(conn)
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

	return GopherResponse{Items: items}
}

func parseGopherLine(line string) GopherItem {
	if len(line) < 1 {
		return GopherItem{
			Type:        "i",
			Display:     " ",
			Selector:    "",
			Host:        "error.host",
			Port:        "1",
			Description: getTypeDescription("i"),
		}
	}

	// Check if line starts with a tab or has no tabs (not proper gopher format)
	// Treat as plain text info line
	if line[0] == '\t' || !strings.Contains(line, "\t") {
		return GopherItem{
			Type:        "i",
			Display:     line,
			Selector:    "",
			Host:        "error.host",
			Port:        "1",
			Description: getTypeDescription("i"),
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
	item.Description = getTypeDescription(itemType)
	return item
}

func getTypeDescription(itemType string) string {
	descriptions := map[string]string{
		"0": "Text file",
		"1": "Directory",
		"2": "CSO phone-book server",
		"3": "Error",
		"4": "BinHex encoded file",
		"5": "DOS binary archive",
		"6": "UNIX uuencoded file",
		"7": "Search server",
		"8": "Telnet session",
		"9": "Binary file",
		"+": "Redundant server",
		"g": "GIF image",
		"I": "Image file",
		"T": "TN3270 session",
		"h": "HTML file",
		"i": "Info line",
		"s": "Sound file",
	}

	if desc, ok := descriptions[itemType]; ok {
		return desc
	}
	return "Unknown type"
}
