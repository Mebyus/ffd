package command

func NewStdHelpTemplate() *Template {
	return &Template{
		Name:        "help",
		Description: "",
		BoolFlags: []BoolFlag{
			{
				Flag: Flag{
					Aliases: map[string]AliasType{
						"v":       SingleChar,
						"version": MultipleChars,
					},
				},
				Default: false,
			},
			{
				Flag: Flag{
					Aliases: map[string]AliasType{
						"h":    SingleChar,
						"help": MultipleChars,
					},
				},
				Default: false,
			},
		},
	}
}
