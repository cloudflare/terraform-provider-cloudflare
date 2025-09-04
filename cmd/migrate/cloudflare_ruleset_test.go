package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestTransformCloudflareRulesetStateJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]interface{}
	}{
		{
			name: "basic rules transformation from indexed to array",
			input: `{
				"resources": [{
					"type": "cloudflare_ruleset",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"id": "abc123",
							"zone_id": "zone1",
							"name": "My ruleset",
							"phase": "http_request_firewall_custom",
							"kind": "zone",
							"rules.#": "2",
							"rules.0.id": "rule1",
							"rules.0.expression": "ip.src eq 1.1.1.1",
							"rules.0.action": "block",
							"rules.0.enabled": true,
							"rules.1.id": "rule2",
							"rules.1.expression": "ip.src eq 2.2.2.2",
							"rules.1.action": "skip",
							"rules.1.action_parameters.#": "1",
							"rules.1.action_parameters.0.ruleset": "current"
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"schema_version": float64(0),
				"attributes": map[string]interface{}{
					"id":      "abc123",
					"zone_id": "zone1",
					"name":    "My ruleset",
					"phase":   "http_request_firewall_custom",
					"kind":    "zone",
					"rules": []interface{}{
						map[string]interface{}{
							"id":         "rule1",
							"expression": "ip.src eq 1.1.1.1",
							"action":     "block",
							"enabled":    true,
						},
						map[string]interface{}{
							"id":         "rule2",
							"expression": "ip.src eq 2.2.2.2",
							"action":     "skip",
							"action_parameters": map[string]interface{}{
								"ruleset": "current",
							},
						},
					},
				},
			},
		},
		{
			name: "headers transformation from list to map",
			input: `{
				"resources": [{
					"type": "cloudflare_ruleset",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"id": "abc123",
							"rules.#": "1",
							"rules.0.action": "rewrite",
							"rules.0.expression": "true",
							"rules.0.action_parameters.#": "1",
							"rules.0.action_parameters.0.headers.#": "2",
							"rules.0.action_parameters.0.headers.0.name": "X-Custom-Header",
							"rules.0.action_parameters.0.headers.0.operation": "set",
							"rules.0.action_parameters.0.headers.0.value": "test-value",
							"rules.0.action_parameters.0.headers.1.name": "X-Another-Header",
							"rules.0.action_parameters.0.headers.1.operation": "remove"
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"schema_version": float64(0),
				"attributes": map[string]interface{}{
					"id": "abc123",
					"rules": []interface{}{
						map[string]interface{}{
							"action":     "rewrite",
							"expression": "true",
							"action_parameters": map[string]interface{}{
								"headers": map[string]interface{}{
									"X-Custom-Header": map[string]interface{}{
										"operation": "set",
										"value":     "test-value",
									},
									"X-Another-Header": map[string]interface{}{
										"operation": "remove",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "log custom fields transformation",
			input: `{
				"resources": [{
					"type": "cloudflare_ruleset",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"id": "abc123",
							"rules.#": "1",
							"rules.0.action": "log_custom_field",
							"rules.0.expression": "true",
							"rules.0.action_parameters.#": "1",
							"rules.0.action_parameters.0.cookie_fields.#": "2",
							"rules.0.action_parameters.0.cookie_fields.0": "session_id",
							"rules.0.action_parameters.0.cookie_fields.1": "user_token",
							"rules.0.action_parameters.0.request_fields.#": "1",
							"rules.0.action_parameters.0.request_fields.0": "cf.bot_score"
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"schema_version": float64(0),
				"attributes": map[string]interface{}{
					"id": "abc123",
					"rules": []interface{}{
						map[string]interface{}{
							"action":     "log_custom_field",
							"expression": "true",
							"action_parameters": map[string]interface{}{
								"cookie_fields": []interface{}{
									map[string]interface{}{"name": "session_id"},
									map[string]interface{}{"name": "user_token"},
								},
								"request_fields": []interface{}{
									map[string]interface{}{"name": "cf.bot_score"},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "nested blocks transformation",
			input: `{
				"resources": [{
					"type": "cloudflare_ruleset",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"id": "abc123",
							"rules.#": "1",
							"rules.0.action": "set_cache_settings",
							"rules.0.expression": "true",
							"rules.0.action_parameters.#": "1",
							"rules.0.action_parameters.0.cache": true,
							"rules.0.action_parameters.0.edge_ttl.#": "1",
							"rules.0.action_parameters.0.edge_ttl.0.mode": "override_origin",
							"rules.0.action_parameters.0.edge_ttl.0.default": 3600,
							"rules.0.action_parameters.0.cache_key.#": "1",
							"rules.0.action_parameters.0.cache_key.0.cache_by_device_type": true,
							"rules.0.action_parameters.0.cache_key.0.custom_key.#": "1",
							"rules.0.action_parameters.0.cache_key.0.custom_key.0.query_string.#": "1",
							"rules.0.action_parameters.0.cache_key.0.custom_key.0.query_string.0.include.#": "2",
							"rules.0.action_parameters.0.cache_key.0.custom_key.0.query_string.0.include.0": "param1",
							"rules.0.action_parameters.0.cache_key.0.custom_key.0.query_string.0.include.1": "param2"
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"schema_version": float64(0),
				"attributes": map[string]interface{}{
					"id": "abc123",
					"rules": []interface{}{
						map[string]interface{}{
							"action":     "set_cache_settings",
							"expression": "true",
							"action_parameters": map[string]interface{}{
								"cache": true,
								"edge_ttl": map[string]interface{}{
									"mode":    "override_origin",
									"default": float64(3600),
								},
								"cache_key": map[string]interface{}{
									"cache_by_device_type": true,
									"custom_key": map[string]interface{}{
										"query_string": map[string]interface{}{
											"include": []interface{}{"param1", "param2"},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ratelimit and logging transformation",
			input: `{
				"resources": [{
					"type": "cloudflare_ruleset",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"id": "abc123",
							"rules.#": "1",
							"rules.0.action": "challenge",
							"rules.0.expression": "true",
							"rules.0.ratelimit.#": "1",
							"rules.0.ratelimit.0.characteristics.#": "2",
							"rules.0.ratelimit.0.characteristics.0": "ip.src",
							"rules.0.ratelimit.0.characteristics.1": "http.request.uri.path",
							"rules.0.ratelimit.0.period": 60,
							"rules.0.ratelimit.0.requests_per_period": 100,
							"rules.0.logging.#": "1",
							"rules.0.logging.0.enabled": true
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"schema_version": float64(0),
				"attributes": map[string]interface{}{
					"id": "abc123",
					"rules": []interface{}{
						map[string]interface{}{
							"action":     "challenge",
							"expression": "true",
							"ratelimit": map[string]interface{}{
								"characteristics":     []interface{}{"ip.src", "http.request.uri.path"},
								"period":              float64(60),
								"requests_per_period": float64(100),
							},
							"logging": map[string]interface{}{
								"enabled": true,
							},
						},
					},
				},
			},
		},
		{
			name: "complex WAF overrides with categories and rules",
			input: `{
				"resources": [{
					"type": "cloudflare_ruleset",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"id": "waf123",
							"zone_id": "zone1",
							"name": "WAF ruleset",
							"phase": "http_request_firewall_managed",
							"kind": "zone",
							"rules.#": "1",
							"rules.0.expression": "ip.src eq 1.1.1.1",
							"rules.0.action": "execute",
							"rules.0.action_parameters.#": "1",
							"rules.0.action_parameters.0.id": "4814384a9e5d4991b9815dcfc25d2f1f",
							"rules.0.action_parameters.0.overrides.#": "1",
							"rules.0.action_parameters.0.overrides.0.action": "log",
							"rules.0.action_parameters.0.overrides.0.enabled": true,
							"rules.0.action_parameters.0.overrides.0.categories.#": "2",
							"rules.0.action_parameters.0.overrides.0.categories.0.category": "language-java",
							"rules.0.action_parameters.0.overrides.0.categories.0.action": "block",
							"rules.0.action_parameters.0.overrides.0.categories.1.category": "language-php",
							"rules.0.action_parameters.0.overrides.0.categories.1.enabled": false,
							"rules.0.action_parameters.0.overrides.0.rules.#": "2",
							"rules.0.action_parameters.0.overrides.0.rules.0.id": "rule1",
							"rules.0.action_parameters.0.overrides.0.rules.0.action": "block",
							"rules.0.action_parameters.0.overrides.0.rules.1.id": "rule2",
							"rules.0.action_parameters.0.overrides.0.rules.1.enabled": false,
							"rules.0.action_parameters.0.overrides.0.rules.1.score_threshold": 40
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"schema_version": float64(0),
				"attributes": map[string]interface{}{
					"id":      "waf123",
					"zone_id": "zone1",
					"name":    "WAF ruleset",
					"phase":   "http_request_firewall_managed",
					"kind":    "zone",
					"rules": []interface{}{
						map[string]interface{}{
							"expression": "ip.src eq 1.1.1.1",
							"action":     "execute",
							"action_parameters": map[string]interface{}{
								"id": "4814384a9e5d4991b9815dcfc25d2f1f",
								"overrides": map[string]interface{}{
									"action":  "log",
									"enabled": true,
									"categories": []interface{}{
										map[string]interface{}{
											"category": "language-java",
											"action":   "block",
										},
										map[string]interface{}{
											"category": "language-php",
											"enabled":  false,
										},
									},
									"rules": []interface{}{
										map[string]interface{}{
											"id":     "rule1",
											"action": "block",
										},
										map[string]interface{}{
											"id":              "rule2",
											"enabled":         false,
											"score_threshold": float64(40),
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "URI rewrite transformation",
			input: `{
				"resources": [{
					"type": "cloudflare_ruleset",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"id": "rewrite123",
							"zone_id": "zone1",
							"name": "URI rewrite",
							"phase": "http_request_transform",
							"kind": "zone",
							"rules.#": "1",
							"rules.0.expression": "http.request.uri.path contains \"/api\"",
							"rules.0.action": "rewrite",
							"rules.0.action_parameters.#": "1",
							"rules.0.action_parameters.0.uri.#": "1",
							"rules.0.action_parameters.0.uri.0.path.#": "1",
							"rules.0.action_parameters.0.uri.0.path.0.value": "/v2/api",
							"rules.0.action_parameters.0.uri.0.query.#": "1",
							"rules.0.action_parameters.0.uri.0.query.0.expression": "concat(\"version=2&\", http.request.uri.query)"
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"schema_version": float64(0),
				"attributes": map[string]interface{}{
					"id":      "rewrite123",
					"zone_id": "zone1",
					"name":    "URI rewrite",
					"phase":   "http_request_transform",
					"kind":    "zone",
					"rules": []interface{}{
						map[string]interface{}{
							"expression": `http.request.uri.path contains "/api"`,
							"action":     "rewrite",
							"action_parameters": map[string]interface{}{
								"uri": map[string]interface{}{
									"path": map[string]interface{}{
										"value": "/v2/api",
									},
									"query": map[string]interface{}{
										"expression": `concat("version=2&", http.request.uri.query)`,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "redirect with from_list and from_value",
			input: `{
				"resources": [{
					"type": "cloudflare_ruleset",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"id": "redirect123",
							"account_id": "account1",
							"name": "Redirect rules",
							"phase": "http_request_redirect",
							"kind": "root",
							"rules.#": "2",
							"rules.0.expression": "ip.src eq 1.1.1.1",
							"rules.0.action": "redirect",
							"rules.0.action_parameters.#": "1",
							"rules.0.action_parameters.0.from_list.#": "1",
							"rules.0.action_parameters.0.from_list.0.key": "http.request.full_uri",
							"rules.0.action_parameters.0.from_list.0.name": "redirect_list",
							"rules.1.expression": "ip.src eq 2.2.2.2",
							"rules.1.action": "redirect",
							"rules.1.action_parameters.#": "1",
							"rules.1.action_parameters.0.from_value.#": "1",
							"rules.1.action_parameters.0.from_value.0.status_code": 301,
							"rules.1.action_parameters.0.from_value.0.preserve_query_string": true,
							"rules.1.action_parameters.0.from_value.0.target_url.#": "1",
							"rules.1.action_parameters.0.from_value.0.target_url.0.expression": "concat(\"https://example.com\", http.request.uri.path)"
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"schema_version": float64(0),
				"attributes": map[string]interface{}{
					"id":         "redirect123",
					"account_id": "account1",
					"name":       "Redirect rules",
					"phase":      "http_request_redirect",
					"kind":       "root",
					"rules": []interface{}{
						map[string]interface{}{
							"expression": "ip.src eq 1.1.1.1",
							"action":     "redirect",
							"action_parameters": map[string]interface{}{
								"from_list": map[string]interface{}{
									"key":  "http.request.full_uri",
									"name": "redirect_list",
								},
							},
						},
						map[string]interface{}{
							"expression": "ip.src eq 2.2.2.2",
							"action":     "redirect",
							"action_parameters": map[string]interface{}{
								"from_value": map[string]interface{}{
									"status_code":          float64(301),
									"preserve_query_string": true,
									"target_url": map[string]interface{}{
										"expression": `concat("https://example.com", http.request.uri.path)`,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "cache settings with edge_ttl status_code_ttl",
			input: `{
				"resources": [{
					"type": "cloudflare_ruleset",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"id": "cache123",
							"zone_id": "zone1",
							"name": "Cache settings",
							"phase": "http_request_cache_settings",
							"kind": "zone",
							"rules.#": "1",
							"rules.0.expression": "http.request.uri.path contains \"/static\"",
							"rules.0.action": "set_cache_settings",
							"rules.0.action_parameters.#": "1",
							"rules.0.action_parameters.0.cache": true,
							"rules.0.action_parameters.0.edge_ttl.#": "1",
							"rules.0.action_parameters.0.edge_ttl.0.mode": "override_origin",
							"rules.0.action_parameters.0.edge_ttl.0.default": 7200,
							"rules.0.action_parameters.0.edge_ttl.0.status_code_ttl.#": "2",
							"rules.0.action_parameters.0.edge_ttl.0.status_code_ttl.0.status_code": 200,
							"rules.0.action_parameters.0.edge_ttl.0.status_code_ttl.0.value": 86400,
							"rules.0.action_parameters.0.edge_ttl.0.status_code_ttl.1.value": 300,
							"rules.0.action_parameters.0.edge_ttl.0.status_code_ttl.1.status_code_range.#": "1",
							"rules.0.action_parameters.0.edge_ttl.0.status_code_ttl.1.status_code_range.0.from": 400,
							"rules.0.action_parameters.0.edge_ttl.0.status_code_ttl.1.status_code_range.0.to": 499,
							"rules.0.action_parameters.0.cache_reserve.#": "1",
							"rules.0.action_parameters.0.cache_reserve.0.eligible": true,
							"rules.0.action_parameters.0.cache_reserve.0.minimum_file_size": 10485760
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"schema_version": float64(0),
				"attributes": map[string]interface{}{
					"id":      "cache123",
					"zone_id": "zone1",
					"name":    "Cache settings",
					"phase":   "http_request_cache_settings",
					"kind":    "zone",
					"rules": []interface{}{
						map[string]interface{}{
							"expression": `http.request.uri.path contains "/static"`,
							"action":     "set_cache_settings",
							"action_parameters": map[string]interface{}{
								"cache": true,
								"edge_ttl": map[string]interface{}{
									"mode":    "override_origin",
									"default": float64(7200),
									"status_code_ttl": []interface{}{
										map[string]interface{}{
											"status_code": float64(200),
											"value":       float64(86400),
										},
										map[string]interface{}{
											"value": float64(300),
											"status_code_range": map[string]interface{}{
												"from": float64(400),
												"to":   float64(499),
											},
										},
									},
								},
								"cache_reserve": map[string]interface{}{
									"eligible":          true,
									"minimum_file_size": float64(10485760),
								},
							},
						},
					},
				},
			},
		},
		{
			name: "multiple rules with different actions",
			input: `{
				"resources": [{
					"type": "cloudflare_ruleset",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"id": "multi123",
							"zone_id": "zone1",
							"name": "Multi-rule",
							"phase": "http_request_firewall_custom",
							"kind": "zone",
							"rules.#": "3",
							"rules.0.expression": "ip.src eq 1.1.1.1",
							"rules.0.action": "block",
							"rules.0.enabled": true,
							"rules.1.expression": "http.user_agent contains \"bot\"",
							"rules.1.action": "challenge",
							"rules.1.ratelimit.#": "1",
							"rules.1.ratelimit.0.characteristics.#": "1",
							"rules.1.ratelimit.0.characteristics.0": "ip.src",
							"rules.1.ratelimit.0.period": 60,
							"rules.1.ratelimit.0.requests_per_period": 10,
							"rules.1.ratelimit.0.mitigation_timeout": 600,
							"rules.2.expression": "ip.src.country eq \"CN\"",
							"rules.2.action": "skip",
							"rules.2.action_parameters.#": "1",
							"rules.2.action_parameters.0.ruleset": "current",
							"rules.2.action_parameters.0.phases.#": "2",
							"rules.2.action_parameters.0.phases.0": "http_ratelimit",
							"rules.2.action_parameters.0.phases.1": "http_request_firewall_managed",
							"rules.2.logging.#": "1",
							"rules.2.logging.0.enabled": true
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"schema_version": float64(0),
				"attributes": map[string]interface{}{
					"id":      "multi123",
					"zone_id": "zone1",
					"name":    "Multi-rule",
					"phase":   "http_request_firewall_custom",
					"kind":    "zone",
					"rules": []interface{}{
						map[string]interface{}{
							"expression": "ip.src eq 1.1.1.1",
							"action":     "block",
							"enabled":    true,
						},
						map[string]interface{}{
							"expression": `http.user_agent contains "bot"`,
							"action":     "challenge",
							"ratelimit": map[string]interface{}{
								"characteristics":     []interface{}{"ip.src"},
								"period":              float64(60),
								"requests_per_period": float64(10),
								"mitigation_timeout":  float64(600),
							},
						},
						map[string]interface{}{
							"expression": `ip.src.country eq "CN"`,
							"action":     "skip",
							"action_parameters": map[string]interface{}{
								"ruleset": "current",
								"phases":  []interface{}{"http_ratelimit", "http_request_firewall_managed"},
							},
							"logging": map[string]interface{}{
								"enabled": true,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Transform the state
			result := transformCloudflareRulesetStateJSON(tt.input, "resources.0.instances.0")

			// Extract the transformed instance
			instance := gjson.Get(result, "resources.0.instances.0")
			require.True(t, instance.Exists(), "Instance should exist after transformation")

			// Convert to map for comparison
			var actual map[string]interface{}
			err := json.Unmarshal([]byte(instance.Raw), &actual)
			require.NoError(t, err, "Should unmarshal transformed instance")

			// Compare the results
			require.Equal(t, tt.expected, actual, "Transformed state should match expected")

			// Verify no indexed keys remain
			attrs := gjson.Get(result, "resources.0.instances.0.attributes")
			attrs.ForEach(func(key, value gjson.Result) bool {
				keyStr := key.String()
				// Check that no indexed rule keys remain
				require.NotContains(t, keyStr, "rules.#", "Should not contain rules count")
				require.NotRegexp(t, `^rules\.\d+`, keyStr, "Should not contain indexed rule keys")
				return true
			})
		})
	}
}

