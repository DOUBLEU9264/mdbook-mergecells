package main

import (
	"os"
	"testing"
)

func Test_processBookContent(t *testing.T) {
	b, _ := os.ReadFile("b.json")
	tests := []struct {
		name string
		arg  []byte
	}{
		{
			name: "1",
			arg:  b,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processBookContent(tt.arg)
		})
	}
}
