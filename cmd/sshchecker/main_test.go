package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseFile(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "users.txt")
	content := " user1 \n\n user2\n\t\nuser3  "
	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	got, err := parseFile(file)
	if err != nil {
		t.Fatalf("parseFile returned error: %v", err)
	}

	want := []string{"user1", "user2", "user3"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}
