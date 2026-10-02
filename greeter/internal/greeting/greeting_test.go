package greeting

import "testing"

func TestNewWithName(t *testing.T) {
	got := New("Ada")
	want := "Hello, Ada!"
	if got.Message != want {
		t.Errorf("New(%q).Message = %q, want %q", "Ada", got.Message, want)
	}
}

func TestNewWithoutName(t *testing.T) {
	got := New("")
	want := "Hello, World!"
	if got.Message != want {
		t.Errorf("New(%q).Message = %q, want %q", "", got.Message, want)
	}
}
