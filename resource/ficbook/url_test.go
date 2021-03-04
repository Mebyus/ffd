package ficbook

import "testing"

func Test_analyze(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name     string
		args     args
		wantBase string
		wantID   string
		wantErr  bool
	}{
		{
			name: "empty string",
			args: args{
				url: "",
			},
			wantBase: "",
			wantID:   "",
			wantErr:  true,
		},
		{
			name: "garbage string",
			args: args{
				url: "fneionwfo",
			},
			wantBase: "",
			wantID:   "",
			wantErr:  true,
		},
		{
			name: "work id missing",
			args: args{
				url: "https://ficbook.net/readfic",
			},
			wantBase: "",
			wantID:   "",
			wantErr:  true,
		},
		{
			name: "just base",
			args: args{
				url: "https://ficbook.net/readfic/10244166",
			},
			wantBase: "https://ficbook.net/readfic/10244166",
			wantID:   "10244166",
			wantErr:  false,
		},
		{
			name: "base with slash at the end",
			args: args{
				url: "https://ficbook.net/readfic/10244166/",
			},
			wantBase: "https://ficbook.net/readfic/10244166",
			wantID:   "10244166",
			wantErr:  false,
		},
		{
			name: "chapter",
			args: args{
				url: "https://ficbook.net/readfic/10244166/26359810#part_content",
			},
			wantBase: "https://ficbook.net/readfic/10244166",
			wantID:   "10244166",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBase, gotID, err := analyze(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("analyze() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotBase != tt.wantBase {
				t.Errorf("analyze() gotBase = %v, want %v", gotBase, tt.wantBase)
			}
			if gotID != tt.wantID {
				t.Errorf("analyze() gotId = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}
