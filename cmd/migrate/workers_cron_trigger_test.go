package main

import (
	"testing"
)

func TestWorkersCronTriggerTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "worker_cron_trigger singular resource rename",
			Config: `resource "cloudflare_worker_cron_trigger" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  cron        = "0 0 * * *"
}`,
			Expected: []string{`resource "cloudflare_workers_cron_trigger" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  cron        = "0 0 * * *"
}`},
		},
		{
			Name: "workers_cron_trigger plural stays same",
			Config: `resource "cloudflare_workers_cron_trigger" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  cron        = "*/5 * * * *"
}`,
			Expected: []string{`resource "cloudflare_workers_cron_trigger" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  cron        = "*/5 * * * *"
}`},
		},
	}

	RunTransformationTests(t, tests, transformFileWithoutImports)
}