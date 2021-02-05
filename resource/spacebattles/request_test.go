package spacebattles

import "testing"

func Test_baseURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name     string
		args     args
		wantBase string
		wantErr  bool
	}{
		{
			name: "empty string",
			args: args{
				url: "",
			},
			wantBase: "",
			wantErr:  true,
		},
		{
			name: "garbage string",
			args: args{
				url: "fneionwfo",
			},
			wantBase: "",
			wantErr:  true,
		},
		{
			name: "thread name missing",
			args: args{
				url: "https://forums.spacebattles.com/threads/",
			},
			wantBase: "",
			wantErr:  true,
		},
		{
			name: "base with slash at the end",
			args: args{
				url: "https://forums.spacebattles.com/threads/crystalized-munchkinry-worm-au-shard-si-fix-it.897992/",
			},
			wantBase: "https://forums.spacebattles.com/threads/crystalized-munchkinry-worm-au-shard-si-fix-it.897992",
			wantErr:  false,
		},
		{
			name: "just base",
			args: args{
				url: "https://forums.spacebattles.com/threads/crystalized-munchkinry-worm-au-shard-si-fix-it.897992",
			},
			wantBase: "https://forums.spacebattles.com/threads/crystalized-munchkinry-worm-au-shard-si-fix-it.897992",
			wantErr:  false,
		},
		{
			name: "reader",
			args: args{
				url: "https://forums.spacebattles.com/threads/crystalized-munchkinry-worm-au-shard-si-fix-it.897992/reader",
			},
			wantBase: "https://forums.spacebattles.com/threads/crystalized-munchkinry-worm-au-shard-si-fix-it.897992",
			wantErr:  false,
		},
		{
			name: "reader at page 3",
			args: args{
				url: "https://forums.spacebattles.com/threads/crystalized-munchkinry-worm-au-shard-si-fix-it.897992/reader/page-3",
			},
			wantBase: "https://forums.spacebattles.com/threads/crystalized-munchkinry-worm-au-shard-si-fix-it.897992",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBase, err := baseURL(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotBase != tt.wantBase {
				t.Errorf("baseURL() = %v, want %v", gotBase, tt.wantBase)
			}
		})
	}
}
