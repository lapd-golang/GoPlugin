package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type urlString string

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func (u urlString) DownloadFile(url string) {
	now := time.Now().String()
	filepath := "./resources" + "image" + now + ".jpg"
	fmt.Println("Download at____", filepath)
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		fmt.Println("\nerror 1:-->", err, 1)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("\nerror 2:-->\n", err, url)
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("\nerror 3:-->\n", err)
	}

	fmt.Printf("\n---->>>Done Download url:%s", url)
	fmt.Printf("\n")
}

var DownloadIDE urlString
