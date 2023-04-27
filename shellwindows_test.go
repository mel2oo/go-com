package com

import "testing"

func TestShellWindowsItem(t *testing.T) {
	if err := NewShellWindows().ShellExecute("calc"); err != nil {
		t.Fatal(err)
	}
}
