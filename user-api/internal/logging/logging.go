package logging

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type ILogger interface {
	InfoLog(logString string)
	ExceptionLog(logString string)
	WarningLog(logString string)
	FatalLog(logString string)
}

func initLogger(dirName string) error {
	if _, err := os.ReadDir(fmt.Sprintf(`logs/%s`, dirName)); err != nil {
		if err := os.Mkdir(fmt.Sprintf(`log/%s`, dirName), 0700); err != nil {
			return err
		}
		if err := os.WriteFile(fmt.Sprintf(`logs/%s/info.log`, dirName), nil, 0700); err != nil {
		}
		if err := os.WriteFile(fmt.Sprintf(`logs/%s/exception.log`, dirName), nil, 0700); err != nil {
		}
		if err := os.WriteFile(fmt.Sprintf(`logs/%s/warning.log`, dirName), nil, 0700); err != nil {
		}
		if err := os.WriteFile(fmt.Sprintf(`logs/%s/fatal.log`, dirName), nil, 0700); err != nil {
		}
	}
	return nil
}

type Logger struct {
	SaveLogsDir     string
	infoLogger      *log.Logger
	exceptionLogger *log.Logger
	warningLogger   *log.Logger
	fatalLogger     *log.Logger
}

func (l *Logger) InfoLog(logString string) {
	d, err := os.Open(fmt.Sprintf(`%s/info.log`, l.SaveLogsDir))
	if err != nil {
		return
	}
	l.infoLogger.SetOutput(d)
	l.infoLogger.Info(logString)
}
func (l *Logger) ExceptionLog(logString string) {
	d, err := os.Open(fmt.Sprintf(`%s/exception.log`, l.SaveLogsDir))
	if err != nil {
		return
	}
	l.exceptionLogger.SetOutput(d)
	l.exceptionLogger.Error(logString)
}
func (l *Logger) WarningLog(logString string) {
	d, err := os.Open(fmt.Sprintf(`%s/warning.log`, l.SaveLogsDir))
	if err != nil {
		return
	}
	l.warningLogger.SetOutput(d)
	l.warningLogger.Warning(logString)
}
func (l *Logger) FatalLog(logString string) {
	d, err := os.Open(fmt.Sprintf(`%s/fatal.log`, l.SaveLogsDir))
	if err != nil {
		return
	}
	l.fatalLogger.SetOutput(d)
	l.fatalLogger.Fatal(logString)
}

func NewLogger(saveLogsDir string) ILogger {
	if err := initLogger(saveLogsDir); err != nil {
		return nil
	}
	return &Logger{
		SaveLogsDir:     saveLogsDir,
		infoLogger:      log.StandardLogger(),
		exceptionLogger: log.StandardLogger(),
		warningLogger:   log.StandardLogger(),
		fatalLogger:     log.StandardLogger(),
	}
}
