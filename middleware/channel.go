package middleware

import (
	"github.com/itzujun/gofgupiao/basic"
	"github.com/itzujun/gofgupiao/res"
)

type Channel struct {
	reqpchan  chan basic.Request  //请求
	respchan  chan basic.Response //结果
	sharechan chan []res.Shares   //股票infos
	ch        chan res.Shares     //单一股票info
}

func NewChannel() *Channel {
	return &Channel{
		make(chan basic.Request, 1),
		make(chan basic.Response, 1),
		make(chan []res.Shares, 10),
		make(chan res.Shares, 10),
	}
}

func (this *Channel) ReqChan() chan basic.Request {
	return this.reqpchan
}

func (this *Channel) RespChan() chan basic.Response {
	return this.respchan
}

func (this *Channel) RespShares() chan []res.Shares {
	return this.sharechan
}

func (this *Channel) RespCh() chan res.Shares {
	return this.ch
}
