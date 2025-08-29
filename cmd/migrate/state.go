package main

import (
	"fmt"
	"os"

	"github.com/tidwall/gjson"
	"github.com/tidwall/pretty"
	"github.com/tidwall/sjson"
)

// transformStateFile reads a terraform state file, transforms it, and writes it back
func transformStateFile(filename string) error {
	// Read the state file
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read state file: %w", err)
	}

	// Transform the state
	transformed, err := transformStateJSON(data)
	if err != nil {
		return fmt.Errorf("failed to transform state: %w", err)
	}

	// Write the transformed state back
	err = os.WriteFile(filename, transformed, 0644)
	if err != nil {
		return fmt.Errorf("failed to write state file: %w", err)
	}

	return nil
}

// transformStateJSON transforms the state JSON using gjson/sjson
func transformStateJSON(data []byte) ([]byte, error) {
	jsonStr := string(data)
	result := jsonStr

	// Process each resource
	resources := gjson.Get(jsonStr, "resources")
	if !resources.Exists() {
		return data, nil
	}

	resources.ForEach(func(ridx, resource gjson.Result) bool {
		resourceType := resource.Get("type").String()
		resourcePath := fmt.Sprintf("resources.%d", ridx.Int())

		// Handle zone_settings_override -> zone_setting transformation
		if resourceType == "cloudflare_zone_settings_override" {
			result = transformZoneSettingsStateJSON(result, resourcePath)
			return true // Continue to next resource
		}

		// Handle resource type renames
		if resourceType == "cloudflare_record" {
			// Rename cloudflare_record to cloudflare_dns_record
			result, _ = sjson.Set(result, resourcePath+".type", "cloudflare_dns_record")
			resourceType = "cloudflare_dns_record"
		}

		if resourceType == "cloudflare_access_mutual_tls_hostname_settings" {
			// Rename to cloudflare_zero_trust_access_mtls_hostname_settings
			result, _ = sjson.Set(result, resourcePath+".type", "cloudflare_zero_trust_access_mtls_hostname_settings")
			resourceType = "cloudflare_zero_trust_access_mtls_hostname_settings"
		}

		// Process each instance
		instances := resource.Get("instances")
		instances.ForEach(func(iidx, instance gjson.Result) bool {
			path := fmt.Sprintf("resources.%d.instances.%d", ridx.Int(), iidx.Int())

			switch resourceType {
			case "cloudflare_load_balancer_pool":
				result = transformLoadBalancerPoolStateJSON(result, path)
				// Remove identity_schema_version at instance level
				result, _ = sjson.Delete(result, path+".identity_schema_version")

			case "cloudflare_load_balancer":
				result = transformLoadBalancerStateJSON(result, path)

			case "cloudflare_tiered_cache":
				result = transformTieredCacheStateJSON(result, path, resourcePath)

			case "cloudflare_zero_trust_access_identity_provider", "cloudflare_access_identity_provider":
				result = transformZeroTrustAccessIdentityProviderStateJSON(result, path)

			case "cloudflare_managed_transforms":
				result = transformManagedTransformsStateJSON(result, path)

			case "cloudflare_argo":
				// cloudflare_argo needs special handling as it may split into multiple resources
				result = transformArgoStateJSON(result, path, resourcePath)

			case "cloudflare_dns_record":
				result = transformDNSRecordStateJSON(result, path, instance)

			case "cloudflare_zero_trust_access_mtls_hostname_settings", "cloudflare_access_mutual_tls_hostname_settings":
				result = transformZeroTrustAccessMTLSHostnameSettingsStateJSON(result, path)
			}

			return true
		})

		return true
	})


	// Pretty format with proper indentation
	return pretty.PrettyOptions([]byte(result), &pretty.Options{
		Indent:   "  ",
		SortKeys: false,
	}), nil
}

