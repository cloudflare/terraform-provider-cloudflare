package main

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestDNSRecordCAATransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "CAA record with numeric flags in data block - content renamed to value",
			Config: `
resource "cloudflare_record" "caa_test" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "CAA"
  ttl     = 3600

  data {
    flags   = 0
    tag     = "issue"
    content = "letsencrypt.org"
  }
}`,
			Expected: []string{`
resource "cloudflare_dns_record" "caa_test" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "CAA"
  ttl     = 3600

  data = {
    flags = 0
    tag   = "issue"
    value = "letsencrypt.org"
  }
}`},
		},
		{
			Name: "CAA record with numeric flags in data attribute map - content renamed to value",
			Config: `
resource "cloudflare_record" "caa_test" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "CAA"
  ttl     = 3600
  data    {
    flags   = 0
    tag     = "issue"
    content = "letsencrypt.org"
  }
}`,
			Expected: []string{`
resource "cloudflare_dns_record" "caa_test" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "CAA"
  ttl     = 3600
  data    = {
    flags = 0
    tag   = "issue"
    value = "letsencrypt.org"
  }
}`},
		},
		{
			Name: "CAA record with flags already as string - content still renamed to value",
			Config: `
resource "cloudflare_record" "caa_test" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "CAA"
  ttl     = 3600

  data {
    flags   = "0"
    tag     = "issue"
    content = "letsencrypt.org"
  }
}`,
			Expected: []string{`
resource "cloudflare_dns_record" "caa_test" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "CAA"
  ttl     = 3600

  data = {
    flags = "0"
    tag   = "issue"
    value = "letsencrypt.org"
  }
}`},
		},
		{
			Name: "Non-CAA record should not be modified",
			Config: `
resource "cloudflare_record" "a_test" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "A"
  ttl     = 3600
  content = "192.168.1.1"
}`,
			Expected: []string{`
resource "cloudflare_dns_record" "a_test" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "A"
  ttl     = 3600
  content = "192.168.1.1"
}`},
		},
		{
			Name: "cloudflare_record (legacy) with CAA type - content renamed to value",
			Config: `
resource "cloudflare_record" "caa_legacy" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "CAA"
  ttl     = 3600

  data {
    flags   = 128
    tag     = "issuewild"
    content = "pki.goog"
  }
}`,
			Expected: []string{`
resource "cloudflare_dns_record" "caa_legacy" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "CAA"
  ttl     = 3600

  data = {
    flags = 128
    tag   = "issuewild"
    value = "pki.goog"
  }
}`},
		},
		{
			Name: "DNS record without TTL - should add TTL with default value",
			Config: `
resource "cloudflare_record" "mx_test" {
  zone_id  = "0da42c8d2132a9ddaf714f9e7c920711"
  name     = "test.example.com"
  type     = "MX"
  content  = "mx.sendgrid.net"
  priority = 10
}`,
			Expected: []string{`
resource "cloudflare_dns_record" "mx_test" {
  zone_id  = "0da42c8d2132a9ddaf714f9e7c920711"
  name     = "test.example.com"
  type     = "MX"
  content  = "mx.sendgrid.net"
  priority = 10
  ttl      = 1
}`},
		},
		{
			Name: "DNS record with existing TTL - should keep existing value",
			Config: `
resource "cloudflare_record" "a_test_ttl" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "A"
  ttl     = 3600
  content = "192.168.1.1"
}`,
			Expected: []string{`
resource "cloudflare_dns_record" "a_test_ttl" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "A"
  ttl     = 3600
  content = "192.168.1.1"
}`},
		},
		{
			Name: "Multiple CAA records in same file - content renamed to value and TTL added",
			Config: `
resource "cloudflare_record" "caa_test1" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test1.example.com"
  type    = "CAA"
  data {
    flags   = 0
    tag     = "issue"
    content = "letsencrypt.org"
  }
}

resource "cloudflare_record" "caa_test2" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test2.example.com"
  type    = "CAA"
  data {
    flags   = 128
    tag     = "issuewild"
    content = "pki.goog"
  }
}`,
			Expected: []string{`
resource "cloudflare_dns_record" "caa_test1" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test1.example.com"
  type    = "CAA"
  ttl     = 1
  data = {
    flags = 0
    tag   = "issue"
    value = "letsencrypt.org"
  }
}

resource "cloudflare_dns_record" "caa_test2" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test2.example.com"
  type    = "CAA"
  ttl     = 1
  data = {
    flags = 128
    tag   = "issuewild"
    value = "pki.goog"
  }
}`},
		},
	}

	RunTransformationTests(t, tests, transformFileDefault)
}

