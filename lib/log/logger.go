package log

import (
	"bufio"
	"coursesheduling/common"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

var logger *logrus.Logger

func InitLog() {
	absPath,_ := filepath.Abs(filepath.Dir(os.Args[0]))
	logPath := filepath.Join(absPath,common.Course,common.LogPath)
	if _, err := os.Stat(logPath); os.IsNotExist(err){
		os.MkdirAll(logPath,os.ModePerm)
	}
	logName := filepath.Join(logPath,"test.log")
	writer, err := rotatelogs.New(
		logName+".%Y%m%d%H",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		//rotatelogs.WithLinkName(logName),

		// WithRotationTime设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Hour),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		rotatelogs.WithMaxAge(time.Hour*24*15),
		//rotatelogs.WithRotationCount(maxRemainCnt),
	)

	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}


	//if 7 > logLevel && logLevel > 0 {
	//	logrus.SetLevel(logLevel)
	//} else {
		logrus.SetLevel(logrus.DebugLevel)
	//}
	wmap := lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}
    logrus.AddHook(lfshook.NewHook(wmap, &logrus.TextFormatter{DisableColors: true}))

	logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.Hooks.Add(lfshook.NewHook(
		wmap,
		&logrus.TextFormatter{DisableColors: true}),
	)
	src, err:= os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Open Src File err", err)
	}
	fWriter := bufio.NewWriter(src)
	logger.SetOutput(fWriter)
	return
}


func Printf(format string,args ...interface{}) {
	logger.Printf(format,args...)
}

func Print(args ...interface{}) {
	logger.Println(args...)
}

func Fatalf(format string,args ...interface{}){
	logger.Fatalf(format,args...)
}

func Fatal(args ...interface{}){
	logger.Fatalln(args...)
}

func Errorf(format string,args ...interface{}){
	logger.Errorf(format,args...)
}

func Error(args ...interface{}){
	logger.Errorln(args...)
}

func Debugf(format string,args ...interface{}){
	logger.Debugf(format,args...)
}

func Debug(args ...interface{}){
	logger.Debugln(args...)
}

func Infof(format string,args ...interface{}){
	logger.Infof(format,args...)
}

func Info(args ...interface{}){
	logger.Infoln(args...)
}