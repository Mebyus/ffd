package command

import "strings"

type AliasType uint8
type flagKind uint8

const (
	NotAnAlias AliasType = iota
	SingleChar
	MultipleChars
)

const (
	notAFlag flagKind = iota
	boolFlag
	valueFlag
)

type Flag struct {
	Aliases     map[string]AliasType
	Description string
}

type BoolFlag struct {
	Flag
	Default bool
}

type ValueFlag struct {
	Flag
	Default string
}

type link struct {
	kind  flagKind
	index int
}

type Template struct {
	Name        string
	Description string
	BoolFlags   []BoolFlag
	ValueFlags  []ValueFlag

	flags map[string]link
}

func (t *Template) prepare() (err error) {
	t.flags = map[string]link{}
	for index, flag := range t.BoolFlags {
		l := link{kind: boolFlag, index: index}
		for k := range flag.Aliases {
			t.flags[k] = l
		}
	}
	for index, flag := range t.ValueFlags {
		l := link{kind: valueFlag, index: index}
		for k := range flag.Aliases {
			t.flags[k] = l
		}
	}
	return
}

// func (t *Template) getFlagLink(flag string) (l link, ok bool) {

// }

func (t *Template) Parse(args []string) (command *Command) {
	_ = t.prepare()
	command = NewCommand(t.Name)
	tp := parser{template: t, command: command}
	more := false
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flag := strings.TrimPrefix(arg, "-")
			more = tp.parseFlag(flag)
		} else {
			if more {
				tp.parseValue(arg)
			} else {
				command.Targets = append(command.Targets, arg)
			}
		}
	}
	return
}
