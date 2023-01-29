package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
)

func main() {
	//query := map[string]interface{}{
	//	"technicianId": 1000000,
	//}
	//ua := map[string]string{
	//	"User-Agent": "PostmanRuntime/7.30.0"}
	//resp, err := Net.Get("https://g.meituan.com/domino/craftsman-app/craftsman-detail.html", query, ua, nil)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(resp)
	////doc := soup.HTMLParse(resp)
	////links := doc.Find("div", "id", "comicLinks").FindAll("a")
	////for _, link := range links {
	////	fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	////}

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(e *colly.Response) {
		body := string(e.Body)
		bodys1 := strings.Split(body, "window.__INITIAL_STATE__ = ")
		bodys2 := bodys1[len(bodys1)-1]
		bodys3 := strings.Split(bodys2, "</script>")
		bodys4 := bodys3[0]
		s6 := strings.TrimSpace(bodys4)
		//jsoniter.UnmarshalFromString(s6)
		fmt.Println(s6)
	})

	c.Visit("https://g.meituan.com/domino/craftsman-app/craftsman-detail.html?technicianId=11728812")

}
