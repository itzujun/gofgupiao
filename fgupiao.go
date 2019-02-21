package main

import "github.com/itzujun/gofgupiao/controller"

func main() {

	url := "http://quote.eastmoney.com/stocklist.html"
	controller.NewController(url)

}
