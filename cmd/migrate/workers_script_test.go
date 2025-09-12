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
		{
			Name: "workers_script with d1_database_binding (should map to d1 type)",
			Config: `resource "cloudflare_workers_script" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name = "my-worker"
  content = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
  
  d1_database_binding {
    name = "MY_DB"
    database_id = "db123"
  }
}`,
			Expected: []string{`resource "cloudflare_workers_script" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
  bindings = [{
    type = "d1"
    id   = "db123"
    name = "MY_DB"
  }]
}`},
		},
		{
			Name: "workers_script with hyperdrive_config_binding (should map to hyperdrive type)",
			Config: `resource "cloudflare_workers_script" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name = "my-worker"
  content = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
  
  hyperdrive_config_binding {
    binding = "HYPERDRIVE"
    id = "hyperdrive123"
  }
}`,
			Expected: []string{`resource "cloudflare_workers_script" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
  bindings = [{
    type = "hyperdrive"
    name = "HYPERDRIVE"
    id   = "hyperdrive123"
  }]
}`},
		},
		{
			Name: "workers_script with webassembly_binding (should generate warning and remove)",
			Config: `resource "cloudflare_workers_script" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name = "my-worker"
  content = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
  
  webassembly_binding {
    name = "WASM_MODULE"
    module = "wasm_bg.wasm"
  }
}`,
			Expected: []string{
				`resource "cloudflare_workers_script" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"

  # MIGRATION WARNING: webassembly_binding is not supported in v5.
  # WebAssembly modules must be bundled into the script content instead.
  # Please update your build process and remove this warning.
}`,
			},
		},
		{
			Name: "workers_script with module=true (should convert to main_module)",
			Config: `resource "cloudflare_workers_script" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name = "my-worker"
  content = "export default { fetch(request) { return new Response('Hello World'); } };"
  module = true
}`,
			Expected: []string{
				`resource "cloudflare_workers_script" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "export default { fetch(request) { return new Response('Hello World'); } };"
  main_module = "worker.js"
}`,
			},
		},
		{
			Name: "workers_script with module=false (should convert to body_part)",
			Config: `resource "cloudflare_workers_script" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name = "my-worker"
  content = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
  module = false
}`,
			Expected: []string{
				`resource "cloudflare_workers_script" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
  body_part   = "worker.js"
}`,
			},
		},
		{
			Name: "workers_script with module=true and existing bindings",
			Config: `resource "cloudflare_workers_script" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name = "my-worker"
  content = "export default { fetch(request) { return new Response('Hello World'); } };"
  module = true
  
  plain_text_binding {
    name = "MY_VAR"
    text = "my-value"
  }
}`,
			Expected: []string{
				`resource "cloudflare_workers_script" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "export default { fetch(request) { return new Response('Hello World'); } };"
  bindings = [{
    type = "plain_text"
    name = "MY_VAR"
    text = "my-value"
  }]
  main_module = "worker.js"
}`,
			},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}
