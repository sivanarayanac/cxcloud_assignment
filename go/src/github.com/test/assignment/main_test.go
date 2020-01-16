package main

import (
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "GetAPIKey()", want: "308871e3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAPIKey(); got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fetchURLInfo(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{args: args{name: "Frozen"}, want: "http://www.omdbapi.com?apikey=308871e3&t=Frozen"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fetchURLInfo(tt.args.name); got != tt.want {
				t.Errorf("fetchURLInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputArg(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{name: "inputArg", want: "Avatar"},
		{name: "inputArg", want: "Frozen"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inputArg(); got != tt.want {
				t.Errorf("inputArg() = %v, want %v", got, tt.want)
			}
		})
	}
}
