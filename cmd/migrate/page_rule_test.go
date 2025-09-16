package main

import (
	"testing"
)

func TestPageRuleConfigTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "transforms cache_key_fields query_string ignore false to include",
			Config: `
resource "cloudflare_page_rule" "test" {
  zone_id = "example.com"
  target  = "example.com/*"

  actions {
    cache_level = "aggressive"

    cache_key_fields {
      query_string {
        ignore = false
      }
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_page_rule" "test" {
  zone_id = "example.com"
  target  = "example.com/*"

  actions = {
    cache_level = "aggressive"

    cache_key_fields = {
      query_string = {
        include = ["*"]
      }
    }
  }
}`},
		},
		{
			Name: "transforms cache_key_fields query_string ignore true to exclude",
			Config: `
resource "cloudflare_page_rule" "test" {
  zone_id = "example.com"
  target  = "example.com/*"

  actions {
    cache_key_fields {
      query_string {
        ignore = true
      }
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_page_rule" "test" {
  zone_id = "example.com"
  target  = "example.com/*"

  actions = {
    cache_key_fields = {
      query_string = {
        exclude = ["*"]
      }
    }
  }
}`},
		},
		{
			Name: "removes empty browser_cache_ttl",
			Config: `
resource "cloudflare_page_rule" "test" {
  zone_id = "example.com"
  target  = "example.com/*"

  actions {
    browser_cache_ttl = ""
    edge_cache_ttl = 0
    cache_level = "aggressive"
  }
}`,
			Expected: []string{`
resource "cloudflare_page_rule" "test" {
  zone_id = "example.com"
  target  = "example.com/*"

  actions = {
    cache_level = "aggressive"
  }
}`},
		},
		{
			Name: "preserves valid numeric values",
			Config: `
resource "cloudflare_page_rule" "test" {
  zone_id = "example.com"
  target  = "example.com/*"

  actions {
    browser_cache_ttl = 3600
    edge_cache_ttl = 7200
    cache_level = "aggressive"
  }
}`,
			Expected: []string{`
resource "cloudflare_page_rule" "test" {
  zone_id = "example.com"
  target  = "example.com/*"

  actions = {
    browser_cache_ttl = 3600
    edge_cache_ttl    = 7200
    cache_level       = "aggressive"
  }
}`},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}

func TestPageRuleStateTransformation(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "fixes empty string browser_cache_ttl in state",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {
								"browser_cache_ttl": "",
								"edge_cache_ttl": "",
								"cache_level": "aggressive"
							}
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {
								"cache_level": "aggressive"
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "transforms cache_key_fields query_string ignore in state",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {
								"cache_key_fields": {
									"query_string": {
										"ignore": false
									}
								}
							}
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {
								"cache_key_fields": {
									"query_string": {
										"include": ["*"]
									}
								}
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "transforms boolean string values to proper booleans",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {
								"always_use_https": "on",
								"cache_deception_armor": "off",
								"cache_level": "aggressive"
							}
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {
								"always_use_https": true,
								"cache_deception_armor": "off",
								"cache_level": "aggressive"
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles forwarding_url empty array",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {
								"forwarding_url": []
							}
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {}
						}
					}]
				}]
			}`,
		},
		{
			Name: "unwraps single-element forwarding_url array",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {
								"forwarding_url": [{
									"status_code": 301,
									"url": "https://new.example.com"
								}]
							}
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {
								"forwarding_url": {
									"status_code": 301,
									"url": "https://new.example.com"
								}
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles cache_ttl_by_status as array instead of object",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {
								"cache_level": "aggressive",
								"cache_ttl_by_status": [{
									"200": "86400",
									"404": "300"
								}],
								"minify": [{
									"html": "on",
									"css": "on",
									"js": "off"
								}]
							}
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {
								"cache_level": "aggressive",
								"cache_ttl_by_status": {
									"200": "86400",
									"404": "300"
								}
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles cache_key_fields as array instead of object",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {
								"cache_level": "aggressive",
								"cache_key_fields": [{
									"query_string": [{
										"ignore": false
									}],
									"cookie": [],
									"header": [{
										"include": ["X-Custom-Header"]
									}]
								}]
							}
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {
								"cache_level": "aggressive",
								"cache_key_fields": {
									"query_string": {
										"include": ["*"]
									},
									"header": {
										"check_presence": [],
										"exclude": [],
										"include": ["X-Custom-Header"]
									}
								}
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles actions as array instead of object",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": [{
								"browser_cache_ttl": "",
								"cache_level": "aggressive",
								"always_use_https": "on"
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {
								"cache_level": "aggressive",
								"always_use_https": true
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "transforms v4 cache_ttl_by_status with codes and ttl fields",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {
								"cache_level": "aggressive",
								"cache_ttl_by_status": [
									{"codes": "200", "ttl": 86400},
									{"codes": "404", "ttl": 300}
								]
							}
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"actions": {
								"cache_level": "aggressive",
								"cache_ttl_by_status": {
									"200": "86400",
									"404": "300"
								}
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "comprehensive page_rule transformation",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"schema_version": 2,
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"priority": 1,
							"actions": {
								"browser_cache_ttl": "",
								"edge_cache_ttl": 0,
								"always_use_https": "on",
								"cache_deception_armor": "off",
								"cache_level": "aggressive",
								"security_level": "high",
								"forwarding_url": [],
								"cache_key_fields": {
									"query_string": {
										"ignore": true
									}
								}
							}
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_page_rule",
					"name": "test",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "rule-123",
							"target": "example.com/*",
							"priority": 1,
							"actions": {
								"always_use_https": true,
								"cache_deception_armor": "off",
								"cache_level": "aggressive",
								"security_level": "high",
								"cache_key_fields": {
									"query_string": {
										"exclude": ["*"]
									}
								}
							}
						}
					}]
				}]
			}`,
		},
	}

	RunFullStateTransformationTests(t, tests)
}