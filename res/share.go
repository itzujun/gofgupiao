package res


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