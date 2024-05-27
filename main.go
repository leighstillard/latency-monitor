package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	url        = "https://www.strategiqai.com.au/"
	outputFile = "latency_results.txt"
	interval   = 5 * time.Minute
)

func checkLatency() {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	latency := time.Since(start).Milliseconds()
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "%s Latency: %dms\n", timestamp, latency)
}

func main() {
	for {
		checkLatency()
		time.Sleep(interval)
	}
}
