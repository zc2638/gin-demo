package logger

import (
	"fmt"
	"github.com/zctod/tool/common/logs"
	"log"
	"os"
	"time"
)

const PATH_LOG = "logs/"

var logger = log.New(nil, "", log.Ldate|log.Ltime|log.Lshortfile)

var loggerSts = map[string][]interface{}{
	"info": {"", nil},
}

func getCurrentPath() string {
	return PATH_LOG + time.Now().Format("2006-01-02")
}

func Info(v... interface{}) {
	if loggerSts["info"][0].(string) != getCurrentPath() {
		if loggerSts["info"][1] != nil {
			_ = loggerSts["info"][1].(*os.File).Close()
		}
		f := logs.ReadPath(getCurrentPath())
		logger.SetOutput(f)
		loggerSts["info"][1] = f
	}
	logger.SetPrefix("[INFO]")
	_ = logger.Output(2, fmt.Sprintln(v...))
}
