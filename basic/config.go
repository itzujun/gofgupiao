package basic

type config struct {
	flag          bool              //是否初始化配置
	Name          string            //爬虫项目名字
	StartUrl      string            //开始访问url
	HttpHead      map[string]string //请求头
	RequestNum    int               //请求数目长度
	RequestMethod string            //爬虫方式
}

var Config *config = &config{flag: true}

func InitConfig() {
	if Config.flag == false {
		return
	}
	Config.HttpHead = make(map[string]string)
	if Config.Name == "" {
		Config.Name = "gofgupiao"
	}
	if Config.RequestNum == 0 {
		Config.RequestNum = 5
	}
	if Config.RequestMethod == "" {
		Config.RequestMethod = "GET"
	}
	Config.flag = true
}
