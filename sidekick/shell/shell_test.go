package shell

import "testing"

func TestLookPath(t *testing.T) {
	type args struct {
		fn string
	}
	tests := []struct {
		name     string
		args     args
		wantPath string
		wantErr  bool
	}{
		{
			name: "ls command",
			args: args{
				"ls",
			},
			wantPath: "/bin/ls",
			wantErr:  false,
		},
		{
			name: "empty",
			args: args{
				"",
			},
			wantPath: "",
			wantErr:  true,
		},
		{
			name: "not-exist command",
			args: args{
				"notfound-notfound-foobar-fizz-bizz-0123456789",
			},
			wantPath: "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPath, err := LookPath(tt.args.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("LookPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPath != tt.wantPath {
				t.Errorf("LookPath() = %v, want %v", gotPath, tt.wantPath)
			}
		})
	}
}
