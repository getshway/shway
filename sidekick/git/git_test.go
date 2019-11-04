package git_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	. "github.com/getshway/shway/sidekick/git"
)

func TestCloneAndPull(t *testing.T) {
	type args struct {
		path string
		url  string
	}
	tests := []struct {
		name         string
		args         args
		wantCloneErr bool
		wantPullErr  bool
	}{
		{
			name: "git protocol",
			args: args{
				path: testDir(),
				url:  "git@github.com:getshway/shway.git",
			},
			wantCloneErr: false,
			wantPullErr:  false,
		},
		{
			name: "https protocol",
			args: args{
				path: testDir(),
				url:  "https://github.com/getshway/shway.git",
			},
			wantCloneErr: false,
			wantPullErr:  false,
		},
		{
			name: "not found repogitory",
			args: args{
				path: testDir(),
				url:  "git@github.com:getshway/notfound.git",
			},
			wantCloneErr: true,
			wantPullErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.MkdirAll(tt.args.path, 0755)
			_, err := Clone(tt.args.path, tt.args.url)
			if (err != nil) != tt.wantCloneErr {
				t.Errorf("Clone() error = %v, wantErr %v", err, tt.wantCloneErr)
				return
			}
			_, err = Pull(tt.args.path)
			if (err != nil) != tt.wantPullErr {
				t.Errorf("Pull() error = %v, wantErr %v", err, tt.wantPullErr)
				return
			}
		})
	}
}

func testDir() string {
	return fmt.Sprintf("/tmp/tmp.%d", time.Now().UnixNano())
}