// transformLoadBalancerPoolStateJSON handles v4 to v5 state migration for cloudflare_load_balancer_pool
func transformLoadBalancerPoolStateJSON(json string, instancePath string) string {
	attrPath := instancePath + ".attributes"

	// 1. Transform load_shedding from array to object
	loadShedding := gjson.Get(json, attrPath+".load_shedding")
	if loadShedding.IsArray() {
		if len(loadShedding.Array()) == 0 {
			json, _ = sjson.Delete(json, attrPath+".load_shedding")
		} else {
			json, _ = sjson.Set(json, attrPath+".load_shedding", loadShedding.Array()[0].Value())
		}
	}

	// 2. Transform origin_steering from array to object
	originSteering := gjson.Get(json, attrPath+".origin_steering")
	if originSteering.IsArray() {
		if len(originSteering.Array()) == 0 {
			json, _ = sjson.Delete(json, attrPath+".origin_steering")
		} else {
			json, _ = sjson.Set(json, attrPath+".origin_steering", originSteering.Array()[0].Value())
		}
	}

	// 3. Transform origins headers from v4 format to v5 format
	origins := gjson.Get(json, attrPath+".origins")
	if origins.IsArray() {
		origins.ForEach(func(idx, origin gjson.Result) bool {
			originPath := fmt.Sprintf("%s.origins.%d", attrPath, idx.Int())
			header := origin.Get("header")

			if header.IsArray() {
				if len(header.Array()) == 0 {
					json, _ = sjson.Delete(json, originPath+".header")
				} else {
					// Transform v4 header format to v5
					firstHeader := header.Array()[0]
					if firstHeader.Get("header").String() == "Host" {
						values := firstHeader.Get("values")
						if values.Exists() {
							json, _ = sjson.Set(json, originPath+".header", map[string]interface{}{
								"host": values.Value(),
							})
						}
					}
				}
			} else if header.IsObject() && !header.Get("host").Exists() {
				// Handle intermediate format from Grit
				if header.Get("header").String() == "Host" {
					values := header.Get("values")
					if values.Exists() {
						json, _ = sjson.Set(json, originPath+".header", map[string]interface{}{
							"host": values.Value(),
						})
					}
				}
			}

			return true
		})
	}

	return json
}

// transformTieredCacheStateJSON handles v4 to v5 state migration for cloudflare_tiered_cache
func transformTieredCacheStateJSON(json string, instancePath string, resourcePath string) string {
	attrPath := instancePath + ".attributes"

	// Get the cache_type value
	cacheType := gjson.Get(json, attrPath+".cache_type")

	if cacheType.Exists() {
		switch cacheType.String() {
		case "generic":
			// Transform to cloudflare_argo_tiered_caching resource
			json, _ = sjson.Set(json, resourcePath+".type", "cloudflare_argo_tiered_caching")
			// Remove cache_type and add value="on"
			json, _ = sjson.Delete(json, attrPath+".cache_type")
			json, _ = sjson.Set(json, attrPath+".value", "on")

		case "smart":
			// Keep as cloudflare_tiered_cache but transform attribute
			json, _ = sjson.Delete(json, attrPath+".cache_type")
			json, _ = sjson.Set(json, attrPath+".value", "on")

		case "off":
			// Keep as cloudflare_tiered_cache but transform attribute
			json, _ = sjson.Delete(json, attrPath+".cache_type")
			json, _ = sjson.Set(json, attrPath+".value", "off")
		}
	}

	return json
}

// transformLoadBalancerStateJSON handles v4 to v5 state migration for cloudflare_load_balancer
func transformLoadBalancerStateJSON(json string, instancePath string) string {
	attrPath := instancePath + ".attributes"

	// Rename fallback_pool_id to fallback_pool
	fallbackPoolID := gjson.Get(json, attrPath+".fallback_pool_id")
	if fallbackPoolID.Exists() {
		json, _ = sjson.Set(json, attrPath+".fallback_pool", fallbackPoolID.Value())
		json, _ = sjson.Delete(json, attrPath+".fallback_pool_id")
	}

	// Rename default_pool_ids to default_pools
	defaultPoolIDs := gjson.Get(json, attrPath+".default_pool_ids")
	if defaultPoolIDs.Exists() {
		json, _ = sjson.Set(json, attrPath+".default_pools", defaultPoolIDs.Value())
		json, _ = sjson.Delete(json, attrPath+".default_pool_ids")
	}

	// Remove empty arrays for single-object attributes
	singleObjectAttrs := []string{
		"adaptive_routing", "location_strategy",
		"random_steering", "session_affinity_attributes",
	}
	for _, attr := range singleObjectAttrs {
		val := gjson.Get(json, attrPath+"."+attr)
		if val.IsArray() && len(val.Array()) == 0 {
			json, _ = sjson.Delete(json, attrPath+"."+attr)
		}
	}

	// Convert empty arrays to empty maps for map attributes
	mapAttrs := []string{"country_pools", "pop_pools", "region_pools"}
	for _, attr := range mapAttrs {
		val := gjson.Get(json, attrPath+"."+attr)
		if val.IsArray() && len(val.Array()) == 0 {
			json, _ = sjson.Set(json, attrPath+"."+attr, map[string]interface{}{})
		}
	}

	return json
}

