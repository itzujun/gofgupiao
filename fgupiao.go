package main

import (
	"github.com/itzujun/gofgupiao/controller"
	"fmt"
)

func main() {
	url := "http://quote.eastmoney.com/stocklist.html"
	fmt.Println("url:", url)
	ctrl := controller.NewController(url)
	ctrl.Go()
}
