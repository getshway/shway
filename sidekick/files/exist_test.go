package files

import (
	"testing"
)

func TestIsExist(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsExist(tt.args.p); got != tt.want {
				t.Errorf("IsExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNotExist(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotExist(tt.args.p); got != tt.want {
				t.Errorf("IsNotExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
