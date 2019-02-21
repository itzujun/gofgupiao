package controller

import (
	"github.com/itzujun/gofgupiao/analyzer"
	"github.com/itzujun/gofgupiao/basic"
	"github.com/itzujun/gofgupiao/downloader"
	"github.com/itzujun/gofgupiao/middleware"
	"sync"
)

var wg sync.WaitGroup

var logger = basic.NewSimpleLogger() // 日志记录器

type Controller struct {
	StartUrl   string                //初始爬行Url
	Downloader downloader.Downloader // 下载器
	Channel    *middleware.Channel   //管道
	WorkPool   *middleware.WorkPool  //工作池
	Parser     analyzer.Analyzer     //解析页函数
}

func NewController(startUrl string, downloader downloader.Downloader,
	channel *middleware.Channel, workPool *middleware.WorkPool, parser analyzer.Analyzer) *Controller {
	return &Controller{StartUrl: startUrl, Downloader: downloader,
		Channel: channel, WorkPool: workPool, Parser: parser}
}

func (ctrl *Controller) Go() {
	basic.Config.StartUrl = ctrl.StartUrl
	basic.InitConfig()

	ctrl.Downloader = downloader.NewDownloader()


}
