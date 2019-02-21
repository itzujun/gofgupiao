package analyzer

import (
	"github.com/pikez/Scrago/basic"
	"net/http"
)

// 接收数据并分析返回结果

type Parser func(httpRes *http.Response) ([]string, []basic.Item)

type GenAnalyzer interface {
	AnalyzeHtml(httpRes *http.Response, parser Parser) ([]string, []basic.Item)
}

type Analyzer struct {
	GenAnalyzer
}

func NewAnalyzer() GenAnalyzer {
	return new(Analyzer)
}

func Anaybase() {

}

//Api解析
func (self *Analyzer) AnalyzeApi(httpRes *http.Response) string {
	return "结果"
}

//用于解析页面
func (self *Analyzer) AnalyzeHtml(httpRes *http.Response, parser Parser) ([]string, []basic.Item) {
	defer httpRes.Body.Close()
	if parser == nil {
		panic("xxx")
	}
	return parser(httpRes)
}
