package controller

import (
	"fmt"
	"github.com/itzujun/gofgupiao/analyzer"
	"github.com/itzujun/gofgupiao/basic"
	"github.com/itzujun/gofgupiao/downloader"
	"github.com/itzujun/gofgupiao/middleware"
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

	fmt.Println("req---: ", prereq)
	logger.Info("开始下载")

	wg.Add(3)
	go ctrl.FirstDown()
	go ctrl.FirstAnalyzer()
	wg.Wait()
	fmt.Println("aaa---")
	SSS := ctrl.Channel.RespShares()
	resp := <-SSS
	fmt.Print("pppppppppppppp---------")
	for _, ch := range resp {
		fmt.Println("ch:", ch)
	}

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
			fmt.Println("访问成功:", res)
			ctrl.Channel.RespChan() <- *res
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
		//for _, ch := range resp {
		//	fmt.Println("ch:", ch)
		//}
		awg.Done()
	}()
	awg.Wait()
	fmt.Println("end-----")
	close(ctrl.Channel.RespShares())
}
