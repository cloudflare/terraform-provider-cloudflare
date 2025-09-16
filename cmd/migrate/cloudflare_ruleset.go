package main

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// isCloudflareRulesetResource checks if a block is a cloudflare_ruleset resource
func isCloudflareRulesetResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 1 &&
		block.Labels()[0] == "cloudflare_ruleset"
}

// transformCloudflareRulesetStateJSON handles v4 to v5 state migration for cloudflare_ruleset
func transformCloudflareRulesetStateJSON(json string, instancePath string) string {
	attrPath := instancePath + ".attributes"
	result := json

	// Set schema_version to 0 for v5
	result, _ = sjson.Set(result, instancePath+".schema_version", 0)

	// Handle rules transformation from indexed format to array
	// First check if rules is already a JSON array
	rulesValue := gjson.Get(json, attrPath+".rules")
	if rulesValue.Exists() && rulesValue.IsArray() {
		// Rules is already a JSON array, need to process action_parameters within it
		var rules []map[string]interface{}
		for _, ruleVal := range rulesValue.Array() {
			if ruleMap, ok := ruleVal.Value().(map[string]interface{}); ok {
				// Process the rule to convert arrays to objects where needed
				ruleMap = convertArraysToObjects(ruleMap)

				// Ensure disable_railgun is completely removed from action_parameters
				if ap, ok := ruleMap["action_parameters"].(map[string]interface{}); ok {
					delete(ap, "disable_railgun")
				}

				rules = append(rules, ruleMap)
			}
		}
		// Set the processed rules back
		result, _ = sjson.Set(result, attrPath+".rules", rules)
		return result
	}

	// Fall back to indexed format processing
	rulesCount := gjson.Get(json, attrPath+`.rules\.#`)
	if rulesCount.Exists() && rulesCount.Int() > 0 {
		var rules []map[string]interface{}

		for i := int64(0); i < rulesCount.Int(); i++ {
			rule := make(map[string]interface{})

			// Copy basic rule attributes
			if val := gjson.Get(json, fmt.Sprintf(`%s.rules\.%d\.id`, attrPath, i)); val.Exists() {
				rule["id"] = val.String()
			}
			if val := gjson.Get(json, fmt.Sprintf(`%s.rules\.%d\.ref`, attrPath, i)); val.Exists() {
				rule["ref"] = val.String()
			}
			if val := gjson.Get(json, fmt.Sprintf(`%s.rules\.%d\.enabled`, attrPath, i)); val.Exists() {
				rule["enabled"] = val.Bool()
			}
			if val := gjson.Get(json, fmt.Sprintf(`%s.rules\.%d\.description`, attrPath, i)); val.Exists() {
				rule["description"] = val.String()
			}
			if val := gjson.Get(json, fmt.Sprintf(`%s.rules\.%d\.expression`, attrPath, i)); val.Exists() {
				rule["expression"] = val.String()
			}
			if val := gjson.Get(json, fmt.Sprintf(`%s.rules\.%d\.action`, attrPath, i)); val.Exists() {
				rule["action"] = val.String()
			}

			// Handle action_parameters transformation
			// First check if action_parameters exists in any format
			actionParamsCount := gjson.Get(json, fmt.Sprintf(`%s.rules\.%d\.action_parameters\.#`, attrPath, i))

			if actionParamsCount.Exists() && actionParamsCount.Int() > 0 {
				// v4 has action_parameters as an indexed list with max 1 item
				// v5 expects it as a single nested object
				actionParams := make(map[string]interface{})

				// Copy all simple action parameter fields
				// Note: disable_railgun was removed in v5, so we exclude it
				simpleFields := []string{
					"additional_cacheable_ports", "automatic_https_rewrites", "bic", "cache",
					"content", "content_type", "disable_apps", "disable_zaraz",
					"disable_rum", "fonts", "email_obfuscation", "host_header", "hotlink_protection",
					"id", "increment", "mirage", "opportunistic_encryption", "origin_cache_control",
					"polish", "products", "read_timeout", "respect_strong_etags",
					"rocket_loader", "rules", "ruleset", "rulesets", "security_level",
					"server_side_excludes", "ssl", "status_code", "sxg", "origin_error_page_passthru",
				}

				for _, field := range simpleFields {
					path := fmt.Sprintf(`%s.rules\.%d\.action_parameters\.0\.%s`, attrPath, i, field)
					if val := gjson.Get(json, path); val.Exists() {
						actionParams[field] = val.Value()
					}
				}

				// Handle array fields that are simple string arrays
				arrayFields := []string{"phases"}
				for _, field := range arrayFields {
					fieldCount := gjson.Get(json, fmt.Sprintf(`%s.rules\.%d\.action_parameters\.0\.%s\.#`, attrPath, i, field))
					if fieldCount.Exists() && fieldCount.Int() > 0 {
						var arr []interface{}
						for j := int64(0); j < fieldCount.Int(); j++ {
							if val := gjson.Get(json, fmt.Sprintf(`%s.rules\.%d\.action_parameters\.0\.%s\.%d`, attrPath, i, field, j)); val.Exists() {
								arr = append(arr, val.String())
							}
						}
						if len(arr) > 0 {
							actionParams[field] = arr
						}
					}
				}

				// Handle headers transformation (from list to map)
				headersCount := gjson.Get(json, fmt.Sprintf(`%s.rules\.%d\.action_parameters\.0\.headers\.#`, attrPath, i))
				if headersCount.Exists() && headersCount.Int() > 0 {
					headers := make(map[string]interface{})
					for j := int64(0); j < headersCount.Int(); j++ {
						headerPath := fmt.Sprintf(`%s.rules\.%d\.action_parameters\.0\.headers\.%d`, attrPath, i, j)
						name := gjson.Get(json, headerPath+`\.name`).String()
						if name != "" {
							header := make(map[string]interface{})
							if val := gjson.Get(json, headerPath+`\.operation`); val.Exists() {
								header["operation"] = val.String()
							}
							if val := gjson.Get(json, headerPath+`\.value`); val.Exists() {
								header["value"] = val.String()
							}
							if val := gjson.Get(json, headerPath+`\.expression`); val.Exists() {
								header["expression"] = val.String()
							}
							headers[name] = header
						}
					}
					if len(headers) > 0 {
						actionParams["headers"] = headers
					}
				}

				// Handle nested single blocks that need to be converted from indexed to object
				singleBlocks := []struct {
					name   string
					fields []string
				}{
					{"algorithms", []string{"name"}},
					{"uri", []string{"origin"}},
					{"matched_data", []string{"public_key"}},
					{"response", []string{"status_code", "content_type", "content"}},
					{"autominify", []string{"html", "css", "js"}},
					{"edge_ttl", []string{"mode", "default"}},
					{"browser_ttl", []string{"mode", "default"}},
					{"serve_stale", []string{"disable_stale_while_updating"}},
					{"cache_key", []string{"cache_by_device_type", "ignore_query_strings_order", "cache_deception_armor"}},
					{"cache_reserve", []string{"eligible", "minimum_file_size"}},
					{"from_list", []string{"name", "key"}},
					{"from_value", []string{"preserve_query_string", "status_code"}},
					{"origin", []string{"host", "port"}},
					{"sni", []string{"value"}},
					{"overrides", []string{"enabled", "action", "sensitivity_level"}},
				}

				for _, block := range singleBlocks {
					blockCount := gjson.Get(json, fmt.Sprintf(`%s.rules\.%d\.action_parameters\.0\.%s\.#`, attrPath, i, block.name))
					if blockCount.Exists() && blockCount.Int() > 0 {
						// Handle nested structures within blocks with special handling
						if block.name == "cache_key" {
							actionParams[block.name] = transformCacheKey(json, attrPath, i)
						} else if block.name == "edge_ttl" || block.name == "browser_ttl" {
							actionParams[block.name] = transformEdgeTTL(json, attrPath, i, block.name)
						} else if block.name == "from_value" {
							actionParams[block.name] = transformFromValue(json, attrPath, i)
						} else if block.name == "overrides" {
							actionParams[block.name] = transformOverrides(json, attrPath, i)
						} else if block.name == "uri" {
							actionParams[block.name] = transformURI(json, attrPath, i)
						} else {
							// Simple blocks with just the listed fields
							blockObj := make(map[string]interface{})
							for _, field := range block.fields {
								fieldPath := fmt.Sprintf(`%s.rules\.%d\.action_parameters\.0\.%s\.0\.%s`, attrPath, i, block.name, field)
								if val := gjson.Get(json, fieldPath); val.Exists() {
									blockObj[field] = val.Value()
								}
							}
							if len(blockObj) > 0 {
								actionParams[block.name] = blockObj
							}
						}
					}
				}

				// Handle cookie_fields, request_fields, response_fields transformation
				// From SetAttribute[String] to ListNestedAttribute with name field
				for _, field := range []string{"cookie_fields", "request_fields", "response_fields"} {
					fieldCount := gjson.Get(json, fmt.Sprintf(`%s.rules\.%d\.action_parameters\.0\.%s\.#`, attrPath, i, field))
					if fieldCount.Exists() && fieldCount.Int() > 0 {
						var fieldList []map[string]interface{}
						for j := int64(0); j < fieldCount.Int(); j++ {
							fieldVal := gjson.Get(json, fmt.Sprintf(`%s.rules\.%d\.action_parameters\.0\.%s\.%d`, attrPath, i, field, j))
							if fieldVal.Exists() {
								fieldList = append(fieldList, map[string]interface{}{
									"name": fieldVal.String(),
								})
							}
						}
						actionParams[field] = fieldList
					}
				}

				rule["action_parameters"] = actionParams
			}

			// Handle ratelimit transformation
			rateLimitCount := gjson.Get(json, fmt.Sprintf(`%s.rules\.%d\.ratelimit\.#`, attrPath, i))
			if rateLimitCount.Exists() && rateLimitCount.Int() > 0 {
				ratelimit := make(map[string]interface{})
				rlPath := fmt.Sprintf(`%s.rules\.%d\.ratelimit\.0`, attrPath, i)

				// Handle characteristics array
				charCount := gjson.Get(json, rlPath+`\.characteristics\.#`)
				if charCount.Exists() && charCount.Int() > 0 {
					var chars []string
					for j := int64(0); j < charCount.Int(); j++ {
						if val := gjson.Get(json, fmt.Sprintf(`%s\.characteristics\.%d`, rlPath, j)); val.Exists() {
							chars = append(chars, val.String())
						}
					}
					ratelimit["characteristics"] = chars
				}

				// Copy simple ratelimit fields
				rlFields := []string{"period", "requests_per_period", "score_per_period",
					"score_response_header_name", "mitigation_timeout", "counting_expression", "requests_to_origin"}
				for _, field := range rlFields {
					if val := gjson.Get(json, rlPath+`\.`+field); val.Exists() {
						ratelimit[field] = val.Value()
					}
				}

				if len(ratelimit) > 0 {
					rule["ratelimit"] = ratelimit
				}
			}

			// Handle exposed_credential_check transformation
			exposedCredCount := gjson.Get(json, fmt.Sprintf(`%s.rules\.%d\.exposed_credential_check\.#`, attrPath, i))
			if exposedCredCount.Exists() && exposedCredCount.Int() > 0 {
				exposedCred := make(map[string]interface{})
				ecPath := fmt.Sprintf(`%s.rules\.%d\.exposed_credential_check\.0`, attrPath, i)

				if val := gjson.Get(json, ecPath+`\.username_expression`); val.Exists() {
					exposedCred["username_expression"] = val.String()
				}
				if val := gjson.Get(json, ecPath+`\.password_expression`); val.Exists() {
					exposedCred["password_expression"] = val.String()
				}

				if len(exposedCred) > 0 {
					rule["exposed_credential_check"] = exposedCred
				}
			}

			// Handle logging transformation
			loggingCount := gjson.Get(json, fmt.Sprintf(`%s.rules\.%d\.logging\.#`, attrPath, i))
			if loggingCount.Exists() && loggingCount.Int() > 0 {
				logging := make(map[string]interface{})
				logPath := fmt.Sprintf(`%s.rules\.%d\.logging\.0`, attrPath, i)

				if val := gjson.Get(json, logPath+`\.enabled`); val.Exists() {
					logging["enabled"] = val.Bool()
				}

				if len(logging) > 0 {
					rule["logging"] = logging
				}
			}

			rules = append(rules, rule)
		}

		// Set the rules as an array
		result, _ = sjson.Set(result, attrPath+".rules", rules)

		// Clean up all indexed keys
		result = cleanupRulesetIndexedKeys(result, attrPath, rulesCount.Int())
	}

	fmt.Println("before")
	result = removeDisableRailgunFromPath(result, instancePath)
	fmt.Println("after")
	return result
}

