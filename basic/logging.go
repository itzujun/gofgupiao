package basic
// 2019.02.22
import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

type Position uint

const (
	POSITION_SINGLE     Position = 1
	POSITION_IN_MANAGER Position = 2
)

func init() {
	log.SetFlags(log.LstdFlags)
}

type Logger interface {
	GetPosition() Position
	SetPosition(pos Position)
	Error(v ...interface{}) string
	Errorf(format string, v ...interface{}) string
	Errorln(v ...interface{}) string
	Fatal(v ...interface{}) string
	Fatalf(format string, v ...interface{}) string
	Fatalln(v ...interface{}) string
	Info(v ...interface{}) string
	Infof(format string, v ...interface{}) string
	Infoln(v ...interface{}) string
	Panic(v ...interface{}) string
	Panicf(format string, v ...interface{}) string
	Panicln(v ...interface{}) string
	Warn(v ...interface{}) string
	Warnf(format string, v ...interface{}) string
	Warnln(v ...interface{}) string
}

func getInvokerLocation(skipNumber int) string {
	pc, file, line, ok := runtime.Caller(skipNumber)
	if !ok {
		return ""
	}
	simpleFileName := ""
	if index := strings.LastIndex(file, "/"); index > 0 {
		simpleFileName = file[index+1 : len(file)]
	}
	funcPath := ""
	funcPtr := runtime.FuncForPC(pc)
	if funcPtr != nil {
		funcPath = funcPtr.Name()
	}
	return fmt.Sprintf("%s : (%s:%d)", funcPath, simpleFileName, line)
}

func generateLogContent(
	logTag LogTag,
	pos Position,
	format string,
	v ...interface{}) string {
	skipNumber := int(pos) + 2
	baseInfo :=
		fmt.Sprintf("%s %s - ", logTag.Prefix(), getInvokerLocation(skipNumber))
	var result string
	if len(format) > 0 {
		result = fmt.Sprintf((baseInfo + format), v...)
	} else {
		vLen := len(v)
		params := make([]interface{}, (vLen + 1))
		params[0] = baseInfo
		for i := 1; i <= vLen; i++ {
			params[i] = v[i-1]
		}
		result = fmt.Sprint(params...)
	}
	return result
}

func NewSimpleLogger() Logger {
	logger := &ConsoleLogger{}
	logger.SetPosition(POSITION_SINGLE)
	return logger
}

func NewLogger(loggers []Logger) Logger {
	for _, logger := range loggers {
		logger.SetPosition(POSITION_IN_MANAGER)
	}
	return &LogManager{loggers: loggers}
}

type ConsoleLogger struct {
	position Position
}

func (logger *ConsoleLogger) GetPosition() Position {
	return logger.position
}

func (logger *ConsoleLogger) SetPosition(pos Position) {
	logger.position = pos
}

