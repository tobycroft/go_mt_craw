package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	jsoniter "github.com/json-iterator/go"
	"main.go/tuuz"
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
		var datas Data
		jsoniter.UnmarshalFromString(s6, &datas)

		var bff BffData
		jsoniter.UnmarshalFromString(datas.BffData[0], &bff)
		fmt.Println(bff.ResponseData[0].Data.Data.AttrValues.Name)
		fmt.Println(bff.ResponseData[0].Data.Data.AttrValues.Skills)
		fmt.Println(bff.ResponseData[0].Data.Data.AttrValues.WorkYears)
		fmt.Println(bff.ResponseData[0].Data.Data.AttrValues.WorkYearsStr)
		fmt.Println(bff.ResponseData[0].Data.Data.TechnicianID)
		fmt.Println(bff.ResponseData[0].Data.Data.AttrValues.PhotoURL)
		fmt.Println(bff.ResponseData[0].Data.Data.ShopIDForFe)

		fmt.Println(bff.ResponseData[0].Data.Data.Share.Title)
		fmt.Println(bff.ResponseData[0].Data.Data.Share.Desc)
		fmt.Println(bff.ResponseData[0].Data.Data.Share.URL)

		db := tuuz.Db().Table("mt_craw")
		db.Where("techid", bff.ResponseData[0].Data.Data.TechnicianID)
		ret, err := db.Find()
		if err != nil {
			panic(err)
		}
		if len(ret) > 0 {
			panic("已经存在")
		}

		db = tuuz.Db().Table("mt_craw")
		data := map[string]interface{}{
			"name":         bff.ResponseData[0].Data.Data.AttrValues.Name,
			"skills":       strings.Join(bff.ResponseData[0].Data.Data.AttrValues.Skills, ","),
			"workyears":    bff.ResponseData[0].Data.Data.AttrValues.WorkYears,
			"workyearsstr": bff.ResponseData[0].Data.Data.AttrValues.WorkYearsStr,
			"techid":       bff.ResponseData[0].Data.Data.TechnicianID,
			"photo":        bff.ResponseData[0].Data.Data.AttrValues.PhotoURL,
			"shopidforfe":  bff.ResponseData[0].Data.Data.ShopIDForFe,
			"title":        bff.ResponseData[0].Data.Data.Share.Title,
			"desc":         bff.ResponseData[0].Data.Data.Share.Desc,
			"url":          bff.ResponseData[0].Data.Data.Share.URL,
		}
		db.Data(data)
		_, err = db.Insert()
		if err != nil {
			panic(err)
		}
	})

	c.Visit("https://g.meituan.com/domino/craftsman-app/craftsman-detail.html?technicianId=11728812")

}

type Data struct {
	BffDatas interface{}
	BffData  []string
}

