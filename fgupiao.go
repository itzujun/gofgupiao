package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

type GInfo struct {
	name string //股票名字
	code string //股票代码
	url  string //股票数据访问网址
}

// 股票爬虫详情
type GResult struct {
	name     string
	code     string
	open     interface{} //开盘
	high     interface{} //最高
	low      interface{} //最低
	close    interface{} //收盘
	volume   interface{} //成交
	preClose interface{} //昨收
}

var baseurl = "http://quote.eastmoney.com/stocklist.html"

var li = list.New()

func startSpider() bool {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", baseurl, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1573.2 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("error:", err.Error())
		return false
	}
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	doc.Find("div.quotebody li").Each(func(i int, s *goquery.Selection) {
		band := s.Find("a").Text()
		if url, exists := s.Find("a").Attr("href"); exists {
			band = ConvertToString(band, "gbk", "utf-8")
			band = strings.Replace(band, ")", "", -1)
			recv := strings.Split(band, "(")
			info := GInfo{recv[0], recv[1], url}
			fmt.Println(info.url, info.name, info.code)
			li.PushBack(info)
		}
	})
	for e := li.Front(); e != nil; e = e.Next() {
		if value, isok := e.Value.(GInfo); isok {
			fmt.Println(value.name, value.code, value.url)
		}
	}
	return true
}

//数据爬虫
func spidDetail(info GInfo) {
	msec := time.Now().UnixNano() / 1e9
	stime := strconv.FormatInt(msec, 10) //时间戳
	liCode := strings.Split(info.url, "/")
	scode := strings.Split(liCode[len(liCode)-1], ".")[0]
	fmt.Println(scode)
	linkurl := "https://gupiao.baidu.com/api/stocks/stockdaybar?from=pc&os_ver=1&cuid=xxx&vv=100&format=json&stock_code=" +
		scode + "&step=3&start=&count=160&fq_type=no&timestamp=" + stime
	client := &http.Client{}
	req, _ := http.NewRequest("GET", linkurl, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1573.2 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("error:", err.Error())
		return
	}
	if resp.StatusCode == 200 {
		respstream, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("resp:", string(respstream))
		recpmap := make(map[string]interface{})
		err = json.Unmarshal(respstream, &recpmap)

		data, ok := recpmap["mashData"]
		if ok == false { //停牌股票不包含数据
			return
		}
		value, _ := data.([]interface{})
		val, _ := value[0].(map[string]interface{})
		kline, _ := val["kline"]

		if kVal, ok := kline.(map[string]interface{}); ok {
			gres := GResult{info.name, info.code, kVal["open"], kVal["high"], kVal["open"],
				kVal["close"], kVal["volume"], kVal["preClose"]}
			fmt.Println("gres:", gres)
			fmt.Println("爬虫成功:", info.name)
		}
	}

}

func main() {

	isOK := startSpider()
	if isOK {
		for e := li.Front(); e != nil; e = e.Next() {
			if value, ok := e.Value.(GInfo); ok {
				fmt.Println(value)
				src := value.url
				lirecp := strings.Split(src, "/")
				res := strings.Split(lirecp[len(lirecp)-1], ".")[0]
				if strings.HasPrefix(res, "sz300") {
					spidDetail(value)

				}
			}
		}
	}

}
