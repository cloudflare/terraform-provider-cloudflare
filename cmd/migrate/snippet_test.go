package main

import (
	"testing"
)

func TestMigrateCloudflareSnippet(t *testing.T) {
	tests := []TestCase{
		{
			Name: "basic snippet migration",
			Config: `
resource "cloudflare_snippet" "test" {
  zone_id     = "abc123"
  name        = "test_snippet"
  main_module = "main.js"
  files {
    name    = "main.js"
    content = "export default {async fetch(request) {return fetch(request)}};"
  }
}`,
			Expected: []string{`resource "cloudflare_snippet" "test" {
  zone_id      = "abc123"
  snippet_name = "test_snippet"
  metadata = {
    main_module = "main.js"
  }
  files = [
    {
      name    = "main.js"
      content = "export default {async fetch(request) {return fetch(request)}};"
    }
  ]
}
`,
			},
		},
		{
			Name: "multiple files migration",
			Config: `
resource "cloudflare_snippet" "test" {
  zone_id     = "abc123"
  name        = "test_snippet"
  main_module = "main.js"
  files {
    name    = "main.js"
    content = "import {helper} from './helper.js';"
  }
  files {
    name    = "helper.js"
    content = "export function helper() {}"
  }
}`,
			Expected: []string{`resource "cloudflare_snippet" "test" {
  zone_id      = "abc123"
  snippet_name = "test_snippet"
  metadata = {
    main_module = "main.js"
  }
  files = [
    {
      name    = "main.js"
      content = "import {helper} from './helper.js';"
    },
    {
      name    = "helper.js"
      content = "export function helper() {}"
    }
  ]
}
`,
			},
		},
		{
			Name: "snippet with heredoc content",
			Config: `
resource "cloudflare_snippet" "test" {
  zone_id     = "abc123"
  name        = "test_snippet"
  main_module = "main.js"
  files {
    name    = "main.js"
    content = <<-EOF
      export default {
        async fetch(request) {
          return fetch(request);
        }
      }
    EOF
  }
}`,
			Expected: []string{`resource "cloudflare_snippet" "test" {
  zone_id      = "abc123"
  snippet_name = "test_snippet"
  metadata = {
    main_module = "main.js"
  }
  files = [
    {
      name    = "main.js"
      content = <<-EOF
      export default {
        async fetch(request) {
          return fetch(request);
        }
      }
    EOF
    }
  ]
}
`,
			},
		},
		{
			Name: "multiple files with heredoc",
			Config: `
resource "cloudflare_snippet" "test" {
  zone_id     = "abc123"
  name        = "test_snippet"
  main_module = "main.js"
  files {
    name    = "main.js"
    content = <<-EOF
      import {helper} from './helper.js';
      export default {
        async fetch(request) {
          return helper(request);
        }
      }
    EOF
  }
  files {
    name    = "helper.js"
    content = <<-HELPER
      export function helper(request) {
        return fetch(request);
      }
    HELPER
  }
}`,
			Expected: []string{`resource "cloudflare_snippet" "test" {
  zone_id      = "abc123"
  snippet_name = "test_snippet"
  metadata = {
    main_module = "main.js"
  }
  files = [
    {
      name    = "main.js"
      content = <<-EOF
      import {helper} from './helper.js';
      export default {
        async fetch(request) {
          return helper(request);
        }
      }
    EOF
    },
    {
      name    = "helper.js"
      content = <<-HELPER
      export function helper(request) {
        return fetch(request);
      }
    HELPER
    }
  ]
}
`,
			},
		},
	}

	RunTransformationTests(t, tests, transformFileDefault)
}

func TestMigrateCloudflareSnippetState(t *testing.T) {
	tests := []struct {
		name          string
		resourceType  string
		instanceState map[string]interface{}
		expected      map[string]interface{}
	}{
		{
			name:         "basic state migration",
			resourceType: "cloudflare_snippet",
			instanceState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id":     "abc123",
					"name":        "test_snippet",
					"main_module": "main.js",
					"files": []interface{}{
						map[string]interface{}{
							"name":    "main.js",
							"content": "export default {}",
						},
					},
					"id": "some-id",
				},
			},
			expected: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id":      "abc123",
					"snippet_name": "test_snippet",
					"metadata": map[string]interface{}{
						"main_module": "main.js",
					},
					"files": []interface{}{
						map[string]interface{}{
							"name":    "main.js",
							"content": "export default {}",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// err := MigrateCloudflareSnippetState(tt.resourceType, tt.instanceState)
			// assert.NoError(t, nil)
			// assert.Equal(t, tt.expected, tt.instanceState)
		})
	}
}