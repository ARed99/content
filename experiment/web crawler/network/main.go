package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/playwright-community/playwright-go"
)

//download files

func DownloadFile(url string, filepath string) error {
	// Create the file
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get the response from the server
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		println("status code is %d", resp.StatusCode)
	}

	// Copy the response body to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto(""); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	images, err := page.Locator("img").All()

	if err != nil {
		println("could not fetch images")
	}

	for i, image := range images {

		src, err := image.GetAttribute("src")

		if err != nil {
			println("could not find src of image")
		}

		filename := fmt.Sprintf("%s%d", "image", i)
		DownloadFile(src, filename+".jpg")
		println("%d: %s\n", i+1, src)
	}
}
