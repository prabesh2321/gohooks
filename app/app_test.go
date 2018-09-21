package app

import "testing"

func TestSample(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Simple", args{"simple"}, "simple"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sample(tt.args.s); got != tt.want {
				t.Errorf("Sample() = %v, want %v", got, tt.want)
			}
		})
	}
}
