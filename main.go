package main

import (
	"fmt"
	Net "github.com/tobycroft/TuuzNet"
)

func main() {
	query := map[string]interface{}{
		"technicianId": 1000000,
	}
	ua := map[string]string{
		"User-Agent": "PostmanRuntime/7.30.0"}
	resp, err := Net.Get("https://g.meituan.com/domino/craftsman-app/craftsman-detail.html", query, ua, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
	//doc := soup.HTMLParse(resp)
	//links := doc.Find("div", "id", "comicLinks").FindAll("a")
	//for _, link := range links {
	//	fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	//}
}
