package middleware

import (
	"github.com/itzujun/gofgupiao/analyzer"
	"github.com/itzujun/gofgupiao/basic"
)

type Channel struct {
	reqpchan  chan basic.Request     //请求
	respchan  chan basic.Response    //结果
	sharechan chan []analyzer.Shares //股票infos
	ch        chan analyzer.Shares   //单一股票info
}

func NewChannel() *Channel {
	return &Channel{
		//make(chan basic.Request, basic.Config.RequestNum),
		make(chan basic.Request, 1),
		//make(chan basic.Response, basic.Config.RequestNum),
		make(chan basic.Response, 1),
		make(chan []analyzer.Shares, 10),
		make(chan analyzer.Shares, 10),
	}
}

func (this *Channel) ReqChan() chan basic.Request {
	return this.reqpchan
}

func (this *Channel) RespChan() chan basic.Response {
	return this.respchan
}

func (this *Channel) RespShares() chan []analyzer.Shares {
	return this.sharechan
}

func (this *Channel) RespCh() chan analyzer.Shares {
	return this.ch
}
