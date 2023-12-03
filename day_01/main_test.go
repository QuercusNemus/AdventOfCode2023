package main

import "testing"

func Test_calibration(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{{
		name: "Correct",
		args: args{
			path: "test_input.txt"},
		want: 142,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calibration(tt.args.path); got != tt.want {
				t.Errorf("calibration() = %v, want %v", got, tt.want)
			}
		})
	}
}
