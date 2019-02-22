package analyzer

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/itzujun/gofgupiao/res"
	"github.com/itzujun/gofgupiao/util"
	"io/ioutil"
	"net/http"
	"strings"
)

type GenAnalyzer interface {
	AnalyzeHtml(httpRes *http.Response) []res.Shares
	AnalyzeApi(httpRes *http.Response, shares res.Shares) *res.SharesRes
}

type Analyzer struct {
	GenAnalyzer
}

func NewAnalyzer() GenAnalyzer {
	return &Analyzer{}
}

//Api解析
func (self *Analyzer) AnalyzeApi(httpResp *http.Response, shares res.Shares) *res.SharesRes {
	fmt.Println("解析Api...")
	//shRes := res.SharesRes{}
	shRes := new(res.SharesRes)
	respstream, _ := ioutil.ReadAll(httpResp.Body)
	recpmap := make(map[string]interface{})
	err := json.Unmarshal(respstream, &recpmap)
	data, ok := recpmap["mashData"]
	if err != nil || ok == false {
		return shRes
	}
	value, _ := data.([]interface{})
	val, _ := value[0].(map[string]interface{})
	kline, _ := val["kline"]
	fmt.Println("kline:", kline)
	if kVal, ok := kline.(map[string]interface{}); ok {
		fmt.Println(shares.Name, shares.Code, kVal["open"], kVal["high"], kVal["open"], kVal["close"], kVal["volume"], kVal["preClose"])
		shRes = &res.SharesRes{Name: shares.Name, Code: shares.Code}
	}
	return shRes
}

//用于解析页面
func (self *Analyzer) AnalyzeHtml(httpRes *http.Response) []res.Shares {
	fmt.Println("解析网页...")
	defer httpRes.Body.Close()
	sh := []res.Shares{}
	doc, _ := goquery.NewDocumentFromReader(httpRes.Body)
	doc.Find("div.quotebody li").Each(func(i int, s *goquery.Selection) {
		band := s.Find("a").Text()
		if url, exists := s.Find("a").Attr("href"); exists {
			band = util.ConvertToString(band, "gbk", "utf-8")
			band = strings.Replace(band, ")", "", -1)
			recv := strings.Split(band, "(")
			liCode := strings.Split(url, "/")
			ApiCode := strings.Split(liCode[len(liCode)-1], ".")[0]
			if strings.HasPrefix(ApiCode, "sz300") {
				sh = append(sh, res.Shares{recv[0], recv[1], url, ApiCode})
			}
		}
	})
	return sh
}
