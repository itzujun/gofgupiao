package basic

import "net/http"

type Request struct {
	httpReq *http.Request
	index   uint32
}

func NewRequest(httpReq *http.Request, index uint32) *Request {
	return &Request{httpReq: httpReq}
}

func (req *Request) GetReq() *http.Request {
	return req.httpReq
}

func (req *Request) GetIndex() uint32 {
	return req.index
}
