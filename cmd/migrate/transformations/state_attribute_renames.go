package transformations

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// StateTransformationConfig represents the YAML configuration for state transformations
type StateTransformationConfig struct {
	Version                   string                            `yaml:"version"`
	Description               string                            `yaml:"description"`
	SchemaVersionReset        SchemaVersionReset                `yaml:"schema_version_reset"`
	StateAttributeRenames     map[string]map[string]interface{} `yaml:"state_attribute_renames"`
	StateAttributeRemovals    map[string][]string               `yaml:"state_attribute_removals"`
	StateSpecialTransformations map[string]SpecialTransformation `yaml:"state_special_transformations"`
	Notes                     []string                          `yaml:"notes"`
}

// SchemaVersionReset defines schema version reset rules
type SchemaVersionReset struct {
	AllCloudflareResources bool `yaml:"all_cloudflare_resources"`
}

// SpecialTransformation defines special transformation rules
type SpecialTransformation struct {
	EmptyToNull         []string                       `yaml:"empty_to_null"`
	FieldTransformations map[string]FieldTransformation `yaml:"field_transformations"`
}

// FieldTransformation defines a field transformation rule
type FieldTransformation struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}

// TerraformState represents the structure of a Terraform state file
type TerraformState struct {
	Version          int                      `json:"version"`
	TerraformVersion string                   `json:"terraform_version"`
	Serial           int                      `json:"serial"`
	Lineage          string                   `json:"lineage"`
	Outputs          map[string]interface{}   `json:"outputs,omitempty"`
	Resources        []TerraformStateResource `json:"resources"`
}

// TerraformStateResource represents a resource in the state file
type TerraformStateResource struct {
	Module    string                   `json:"module,omitempty"`
	Mode      string                   `json:"mode"`
	Type      string                   `json:"type"`
	Name      string                   `json:"name"`
	Provider  string                   `json:"provider"`
	Instances []TerraformStateInstance `json:"instances"`
}

// TerraformStateInstance represents an instance of a resource
type TerraformStateInstance struct {
	SchemaVersion       int                    `json:"schema_version"`
	Attributes          map[string]interface{} `json:"attributes"`
	SensitiveAttributes []interface{}         `json:"sensitive_attributes,omitempty"`
	Private             string                 `json:"private,omitempty"`
	Dependencies        []string               `json:"dependencies,omitempty"`
	CreateBeforeDestroy bool                   `json:"create_before_destroy,omitempty"`
	IndexKey            interface{}           `json:"index_key,omitempty"` // For for_each and count resources
}

// LoadStateConfig loads the state transformation configuration from a YAML file
func LoadStateConfig(filename string) (*StateTransformationConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config StateTransformationConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &config, nil
}

// TransformStateFile transforms a Terraform state file according to the configuration
func TransformStateFile(config *StateTransformationConfig, inputPath, outputPath string) error {
	// Read the state file
	stateData, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read state file: %w", err)
	}

	// Parse the state file
	var state TerraformState
	if err := json.Unmarshal(stateData, &state); err != nil {
		return fmt.Errorf("failed to parse state file: %w", err)
	}

	// Apply transformations
	for i := range state.Resources {
		resource := &state.Resources[i]
		
		// Only process managed resources
		if resource.Mode != "managed" {
			continue
		}

		// Check if this is a Cloudflare resource
		if !isCloudflareResource(resource.Type) {
			continue
		}

		// Apply schema version reset if configured
		if config.SchemaVersionReset.AllCloudflareResources {
			for j := range resource.Instances {
				resource.Instances[j].SchemaVersion = 0
			}
		}

		// Apply attribute renames
		if renames, exists := config.StateAttributeRenames[resource.Type]; exists {
			for j := range resource.Instances {
				applyStateAttributeRenames(&resource.Instances[j], renames)
			}
		}

		// Apply attribute removals
		if removals, exists := config.StateAttributeRemovals[resource.Type]; exists {
			for j := range resource.Instances {
				applyStateAttributeRemovals(&resource.Instances[j], removals)
			}
		}

		// Apply special transformations
		if special, exists := config.StateSpecialTransformations[resource.Type]; exists {
			for j := range resource.Instances {
				applySpecialTransformations(&resource.Instances[j], special)
			}
		}

		// Always apply page rule specific transformations regardless of config
		if resource.Type == "cloudflare_page_rule" {
			for j := range resource.Instances {
				applyPageRuleTransformations(&resource.Instances[j])
			}
		}
	}

	// Write the transformed state file
	output, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal state: %w", err)
	}

	if outputPath == "" || outputPath == inputPath {
		outputPath = inputPath
	}

	if err := os.WriteFile(outputPath, output, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}

// isCloudflareResource checks if a resource type is a Cloudflare resource
func isCloudflareResource(resourceType string) bool {
	return len(resourceType) > 10 && resourceType[:10] == "cloudflare" && resourceType[10] == '_'
}

