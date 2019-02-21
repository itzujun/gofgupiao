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

func NewController(startUrl string, downloader *downloader.Downloader,
	channel *middleware.Channel, workPool *middleware.WorkPool, parser analyzer.GenAnalyzer) *Controller {
	return &Controller{StartUrl: startUrl, Downloader: downloader,
		Channel: channel, WorkPool: workPool, Parser: parser}
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
	//basereq := basic.NewRequest(prereq, 0)
	//ctrl.Channel.ReqChan() <- *basereq
	fmt.Println("req:", prereq)

	wg.Add(1)

}

func (ctrl *Controller) FirstDown() {
	defer wg.Done()
	dwg := new(sync.WaitGroup)
	dwg.Add(1)
	ctrl.WorkPool.Pool(1, func() {
		for req := range ctrl.Channel.ReqChan() {
			res := ctrl.Downloader.Download(&req)
			if res != nil {
				ctrl.Channel.RespChan() <- *res
			}
		}
		dwg.Done()
	})
	dwg.Wait()
	close(ctrl.Channel.RespChan())
}

func (ctrl *Controller) FirstAnalyzer() {
	defer wg.Done()
	awg := new(sync.WaitGroup)
	awg.Add(1)
	ctrl.WorkPool.Pool(1, func() {
		for res := range ctrl.Channel.RespChan() {
			// 解析html页面
			resp := ctrl.Parser.AnalyzeHtml(res.GetRes())
			ctrl.Channel.RespChan() <- resp
		}
	})
}
