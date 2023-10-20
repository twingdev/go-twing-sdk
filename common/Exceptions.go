package common

type IError interface {
	ThrowRecoverableError(msg string)
	ThrowFatalError(msg string)
	ThrowNotImplemented(msg string)
}

type Error struct {
	msg string
}

func (err Error) ThrowRecoverableError(msg string) {
	l := GetLogger()
	l.Warn(msg)

}

func (err Error) ThrowFatalError(msg string) {
	l := GetLogger()
	l.Fatal(msg)
}

func (err Error) ThrowNotImplemented(msg string) {
	err.ThrowFatalError(msg)
}

func InitExceptionManager() Error {
	return Error{}
}

var ExceptionManager Error = InitExceptionManager()
