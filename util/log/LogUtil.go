package log

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"time"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"path"
	"com.dingjinlin.getBing/util/file"
)

const logDir = "logs"
const logName = "log"
const MaxAge = time.Hour * 24 * 7
const RotationTime = time.Hour * 24

var log *logrus.Logger

func GetLoggerInstance() *logrus.Logger {
	if log == nil {
		log = configLocalFilesystemLogger(logDir, logName, MaxAge, RotationTime)
	}

	return log
}

func configLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) *logrus.Logger {
	filePath,err := file.CreateDir(logPath)
	baseLogPath := path.Join(filePath, logFileName)
	CheckError(err)
	logFilePath := baseLogPath + ".%Y%m%d%H%M"

	writer, err := rotatelogs.New(
		logFilePath,
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	pathMap := lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}
	lfHook := lfshook.NewHook(pathMap, &logrus.JSONFormatter{}, )
	var Log *logrus.Logger
	Log = logrus.New()
	Log.Hooks.Add(lfHook)

	return Log
}

func CheckError(err error) {
	if err != nil {
		log := GetLoggerInstance()
		log.Error(err)
	}
}