// Helper function to transform cache_key structure
func transformCacheKey(json string, attrPath string, ruleIdx int64) map[string]interface{} {
	cacheKey := make(map[string]interface{})
	ckPath := fmt.Sprintf(`%s.rules\.%d\.action_parameters\.0\.cache_key\.0`, attrPath, ruleIdx)

	// Simple fields
	fields := []string{"cache_by_device_type", "ignore_query_strings_order", "cache_deception_armor"}
	for _, field := range fields {
		if val := gjson.Get(json, ckPath+`\.`+field); val.Exists() {
			cacheKey[field] = val.Bool()
		}
	}

	// Handle custom_key
	customKeyCount := gjson.Get(json, ckPath+`\.custom_key\.#`)
	if customKeyCount.Exists() && customKeyCount.Int() > 0 {
		customKey := make(map[string]interface{})
		ckCustomPath := ckPath + `\.custom_key\.0`

		// Handle query_string
		qsCount := gjson.Get(json, ckCustomPath+`\.query_string\.#`)
		if qsCount.Exists() && qsCount.Int() > 0 {
			qs := make(map[string]interface{})
			qsPath := ckCustomPath + `\.query_string\.0`

			// Handle include array - transform to v5 structure
			includeCount := gjson.Get(json, qsPath+`\.include\.#`)
			if includeCount.Exists() && includeCount.Int() > 0 {
				var includes []string
				for j := int64(0); j < includeCount.Int(); j++ {
					if val := gjson.Get(json, fmt.Sprintf(`%s\.include\.%d`, qsPath, j)); val.Exists() {
						includes = append(includes, val.String())
					}
				}
				// Transform to v5 structure: { list = [...] } or { all = true }
				if len(includes) == 1 && includes[0] == "*" {
					// Special case: ["*"] becomes { all = true }
					qs["include"] = map[string]interface{}{"all": true}
				} else if len(includes) > 0 {
					// Regular case: ["param1", "param2"] becomes { list = ["param1", "param2"] }
					qs["include"] = map[string]interface{}{"list": includes}
				}
			}

			// Handle exclude array - transform to v5 structure
			excludeCount := gjson.Get(json, qsPath+`\.exclude\.#`)
			if excludeCount.Exists() && excludeCount.Int() > 0 {
				var excludes []string
				for j := int64(0); j < excludeCount.Int(); j++ {
					if val := gjson.Get(json, fmt.Sprintf(`%s\.exclude\.%d`, qsPath, j)); val.Exists() {
						excludes = append(excludes, val.String())
					}
				}
				// Transform to v5 structure: { list = [...] } or { all = true }
				if len(excludes) == 1 && excludes[0] == "*" {
					// Special case: ["*"] becomes { all = true }
					qs["exclude"] = map[string]interface{}{"all": true}
				} else if len(excludes) > 0 {
					// Regular case: ["param1", "param2"] becomes { list = ["param1", "param2"] }
					qs["exclude"] = map[string]interface{}{"list": excludes}
				}
			}

			if len(qs) > 0 {
				customKey["query_string"] = qs
			}
		}

		// Handle other custom_key fields (header, cookie, user, host)
		// Similar pattern for each...

		if len(customKey) > 0 {
			cacheKey["custom_key"] = customKey
		}
	}

	return cacheKey
}

