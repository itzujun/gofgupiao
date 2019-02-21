package middleware

import (
	"github.com/itzujun/gofgupiao/basic"
)

type Channel struct {
	reqpchan chan basic.Request  //请求
	respchan chan basic.Response //结果
}

func NewChannel() *Channel {
	return &Channel{
		make(chan basic.Request, basic.Config.RequestNum),
		make(chan basic.Response, basic.Config.RequestNum),
	}
}

func (this *Channel) ReqChan() chan basic.Request {
	return this.reqpchan
}

func (this *Channel) RespChan() chan basic.Response {
	return this.respchan
}
