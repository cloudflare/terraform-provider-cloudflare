package main

import (
	"testing"
)

func TestWorkersScriptTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "workers_script name to script_name",
			Config: `resource "cloudflare_workers_script" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name = "my-worker"
  content = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
}`,
			Expected: []string{`resource "cloudflare_workers_script" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
}`},
		},
		{
			Name: "workers_script with plain_text_binding",
			Config: `resource "cloudflare_workers_script" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name = "my-worker"
  content = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
  
  plain_text_binding {
    name = "MY_VAR"
    text = "my-value"
  }
}`,
			Expected: []string{`resource "cloudflare_workers_script" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
  bindings = [{
    type = "plain_text"
    name = "MY_VAR"
    text = "my-value"
  }]
}`},
		},
		{
			Name: "workers_script with multiple binding types",
			Config: `resource "cloudflare_workers_script" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name = "my-worker"
  content = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
  
  plain_text_binding {
    name = "MY_VAR"
    text = "my-value"
  }
  
  kv_namespace_binding {
    name = "MY_KV"
    namespace_id = "abc123"
  }
}`,
			Expected: []string{`resource "cloudflare_workers_script" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
  bindings = [{
    type = "plain_text"
    name = "MY_VAR"
    text = "my-value"
    },
    {
      type         = "kv_namespace"
      name         = "MY_KV"
      namespace_id = "abc123"
  }]
}`},
		},
		{
			Name: "worker_script (singular) with name",
			Config: `resource "cloudflare_worker_script" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name = "my-worker"
  content = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
}`,
			Expected: []string{`resource "cloudflare_workers_script" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
}`},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}