// removeDisableRailgunFromPath removes disable_railgun from a specific resource instance path
func removeDisableRailgunFromPath(json string, instancePath string) string {
	result := json

	// Get the attributes at this path
	attrPath := instancePath + ".attributes"

	// Check if rules exist
	rules := gjson.Get(json, attrPath+".rules")
	if !rules.Exists() || !rules.IsArray() {
		return result
	}

	fmt.Printf("\nRULES:::::::\n%+v\n\n", rules)

	// Process each rule to remove disable_railgun from action_parameters
	rulesArray := rules.Array()
	for i := range rulesArray {
		actionParamsPath := fmt.Sprintf("%s.rules.%d.action_parameters.disable_railgun", attrPath, i)
		if gjson.Get(json, actionParamsPath).Exists() {
			fmt.Printf("Removing disable_railgun from %s\n", actionParamsPath)
			result, _ = sjson.Delete(result, actionParamsPath)
		}
	}

	return result
}

// Helper function to transform edge_ttl and browser_ttl structures
func transformEdgeTTL(json string, attrPath string, ruleIdx int64, blockName string) map[string]interface{} {
	ttl := make(map[string]interface{})
	ttlPath := fmt.Sprintf(`%s.rules\.%d\.action_parameters\.0\.%s\.0`, attrPath, ruleIdx, blockName)

	if val := gjson.Get(json, ttlPath+`\.mode`); val.Exists() {
		ttl["mode"] = val.String()
	}
	if val := gjson.Get(json, ttlPath+`\.default`); val.Exists() {
		ttl["default"] = val.Float()
	}

	// Handle status_code_ttl array
	sctCount := gjson.Get(json, ttlPath+`\.status_code_ttl\.#`)
	if sctCount.Exists() && sctCount.Int() > 0 {
		var sctList []map[string]interface{}
		for j := int64(0); j < sctCount.Int(); j++ {
			sct := make(map[string]interface{})
			sctPath := fmt.Sprintf(`%s\.status_code_ttl\.%d`, ttlPath, j)

			if val := gjson.Get(json, sctPath+`\.status_code`); val.Exists() {
				sct["status_code"] = val.Float()
			}
			if val := gjson.Get(json, sctPath+`\.value`); val.Exists() {
				sct["value"] = val.Float()
			}

			// Handle status_code_range
			scrCount := gjson.Get(json, sctPath+`\.status_code_range\.#`)
			if scrCount.Exists() && scrCount.Int() > 0 {
				scr := make(map[string]interface{})
				scrPath := sctPath + `\.status_code_range\.0`
				if val := gjson.Get(json, scrPath+`\.from`); val.Exists() {
					scr["from"] = val.Float()
				}
				if val := gjson.Get(json, scrPath+`\.to`); val.Exists() {
					scr["to"] = val.Float()
				}
				if len(scr) > 0 {
					sct["status_code_range"] = scr
				}
			}

			if len(sct) > 0 {
				sctList = append(sctList, sct)
			}
		}
		if len(sctList) > 0 {
			ttl["status_code_ttl"] = sctList
		}
	}

	return ttl
}

