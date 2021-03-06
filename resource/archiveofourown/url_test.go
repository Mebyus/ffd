package archiveofourown

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
				url: "https://archiveofourown.org/works/",
			},
			wantBase: "",
			wantID:   "",
			wantErr:  true,
		},
		{
			name: "navigate with slash at the end",
			args: args{
				url: "https://archiveofourown.org/works/29288139/navigate/",
			},
			wantBase: "https://archiveofourown.org/works/29288139",
			wantID:   "29288139",
			wantErr:  false,
		},
		{
			name: "navigate",
			args: args{
				url: "https://archiveofourown.org/works/29288139/navigate",
			},
			wantBase: "https://archiveofourown.org/works/29288139",
			wantID:   "29288139",
			wantErr:  false,
		},
		{
			name: "just base",
			args: args{
				url: "https://archiveofourown.org/works/29288139",
			},
			wantBase: "https://archiveofourown.org/works/29288139",
			wantID:   "29288139",
			wantErr:  false,
		},
		{
			name: "chapter",
			args: args{
				url: "https://archiveofourown.org/works/29288139/chapters/71981736",
			},
			wantBase: "https://archiveofourown.org/works/29288139",
			wantID:   "29288139",
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
				t.Errorf("analyze() gotName = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}
