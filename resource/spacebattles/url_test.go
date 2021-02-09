package spacebattles

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
				url: "fneionwfo",
			},
			wantBase: "",
			wantName: "",
			wantErr:  true,
		},
		{
			name: "thread name missing",
			args: args{
				url: "https://forums.spacebattles.com/threads/",
			},
			wantBase: "",
			wantName: "",
			wantErr:  true,
		},
		{
			name: "base with slash at the end",
			args: args{
				url: "https://forums.spacebattles.com/threads/crystalized-munchkinry-worm-au-shard-si-fix-it.897992/",
			},
			wantBase: "https://forums.spacebattles.com/threads/crystalized-munchkinry-worm-au-shard-si-fix-it.897992",
			wantName: "crystalized-munchkinry-worm-au-shard-si-fix-it",
			wantErr:  false,
		},
		{
			name: "just base",
			args: args{
				url: "https://forums.spacebattles.com/threads/crystalized-munchkinry-worm-au-shard-si-fix-it.897992",
			},
			wantBase: "https://forums.spacebattles.com/threads/crystalized-munchkinry-worm-au-shard-si-fix-it.897992",
			wantName: "crystalized-munchkinry-worm-au-shard-si-fix-it",
			wantErr:  false,
		},
		{
			name: "reader",
			args: args{
				url: "https://forums.spacebattles.com/threads/crystalized-munchkinry-worm-au-shard-si-fix-it.897992/reader",
			},
			wantBase: "https://forums.spacebattles.com/threads/crystalized-munchkinry-worm-au-shard-si-fix-it.897992",
			wantName: "crystalized-munchkinry-worm-au-shard-si-fix-it",
			wantErr:  false,
		},
		{
			name: "reader at page 3",
			args: args{
				url: "https://forums.spacebattles.com/threads/crystalized-munchkinry-worm-au-shard-si-fix-it.897992/reader/page-3",
			},
			wantBase: "https://forums.spacebattles.com/threads/crystalized-munchkinry-worm-au-shard-si-fix-it.897992",
			wantName: "crystalized-munchkinry-worm-au-shard-si-fix-it",
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
