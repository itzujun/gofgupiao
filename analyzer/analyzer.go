package analyzer

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/itzujun/gofgupiao/analyzer"
	"github.com/itzujun/gofgupiao/util"
	"io/ioutil"
	"net/http"
	"strings"
)

//type Shares struct {
//	Name    string
//	Code    string
//	Url     string
//	ApiCode string
//}
//
//type SharesRes struct {
//	Name     string
//	Code     string
//	Open     string
//	High     string
//	Close    string
//	Volume   string
//	PreClose string
//}

type GenAnalyzer interface {
	AnalyzeHtml(httpRes *http.Response) []Shares
	AnalyzeApi(httpRes *http.Response, shares analyzer.Shares) SharesRes
}

type Analyzer struct {
	GenAnalyzer
}

func NewAnalyzer() GenAnalyzer {
	return new(Analyzer)
}

//Api解析
func (self *Analyzer) AnalyzeApi(httpResp *http.Response, shares analyzer.Shares) SharesRes {
	shRes := SharesRes{}
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
	if kVal, ok := kline.(map[string]interface{}); ok {
		fmt.Println(shares.Name, shares.Code, kVal["open"], kVal["high"], kVal["open"], kVal["close"], kVal["volume"], kVal["preClose"])
		shRes = SharesRes{Name: shares.Name, Code: shares.Code}
	}
	return shRes
}

//用于解析页面
func (self *Analyzer) AnalyzeHtml(httpRes *http.Response) []Shares {
	fmt.Println("解析网页...")
	defer httpRes.Body.Close()
	sh := []Shares{}
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
				sh = append(sh, Shares{recv[0], recv[1], url, ApiCode})
			}
		}
	})
	return sh
}
