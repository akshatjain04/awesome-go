package main

import (
	"io/ioutil"
	"os"
	"testing"
	"path/filepath"
)

func TestdropCreateDir(t *testing.T) {
	tests := []struct {
		name       string
		setup      func() (string, error)
		assert     func(t *testing.T, path string, err error)
	}{
		{
			name: "Scenario 1: Successful removal and creation of directory",
			setup: func() (string, error) {
				path, err := ioutil.TempDir("", "test")
				return path, err
			},
			assert: func(t *testing.T, path string, err error) {
				_, err = os.Stat(path)
				if err != nil {
					t.Errorf("Expected directory to exist, but got error: %v", err)
				}
			},
		},
		{
			name: "Scenario 2: Directory does not exist",
			setup: func() (string, error) {
				path, err := ioutil.TempDir("", "test")
				if err != nil {
					return "", err
				}
				err = os.Remove(path)
				return path, err
			},
			assert: func(t *testing.T, path string, err error) {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}
				_, err = os.Stat(path)
				if err == nil {
					t.Errorf("Expected directory to not exist, but it does")
				}
			},
		},
		{
			name: "Scenario 3: Insufficient permissions to remove directory",
			setup: func() (string, error) {
				path, err := ioutil.TempDir("", "test")
				if err != nil {
					return "", err
				}
				err = os.Chmod(path, 0400)
				return path, err
			},
			assert: func(t *testing.T, path string, err error) {
				if err == nil {
					t.Errorf("Expected error, but got none")
				}
			},
		},
		{
			name: "Scenario 4: Insufficient permissions to create directory",
			setup: func() (string, error) {
				path, err := ioutil.TempDir("", "test")
				if err != nil {
					return "", err
				}
				err = os.Remove(path)
				if err != nil {
					return "", err
				}
				err = os.Chmod(filepath.Dir(path), 0400)
				return path, err
			},
			assert: func(t *testing.T, path string, err error) {
				if err == nil {
					t.Errorf("Expected error, but got none")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, setupErr := tt.setup()
			if setupErr != nil {
				t.Fatalf("Failed to setup test: %v", setupErr)
			}
			err := dropCreateDir(path)
			tt.assert(t, path, err)
		})
	}
}

func dropCreateDir(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return err
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return nil
}