// applyStateAttributeRenames applies attribute renames to a state instance
func applyStateAttributeRenames(instance *TerraformStateInstance, renames map[string]interface{}) {
	if instance.Attributes == nil {
		return
	}

	// Handle regular renames
	for oldName, newValue := range renames {
		// Skip special cases
		if oldName == "id_duplicate_as" {
			continue
		}

		if val, exists := instance.Attributes[oldName]; exists {
			// Check if it's a simple rename
			if newName, ok := newValue.(string); ok {
				instance.Attributes[newName] = val
				delete(instance.Attributes, oldName)
			}
		}
	}

	// Handle id duplication special cases
	if dupTarget, exists := renames["id_duplicate_as"]; exists {
		if targetName, ok := dupTarget.(string); ok {
			if id, exists := instance.Attributes["id"]; exists {
				instance.Attributes[targetName] = id
			}
		}
	}
}

// applyStateAttributeRemovals removes specified attributes from a state instance
func applyStateAttributeRemovals(instance *TerraformStateInstance, removals []string) {
	if instance.Attributes == nil {
		return
	}

	for _, attrName := range removals {
		delete(instance.Attributes, attrName)
	}
}

// applyPageRuleTransformations applies page rule specific transformations
func applyPageRuleTransformations(instance *TerraformStateInstance) {
	if instance.Attributes == nil {
		return
	}

	// Handle top-level cache_key_fields if it exists
	if cacheKeyFields, exists := instance.Attributes["cache_key_fields"]; exists {
		instance.Attributes["cache_key_fields"] = transformCacheKeyFields(cacheKeyFields)
	}

	// Handle actions nested attributes
	if actionsVal, exists := instance.Attributes["actions"]; exists {
		// In v4, actions is an array with a single map. In v5, it's just a map.
		// First unwrap if it's an array
		unwrappedActions := unwrapSingleElementArray(actionsVal)

		if actions, ok := unwrappedActions.(map[string]interface{}); ok {
			// Store the unwrapped version
			instance.Attributes["actions"] = actions
			// Remove unsupported attributes from actions in v5
			delete(actions, "disable_railgun")
			delete(actions, "minify")
			delete(actions, "server_side_exclude")

			// Handle forwarding_url transformation inside actions
			if fwdUrl, exists := actions["forwarding_url"]; exists {
				actions["forwarding_url"] = transformForwardingURL(fwdUrl)
			}

			// Handle cache_key_fields transformation inside actions
			if cacheKeyFields, exists := actions["cache_key_fields"]; exists {
				actions["cache_key_fields"] = transformCacheKeyFields(cacheKeyFields)
			}

			// Handle cache_ttl_by_status transformation inside actions
			// In v4 it was an array of {codes, ttl} objects, in v5 it's a map[status_code]ttl
			if cacheTTL, exists := actions["cache_ttl_by_status"]; exists {
				actions["cache_ttl_by_status"] = transformCacheTTLByStatus(cacheTTL)
			}

			// Clean up empty values in actions to prevent parsing errors
			for actionKey, actionVal := range actions {
				// Handle browser_cache_ttl and other numeric fields
				if actionKey == "browser_cache_ttl" || actionKey == "edge_cache_ttl" {
					// Convert empty strings or invalid values to null
					if str, ok := actionVal.(string); ok {
						if str == "" || str == "0" {
							actions[actionKey] = nil
						}
					} else if num, ok := actionVal.(float64); ok && num == 0 {
						actions[actionKey] = nil
					} else if num, ok := actionVal.(int); ok && num == 0 {
						actions[actionKey] = nil
					}
				}
			}
		}
	}
}

// applySpecialTransformations applies special transformations to a state instance
func applySpecialTransformations(instance *TerraformStateInstance, special SpecialTransformation) {
	if instance.Attributes == nil {
		return
	}

	// Handle empty to null transformations
	for _, pattern := range special.EmptyToNull {
		for key, val := range instance.Attributes {
			if shouldConvertToNull(val, pattern) {
				instance.Attributes[key] = nil
			}
		}
	}

	// Handle field-specific transformations
	for fieldName := range special.FieldTransformations {
		if fieldName == "forwarding_url" {
			if val, exists := instance.Attributes["forwarding_url"]; exists {
				instance.Attributes["forwarding_url"] = transformForwardingURL(val)
			}
		} else if fieldName == "actions" {
			if val, exists := instance.Attributes["actions"]; exists {
				instance.Attributes["actions"] = unwrapSingleElementArray(val)
			}
		}
	}

	// Apply empty to null transformations for values within actions
	if actionsVal, exists := instance.Attributes["actions"]; exists {
		if actions, ok := actionsVal.(map[string]interface{}); ok {
			for actionKey, actionVal := range actions {
				for _, pattern := range special.EmptyToNull {
					if shouldConvertToNull(actionVal, pattern) {
						actions[actionKey] = nil
					}
				}
			}
		}
	}
}

