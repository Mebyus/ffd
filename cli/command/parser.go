package command

type parser struct {
	template *Template
	command  *Command
	lastLink link
}

func (p *parser) parseFlag(flag string) (more bool) {
	l, ok := p.template.flags[flag]
	if ok {
		if l.kind == boolFlag {
			p.command.setBool(p.template.BoolFlags[l.index])
		}
		p.lastLink = l
	}
	return
}

func (p *parser) parseValue(value string) {
	return
}