func TestDNSRecordStateTransformation(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "CAA record v4 format with array data and numeric flags",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "caa_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "CAA",
							"content": "0 issue letsencrypt.org",
							"data": [{
								"flags": 0,
								"tag": "issue",
								"content": "letsencrypt.org"
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "caa_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "CAA",
							"ttl": 1,
							"content": "0 issue letsencrypt.org",
							"created_on": "2024-01-01T00:00:00Z",
							"modified_on": "2024-01-01T00:00:00Z",
							"data": {
								"flags": {
									"value": 0,
									"type": "number"
								},
								"tag": "issue",
								"value": "letsencrypt.org"
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "CAA record with numeric flags 128",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "caa_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "CAA",
							"content": "128 issuewild pki.goog",
							"data": [{
								"flags": 128,
								"tag": "issuewild",
								"content": "pki.goog"
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "caa_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "CAA",
							"ttl": 1,
							"content": "128 issuewild pki.goog",
							"created_on": "2024-01-01T00:00:00Z",
							"modified_on": "2024-01-01T00:00:00Z",
							"data": {
								"flags": {
									"value": 128,
									"type": "number"
								},
								"tag": "issuewild",
								"value": "pki.goog"
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "CAA record already migrated to object format",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "caa_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "CAA",
							"data": {
								"flags": {
									"value": 0,
									"type": "number"
								},
								"tag": "issue",
								"value": "letsencrypt.org"
							}
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "caa_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "CAA",
							"ttl": 1,
							"created_on": "2024-01-01T00:00:00Z",
							"modified_on": "2024-01-01T00:00:00Z",
							"data": {
								"flags": {
									"value": 0,
									"type": "number"
								},
								"tag": "issue",
								"value": "letsencrypt.org"
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "Simple A record should set data field to null",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "a_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "A",
							"content": "192.168.1.1",
							"data": []
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "a_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "A",
							"ttl": 1,
							"content": "192.168.1.1",
							"created_on": "2024-01-01T00:00:00Z",
							"modified_on": "2024-01-01T00:00:00Z"
						}
					}]
				}]
			}`,
		},
		{
			Name: "cloudflare_record (legacy) renamed to cloudflare_dns_record",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_record",
					"name": "caa_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "CAA",
							"content": "0 issue letsencrypt.org",
							"data": [{
								"flags": 0,
								"tag": "issue",
								"content": "letsencrypt.org"
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "caa_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "CAA",
							"ttl": 1,
							"content": "0 issue letsencrypt.org",
							"created_on": "2024-01-01T00:00:00Z",
							"modified_on": "2024-01-01T00:00:00Z",
							"data": {
								"flags": {
									"value": 0,
									"type": "number"
								},
								"tag": "issue",
								"value": "letsencrypt.org"
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "SRV record with array data migration",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "srv_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "_sip._tcp.example.com",
							"type": "SRV",
							"data": [{
								"priority": 10,
								"weight": 60,
								"port": 5060,
								"target": "sipserver.example.com"
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "srv_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "_sip._tcp.example.com",
							"type": "SRV",
							"priority": 10,
							"ttl": 1,
							"created_on": "2024-01-01T00:00:00Z",
							"modified_on": "2024-01-01T00:00:00Z",
							"data": {
								"priority": 10,
								"weight": 60,
								"port": 5060,
								"target": "sipserver.example.com"
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "Record with value field renamed to content",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_record",
					"name": "a_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "A",
							"value": "192.168.1.1",
							"hostname": "test.example.com",
							"allow_overwrite": true
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "a_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "A",
							"ttl": 1,
							"content": "192.168.1.1",
							"created_on": "2024-01-01T00:00:00Z",
							"modified_on": "2024-01-01T00:00:00Z"
						}
					}]
				}]
			}`,
		},
	}

	RunFullStateTransformationTests(t, tests)
}

