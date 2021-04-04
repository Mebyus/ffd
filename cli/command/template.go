package command

import (
	"fmt"
	"strings"
)

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
	if t.Name == "" {
		err = fmt.Errorf("command must have a non-empty name")
		return
	}
	if strings.HasPrefix(t.Name, "-") {
		err = fmt.Errorf("command name [ %s ] cannot start with \"-\"", t.Name)
		return
	}
	if t.Name == "help" || t.Name == "version" {
		err = fmt.Errorf("command name [ %s ] is reserved", t.Name)
		return
	}
	t.flags = map[string]link{}
	for index, flag := range t.BoolFlags {
		l := link{kind: boolFlag, index: index}
		for k := range flag.Aliases {
			if strings.HasPrefix(k, "-") {
				err = fmt.Errorf("bool flag alias [ %s ] cannot start with prefix", k)
				return
			}
			_, ok := t.flags[k]
			if ok {
				err = fmt.Errorf("multiple occurences of [ %s ] alias", k)
				return
			}
			t.flags[k] = l
		}
	}
	for index, flag := range t.ValueFlags {
		l := link{kind: valueFlag, index: index}
		for k := range flag.Aliases {
			if strings.HasPrefix(k, "-") {
				err = fmt.Errorf("value flag alias [ %s ] cannot start with prefix", k)
				return
			}
			_, ok := t.flags[k]
			if ok {
				err = fmt.Errorf("multiple occurences of [ %s ] alias", k)
				return
			}
			t.flags[k] = l
		}
	}
	return
}

func (t *Template) Parse(args []string) (command *Command, err error) {
	command = NewCommand(t.Name)
	tp := newParser(t, command)
	for _, arg := range args {
		err = tp.parseNext(arg)
		if err != nil {
			return
		}
	}
	tp.setDefaults()
	return
}

func (p *parser) parseNext(arg string) (err error) {
	if strings.HasPrefix(arg, "-") {
		flag := strings.TrimPrefix(arg, "-")
		err = p.parseFlag(flag)
		if err != nil {
			return
		}
	} else {
		if p.more {
			p.parseValue(arg)
		} else {
			p.command.Targets = append(p.command.Targets, arg)
		}
	}
	return
}
