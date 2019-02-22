package util
// 2019.02.22
import (
	"github.com/axgle/mahonia"
	"strconv"
	"time"
)

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func GetTimeStap() string {
	msec := time.Now().UnixNano() / 1e9
	return strconv.FormatInt(msec, 10) //时间戳
}
