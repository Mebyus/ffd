package document

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestEqual(t *testing.T) {
	// in this section we feed the same tree
	// to the function
	type argsSame struct {
		path string // path to file with html
	}
	testsSame := []struct {
		name string
		args argsSame
		want bool
	}{
		{
			name: "same 1_ab.html",
			args: argsSame{
				path: "test_data/equal_true/1_ab.html",
			},
			want: true,
		},
	}
	for _, tt := range testsSame {
		t.Run(tt.name, func(t *testing.T) {
			r, err := os.Open(tt.args.path)
			if err != nil {
				t.Errorf("Failed to load test data: %v", err)
			}
			defer r.Close()
			a, err := html.Parse(r)
			if err != nil {
				t.Errorf("Failed to load test data: %v", err)
			}
			if got := Equal(a, a); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
