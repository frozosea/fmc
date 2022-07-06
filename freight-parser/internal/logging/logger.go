package logging

type ILogger interface {
	InfoLog(logString string)
	ExceptionLog(logString string)
	WarningLog(logString string)
	FatalLog(logString string)
}

//TODO create logger
type logger struct {
}

func NewLogger() *logger {
	return &logger{}
}
