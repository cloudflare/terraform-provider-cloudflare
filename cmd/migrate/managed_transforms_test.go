package main

import (
	"testing"
)

func TestManagedTransformsTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "adds missing managed_response_headers",
			Config: `resource "cloudflare_managed_transforms" "test" {
  zone_id = "abc123"
  managed_request_headers = [
    {
      id      = "add_true_client_ip_headers"
      enabled = true
    }
  ]
}`,
			Expected: []string{`resource "cloudflare_managed_transforms" "test" {
  zone_id                  = "abc123"
  managed_request_headers  = [
    {
      id      = "add_true_client_ip_headers"
      enabled = true
    }
  ]
  managed_response_headers = []
}`},
		},
		{
			Name: "adds missing managed_request_headers",
			Config: `resource "cloudflare_managed_transforms" "test" {
  zone_id = "abc123"
  managed_response_headers = [
    {
      id      = "add_security_headers"
      enabled = true
    }
  ]
}`,
			Expected: []string{`resource "cloudflare_managed_transforms" "test" {
  zone_id                  = "abc123"
  managed_response_headers = [
    {
      id      = "add_security_headers"
      enabled = true
    }
  ]
  managed_request_headers  = []
}`},
		},
		{
			Name: "adds both missing attributes",
			Config: `resource "cloudflare_managed_transforms" "test" {
  zone_id = "abc123"
}`,
			Expected: []string{`resource "cloudflare_managed_transforms" "test" {
  zone_id                  = "abc123"
  managed_request_headers  = []
  managed_response_headers = []
}`},
		},
		{
			Name: "preserves existing attributes",
			Config: `resource "cloudflare_managed_transforms" "test" {
  zone_id = "abc123"
  managed_request_headers = [
    {
      id      = "add_true_client_ip_headers"
      enabled = true
    }
  ]
  managed_response_headers = [
    {
      id      = "add_security_headers"
      enabled = true
    }
  ]
}`,
			Expected: []string{`resource "cloudflare_managed_transforms" "test" {
  zone_id                  = "abc123"
  managed_request_headers  = [
    {
      id      = "add_true_client_ip_headers"
      enabled = true
    }
  ]
  managed_response_headers = [
    {
      id      = "add_security_headers"
      enabled = true
    }
  ]
}`},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}