// Helper function to transform from_value structure
func transformFromValue(json string, attrPath string, ruleIdx int64) map[string]interface{} {
	fromValue := make(map[string]interface{})
	fvPath := fmt.Sprintf(`%s.rules\.%d\.action_parameters\.0\.from_value\.0`, attrPath, ruleIdx)

	if val := gjson.Get(json, fvPath+`\.preserve_query_string`); val.Exists() {
		fromValue["preserve_query_string"] = val.Bool()
	}
	if val := gjson.Get(json, fvPath+`\.status_code`); val.Exists() {
		fromValue["status_code"] = val.Float()
	}

	// Handle target_url
	tuCount := gjson.Get(json, fvPath+`\.target_url\.#`)
	if tuCount.Exists() && tuCount.Int() > 0 {
		targetUrl := make(map[string]interface{})
		tuPath := fvPath + `\.target_url\.0`
		if val := gjson.Get(json, tuPath+`\.value`); val.Exists() {
			targetUrl["value"] = val.String()
		}
		if val := gjson.Get(json, tuPath+`\.expression`); val.Exists() {
			targetUrl["expression"] = val.String()
		}
		if len(targetUrl) > 0 {
			fromValue["target_url"] = targetUrl
		}
	}

	return fromValue
}

// Helper function to transform overrides structure
func transformOverrides(json string, attrPath string, ruleIdx int64) map[string]interface{} {
	overrides := make(map[string]interface{})
	ovPath := fmt.Sprintf(`%s.rules\.%d\.action_parameters\.0\.overrides\.0`, attrPath, ruleIdx)

	// Simple fields
	if val := gjson.Get(json, ovPath+`\.enabled`); val.Exists() {
		overrides["enabled"] = val.Bool()
	}
	if val := gjson.Get(json, ovPath+`\.action`); val.Exists() {
		overrides["action"] = val.String()
	}
	if val := gjson.Get(json, ovPath+`\.sensitivity_level`); val.Exists() {
		overrides["sensitivity_level"] = val.String()
	}

	// Handle categories array
	catCount := gjson.Get(json, ovPath+`\.categories\.#`)
	if catCount.Exists() && catCount.Int() > 0 {
		var categories []map[string]interface{}
		for j := int64(0); j < catCount.Int(); j++ {
			cat := make(map[string]interface{})
			catPath := fmt.Sprintf(`%s\.categories\.%d`, ovPath, j)

			if val := gjson.Get(json, catPath+`\.category`); val.Exists() {
				cat["category"] = val.String()
			}
			if val := gjson.Get(json, catPath+`\.action`); val.Exists() {
				cat["action"] = val.String()
			}
			if val := gjson.Get(json, catPath+`\.enabled`); val.Exists() {
				cat["enabled"] = val.Bool()
			}

			if len(cat) > 0 {
				categories = append(categories, cat)
			}
		}
		if len(categories) > 0 {
			overrides["categories"] = categories
		}
	}

	// Handle rules array
	rulesCount := gjson.Get(json, ovPath+`\.rules\.#`)
	if rulesCount.Exists() && rulesCount.Int() > 0 {
		var rulesList []map[string]interface{}
		for j := int64(0); j < rulesCount.Int(); j++ {
			ruleItem := make(map[string]interface{})
			rulePath := fmt.Sprintf(`%s\.rules\.%d`, ovPath, j)

			if val := gjson.Get(json, rulePath+`\.id`); val.Exists() {
				ruleItem["id"] = val.String()
			}
			if val := gjson.Get(json, rulePath+`\.action`); val.Exists() {
				ruleItem["action"] = val.String()
			}
			if val := gjson.Get(json, rulePath+`\.enabled`); val.Exists() {
				ruleItem["enabled"] = val.Bool()
			}
			if val := gjson.Get(json, rulePath+`\.score_threshold`); val.Exists() {
				ruleItem["score_threshold"] = val.Float()
			}
			if val := gjson.Get(json, rulePath+`\.sensitivity_level`); val.Exists() {
				ruleItem["sensitivity_level"] = val.String()
			}

			if len(ruleItem) > 0 {
				rulesList = append(rulesList, ruleItem)
			}
		}
		if len(rulesList) > 0 {
			overrides["rules"] = rulesList
		}
	}

	return overrides
}

