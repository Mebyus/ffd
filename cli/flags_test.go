package cli

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name  string
		args  args
		wantC *Command
	}{
		{
			name: "nil args",
			args: args{
				args: nil,
			},
			wantC: &Command{
				Name:   "help",
				Target: "",
				Flags:  map[string]string{},
			},
		},
		{
			name: "empty args slice",
			args: args{
				args: []string{},
			},
			wantC: &Command{
				Name:   "help",
				Target: "",
				Flags:  map[string]string{},
			},
		},
		{
			name: "help command with no flags",
			args: args{
				args: []string{"help"},
			},
			wantC: &Command{
				Name:   "help",
				Target: "",
				Flags:  map[string]string{},
			},
		},
		{
			name: "add command with no flags",
			args: args{
				args: []string{"add"},
			},
			wantC: &Command{
				Name:   "add",
				Target: "",
				Flags:  map[string]string{},
			},
		},
		{
			name: "add command with target",
			args: args{
				args: []string{"add", "my_target"},
			},
			wantC: &Command{
				Name:   "add",
				Target: "my_target",
				Flags:  map[string]string{},
			},
		},
		{
			name: "clean command with separated flags",
			args: args{
				args: []string{"add", "-h", "-s"},
			},
			wantC: &Command{
				Name:   "add",
				Target: "",
				Flags: map[string]string{
					"h": "",
					"s": "",
				},
			},
		},
		{
			name: "clean command with combined flags",
			args: args{
				args: []string{"add", "-hs"},
			},
			wantC: &Command{
				Name:   "add",
				Target: "",
				Flags: map[string]string{
					"h": "",
					"s": "",
				},
			},
		},
		{
			name: "parse command with named flag",
			args: args{
				args: []string{"parse", "--res=sb"},
			},
			wantC: &Command{
				Name:   "parse",
				Target: "",
				Flags: map[string]string{
					"res": "sb",
				},
			},
		},
		{
			name: "download command with target and flag",
			args: args{
				args: []string{"download", "https://site.com/route", "-s"},
			},
			wantC: &Command{
				Name:   "download",
				Target: "https://site.com/route",
				Flags: map[string]string{
					"s": "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := Parse(tt.args.args); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("Parse() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
