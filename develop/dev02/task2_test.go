package main

import (
	"testing"
)

func TestTask2(t *testing.T) {
	tests := []struct {
		str     string
		wantErr bool
		wantRes string
	}{
		{
			str: "a4bc2d5e", wantErr: false, wantRes: "aaaabccddddde",
		}, {
			str: "abcd", wantErr: false, wantRes: "abcd",
		}, {
			str: "45", wantErr: true, wantRes: "",
		}, {
			str: "", wantErr: false, wantRes: "",
		},
	}
	for _, tt := range tests {
		s, err := unpack(tt.str)
		if err != nil && !tt.wantErr {
			t.Errorf("was error but error was not expected: %v", err)
		}
		if err == nil && tt.wantErr {
			t.Errorf("was not error but error was expected: ")
		}
		if s != tt.wantRes {
			t.Errorf("was not wanted result")
		}
	}
}