// Helper function to transform URI structure
func transformURI(json string, attrPath string, ruleIdx int64) map[string]interface{} {
	uri := make(map[string]interface{})
	uriPath := fmt.Sprintf(`%s.rules\.%d\.action_parameters\.0\.uri\.0`, attrPath, ruleIdx)

	if val := gjson.Get(json, uriPath+`\.origin`); val.Exists() {
		uri["origin"] = val.Bool()
	}

	// Handle path
	pathCount := gjson.Get(json, uriPath+`\.path\.#`)
	if pathCount.Exists() && pathCount.Int() > 0 {
		path := make(map[string]interface{})
		pPath := uriPath + `\.path\.0`
		if val := gjson.Get(json, pPath+`\.value`); val.Exists() {
			path["value"] = val.String()
		}
		if val := gjson.Get(json, pPath+`\.expression`); val.Exists() {
			path["expression"] = val.String()
		}
		if len(path) > 0 {
			uri["path"] = path
		}
	}

	// Handle query
	queryCount := gjson.Get(json, uriPath+`\.query\.#`)
	if queryCount.Exists() && queryCount.Int() > 0 {
		query := make(map[string]interface{})
		qPath := uriPath + `\.query\.0`
		if val := gjson.Get(json, qPath+`\.value`); val.Exists() {
			query["value"] = val.String()
		}
		if val := gjson.Get(json, qPath+`\.expression`); val.Exists() {
			query["expression"] = val.String()
		}
		if len(query) > 0 {
			uri["query"] = query
		}
	}

	return uri
}

