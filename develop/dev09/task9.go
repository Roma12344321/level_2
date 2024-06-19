package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := flag.String("url", "", "URL to download")
	output := flag.String("output", "", "Output file path")
	flag.Parse()
	if *url == "" || *output == "" {
		fmt.Println("Usage: ./task9 -url <URL> -output <output_file>")
		return
	}
	resp, err := http.Get(*url)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Server returned status %v\n", resp.Status)
		return
	}
	file, err := os.Create(*output)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer func(outFile *os.File) {
		_ = outFile.Close()
	}(file)
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Error copying response to file: %v\n", err)
		return
	}
	fmt.Printf("Downloaded %s to %s\n", *url, *output)
}
