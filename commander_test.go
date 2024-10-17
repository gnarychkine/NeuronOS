package main

import "testing"

func TestGetSystemInfo(t *testing.T) {
	cmdr := NewCommander()
	info, err := cmdr.GetSystemInfo()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if info.Hostname == "" {
		t.Error("Expected hostname to be non-empty")
	}

	if info.IPAddress == "" {
		t.Error("Expected IP address to be non-empty")
	}
}

func TestPing(t *testing.T) {
	cmdr := NewCommander()
	pr, err := cmdr.Ping("google.com")
	_ = pr

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !pr.Successful {
		t.Fatalf("Number of received packets is 0")
	}
}
