package main

import (
	"testing"
)

func TestWorkersRouteTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "workers_route script_name to script",
			Config: `resource "cloudflare_workers_route" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern = "example.com/*"
  script_name = "my-worker"
}`,
			Expected: []string{`resource "cloudflare_workers_route" "example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern     = "example.com/*"
  script      = "my-worker"
}`},
		},
		{
			Name: "workers_route with no script_name",
			Config: `resource "cloudflare_workers_route" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"  
  pattern = "example.com/*"
}`,
			Expected: []string{`resource "cloudflare_workers_route" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern = "example.com/*"
}`},
		},
		{
			Name: "worker_route (singular) with script_name",
			Config: `resource "cloudflare_worker_route" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern = "example.com/*"
  script_name = "my-worker"
}`,
			Expected: []string{`resource "cloudflare_workers_route" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern = "example.com/*"
  script  = "my-worker"
}`},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}
