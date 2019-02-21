package downloader

import (
	"github.com/itzujun/gofgupiao/basic"
	"net/http"
)

type Downloader struct {
	client *http.Client
}

func NewDownloader() *Downloader {
	return &Downloader{}
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