func (logger *ConsoleLogger) Error(v ...interface{}) string {
	content := generateLogContent(getErrorLogTag(), logger.GetPosition(), "", v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Errorf(format string, v ...interface{}) string {
	content := generateLogContent(getErrorLogTag(), logger.GetPosition(), format, v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Errorln(v ...interface{}) string {
	content := generateLogContent(getErrorLogTag(), logger.GetPosition(), "", v...)
	log.Println(content)
	return content
}

func (logger *ConsoleLogger) Fatal(v ...interface{}) string {
	content := generateLogContent(getFatalLogTag(), logger.GetPosition(), "", v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Fatalf(format string, v ...interface{}) string {
	content := generateLogContent(getFatalLogTag(), logger.GetPosition(), format, v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Fatalln(v ...interface{}) string {
	content := generateLogContent(getFatalLogTag(), logger.GetPosition(), "", v...)
	log.Println(content)
	return content
}

func (logger *ConsoleLogger) Info(v ...interface{}) string {
	content := generateLogContent(getInfoLogTag(), logger.GetPosition(), "", v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Infof(format string, v ...interface{}) string {
	content := generateLogContent(getInfoLogTag(), logger.GetPosition(), format, v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Infoln(v ...interface{}) string {
	content := generateLogContent(getInfoLogTag(), logger.GetPosition(), "", v...)
	log.Println(content)
	return content
}

func (logger *ConsoleLogger) Panic(v ...interface{}) string {
	content := generateLogContent(getPanicLogTag(), logger.GetPosition(), "", v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Panicf(format string, v ...interface{}) string {
	content := generateLogContent(getPanicLogTag(), logger.GetPosition(), format, v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Panicln(v ...interface{}) string {
	content := generateLogContent(getPanicLogTag(), logger.GetPosition(), "", v...)
	log.Println(content)
	return content
}

func (logger *ConsoleLogger) Warn(v ...interface{}) string {
	content := generateLogContent(getWarnLogTag(), logger.GetPosition(), "", v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Warnf(format string, v ...interface{}) string {
	content := generateLogContent(getWarnLogTag(), logger.GetPosition(), format, v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Warnln(v ...interface{}) string {
	content := generateLogContent(getWarnLogTag(), logger.GetPosition(), "", v...)
	log.Println(content)
	return content
}

type LogManager struct {
	loggers []Logger
}

func (logger *LogManager) GetPosition() Position {
	return POSITION_SINGLE
}

func (logger *LogManager) SetPosition(pos Position) {}

func (self *LogManager) Error(v ...interface{}) string {
	var content string
	for _, logger := range self.loggers {
		content = logger.Error(v...)
	}
	return content
}

func (self *LogManager) Errorf(format string, v ...interface{}) string {
	var content string
	for _, logger := range self.loggers {
		content = logger.Errorf(format, v...)
	}
	return content
}

func (self *LogManager) Errorln(v ...interface{}) string {
	var content string
	for _, logger := range self.loggers {
		content = logger.Errorln(v...)
	}
	return content
}

func (self *LogManager) Fatal(v ...interface{}) string {
	var content string
	for _, logger := range self.loggers {
		content = logger.Fatal(v...)
	}
	return content
}

func (self *LogManager) Fatalf(format string, v ...interface{}) string {
	var content string
	for _, logger := range self.loggers {
		content = logger.Fatalf(format, v...)
	}
	return content
}

func (self *LogManager) Fatalln(v ...interface{}) string {
	var content string
	for _, logger := range self.loggers {
		content = logger.Fatalln(v...)
	}
	return content
}

func (self *LogManager) Info(v ...interface{}) string {
	var content string
	for _, logger := range self.loggers {
		content = logger.Info(v...)
	}
	return content
}

func (self *LogManager) Infof(format string, v ...interface{}) string {
	var content string
	for _, logger := range self.loggers {
		content = logger.Infof(format, v...)
	}
	return content
}

func (self *LogManager) Infoln(v ...interface{}) string {
	var content string
	for _, logger := range self.loggers {
		content = logger.Infoln(v...)
	}
	return content
}

func (self *LogManager) Panic(v ...interface{}) string {
	var content string
	for _, logger := range self.loggers {
		content = logger.Panic(v...)
	}
	return content
}

func (self *LogManager) Panicf(format string, v ...interface{}) string {
	var content string
	for _, logger := range self.loggers {
		content = logger.Panicf(format, v...)
	}
	return content
}

func (self *LogManager) Panicln(v ...interface{}) string {
	var content string
	for _, logger := range self.loggers {
		content = logger.Panicln(v...)
	}
	return content
}

func (self *LogManager) Warn(v ...interface{}) string {
	var content string
	for _, logger := range self.loggers {
		content = logger.Warn(v...)
	}
	return content
}

func (self *LogManager) Warnf(format string, v ...interface{}) string {
	var content string
	for _, logger := range self.loggers {
		content = logger.Warnf(format, v...)
	}
	return content
}

func (self *LogManager) Warnln(v ...interface{}) string {
	var content string
	for _, logger := range self.loggers {
		content = logger.Warnln(v...)
	}
	return content
}

const (
	ERROR_LOG_KEY = "ERROR"
	FATAL_LOG_KEY = "FATAL"
	INFO_LOG_KEY  = "INFO"
	PANIC_LOG_KEY = "PANIC"
	WARN_LOG_KEY  = "WARN"
)

type LogTag struct {
	name   string
	prefix string
}

func (self *LogTag) Name() string {
	return self.name
}

func (self *LogTag) Prefix() string {
	return self.prefix
}

var logTagMap map[string]LogTag = map[string]LogTag{
	ERROR_LOG_KEY: LogTag{name: ERROR_LOG_KEY, prefix: "[" + ERROR_LOG_KEY + "]"},
	FATAL_LOG_KEY: LogTag{name: FATAL_LOG_KEY, prefix: "[" + FATAL_LOG_KEY + "]"},
	INFO_LOG_KEY:  LogTag{name: INFO_LOG_KEY, prefix: "[" + INFO_LOG_KEY + "]"},
	PANIC_LOG_KEY: LogTag{name: PANIC_LOG_KEY, prefix: "[" + PANIC_LOG_KEY + "]"},
	WARN_LOG_KEY:  LogTag{name: WARN_LOG_KEY, prefix: "[" + WARN_LOG_KEY + "]"},
}

func getErrorLogTag() LogTag {
	return logTagMap[ERROR_LOG_KEY]
}

func getFatalLogTag() LogTag {
	return logTagMap[FATAL_LOG_KEY]
}

func getInfoLogTag() LogTag {
	return logTagMap[INFO_LOG_KEY]
}

func getPanicLogTag() LogTag {
	return logTagMap[PANIC_LOG_KEY]
}

func getWarnLogTag() LogTag {
	return logTagMap[WARN_LOG_KEY]
}
