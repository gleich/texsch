package config

import (
	"os"
	"strings"
	"testing"
)

// TestPath ... Test the Path() function
func TestPath(t *testing.T) {
	path := Path()
	if !strings.Contains(path, "texsch") {
		t.Error("Path doesn't contain  texsch")
	}
}

// Create ... Test for the Create function
func TestCreate(t *testing.T) {
	path := Path()
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		err := os.Remove(path)
		if err != nil {
			t.Error(err)
		}
	}
	Create(path, true)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		t.Error("Path does not exist")
	}
	Create(path, false)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		t.Error("Path does not exist")
	}
}
