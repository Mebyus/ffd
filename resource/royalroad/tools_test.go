package royalroad

import "testing"

func Test_analyze(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name     string
		args     args
		wantBase string
		wantName string
		wantErr  bool
	}{
		{
			name: "empty string",
			args: args{
				url: "",
			},
			wantBase: "",
			wantName: "",
			wantErr:  true,
		},
		{
			name: "garbage string",
			args: args{
				url: "frefervxa",
			},
			wantBase: "",
			wantName: "",
			wantErr:  true,
		},
		{
			name: "fic name missing",
			args: args{
				url: "https://www.royalroad.com/fiction/25225",
			},
			wantBase: "",
			wantName: "",
			wantErr:  true,
		},
		{
			name: "just base",
			args: args{
				url: "https://www.royalroad.com/fiction/25225/delve",
			},
			wantBase: "https://www.royalroad.com/fiction/25225/delve",
			wantName: "delve",
			wantErr:  false,
		},
		{
			name: "base with slash at the end",
			args: args{
				url: "https://www.royalroad.com/fiction/25225/delve/",
			},
			wantBase: "https://www.royalroad.com/fiction/25225/delve",
			wantName: "delve",
			wantErr:  false,
		},
		{
			name: "chapter",
			args: args{
				url: "https://www.royalroad.com/fiction/25225/delve/chapter/368849/010-broke",
			},
			wantBase: "https://www.royalroad.com/fiction/25225/delve",
			wantName: "delve",
			wantErr:  false,
		},
		{
			name: "chapter with slash at the end",
			args: args{
				url: "https://www.royalroad.com/fiction/25225/delve/chapter/368849/010-broke/",
			},
			wantBase: "https://www.royalroad.com/fiction/25225/delve",
			wantName: "delve",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBase, gotName, err := analyze(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("analyze() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotBase != tt.wantBase {
				t.Errorf("analyze() gotBase = %v, want %v", gotBase, tt.wantBase)
			}
			if gotName != tt.wantName {
				t.Errorf("analyze() gotName = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}
