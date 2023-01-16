package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
)

func main() {
	resp, err := soup.Get("https://g.meituan.com/domino/craftsman-app/craftsman-detail.html?technicianId=1000000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resp)
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "id", "comicLinks").FindAll("a")
	for _, link := range links {
		fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	}
}
