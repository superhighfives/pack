package main

import (
	"bytes"
	"io/fs"
	"strings"
	"testing"

	"github.com/fatih/color"
)

func TestListScripts(t *testing.T) {
	// Force deterministic, ANSI-free output for assertions.
	color.NoColor = true

	tests := []struct {
		name     string
		data     string
		readErr  error
		wantCode int
		wantOut  []string // substrings expected on stdout
		wantErr  []string // substrings expected on stderr
	}{
		{
			name:     "missing file",
			readErr:  fs.ErrNotExist,
			wantCode: 1,
			wantErr:  []string{"No package.json found"},
		},
		{
			name:     "unreadable file",
			readErr:  fs.ErrPermission,
			wantCode: 1,
			wantErr:  []string{"Couldn't read package.json", "permission denied"},
		},
		{
			name:     "invalid json",
			data:     "{ not json",
			wantCode: 1,
			wantErr:  []string{"isn't valid JSON"},
		},
		{
			name:     "scripts is a string",
			data:     `{"scripts":"foo"}`,
			wantCode: 1,
			wantErr:  []string{`"scripts" field in package.json isn't an object`},
		},
		{
			name:     "scripts is an array",
			data:     `{"scripts":[]}`,
			wantCode: 1,
			wantErr:  []string{"isn't an object"},
		},
		{
			name:     "no scripts key",
			data:     `{"name":"x"}`,
			wantCode: 1,
			wantErr:  []string{"No script commands found"},
		},
		{
			name:     "empty scripts object",
			data:     `{"scripts":{}}`,
			wantCode: 1,
			wantErr:  []string{"No script commands found"},
		},
		{
			name:     "valid scripts",
			data:     `{"scripts":{"build":"go build","test":"go test ./..."}}`,
			wantCode: 0,
			wantOut: []string{
				"Available script commands in package.json",
				"build go build",
				"test go test ./...",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer

			var data []byte
			if tt.readErr == nil {
				data = []byte(tt.data)
			}

			code := listScripts(&stdout, &stderr, data, tt.readErr)
			if code != tt.wantCode {
				t.Errorf("exit code = %d, want %d", code, tt.wantCode)
			}

			for _, want := range tt.wantOut {
				if !strings.Contains(stdout.String(), want) {
					t.Errorf("stdout = %q, want substring %q", stdout.String(), want)
				}
			}
			for _, want := range tt.wantErr {
				if !strings.Contains(stderr.String(), want) {
					t.Errorf("stderr = %q, want substring %q", stderr.String(), want)
				}
			}

			// Error cases must not leak the header onto stdout.
			if tt.wantCode != 0 && stdout.Len() != 0 {
				t.Errorf("stdout = %q, want empty on error", stdout.String())
			}
		})
	}
}
