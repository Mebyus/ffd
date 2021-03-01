package cli

import "fmt"

type Executor interface {
	Execute() (err error)
}

type SimpleExecutor struct {
	execute func(c *Command) (err error)
	c       *Command
}

// sem - SimpleExecutor Map
var sem = map[string]func(c *Command) (err error){
	"add":      add,
	"download": download,
	"parse":    parse,
	"help":     help,
	"remove":   remove,
	"check":    check,
	"suppress": suppress,
	"list":     list,
	"tidy":     tidy,
	"clean":    clean,
	"bookmark": bookmark,
}

func CreateExecutor(c *Command) (e Executor, err error) {
	if c == nil {
		err = fmt.Errorf("empty command")
		return
	}
	execute, ok := sem[c.Name]
	if !ok {
		err = fmt.Errorf("unknown command \"%s\"", c.Name)
		return
	}
	e = &SimpleExecutor{
		execute: execute,
		c:       c,
	}
	return
}

func (se *SimpleExecutor) Execute() (err error) {
	if se == nil {
		err = fmt.Errorf("empty executor")
		return
	}
	if se.c == nil || se.execute == nil {
		err = fmt.Errorf("uninitialized executor")
		return
	}
	err = se.execute(se.c)
	return
}
