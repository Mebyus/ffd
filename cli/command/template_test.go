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
		wantPrepErr bool
		wantErr     bool
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
			wantPrepErr: false,
			wantErr:     false,
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
			wantPrepErr: false,
			wantErr:     false,
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
			wantPrepErr: false,
			wantErr:     false,
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
			wantPrepErr: false,
			wantErr:     false,
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
			wantPrepErr: false,
			wantErr:     false,
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
			wantPrepErr: false,
			wantErr:     false,
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
			wantPrepErr: false,
			wantErr:     false,
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
			wantPrepErr: false,
			wantErr:     false,
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
			wantPrepErr: false,
			wantErr:     false,
		},
		{
			name: "realistic download command",
			fields: fields{
				Name: "download",
				ValueFlags: []ValueFlag{
					{
						Flag: Flag{
							Aliases: map[string]AliasType{
								"o":   SingleChar,
								"out": MultipleChars,
							},
						},
						Default: "",
					},
					{
						Flag: Flag{
							Aliases: map[string]AliasType{
								"f":      SingleChar,
								"format": MultipleChars,
							},
						},
						Default: "txt",
					},
				},
				BoolFlags: []BoolFlag{
					{
						Flag: Flag{
							Aliases: map[string]AliasType{
								"s":           SingleChar,
								"save-source": MultipleChars,
							},
						},
						Default: false,
					},
				},
			},
			args: args{
				args: []string{"-o", "output/dir", "--format", "fb2", "--save-source",
					"https://archiveofourown.org/works/29288139",
				},
			},
			wantCommand: &Command{
				Name:    "download",
				Targets: []string{"https://archiveofourown.org/works/29288139"},
				BoolFlags: map[string]bool{
					"s":           true,
					"save-source": true,
				},
				ValueFlags: map[string]string{
					"o":      "output/dir",
					"out":    "output/dir",
					"f":      "fb2",
					"format": "fb2",
				},
			},
			wantPrepErr: false,
			wantErr:     false,
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
			err := tr.prepare()
			if (err != nil) != tt.wantPrepErr {
				t.Errorf("Template.prepare() error = %v, wantPrepErr %v", err, tt.wantPrepErr)
				return
			}
			gotCommand, err := tr.Parse(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Template.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCommand, tt.wantCommand) {
				t.Errorf("Template.Parse() = %v, want %v", gotCommand, tt.wantCommand)
			}
		})
	}
}