type BffData struct {
	IsOK    string `json:"isOK"`
	Context struct {
		JsContext struct {
		} `json:"jsContext"`
		WarmUpSeq             interface{} `json:"warmUpSeq"`
		UseAsWarmUpData       bool        `json:"useAsWarmUpData"`
		TimelyWarmUp          bool        `json:"timelyWarmUp"`
		AllowCrossEnvCalls    interface{} `json:"allowCrossEnvCalls"`
		FuncScriptExecTimeout interface{} `json:"funcScriptExecTimeout"`
		ScriptID              int         `json:"scriptId"`
		HeadersMap            struct {
		} `json:"headersMap"`
		ParamsMap struct {
		} `json:"paramsMap"`
		Headers struct {
			MTSIRemoteIP     string `json:"MTSI-remote-ip"`
			UpstreamName     string `json:"$upstream_name"`
			MTSIFlowStrategy string `json:"MTSI-flow-strategy"`
			MTSIScore        string `json:"MTSI-score"`
			Accept           string `json:"Accept"`
			SafaInternal     string `json:"safa-internal"`
			MTSIFlag         string `json:"MTSI-flag"`
			IP               string `json:"ip"`
			XForwardedProto  string `json:"X-Forwarded-Proto"`
			UserAgent        string `json:"User-Agent"`
			Host             string `json:"Host"`
			AcceptEncoding   string `json:"Accept-Encoding"`
			MTSIChecked      string `json:"MTSI-checked"`
			XForwardedFor    string `json:"X-Forwarded-For"`
			PostmanToken     string `json:"Postman-Token"`
			Href             string `json:"href"`
			XRealIP          string `json:"X-Real-IP"`
			MTSIRequestCode  string `json:"MTSI-request-code"`
		} `json:"headers"`
		Params struct {
			PageIdentifier string `json:"page_identifier"`
			Env            struct {
				ClientVersion interface{} `json:"clientVersion"`
			} `json:"env"`
			TechnicianID string `json:"technicianId"`
		} `json:"params"`
		ReqSeqInFuncScript interface{} `json:"reqSeqInFuncScript"`
		ExecParamTypeEnum  struct {
		} `json:"execParamTypeEnum"`
		ExecTotalConsume int  `json:"execTotalConsume"`
		StressTestReq    bool `json:"stressTestReq"`
	} `json:"context"`
	Store struct {
		Client bool `json:"client"`
		Env    struct {
			IsBrowser bool   `json:"isBrowser"`
			IsNode    bool   `json:"isNode"`
			Host      string `json:"host"`
			Href      string `json:"href"`
			Query     struct {
				TechnicianID string `json:"technicianId"`
			} `json:"query"`
			Ua              string `json:"ua"`
			IsBeta          bool   `json:"isBeta"`
			IsLocal         bool   `json:"isLocal"`
			IsX             bool   `json:"isX"`
			IsIOS           bool   `json:"isIOS"`
			IsAndroid       bool   `json:"isAndroid"`
			Os              string `json:"os"`
			IsWX            bool   `json:"isWX"`
			IsWXApp         bool   `json:"isWXApp"`
			IsMtWXApp       bool   `json:"isMtWXApp"`
			IsDpWXApp       bool   `json:"isDpWXApp"`
			IsGroupWXApp    bool   `json:"isGroupWXApp"`
			Version         string `json:"version"`
			Type            string `json:"type"`
			IsDpmerchant    bool   `json:"isDpmerchant"`
			IsMerchantApp   bool   `json:"isMerchantApp"`
			IsDPApp         bool   `json:"isDPApp"`
			IsMTApp         bool   `json:"isMTApp"`
			IsWMApp         bool   `json:"isWMApp"`
			IsKLApp         bool   `json:"isKLApp"`
			IsErpbossproApp bool   `json:"isErpbossproApp"`
			IsWMBApp        bool   `json:"isWMBApp"`
			IsApp           bool   `json:"isApp"`
			IsMT            bool   `json:"isMT"`
			IsMTURL         bool   `json:"isMTUrl"`
			IsRainbow       bool   `json:"isRainbow"`
			UUID            string `json:"uuid"`
			Dpid            string `json:"dpid"`
			Cookie          struct {
			} `json:"cookie"`
			UseVConsole bool `json:"useVConsole"`
			Geo         struct {
			} `json:"geo"`
		} `json:"env"`
	} `json:"store"`
	Geo          bool `json:"geo"`
	ResponseData []struct {
		Code int `json:"code"`
		Data struct {
			TraceID string `json:"traceId"`
			Code    int    `json:"code"`
			Data    struct {
				StatisticValues struct {
				} `json:"statisticValues"`
				HeadActionPoints interface{} `json:"headActionPoints"`
				ExtValues        struct {
					CardInfo struct {
					} `json:"cardInfo"`
				} `json:"extValues"`
				TechCategoryID int `json:"techCategoryId"`
				Share          struct {
					Image string `json:"image"`
					Title string `json:"title"`
					URL   string `json:"url"`
					Desc  string `json:"desc"`
				} `json:"share"`
				ShopIDForFe string `json:"shopIdForFe"`
				AttrValues  struct {
					WorkYearsStr string      `json:"workYearsStr"`
					Skills       []string    `json:"skills"`
					Name         string      `json:"name"`
					JobNumber    int         `json:"jobNumber"`
					PhotoURL     string      `json:"photoUrl"`
					Summary      string      `json:"summary"`
					WorkYears    int         `json:"workYears"`
					JobNumberStr interface{} `json:"jobNumberStr"`
				} `json:"attrValues"`
				TechnicianID      int    `json:"technicianId"`
				ShopCategoryForFe string `json:"shopCategoryForFe"`
			} `json:"data"`
		} `json:"data"`
		Msg string `json:"msg"`
	} `json:"responseData"`
	RequestRepairs []struct {
		Type   string `json:"type"`
		Params []struct {
			TechnicianID string `json:"technicianId"`
		} `json:"params"`
	} `json:"requestRepairs"`
}