// convertArraysToObjects converts single-element arrays to objects for SingleNestedAttribute fields
func convertArraysToObjects(data map[string]interface{}) map[string]interface{} {
	// Remove deprecated disable_railgun attribute if present in action_parameters
	if data["action"] != nil && data["action_parameters"] != nil {
		if ap, ok := data["action_parameters"].(map[string]interface{}); ok {
			delete(ap, "disable_railgun")
		}
	}

	// Fields that should be objects (SingleNestedAttribute), not arrays
	singleNestedFields := map[string]bool{
		"action_parameters":        true,
		"response":                 true,
		"matched_data":             true,
		"overrides":                true,
		"from_list":                true,
		"from_value":               true,
		"target_url":               true,
		"uri":                      true,
		"path":                     true,
		"query":                    true,
		"origin":                   true,
		"sni":                      true,
		"autominify":               true,
		"browser_ttl":              true,
		"edge_ttl":                 true,
		"cache_key":                true,
		"custom_key":               true,
		"cookie":                   true,
		"header":                   true,
		"host":                     true,
		"query_string":             true,
		"user":                     true,
		"serve_stale":              true,
		"cache_reserve":            true,
		"exposed_credential_check": true,
		"logging":                  true,
		"ratelimit":                true,
		"status_code_range":        true,
	}

	// Process each field in the data
	for key, value := range data {
		// Special handling for headers - convert from list to map
		if key == "headers" {
			if arr, isArray := value.([]interface{}); isArray {
				headersMap := make(map[string]interface{})
				for _, header := range arr {
					if headerObj, isMap := header.(map[string]interface{}); isMap {
						if name, hasName := headerObj["name"].(string); hasName {
							// Remove name from the header object as it becomes the key
							delete(headerObj, "name")
							headersMap[name] = headerObj
						}
					}
				}
				if len(headersMap) > 0 {
					data[key] = headersMap
				} else {
					delete(data, key)
				}
			}
			continue
		}

		// Special handling for query_string fields (include/exclude)
		// These should be wrapped in {list: [...]} or {all: true}
		if key == "query_string" {
			if obj, isMap := value.(map[string]interface{}); isMap {
				// Process include field
				if includeVal, hasInclude := obj["include"]; hasInclude {
					if arr, isArray := includeVal.([]interface{}); isArray {
						// Transform array to v5 structure
						if len(arr) == 1 && arr[0] == "*" {
							obj["include"] = map[string]interface{}{"all": true}
						} else if len(arr) > 0 {
							obj["include"] = map[string]interface{}{"list": arr}
						}
					}
				}
				// Process exclude field
				if excludeVal, hasExclude := obj["exclude"]; hasExclude {
					if arr, isArray := excludeVal.([]interface{}); isArray {
						// Transform array to v5 structure
						if len(arr) == 1 && arr[0] == "*" {
							obj["exclude"] = map[string]interface{}{"all": true}
						} else if len(arr) > 0 {
							obj["exclude"] = map[string]interface{}{"list": arr}
						}
					}
				}
				// Recursively process the rest of query_string
				data[key] = convertArraysToObjects(obj)
			} else if arr, isArray := value.([]interface{}); isArray && len(arr) > 0 {
				// If query_string itself is an array, convert to object first
				if obj, isMap := arr[0].(map[string]interface{}); isMap {
					// Process include/exclude fields within
					if includeVal, hasInclude := obj["include"]; hasInclude {
						if includeArr, isArray := includeVal.([]interface{}); isArray {
							if len(includeArr) == 1 && includeArr[0] == "*" {
								obj["include"] = map[string]interface{}{"all": true}
							} else if len(includeArr) > 0 {
								obj["include"] = map[string]interface{}{"list": includeArr}
							}
						}
					}
					if excludeVal, hasExclude := obj["exclude"]; hasExclude {
						if excludeArr, isArray := excludeVal.([]interface{}); isArray {
							if len(excludeArr) == 1 && excludeArr[0] == "*" {
								obj["exclude"] = map[string]interface{}{"all": true}
							} else if len(excludeArr) > 0 {
								obj["exclude"] = map[string]interface{}{"list": excludeArr}
							}
						}
					}
					data[key] = convertArraysToObjects(obj)
				}
			}
			continue
		}

		// Check if this field should be converted from array to object
		if singleNestedFields[key] {
			if arr, isArray := value.([]interface{}); isArray {
				// Convert array to object
				if len(arr) > 0 {
					if obj, isMap := arr[0].(map[string]interface{}); isMap {
						// Recursively process nested objects
						data[key] = convertArraysToObjects(obj)
					} else {
						data[key] = arr[0]
					}
				} else {
					// Empty array, remove it
					delete(data, key)
				}
			} else if obj, isMap := value.(map[string]interface{}); isMap {
				// Recursively process nested objects
				data[key] = convertArraysToObjects(obj)
			}
		} else if obj, isMap := value.(map[string]interface{}); isMap {
			// Even if not in the list, recursively process nested objects
			data[key] = convertArraysToObjects(obj)
		} else if arr, isArray := value.([]interface{}); isArray {
			// Process arrays that aren't in singleNestedFields
			// This handles arrays of objects like status_code_ttl
			var processedArr []interface{}
			for _, item := range arr {
				if itemMap, isMap := item.(map[string]interface{}); isMap {
					// Recursively process objects within arrays
					processedArr = append(processedArr, convertArraysToObjects(itemMap))
				} else {
					processedArr = append(processedArr, item)
				}
			}
			data[key] = processedArr
		}
	}

	return data
}

// cleanupRulesetIndexedKeys removes all indexed format keys for ruleset rules
func cleanupRulesetIndexedKeys(json string, attrPath string, rulesCount int64) string {
	result := json

	// Get the current attributes object
	attrs := gjson.Get(json, attrPath)
	if !attrs.Exists() || !attrs.IsObject() {
		return result
	}

	// Create a new attributes map without indexed keys
	attrsMap := make(map[string]interface{})
	attrs.ForEach(func(key, value gjson.Result) bool {
		keyStr := key.String()

		// Skip any key that starts with "rules." (indexed keys)
		if len(keyStr) >= 6 && keyStr[:6] == "rules." {
			return true // Skip all rules.* keys
		}

		// Keep all other attributes except the transformed rules
		if keyStr != "rules" {
			attrsMap[keyStr] = value.Value()
		}
		return true
	})

	// Add the transformed rules array back (it was set earlier)
	rules := gjson.Get(result, attrPath+".rules")
	if rules.Exists() {
		attrsMap["rules"] = rules.Value()
	}

	// Set the cleaned attributes back
	result, _ = sjson.Set(result, attrPath, attrsMap)
	return result
}
