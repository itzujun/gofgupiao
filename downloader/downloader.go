package downloader

import (
	"github.com/itzujun/gofgupiao/basic"
	"net/http"
)

type GenDownloader interface {
	Download(req *basic.Request) *basic.Response
}

type Downloader struct {
	client *http.Client
}

func NewDownloader() GenDownloader {
	return &Downloader{&http.Client{}}
}

func (this *Downloader) Download(req *basic.Request) *basic.Response {
	for k, v := range basic.Config.HttpHead {
		req.GetReq().Header.Set(k, v)
	}
	httpRes, err := this.client.Do(req.GetReq())
	if err != nil {
		return nil
	}
	return basic.NewResponse(httpRes, req.GetIndex())
}
