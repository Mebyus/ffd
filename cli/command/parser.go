package command

import "strings"

type parser struct {
	template *Template
	command  *Command
	lastLink link
	more     bool
}

func (p *parser) parseFlag(flag string) (more bool, err error) {
	if strings.HasPrefix(flag, "-") {
		more, err = p.parseMultiCharFlag(strings.TrimPrefix(flag, "-"))
	} else {
		if len(flag) == 1 {
			more = p.parseSingleCharFlag(flag)
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
		} else if l.kind == valueFlag {
			more = true
			p.lastLink = l
		}
	}
	return
}

func (p *parser) parseMultiCharFlag(flag string) (more bool, err error) {
	l, ok := p.template.flags[flag]
	if ok {
		if l.kind == boolFlag {
			p.command.setBool(p.template.BoolFlags[l.index])
		} else if l.kind == valueFlag {
			more = true
			p.lastLink = l
		}
	}
	return
}

func (p *parser) parseValue(value string) {
	p.command.setValue(p.template.ValueFlags[p.lastLink.index], value)
	p.lastLink = link{}
}
