package main

import (
	"testing"
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
