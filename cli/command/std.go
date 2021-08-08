package command

import "fmt"

var stdDispatcher = NewDispatcher("", "", "")

func Init(name, version, description string) {
	SetName(name)
	SetStringVersion(version)
	SetDescription(description)
}

func Register(template *Template, executor Executor) {
	stdDispatcher.Register(template, executor)
}

func Dispatch(args []string) (err error) {
	return stdDispatcher.Dispatch(args)
}

func SetVersion(version fmt.Stringer) {
	stdDispatcher.SetVersion(version)
}

func SetStringVersion(version string) {
	stdDispatcher.SetStringVersion(version)
}

func SetName(name string) {
	stdDispatcher.SetName(name)
}

func SetDescription(description string) {
	stdDispatcher.SetDescription(description)
}
