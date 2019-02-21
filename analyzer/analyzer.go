package analyzer

import (
	"github.com/pikez/Scrago/basic"
	"net/http"
)

// 接收数据并分析返回结果

type Parser func(httpRes *http.Response) ([]string, []basic.Item)

type GenAnalyzer interface {
	Analyze(httpRes *http.Response, parser Parser) ([]string, []basic.Item)
}

type Analyzer struct {
	linklist []string
	itemlist []basic.Item
}

func NewAnalyzer() GenAnalyzer {
	return &Analyzer{
		make([]string, 0),
		make([]basic.Item, 0),
	}

}

//用于解析页面
func (self *Analyzer) Analyze(httpRes *http.Response, parser Parser) ([]string, []basic.Item) {
	defer httpRes.Body.Close()
	if parser == nil {
		panic("xxx")
	}
	return parser(httpRes)
}
