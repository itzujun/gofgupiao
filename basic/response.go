package basic
// 2019.02.22
import "net/http"

type Response struct {
	httpRes *http.Response
	index   uint32
}

func NewResponse(httpRes *http.Response, index uint32) *Response {
	return &Response{httpRes: httpRes, index: index}
}

func (res *Response) GetRes() *http.Response {
	return res.httpRes
}

func (res *Response) GetIndex() uint32 {
	return res.index
}
