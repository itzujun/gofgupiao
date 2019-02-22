package res

import "strings"

type Shares struct {
	Name    string
	Code    string
	Url     string
	ApiCode string
}

type SharesRes struct {
	Name     string
	Code     string
	Open     string
	High     string
	Close    string
	Volume   string
	PreClose string
}

func NewShares() *Shares {
	return &Shares{}
}

func NewSharesRes() *SharesRes {
	return &SharesRes{}
}

func (sh *Shares) GetLinkCode() string {
	liCode := strings.Split(sh.Url, "/")
	return strings.Split(liCode[len(liCode)-1], ".")[0]

}
