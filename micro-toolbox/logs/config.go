package logs

import (
	"io"
	"log"
	"os"
	"time"
)

var (
	Log * log.Logger
)

func init(){
	logFile,err:=os.OpenFile("logs.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil {
	log.Fatalln("打开日志文件失败：", err)
	}
	Log = log.New(io.MultiWriter(os.Stderr,logFile),"Log:",log.Ldate | log.Ltime | log.Lshortfile)
	Log.Println("日志开始",time.Now())
}