func TestDNSRecordStateTransformationWithComputedFields(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "TXT record with all computed fields from v4",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [{
					"type": "cloudflare_record",
					"name": "txt_spf",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "7c2d6320347f97de16dd2569e1fcd6b5",
							"name": "static.example.com.terraform.cfapi.net",
							"type": "TXT",
							"value": "v=spf1 include:_spf.mx.cloudflare.net ~all",
							"ttl": 1,
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"created_on": "2025-08-26T05:22:13.523335Z",
							"modified_on": "2025-08-26T05:22:13.523335Z",
							"proxied": false,
							"meta": "{}",
							"tags": ["tf-applied"],
							"tags_modified_on": "2025-08-26T05:22:13Z",
							"settings": {
								"flatten_cname": null,
								"ipv4_only": null,
								"ipv6_only": null
							}
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "txt_spf",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "7c2d6320347f97de16dd2569e1fcd6b5",
							"name": "static.example.com.terraform.cfapi.net",
							"type": "TXT",
							"content": "v=spf1 include:_spf.mx.cloudflare.net ~all",
							"ttl": 1,
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"created_on": "2025-08-26T05:22:13.523335Z",
							"modified_on": "2025-08-26T05:22:13.523335Z",
							"proxied": false,
							"tags": ["tf-applied"],
							"tags_modified_on": "2025-08-26T05:22:13Z"
						}
					}]
				}]
			}`,
		},
		{
			Name: "A record with missing computed fields - should add them",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_record",
					"name": "a_test",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "test-id",
							"name": "test.example.com",
							"type": "A",
							"value": "192.168.1.1",
							"ttl": 3600,
							"zone_id": "test-zone"
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "a_test",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "test-id",
							"name": "test.example.com",
							"type": "A",
							"content": "192.168.1.1",
							"ttl": 3600,
							"zone_id": "test-zone",
							"created_on": "2024-01-01T00:00:00Z",
							"modified_on": "2024-01-01T00:00:00Z"
						}
					}]
				}]
			}`,
		},
		{
			Name: "CAA record with all computed fields",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "caa_google",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "47480c33c49b0240b17dc9168d4442dd",
							"name": "static.example.com.terraform.cfapi.net",
							"type": "CAA",
							"content": "0 issue \"pki.goog\"",
							"ttl": 1,
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"data": [{
								"flags": 0,
								"tag": "issue",
								"value": "pki.goog"
							}],
							"created_on": "2025-08-26T05:21:58Z",
							"modified_on": "2025-08-26T05:21:58Z",
							"proxied": false,
							"meta": "{}",
							"comment": null,
							"comment_modified_on": null,
							"tags": ["tf-applied"],
							"tags_modified_on": "2025-08-26T05:21:58Z",
							"settings": {
								"flatten_cname": null,
								"ipv4_only": null,
								"ipv6_only": null
							}
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "caa_google",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "47480c33c49b0240b17dc9168d4442dd",
							"name": "static.example.com.terraform.cfapi.net",
							"type": "CAA",
							"content": "0 issue \"pki.goog\"",
							"ttl": 1,
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"data": {
								"flags": {
									"value": 0,
									"type": "number"
								},
								"tag": "issue",
								"value": "pki.goog"
							},
							"created_on": "2025-08-26T05:21:58Z",
							"modified_on": "2025-08-26T05:21:58Z",
							"proxied": false,
							"comment": null,
							"comment_modified_on": null,
							"tags": ["tf-applied"],
							"tags_modified_on": "2025-08-26T05:21:58Z"
						}
					}]
				}]
			}`,
		},
	}

	RunFullStateTransformationTests(t, tests)
}

func TestIsDNSRecordResource(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name: "cloudflare_dns_record resource",
			input: `resource "cloudflare_dns_record" "test" {
  zone_id = "test"
  name = "test"
  type = "A"
  value = "1.1.1.1"
  ttl = 300
}`,
			expected: true,
		},
		{
			name: "cloudflare_record resource (old name)",
			input: `resource "cloudflare_record" "test" {
  zone_id = "test"
  name = "test"
  type = "A"
  value = "1.1.1.1"
  ttl = 300
}`,
			expected: true,
		},
		{
			name: "non-dns-record resource",
			input: `resource "cloudflare_zone" "test" {
  zone = "example.com"
}`,
			expected: false,
		},
		{
			name: "data source not resource",
			input: `data "cloudflare_dns_record" "test" {
  zone_id = "test"
}`,
			expected: false,
		},
		{
			name: "resource with single label",
			input: `resource "cloudflare_dns_record" {
  zone_id = "test"
}`,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("Failed to parse input: %v", diags.Error())
			}

			blocks := file.Body().Blocks()
			if len(blocks) != 1 {
				t.Fatalf("Expected 1 block, got %d", len(blocks))
			}

			result := isDNSRecordResource(blocks[0])
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestProcessDNSRecordConfig(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name: "rename cloudflare_record to cloudflare_dns_record",
			input: `resource "cloudflare_record" "test" {
  zone_id = "test"
  name = "test"
  type = "A"
  value = "1.1.1.1"
}`,
			expected: []string{
				`resource "cloudflare_dns_record" "test"`,
				`ttl     = 1`,
			},
		},
		{
			name: "add missing TTL attribute",
			input: `resource "cloudflare_dns_record" "test" {
  zone_id = "test"
  name = "test"
  type = "A"
  value = "1.1.1.1"
}`,
			expected: []string{
				`resource "cloudflare_dns_record" "test"`,
				`ttl     = 1`,
			},
		},
		{
			name: "keep existing TTL",
			input: `resource "cloudflare_dns_record" "test" {
  zone_id = "test"
  name = "test"
  type = "A"
  value = "1.1.1.1"
  ttl = 3600
}`,
			expected: []string{
				`ttl     = 3600`,
			},
		},
		{
			name: "handle multiple DNS records",
			input: `resource "cloudflare_record" "record1" {
  zone_id = "test"
  name = "test1"
  type = "A"
  value = "1.1.1.1"
}