// shouldConvertToNull checks if a value should be converted to null based on pattern
func shouldConvertToNull(val interface{}, pattern string) bool {
	switch pattern {
	case `""`:
		if str, ok := val.(string); ok && str == "" {
			return true
		}
	case "[]":
		if arr, ok := val.([]interface{}); ok && len(arr) == 0 {
			return true
		}
	case "false":
		if b, ok := val.(bool); ok && !b {
			return true
		}
	case "0":
		if num, ok := val.(float64); ok && num == 0 {
			return true
		}
		if num, ok := val.(int); ok && num == 0 {
			return true
		}
	}
	return false
}

// transformForwardingURL handles the forwarding_url transformation
func transformForwardingURL(val interface{}) interface{} {
	// Check if it's an empty array
	if arr, ok := val.([]interface{}); ok {
		if len(arr) == 0 {
			return nil
		}
		// If single element array, unwrap it
		if len(arr) == 1 {
			return arr[0]
		}
	}
	return val
}

// unwrapSingleElementArray unwraps single-element arrays
func unwrapSingleElementArray(val interface{}) interface{} {
	if arr, ok := val.([]interface{}); ok && len(arr) == 1 {
		return arr[0]
	}
	return val
}

// transformCacheKeyFields transforms cache_key_fields from v4 to v5 format
// v4: array with nested arrays in subfields
// v5: object with nested objects in subfields
func transformCacheKeyFields(val interface{}) interface{} {
	// First unwrap the outer array if it's a single element array
	unwrapped := unwrapSingleElementArray(val)

	// Now process the cache_key_fields object
	if obj, ok := unwrapped.(map[string]interface{}); ok {
		// Process each nested field that might be an array
		for key, fieldVal := range obj {
			switch key {
			case "cookie", "header", "host", "user":
				// These are nested objects that might be wrapped in arrays
				// The unwrapping preserves the internal array order since we're just removing the outer array wrapper
				unwrappedField := unwrapSingleElementArray(fieldVal)
				obj[key] = unwrappedField
			case "query_string":
				// Special handling for query_string to convert ignore field
				qsObj := unwrapSingleElementArray(fieldVal)
				if qs, ok := qsObj.(map[string]interface{}); ok {
					// Convert ignore field to include/exclude
					if ignoreVal, hasIgnore := qs["ignore"]; hasIgnore {
						delete(qs, "ignore")

						// ignore = true becomes exclude = ["*"]
						// ignore = false becomes include = ["*"]
						ignore := false
						switch v := ignoreVal.(type) {
						case bool:
							ignore = v
						case string:
							ignore = v == "true"
						}

						if ignore {
							// ignore = true means exclude all
							qs["exclude"] = []interface{}{"*"}
							// Remove include if we're setting exclude
							delete(qs, "include")
						} else {
							// ignore = false means include all query strings
							// According to test expectation, this should be include = ["*"]
							qs["include"] = []interface{}{"*"}
							qs["exclude"] = []interface{}{}
						}
					}
					// Note: If there's no ignore field, we preserve existing include/exclude arrays as-is
					// They're already arrays and don't need transformation
					obj[key] = qs
				} else {
					obj[key] = qsObj
				}
			}
		}
		return obj
	}

	return unwrapped
}

// transformCacheTTLByStatus transforms cache_ttl_by_status from v4 array format to v5 map format
// v4: [{"codes": "200", "ttl": 86400}, {"codes": "404", "ttl": 300}]
// v5: {"200": "86400", "404": "300"} - keys are strings, values are strings
func transformCacheTTLByStatus(val interface{}) interface{} {
	// Check if it's an array
	arr, ok := val.([]interface{})
	if !ok {
		return val
	}

	// Convert array of objects to map
	result := make(map[string]interface{})
	for _, item := range arr {
		obj, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		// Extract codes and ttl
		codes, hasCode := obj["codes"]
		ttl, hasTTL := obj["ttl"]
		if !hasCode || !hasTTL {
			continue
		}

		// Convert codes to string if needed
		codeStr := ""
		switch c := codes.(type) {
		case string:
			codeStr = c
		case float64:
			codeStr = fmt.Sprintf("%.0f", c)
		case int:
			codeStr = fmt.Sprintf("%d", c)
		default:
			continue
		}

		// Convert ttl to string (v5 stores as strings in state)
		ttlStr := ""
		switch t := ttl.(type) {
		case string:
			ttlStr = t
		case float64:
			ttlStr = fmt.Sprintf("%.0f", t)
		case int:
			ttlStr = fmt.Sprintf("%d", t)
		default:
			continue
		}

		// Add to result map
		result[codeStr] = ttlStr
	}

	// Return the map, or nil if empty
	if len(result) == 0 {
		return nil
	}
	return result
}

// StateTransformer handles state file transformations
type StateTransformer struct {
	config *StateTransformationConfig
}

// NewStateTransformer creates a new state transformer
func NewStateTransformer(configPath string) (*StateTransformer, error) {
	config, err := LoadStateConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &StateTransformer{
		config: config,
	}, nil
}

// TransformFile transforms a state file
func (st *StateTransformer) TransformFile(inputPath, outputPath string) error {
	return TransformStateFile(st.config, inputPath, outputPath)
}