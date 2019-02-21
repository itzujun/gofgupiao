package controller

import (
	"fmt"
	"github.com/itzujun/gofgupiao/analyzer"
	"github.com/itzujun/gofgupiao/basic"
	"github.com/itzujun/gofgupiao/downloader"
	"github.com/itzujun/gofgupiao/middleware"
	"github.com/itzujun/gofgupiao/util"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

var logger = basic.NewSimpleLogger() // 日志记录器

type Controller struct {
	StartUrl   string                   //初始爬行Url
	Downloader downloader.GenDownloader // 下载器
	Channel    *middleware.Channel      //管道
	WorkPool   *middleware.WorkPool     //工作池
	Parser     analyzer.GenAnalyzer     //解析页函数
}

func NewController(startUrl string) *Controller {
	return &Controller{StartUrl: startUrl}
}

func (ctrl *Controller) Go() {
	basic.Config.StartUrl = ctrl.StartUrl
	basic.InitConfig()
	ctrl.Downloader = downloader.NewDownloader()
	ctrl.Channel = middleware.NewChannel()
	ctrl.WorkPool = middleware.NewWorkPool()
	ctrl.Parser = analyzer.NewAnalyzer()
	prereq, err := http.NewRequest(basic.Config.RequestMethod, basic.Config.StartUrl, nil)
	if err != nil {
		return
	}
	basereq := basic.NewRequest(prereq, 0)
	ctrl.Channel.ReqChan() <- *basereq
	wg.Add(2)
	go ctrl.FirstDown()
	go ctrl.FirstAnalyzer()
	wg.Wait()

	respshares := ctrl.Channel.RespShares()
	resp := <-respshares

	for _, ch := range resp {
		fmt.Println("ch:", ch)
	}

}

func (ctrl *Controller) FeedDown(task chan analyzer.Shares, chs []analyzer.Shares) { //添加任务

	for _, req := range chs {
		task <- req
	}
}

//func (ctrl *Controller) DoDown(chs chan analyzer.Shares) { //执行任务
func (ctrl *Controller) DoDown(ch chan analyzer.Shares) { //执行任务
	for {
		shares, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("info", shares)
		linkurl := "https://gupiao.baidu.com/api/stocks/stockdaybar?from=pc&os_ver=1&cuid=xxx&vv=100&format=json&stock_code=" +
			shares.ApiCode + "&step=3&start=&count=160&fq_type=no&timestamp=" + util.GetTimeStap()
		fmt.Println(linkurl)
		req, err := http.NewRequest(basic.Config.RequestMethod, linkurl, nil)
		if err != nil {
			break
		}
		basereq := basic.NewRequest(req, 0)
		resp := ctrl.Downloader.Download(basereq)
		if resp.GetRes().StatusCode != 200 {
			continue
		}
		info := ctrl.Parser.AnalyzeApi(resp.GetRes(), shares)
		fmt.Println("info:", info)
	}

	//linkurl := "https://gupiao.baidu.com/api/stocks/stockdaybar?from=pc&os_ver=1&cuid=xxx&vv=100&format=json&stock_code=" +
	//	info.GetLinkCode(sh) + "&step=3&start=&count=160&fq_type=no&timestamp=" + info.GetTimeStack()
	//req, err := http.NewRequest("GET", linkurl, nil)
	//resp, err := info.client.Do(req)
	//if err != nil || resp.StatusCode != 200 {
	//	fmt.Println("error:", err.Error())
	//	return
	//}
	//respstream, _ := ioutil.ReadAll(resp.Body)
	//recpmap := make(map[string]interface{})
	//err = json.Unmarshal(respstream, &recpmap)
	//data, ok := recpmap["mashData"]
	//if ok == false { //停牌股票不包含数据
	//	return
	//}
	//value, _ := data.([]interface{})
	//val, _ := value[0].(map[string]interface{})
	//kline, _ := val["kline"]
	//if kVal, ok := kline.(map[string]interface{}); ok {
	//	fmt.Println(sh.name, sh.code, kVal["open"], kVal["high"], kVal["open"], kVal["close"], kVal["volume"], kVal["preClose"])
	//}
	//fmt.Println("爬虫ok:", sh.name)

}

func (ctrl *Controller) GoDowndetail() {
	defer wg.Done()
	fmt.Println("GoDown...")

}

func (ctrl *Controller) FirstDown() {
	logger.Info("FirstDown...")
	defer wg.Done()
	dwg := new(sync.WaitGroup)
	dwg.Add(1)
	go func() {
		req := <-ctrl.Channel.ReqChan()
		res := ctrl.Downloader.Download(&req)
		if res != nil {
			ctrl.Channel.RespChan() <- *res //访问成功
		}
		dwg.Done()
	}()
	dwg.Wait()
	go func() {}()
	close(ctrl.Channel.RespChan())
}

func (ctrl *Controller) FirstAnalyzer() {
	fmt.Print("FirstAnalyzer")
	defer wg.Done()
	awg := new(sync.WaitGroup)
	awg.Add(1)
	go func() {
		res := <-ctrl.Channel.RespChan()
		resp := ctrl.Parser.AnalyzeHtml(res.GetRes())
		fmt.Println("解析结果:", resp)
		ctrl.Channel.RespShares() <- resp
		awg.Done()
	}()
	awg.Wait()
	fmt.Println("end-----")
	close(ctrl.Channel.RespShares())
}
