package main

import (
	"testing"
)

func TestWorkersReferencesTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "workers cross-resource references",
			Config: `resource "cloudflare_worker_script" "my_script" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "my-worker"
  content    = "// worker code"
}

resource "cloudflare_worker_route" "my_route" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern     = "example.cfapi.net/*"
  script_name = cloudflare_worker_script.my_script.name
}

resource "cloudflare_worker_cron_trigger" "my_cron" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = cloudflare_worker_script.my_script.name
  cron        = "0 0 * * *"
}`,
			Expected: []string{`resource "cloudflare_workers_script" "my_script" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "// worker code"
}

resource "cloudflare_workers_route" "my_route" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern = "example.cfapi.net/*"
  script  = cloudflare_workers_script.my_script.script_name
}

resource "cloudflare_workers_cron_trigger" "my_cron" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = cloudflare_workers_script.my_script.script_name
  cron        = "0 0 * * *"
}`},
		},
		{
			Name: "workers references with mixed v4/v5 syntax",
			Config: `resource "cloudflare_workers_script" "my_script" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "// worker code"
}

resource "cloudflare_workers_route" "my_route" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern = "example.cfapi.net/*"
  script  = cloudflare_workers_script.my_script.name
}`,
			Expected: []string{`resource "cloudflare_workers_script" "my_script" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "// worker code"
}

resource "cloudflare_workers_route" "my_route" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern = "example.cfapi.net/*"
  script  = cloudflare_workers_script.my_script.script_name
}`},
		},
	}

	RunTransformationTests(t, tests, transformFileWithoutImports)
}