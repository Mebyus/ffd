package samlib

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
			name: "base url (shtml)",
			args: args{
				url: "http://samlib.ru/t/tokmakow_k_d/worm_esp.shtml",
			},
			wantBase: "http://samlib.ru/t/tokmakow_k_d/worm_esp.shtml",
			wantName: "worm-esp_tokmakow-k-d",
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
