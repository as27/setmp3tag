package main

import "testing"

func Test_useFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"use",
			args{"abc.mp3"},
			true,
		},
		{
			"not use",
			args{"abc.txt"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := useFile(tt.args.fileName); got != tt.want {
				t.Errorf("useFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