resource "cloudflare_dns_record" "record2" {
  zone_id = "test"
  name = "test2"
  type = "AAAA"
  value = "::1"
  ttl = 300
}`,
			expected: []string{
				`resource "cloudflare_dns_record" "record1"`,
				`resource "cloudflare_dns_record" "record2"`,
				`ttl     = 300`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("Failed to parse input: %v", diags.Error())
			}

			err := ProcessDNSRecordConfig(file)
			assert.NoError(t, err)

			output := string(hclwrite.Format(file.Bytes()))
			for _, exp := range tt.expected {
				assert.Contains(t, output, exp)
			}
		})
	}
}

func TestTransformDNSRecordStateJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		path     string
		expected string
	}{
		{
			name: "add missing TTL in state",
			input: `{
				"type": "cloudflare_dns_record",
				"name": "test",
				"attributes": {
					"zone_id": "test-zone",
					"name": "test.example.com",
					"type": "A",
					"value": "1.1.1.1"
				}
			}`,
			path:     "resources.0.instances.0",
			expected: `"ttl":1`,
		},
		{
			name: "handle CAA record data transformation in state",
			input: `{
				"type": "cloudflare_dns_record",
				"name": "caa",
				"attributes": {
					"zone_id": "test-zone",
					"name": "example.com",
					"type": "CAA",
					"data": {
						"flags": "0",
						"tag": "issue",
						"content": "letsencrypt.org"
					}
				}
			}`,
			path:     "resources.0.instances.0",
			expected: `"value":"letsencrypt.org"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instance := gjson.Parse(tt.input)
			result := transformDNSRecordStateJSON(tt.input, tt.path, instance)
			
			assert.Contains(t, result, tt.expected)
		})
	}
}

func TestDNSRecordComplexTransformations(t *testing.T) {
	tests := []TestCase{
		{
			Name: "SRV record with data block",
			Config: `resource "cloudflare_dns_record" "srv" {
  zone_id = "test"
  name = "_service._proto"
  type = "SRV"
  
  data {
    priority = 10
    weight = 60
    port = 5060
    target = "srv.example.com"
  }
}`,
			Expected: []string{
				`resource "cloudflare_dns_record" "srv"`,
				`type = "SRV"`,
				`ttl  = 1`,
			},
		},
		{
			Name: "MX record with priority",
			Config: `resource "cloudflare_dns_record" "mx" {
  zone_id = "test"
  name = "@"
  type = "MX"
  value = "mail.example.com"
  priority = 10
}`,
			Expected: []string{
				`resource "cloudflare_dns_record" "mx"`,
				`priority = 10`,
				`ttl      = 1`,
			},
		},
		{
			Name: "TXT record with long value",
			Config: `resource "cloudflare_record" "txt" {
  zone_id = "test"
  name = "_dmarc"
  type = "TXT"
  value = "v=DMARC1; p=none; rua=mailto:dmarc@example.com; ruf=mailto:dmarc@example.com; sp=none; adkim=r; aspf=r"
}`,
			Expected: []string{
				`resource "cloudflare_dns_record" "txt"`,
				`type = "TXT"`,
				`ttl  = 1`,
			},
		},
		{
			Name: "CNAME record with proxied flag",
			Config: `resource "cloudflare_record" "cname" {
  zone_id = "test"
  name = "www"
  type = "CNAME"
  value = "example.com"
  proxied = true
}`,
			Expected: []string{
				`resource "cloudflare_dns_record" "cname"`,
				`proxied = true`,
				`ttl     = 1`,
			},
		},
	}
	
	RunTransformationTests(t, tests, transformFileDefault)
}

