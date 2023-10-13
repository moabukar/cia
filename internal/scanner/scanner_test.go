package scanner

import (
	"testing"
)

type MockCommandRunner struct {
	Output []byte
	Err    error
}

func (m *MockCommandRunner) RunCommand(name string, arg ...string) ([]byte, error) {
	return m.Output, m.Err
}

func TestScanImage(t *testing.T) {
	tests := []struct {
		name       string
		image      string
		options    ScanOptions
		mockRunner CommandRunner
		expectErr  bool
	}{
		{
			name:    "successful scan",
			image:   "nginx",
			options: ScanOptions{Format: "json", SkipPull: false},
			mockRunner: &MockCommandRunner{
				Output: []byte("some output"),
				Err:    nil,
			},
			expectErr: false,
		},
		// Add more test cases as needed.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ScanImage(tt.image, tt.options, tt.mockRunner)
			if (err != nil) != tt.expectErr {
				t.Errorf("ScanImage() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}
