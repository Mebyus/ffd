package command

var stdDispatcher = NewDispatcher()

func Register(template *Template, executor Executor) {
	stdDispatcher.Register(template, executor)
}

func Dispatch(args []string) (err error) {
	return stdDispatcher.Dispatch(args)
}

func init() {
	// Register(NewStdHelpTemplate(), NewStdHelper())
}