func TestDNSRecordEdgeCases(t *testing.T) {
	tests := []TestCase{
		{
			Name: "record with computed TTL",
			Config: `resource "cloudflare_dns_record" "computed" {
  zone_id = "test"
  name = "test"
  type = "A"
  value = "1.1.1.1"
  ttl = var.dns_ttl
}`,
			Expected: []string{
				`resource "cloudflare_dns_record" "computed"`,
				`ttl     = var.dns_ttl`,
			},
		},
		{
			Name: "record with allow_overwrite",
			Config: `resource "cloudflare_record" "overwrite" {
  zone_id = "test"
  name = "test"
  type = "A"
  value = "1.1.1.1"
  allow_overwrite = true
}`,
			Expected: []string{
				`resource "cloudflare_dns_record" "overwrite"`,
				`ttl     = 1`,
			},
		},
		{
			Name: "multiple records in same file",
			Config: `resource "cloudflare_record" "a1" {
  zone_id = "test"
  name = "test1"
  type = "A"
  value = "1.1.1.1"
}

resource "cloudflare_record" "a2" {
  zone_id = "test"
  name = "test2"
  type = "A"
  value = "1.1.1.2"
  ttl = 3600
}

resource "cloudflare_dns_record" "a3" {
  zone_id = "test"
  name = "test3"
  type = "A"
  value = "1.1.1.3"
}`,
			Expected: []string{
				`resource "cloudflare_dns_record" "a1"`,
				`resource "cloudflare_dns_record" "a2"`,
				`resource "cloudflare_dns_record" "a3"`,
			},
		},
	}
	
	RunTransformationTests(t, tests, transformFileDefault)
}

func TestDNSRecordWithComments(t *testing.T) {
	tests := []TestCase{
		{
			Name: "preserve comments during transformation",
			Config: `# Main A record for website
resource "cloudflare_record" "main" {
  zone_id = "test"
  name = "@"  # Root domain
  type = "A"
  value = "1.1.1.1" # Cloudflare IP
  # TTL will be added automatically
}`,
			Expected: []string{
				`resource "cloudflare_dns_record" "main"`,
				`# Root domain`,
				`# TTL will be added automatically`,
			},
		},
	}
	
	RunTransformationTests(t, tests, transformFileDefault)
}

func TestDNSRecordDataBlockTransformations(t *testing.T) {
	tests := []TestCase{
		{
			Name: "CAA record with data block string flags",
			Config: `resource "cloudflare_dns_record" "caa" {
  zone_id = "test"
  name = "example.com"
  type = "CAA"
  
  data {
    flags = "0"
    tag = "issue"
    content = "letsencrypt.org"
  }
}`,
			Expected: []string{
				`flags = "0"`,
				`value = "letsencrypt.org"`,
			},
		},
		{
			Name: "SRV record with all data fields",
			Config: `resource "cloudflare_dns_record" "srv" {
  zone_id = "test"
  name = "_sip._tcp"
  type = "SRV"
  
  data {
    priority = 10
    weight = 60
    port = 5060
    target = "sipserver.example.com"
    name = "_sip._tcp"
    proto = "_tcp"
    service = "_sip"
  }
}`,
			Expected: []string{
				`resource "cloudflare_dns_record" "srv"`,
				`priority = 10`,
				`weight   = 60`,
			},
		},
	}
	
	RunTransformationTests(t, tests, transformFileDefault)
}

func TestDNSRecordProxiedAndTTLInteraction(t *testing.T) {
	tests := []TestCase{
		{
			Name: "proxied record with TTL should keep both",
			Config: `resource "cloudflare_dns_record" "proxied" {
  zone_id = "test"
  name = "www"
  type = "CNAME"
  value = "example.com"
  proxied = true
  ttl = 3600
}`,
			Expected: []string{
				`proxied = true`,
				`ttl     = 3600`,
			},
		},
		{
			Name: "non-proxied record without TTL gets default",
			Config: `resource "cloudflare_dns_record" "not_proxied" {
  zone_id = "test"
  name = "mail"
  type = "A"
  value = "1.1.1.1"
  proxied = false
}`,
			Expected: []string{
				`proxied = false`,
				`ttl     = 1`,
			},
		},
	}
	
	RunTransformationTests(t, tests, transformFileDefault)
}
