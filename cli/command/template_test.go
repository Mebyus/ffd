package command

import (
	"reflect"
	"testing"
)

func TestTemplate_Parse(t *testing.T) {
	type fields struct {
		Name        string
		Description string
		BoolFlags   []BoolFlag
		ValueFlags  []ValueFlag
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantCommand *Command
	}{
		{
			name: "template with no flags, nil args",
			fields: fields{
				Name: "list",
			},
			args: args{
				args: nil,
			},
			wantCommand: &Command{
				Name:       "list",
				Targets:    []string{},
				BoolFlags:  map[string]bool{},
				ValueFlags: map[string]string{},
			},
		},
		{
			name: "template with no flags, one target in args",
			fields: fields{
				Name: "download",
			},
			args: args{
				args: []string{"https://archiveofourown.org/works/29288139"},
			},
			wantCommand: &Command{
				Name:       "download",
				Targets:    []string{"https://archiveofourown.org/works/29288139"},
				BoolFlags:  map[string]bool{},
				ValueFlags: map[string]string{},
			},
		},
		{
			name: "template with one bool flag (one alias), empty args",
			fields: fields{
				Name: "list",
				BoolFlags: []BoolFlag{
					{
						Flag: Flag{
							Aliases: map[string]AliasType{
								"a": SingleChar,
							},
						},
						Default: false,
					},
				},
			},
			args: args{
				args: []string{},
			},
			wantCommand: &Command{
				Name:       "list",
				Targets:    []string{},
				BoolFlags:  map[string]bool{},
				ValueFlags: map[string]string{},
			},
		},
		{
			name: "template with one bool flag (one alias), args with this flag",
			fields: fields{
				Name: "list",
				BoolFlags: []BoolFlag{
					{
						Flag: Flag{
							Aliases: map[string]AliasType{
								"a": SingleChar,
							},
						},
						Default: false,
					},
				},
			},
			args: args{
				args: []string{"-a"},
			},
			wantCommand: &Command{
				Name:    "list",
				Targets: []string{},
				BoolFlags: map[string]bool{
					"a": true,
				},
				ValueFlags: map[string]string{},
			},
		},
		{
			name: "template with one bool flag (one alias and default = true), args with this flag",
			fields: fields{
				Name: "list",
				BoolFlags: []BoolFlag{
					{
						Flag: Flag{
							Aliases: map[string]AliasType{
								"a": SingleChar,
							},
						},
						Default: true,
					},
				},
			},
			args: args{
				args: []string{"-a"},
			},
			wantCommand: &Command{
				Name:    "list",
				Targets: []string{},
				BoolFlags: map[string]bool{
					"a": false,
				},
				ValueFlags: map[string]string{},
			},
		},
		{
			name: "template with one bool flag (two aliases), args with this flag",
			fields: fields{
				Name: "list",
				BoolFlags: []BoolFlag{
					{
						Flag: Flag{
							Aliases: map[string]AliasType{
								"a": SingleChar,
								"A": SingleChar,
							},
						},
						Default: false,
					},
				},
			},
			args: args{
				args: []string{"-a"},
			},
			wantCommand: &Command{
				Name:    "list",
				Targets: []string{},
				BoolFlags: map[string]bool{
					"a": true,
					"A": true,
				},
				ValueFlags: map[string]string{},
			},
		},
		{
			name: "template with two bool flags, args with this flags",
			fields: fields{
				Name: "list",
				BoolFlags: []BoolFlag{
					{
						Flag: Flag{
							Aliases: map[string]AliasType{
								"a": SingleChar,
								"A": SingleChar,
							},
						},
						Default: false,
					},
					{
						Flag: Flag{
							Aliases: map[string]AliasType{
								"b": SingleChar,
								"B": SingleChar,
								"r": SingleChar,
							},
						},
						Default: false,
					},
				},
			},
			args: args{
				args: []string{"-aB"},
			},
			wantCommand: &Command{
				Name:    "list",
				Targets: []string{},
				BoolFlags: map[string]bool{
					"a": true,
					"A": true,
					"b": true,
					"B": true,
					"r": true,
				},
				ValueFlags: map[string]string{},
			},
		},
		{
			name: "template with one bool flag (two aliases), args with multichar flag",
			fields: fields{
				Name: "list",
				BoolFlags: []BoolFlag{
					{
						Flag: Flag{
							Aliases: map[string]AliasType{
								"a":   SingleChar,
								"all": MultipleChars,
							},
						},
						Default: false,
					},
				},
			},
			args: args{
				args: []string{"--all"},
			},
			wantCommand: &Command{
				Name:    "list",
				Targets: []string{},
				BoolFlags: map[string]bool{
					"a":   true,
					"all": true,
				},
				ValueFlags: map[string]string{},
			},
		},
		{
			name: "template with one value flag (one alias)",
			fields: fields{
				Name: "download",
				ValueFlags: []ValueFlag{
					{
						Flag: Flag{
							Aliases: map[string]AliasType{
								"o": SingleChar,
							},
						},
						Default: "",
					},
				},
			},
			args: args{
				args: []string{"-o", "output/dir"},
			},
			wantCommand: &Command{
				Name:      "download",
				Targets:   []string{},
				BoolFlags: map[string]bool{},
				ValueFlags: map[string]string{
					"o": "output/dir",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Template{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				BoolFlags:   tt.fields.BoolFlags,
				ValueFlags:  tt.fields.ValueFlags,
			}
			if gotCommand := tr.Parse(tt.args.args); !reflect.DeepEqual(gotCommand, tt.wantCommand) {
				t.Errorf("Template.Parse() = %v, want %v", gotCommand, tt.wantCommand)
			}
		})
	}
}
