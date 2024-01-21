package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {

	c := colly.NewCollector()
	count := 0

	/* c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")

		fmt.Printf("link found: %q -> %s\n", h.Text, link)

		c.Visit(h.Request.AbsoluteURL(link))
	})
	*/

	c.OnHTML("img", func(h *colly.HTMLElement) {
		count = count + 1
		name := "image"
		filename := fmt.Sprintf("%s%d", name, count) + ".jpg"
		image := h.Attr("src")
		print(image)

		resp, err := http.Get(image)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		out, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println("File downloaded successfully: ", filename)

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})

	c.Visit("https://in.pinterest.com/piriltiyildiz/all-type-monkey/")
}
