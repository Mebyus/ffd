package command

import "fmt"

var stdDispatcher = NewDispatcher()

func Register(template *Template, executor Executor) {
	stdDispatcher.Register(template, executor)
}

func Dispatch(args []string) (err error) {
	return stdDispatcher.Dispatch(args)
}

func SetVersion(version fmt.Stringer) {
	stdDispatcher.SetVersion(version)
}
