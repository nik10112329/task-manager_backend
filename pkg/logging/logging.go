package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type writerHook struct {
	Writter   []io.Writer
	LogLevels []logrus.Level
}
type Logger struct {
	*logrus.Entry
}

var e *logrus.Entry

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writter {
		w.Write([]byte(line))
	}
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func GetLogger() *Logger {
	return &Logger{e}
}

func (logger *Logger) GetLoggerWithField(k string, v interface{}) *Logger {
	return &Logger{logger.WithField(k, v)}
}

func init() {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", fileName, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}
	err := os.Mkdir("logs", 0644)
	if err != nil {
		if !os.IsExist(err) {
			panic(err)
		}
	}

	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)

	if err != nil {
		panic(err)
	}

	logger.SetOutput(io.Discard)

	logger.AddHook(&writerHook{
		Writter:   []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	logger.SetLevel(logrus.TraceLevel)

	e = logger.WithFields(logrus.Fields{})
}
