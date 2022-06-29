package logging

type ILogger interface {
	ExceptionLog(logString string)
	InfoLog(logString string)
	FatalLog(logString string)
}
