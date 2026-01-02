package utils

import (
	"os"
	"testing"
)

func TestIsTestResource(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid test resource", "cftftestabcdefghij", true},
		{"valid test resource with suffix", "cftftestabcdefghijupdated", true},
		{"test resource with prefix", "my-tf-pool-basic-cftftestabcdefghij", true},
		{"test resource with multiple prefixes", "tf-testacc-lb-cftftestxyz", true},
		{"legacy test resource", "tf-acctest-abcdefghij", false},
		{"production resource", "my-production-resource", false},
		{"empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsTestResource(tt.input)
			if result != tt.expected {
				t.Errorf("IsTestResource(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestShouldSweepResource(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"test resource", "cftftestabcdefghij", true},
		{"test resource with suffix", "cftftestbasic", true},
		{"test resource with prefix", "my-tf-pool-basic-cftftestabcdefghij", true},
		{"test resource with load balancer prefix", "tf-testacc-lb-cftftestxyz", true},
		{"production resource", "my-production-resource", false},
		{"empty string", "", false},
		{"similar but not matching", "cf-test-resource", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ShouldSweepResource(tt.input)
			if result != tt.expected {
				t.Errorf("ShouldSweepResource(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestShouldSweepResource_DangerMode(t *testing.T) {
	// Save original env var and restore after test
	originalValue := os.Getenv("SWEEP_DANGEROUSLY_DELETE_ALL")
	defer func() {
		if originalValue == "" {
			os.Unsetenv("SWEEP_DANGEROUSLY_DELETE_ALL")
		} else {
			os.Setenv("SWEEP_DANGEROUSLY_DELETE_ALL", originalValue)
		}
	}()

	// Set danger mode
	os.Setenv("SWEEP_DANGEROUSLY_DELETE_ALL", "true")

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"production resource in danger mode", "my-production-resource", true},
		{"random resource in danger mode", "anything-goes", true},
		{"empty string in danger mode", "", true},
		{"test resource in danger mode", "cftftestabc", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ShouldSweepResource(tt.input)
			if result != tt.expected {
				t.Errorf("ShouldSweepResource(%q) in danger mode = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}

	// Verify normal mode still works after unsetting
	os.Unsetenv("SWEEP_DANGEROUSLY_DELETE_ALL")
	if ShouldSweepResource("my-production-resource") {
		t.Error("ShouldSweepResource('my-production-resource') should return false after unsetting danger mode")
	}
}
