package files

import "testing"

func TestMakeDirIfNotExists(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MakeDirIfNotExists(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("MakeDirIfNotExists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
