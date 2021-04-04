package command

import (
	"fmt"
	"strings"
)

type parser struct {
	template      *Template
	command       *Command
	lastLink      link
	more          bool
	boolFlagUsed  []bool
	valueFlagUsed []bool
}

func newParser(template *Template, command *Command) *parser {
	return &parser{
		template:      template,
		command:       command,
		boolFlagUsed:  make([]bool, len(template.BoolFlags)),
		valueFlagUsed: make([]bool, len(template.ValueFlags)),
	}
}

func (p *parser) parseFlag(flag string) (err error) {
	if strings.HasPrefix(flag, "-") {
		p.more, err = p.parseMultiCharFlag(strings.TrimPrefix(flag, "-"))
	} else {
		if len(flag) == 1 {
			p.more = p.parseSingleCharFlag(flag)
		} else {
			p.parseSetOfSingleCharFlags(flag)
		}
	}

	return
}

func (p *parser) parseSetOfSingleCharFlags(set string) {
	for _, flag := range strings.Split(set, "") {
		p.parseSingleCharFlag(flag)
	}
}

func (p *parser) parseSingleCharFlag(flag string) (more bool) {
	l, ok := p.template.flags[flag]
	if ok {
		if l.kind == boolFlag {
			p.command.setBool(p.template.BoolFlags[l.index])
			p.boolFlagUsed[l.index] = true
		} else if l.kind == valueFlag {
			more = true
			p.lastLink = l
		}
	}
	return
}

func (p *parser) parseValueEqualFlag(flagAndValue string) (err error) {
	split := strings.SplitN(flagAndValue, "=", 2)
	flag := split[0]
	value := split[1]
	if value == "" {
		err = fmt.Errorf("cannot set empty value for [ %s ] flag", flag)
		return
	}
	l, ok := p.template.flags[flag]
	if ok {
		if l.kind == boolFlag {
			err = fmt.Errorf("cannot set custom value for boolean flag")
			return
		} else if l.kind == valueFlag {
			p.command.setValue(p.template.ValueFlags[l.index], value)
			p.valueFlagUsed[l.index] = true
		}
	}
	return
}

func (p *parser) parseMultiCharFlag(flag string) (more bool, err error) {
	if strings.Contains(flag, "=") {
		err = p.parseValueEqualFlag(flag)
		return
	}
	l, ok := p.template.flags[flag]
	if ok {
		if l.kind == boolFlag {
			p.command.setBool(p.template.BoolFlags[l.index])
			p.boolFlagUsed[l.index] = true
		} else if l.kind == valueFlag {
			more = true
			p.lastLink = l
		}
	}
	return
}

func (p *parser) parseValue(value string) {
	p.command.setValue(p.template.ValueFlags[p.lastLink.index], value)
	p.valueFlagUsed[p.lastLink.index] = true
	p.lastLink = link{}
	p.more = false
}

func (p *parser) setDefaults() {
	for i, used := range p.boolFlagUsed {
		if !used {
			p.command.setBoolDefault(p.template.BoolFlags[i])
		}
	}
	for i, used := range p.valueFlagUsed {
		if !used {
			p.command.setValueDefault(p.template.ValueFlags[i])
		}
	}
}
