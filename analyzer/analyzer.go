package analyzer

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/itzujun/gofgupiao/util"
	"net/http"
	"strings"
)

// 接收数据并分析返回结果

type Shares struct {
	Name    string //股票名字
	Code    string //股票代码
	Url     string //访问api网页地址
	ApiCode string //股票访问diamante
}

type GenAnalyzer interface {
	AnalyzeHtml(httpRes *http.Response) []Shares
	AnalyzeApi(httpRes *http.Response) string
}

type Analyzer struct {
	GenAnalyzer
}

func NewAnalyzer() GenAnalyzer {
	return new(Analyzer)
}

//Api解析
func (self *Analyzer) AnalyzeApi(httpRes *http.Response) string {
	return "结果"
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
			sh = append(sh, Shares{recv[0], recv[1], url, ApiCode})
		}
	})
	return sh
}