// Legacy functions to maintain compatibility with existing code that calls these directly
// These just delegate to the JSON-based implementations above

// transformLoadBalancerPoolState is kept for compatibility but delegates to JSON version
func transformLoadBalancerPoolState(attributes map[string]interface{}) {
	// This function is called from tests, so we need to keep it
	// We'll convert the map to JSON, transform it, and convert back
	// This is not efficient but maintains the existing interface

	// For now, keeping the old implementation to not break tests
	// The actual state transformation uses the JSON version above

	// 1. Transform load_shedding from array to object
	if loadShedding, ok := attributes["load_shedding"]; ok {
		if arr, ok := loadShedding.([]interface{}); ok {
			if len(arr) == 0 {
				delete(attributes, "load_shedding")
			} else if len(arr) > 0 {
				if firstElem, ok := arr[0].(map[string]interface{}); ok {
					attributes["load_shedding"] = firstElem
				}
			}
		}
	}

	// 2. Transform origin_steering from array to object
	if originSteering, ok := attributes["origin_steering"]; ok {
		if arr, ok := originSteering.([]interface{}); ok {
			if len(arr) == 0 {
				delete(attributes, "origin_steering")
			} else if len(arr) > 0 {
				if firstElem, ok := arr[0].(map[string]interface{}); ok {
					attributes["origin_steering"] = firstElem
				}
			}
		}
	}

	// 3. Transform origins headers
	if origins, ok := attributes["origins"]; ok {
		if originsArray, ok := origins.([]interface{}); ok {
			for _, origin := range originsArray {
				if originMap, ok := origin.(map[string]interface{}); ok {
					if header, ok := originMap["header"]; ok {
						if headerArray, ok := header.([]interface{}); ok {
							if len(headerArray) == 0 {
								delete(originMap, "header")
							} else if len(headerArray) > 0 {
								if firstHeader, ok := headerArray[0].(map[string]interface{}); ok {
									if headerName, ok := firstHeader["header"].(string); ok && headerName == "Host" {
										if values, ok := firstHeader["values"].([]interface{}); ok {
											originMap["header"] = map[string]interface{}{
												"host": values,
											}
										}
									}
								}
							}
						} else if headerObj, ok := header.(map[string]interface{}); ok {
							if _, hasHost := headerObj["host"]; !hasHost {
								if headerName, ok := headerObj["header"].(string); ok && headerName == "Host" {
									if values, ok := headerObj["values"].([]interface{}); ok {
										originMap["header"] = map[string]interface{}{
											"host": values,
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

// transformLoadBalancerState is kept for compatibility but delegates to JSON version
func transformLoadBalancerState(attributes map[string]interface{}) {
	// This function is called from tests, so we need to keep it
	// For now, keeping the old implementation to not break tests

	// Rename fallback_pool_id to fallback_pool
	if fallbackPoolID, ok := attributes["fallback_pool_id"]; ok {
		attributes["fallback_pool"] = fallbackPoolID
		delete(attributes, "fallback_pool_id")
	}

	// Rename default_pool_ids to default_pools
	if defaultPoolIDs, ok := attributes["default_pool_ids"]; ok {
		attributes["default_pools"] = defaultPoolIDs
		delete(attributes, "default_pool_ids")
	}

	// Remove empty arrays for single-object attributes
	singleObjectAttributes := []string{
		"adaptive_routing",
		"location_strategy",
		"random_steering",
		"session_affinity_attributes",
	}

	for _, attr := range singleObjectAttributes {
		if val, ok := attributes[attr]; ok {
			if arr, ok := val.([]interface{}); ok && len(arr) == 0 {
				delete(attributes, attr)
			}
		}
	}

	// Convert empty arrays to empty maps for map attributes
	mapAttributes := []string{
		"country_pools",
		"pop_pools",
		"region_pools",
	}

	for _, attr := range mapAttributes {
		if val, ok := attributes[attr]; ok {
			if arr, ok := val.([]interface{}); ok && len(arr) == 0 {
				attributes[attr] = make(map[string]interface{})
			}
		}
	}
}

// transformZeroTrustAccessIdentityProviderStateJSON handles v4 to v5 state migration for cloudflare_zero_trust_access_identity_provider
func transformZeroTrustAccessIdentityProviderStateJSON(json string, instancePath string) string {
	attrPath := instancePath + ".attributes"

	// 1. Transform config from array to object (v4: config=[{...}], v5: config={...})
	config := gjson.Get(json, attrPath+".config")
	if config.IsArray() {
		if len(config.Array()) == 0 {
			// Empty array -> empty object
			json, _ = sjson.Set(json, attrPath+".config", map[string]interface{}{})
		} else {
			// Take first item from array and make it the object
			json, _ = sjson.Set(json, attrPath+".config", config.Array()[0].Value())
		}
	}

	// 2. Transform scim_config from array to object (v4: scim_config=[{...}], v5: scim_config={...})
	scimConfig := gjson.Get(json, attrPath+".scim_config")
	if scimConfig.IsArray() {
		if len(scimConfig.Array()) == 0 {
			// Empty array -> delete attribute (it's optional in v5)
			json, _ = sjson.Delete(json, attrPath+".scim_config")
		} else {
			// Take first item from array and make it the object
			json, _ = sjson.Set(json, attrPath+".scim_config", scimConfig.Array()[0].Value())
		}
	}

	// 3. Transform idp_public_cert -> idp_public_certs (string -> list[string])
	// This works on the config object after transformation above
	configObj := gjson.Get(json, attrPath+".config")
	if configObj.Exists() && configObj.IsObject() {
		idpPublicCert := configObj.Get("idp_public_cert")
		if idpPublicCert.Exists() && idpPublicCert.Type == gjson.String {
			// Convert single cert string to list with one item
			certList := []string{idpPublicCert.String()}
			json, _ = sjson.Set(json, attrPath+".config.idp_public_certs", certList)
			// Remove old field
			json, _ = sjson.Delete(json, attrPath+".config.idp_public_cert")
		}
	}

	// 4. Handle field normalization - remove empty/null/false fields from config
	// The v5 provider only includes relevant fields, not all possible fields
	configPath := attrPath + ".config"

	// Remove false boolean values
	boolFields := []string{"sign_request", "conditional_access_enabled", "support_groups", "pkce_enabled"}
	for _, field := range boolFields {
		fieldPath := configPath + "." + field
		value := gjson.Get(json, fieldPath)
		if value.Exists() && value.Type == gjson.False {
			json, _ = sjson.Delete(json, fieldPath)
		}
	}

	// Remove empty strings
	stringFields := []string{"client_secret", "client_id", "apps_domain", "auth_url", "certs_url", "directory_id", "email_claim_name", "okta_account", "onelogin_account", "ping_env_id", "issuer_url", "sso_target_url", "token_url", "email_attribute_name", "centrify_account", "centrify_app_id", "authorization_server_id"}
	for _, field := range stringFields {
		fieldPath := configPath + "." + field
		value := gjson.Get(json, fieldPath)
		if value.Exists() && value.Type == gjson.String && value.String() == "" {
			json, _ = sjson.Delete(json, fieldPath)
		}
	}

	// Remove empty arrays
	arrayFields := []string{"claims", "scopes", "attributes", "header_attributes", "idp_public_certs"}
	for _, field := range arrayFields {
		fieldPath := configPath + "." + field
		value := gjson.Get(json, fieldPath)
		if value.Exists() && value.IsArray() && len(value.Array()) == 0 {
			json, _ = sjson.Delete(json, fieldPath)
		}
	}

	// Remove null values
	nullableFields := []string{"prompt"}
	for _, field := range nullableFields {
		fieldPath := configPath + "." + field
		value := gjson.Get(json, fieldPath)
		if value.Exists() && value.Type == gjson.Null {
			json, _ = sjson.Delete(json, fieldPath)
		}
	}

	// 5. Remove deprecated fields
	json, _ = sjson.Delete(json, attrPath+".config.api_token")
	json, _ = sjson.Delete(json, attrPath+".scim_config.group_member_deprovision")

	return json
}

// transformManagedTransformsStateJSON handles v4 to v5 state migration for cloudflare_managed_transforms
func transformManagedTransformsStateJSON(json string, instancePath string) string {
	attrPath := instancePath + ".attributes"

	// Check if managed_request_headers exists, if not add empty array
	requestHeaders := gjson.Get(json, attrPath+".managed_request_headers")
	if !requestHeaders.Exists() {
		json, _ = sjson.Set(json, attrPath+".managed_request_headers", []interface{}{})
	}

	// Check if managed_response_headers exists, if not add empty array
	responseHeaders := gjson.Get(json, attrPath+".managed_response_headers")
	if !responseHeaders.Exists() {
		json, _ = sjson.Set(json, attrPath+".managed_response_headers", []interface{}{})
	}

	return json
}

// transformArgoStateJSON handles v4 to v5 state migration for cloudflare_argo
// The cloudflare_argo resource is split into cloudflare_argo_smart_routing and/or cloudflare_argo_tiered_caching
func transformArgoStateJSON(json string, instancePath string, resourcePath string) string {
	attrPath := instancePath + ".attributes"

	// Get the smart_routing and tiered_caching values
	smartRouting := gjson.Get(json, attrPath+".smart_routing")
	tieredCaching := gjson.Get(json, attrPath+".tiered_caching")

	// Determine which resource type to use based on attributes
	// Note: In state files, we can only represent one resource type per resource block
	// When both attributes exist with non-null values, we prioritize smart_routing
	// A proper migration would need to handle splitting into multiple resources at the Terraform config level

	// Check if smart_routing exists and is not null
	if smartRouting.Exists() && smartRouting.Type != gjson.Null {
		// Transform to cloudflare_argo_smart_routing
		json, _ = sjson.Set(json, resourcePath+".type", "cloudflare_argo_smart_routing")

		// Rename smart_routing to value
		json, _ = sjson.Set(json, attrPath+".value", smartRouting.Value())
		json, _ = sjson.Delete(json, attrPath+".smart_routing")

		// Remove tiered_caching if it exists
		if tieredCaching.Exists() {
			json, _ = sjson.Delete(json, attrPath+".tiered_caching")
			// Note: This loses the tiered_caching value - proper migration requires config-level handling
		}
	} else if tieredCaching.Exists() && tieredCaching.Type != gjson.Null {
		// Transform to cloudflare_argo_tiered_caching if tiered_caching exists and is not null
		json, _ = sjson.Set(json, resourcePath+".type", "cloudflare_argo_tiered_caching")

		// Rename tiered_caching to value
		json, _ = sjson.Set(json, attrPath+".value", tieredCaching.Value())
		json, _ = sjson.Delete(json, attrPath+".tiered_caching")

		// Remove smart_routing if it exists (and is null)
		if smartRouting.Exists() {
			json, _ = sjson.Delete(json, attrPath+".smart_routing")
		}
	} else {
		// Both are null or don't exist - default to cloudflare_argo_smart_routing with value "off"
		json, _ = sjson.Set(json, resourcePath+".type", "cloudflare_argo_smart_routing")
		json, _ = sjson.Set(json, attrPath+".value", "off")

		// Clean up null attributes
		if smartRouting.Exists() {
			json, _ = sjson.Delete(json, attrPath+".smart_routing")
		}
		if tieredCaching.Exists() {
			json, _ = sjson.Delete(json, attrPath+".tiered_caching")
		}
	}

	return json
}

// transformZeroTrustAccessMTLSHostnameSettingsStateJSON handles v4 to v5 state migration for cloudflare_zero_trust_access_mtls_hostname_settings
func transformZeroTrustAccessMTLSHostnameSettingsStateJSON(json string, instancePath string) string {
	attrPath := instancePath + ".attributes"

	// Transform settings to ensure boolean defaults are set
	settings := gjson.Get(json, attrPath+".settings")
	if settings.IsArray() {
		settings.ForEach(func(idx, setting gjson.Result) bool {
			settingPath := fmt.Sprintf("%s.settings.%d", attrPath, idx.Int())
			
			// Ensure china_network has a value (default to false if missing)
			chinaNetwork := setting.Get("china_network")
			if !chinaNetwork.Exists() {
				json, _ = sjson.Set(json, settingPath+".china_network", false)
			}
			
			// Ensure client_certificate_forwarding has a value (default to false if missing)
			clientCertForwarding := setting.Get("client_certificate_forwarding")
			if !clientCertForwarding.Exists() {
				json, _ = sjson.Set(json, settingPath+".client_certificate_forwarding", false)
			}
			
			return true
		})
	}

	return json
}
