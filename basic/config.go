package basic
// 2019.02.22
type config struct {
	flag          bool
	Name          string
	StartUrl      string
	HttpHead      map[string]string
	RequestNum    int
	RequestMethod string
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
		Config.RequestNum = 50
	}
	if Config.RequestMethod == "" {
		Config.RequestMethod = "GET"
	}
	Config.flag = true
}
