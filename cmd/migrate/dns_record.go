package main

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// isDNSRecordResource checks if a block is a DNS record resource
func isDNSRecordResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		(block.Labels()[0] == "cloudflare_dns_record" || block.Labels()[0] == "cloudflare_record")
}

// ProcessDNSRecordConfig processes cloudflare_dns_record configurations to fix CAA record issues and ensure TTL is present
func ProcessDNSRecordConfig(file *hclwrite.File) error {
	body := file.Body()

	for _, block := range body.Blocks() {
		if block.Type() != "resource" {
			continue
		}

		labels := block.Labels()
		if len(labels) < 2 {
			continue
		}

		resourceType := labels[0]
		if resourceType != "cloudflare_dns_record" && resourceType != "cloudflare_record" {
			continue
		}

		// Rename cloudflare_record to cloudflare_dns_record
		if resourceType == "cloudflare_record" {
			labels[0] = "cloudflare_dns_record"
			block.SetLabels(labels)
		}

		// Ensure TTL is present for v5 (required field)
		ttlAttr := block.Body().GetAttribute("ttl")
		if ttlAttr == nil {
			// TTL is missing - add it with default value of 1 (automatic)
			// Create the TTL token with value 1
			ttlToken := &hclwrite.Token{
				Type:  hclsyntax.TokenNumberLit,
				Bytes: []byte("1"),
			}
			block.Body().SetAttributeRaw("ttl", hclwrite.Tokens{ttlToken})
		}

		// Get the record type first
		typeAttr := block.Body().GetAttribute("type")
		if typeAttr == nil {
			continue
		}

		// Extract the record type value
		typeTokens := typeAttr.Expr().BuildTokens(nil)
		var recordType string
		for _, token := range typeTokens {
			if token.Type == hclsyntax.TokenQuotedLit || token.Type == hclsyntax.TokenIdent {
				recordType = strings.Trim(string(token.Bytes), "\"")
				break
			}
		}

		// For simple record types, rename value to content
		// Simple types don't have data blocks
		simpleTypes := map[string]bool{
			"A": true, "AAAA": true, "CNAME": true, "MX": true,
			"NS": true, "PTR": true, "TXT": true, "OPENPGPKEY": true,
		}
		if simpleTypes[recordType] {
			// Rename value to content for simple record types
			valueAttr := block.Body().GetAttribute("value")
			if valueAttr != nil {
				// Get the value tokens
				valueTokens := valueAttr.Expr().BuildTokens(nil)
				// Set as content
				block.Body().SetAttributeRaw("content", valueTokens)
				// Remove value attribute
				block.Body().RemoveAttribute("value")
			}
		}

		// Remove deprecated attributes
		// allow_overwrite was removed in v5
		if allowOverwrite := block.Body().GetAttribute("allow_overwrite"); allowOverwrite != nil {
			block.Body().RemoveAttribute("allow_overwrite")
		}

		// hostname was removed in v5
		if hostname := block.Body().GetAttribute("hostname"); hostname != nil {
			block.Body().RemoveAttribute("hostname")
		}

		// Check if this record has any data blocks to process
		hasDataBlock := false
		for _, b := range block.Body().Blocks() {
			if b.Type() == "data" {
				hasDataBlock = true
				break
			}
		}

		// Skip if no data blocks to process
		if !hasDataBlock {
			continue
		}

		// For SRV, MX, and URI records, extract priority from data block if present
		if recordType == "SRV" || recordType == "MX" || recordType == "URI" {
			for _, dataBlock := range block.Body().Blocks() {
				if dataBlock.Type() != "data" {
					continue
				}

				// Check if priority exists in data block
				if priorityAttr := dataBlock.Body().GetAttribute("priority"); priorityAttr != nil {
					// Copy priority to root level if not already there
					if block.Body().GetAttribute("priority") == nil {
						// Copy the priority value to the root level (keep it in both places)
						// Get the actual tokens representing the value
						priorityTokens := priorityAttr.Expr().BuildTokens(nil)
						block.Body().SetAttributeRaw("priority", priorityTokens)
					}
					// Keep priority in the data block as well
				}
			}
		}

		var dataBlocksToRemove []*hclwrite.Block
		for _, dataBlock := range block.Body().Blocks() {
			if dataBlock.Type() != "data" {
				continue
			}

			// Convert data block to attribute format
			// Build the object expression from the block's attributes
			var objTokens hclwrite.Tokens
			objTokens = append(objTokens, &hclwrite.Token{
				Type:  hclsyntax.TokenOBrace,
				Bytes: []byte("{"),
			})
			objTokens = append(objTokens, &hclwrite.Token{
				Type:  hclsyntax.TokenNewline,
				Bytes: []byte("\n"),
			})

			// Process attributes from the data block in a consistent order
			// First collect all attributes
			attrs := dataBlock.Body().Attributes()
			processedAttrs := make(map[string]bool)

			// Define the preferred order based on record type
			var attrOrder []string
			switch recordType {
			case "CAA":
				attrOrder = []string{"flags", "tag", "value"}
			case "SRV":
				// Priority is kept in both root level and data
				attrOrder = []string{"priority", "weight", "port", "target", "service"}
			case "URI":
				attrOrder = []string{"weight", "target"}
			default:
				attrOrder = []string{}
			}

			// Process attributes in the defined order
			for _, attrName := range attrOrder {
				var attr *hclwrite.Attribute
				var finalName string
				var origName string

				// Special handling for value/content renaming
				if attrName == "value" {
					// For CAA records, rename content to value
					if contentAttr, exists := attrs["content"]; exists && recordType == "CAA" {
						attr = contentAttr
						finalName = "value"
						origName = "content"
					} else if valueAttr, exists := attrs["value"]; exists {
						// Already has value attribute
						attr = valueAttr
						finalName = "value"
						origName = "value"
					}
				} else {
					// For other attributes, use as-is if they exist
					if a, exists := attrs[attrName]; exists {
						attr = a
						finalName = attrName
						origName = attrName
					}
				}

				if attr == nil {
					continue
				}

				// Mark as processed
				processedAttrs[origName] = true

				// Add indentation
				objTokens = append(objTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenIdent,
					Bytes: []byte("    "),
				})

				// Add attribute name
				objTokens = append(objTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenIdent,
					Bytes: []byte(finalName),
				})

				// Add equals sign
				objTokens = append(objTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenEqual,
					Bytes: []byte(" = "),
				})

				// Add the attribute value
				objTokens = append(objTokens, attr.Expr().BuildTokens(nil)...)

				// Add newline
				objTokens = append(objTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenNewline,
					Bytes: []byte("\n"),
				})
			}

			// Process any remaining attributes not in the predefined order
			for name, attr := range attrs {
				// Skip if already processed
				if processedAttrs[name] {
					continue
				}

				// Add indentation
				objTokens = append(objTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenIdent,
					Bytes: []byte("    "),
				})

				// Add attribute name
				objTokens = append(objTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenIdent,
					Bytes: []byte(name),
				})

				// Add equals sign
				objTokens = append(objTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenEqual,
					Bytes: []byte(" = "),
				})

				// Add the attribute value
				objTokens = append(objTokens, attr.Expr().BuildTokens(nil)...)

				// Add newline
				objTokens = append(objTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenNewline,
					Bytes: []byte("\n"),
				})
			}

			// Add indentation before closing brace
			objTokens = append(objTokens, &hclwrite.Token{
				Type:  hclsyntax.TokenIdent,
				Bytes: []byte("  "),
			})
			objTokens = append(objTokens, &hclwrite.Token{
				Type:  hclsyntax.TokenCBrace,
				Bytes: []byte("}"),
			})

			// Set the data attribute with the object expression
			block.Body().SetAttributeRaw("data", objTokens)

			// Mark the block for removal
			dataBlocksToRemove = append(dataBlocksToRemove, dataBlock)
		}

		// Remove the data blocks after converting to attributes
		for _, dataBlock := range dataBlocksToRemove {
			block.Body().RemoveBlock(dataBlock)
		}

		// Also handle data as an attribute (not a block)
		dataAttr := block.Body().GetAttribute("data")
		if dataAttr != nil && recordType == "CAA" {
			// This is more complex as we need to parse and modify the map
			// For now, we'll use a simple approach
			expr := dataAttr.Expr()
			tokens := expr.BuildTokens(nil)

			// Look for pattern to replace:
			// content = -> value =
			// Process tokens to handle transformations
			newTokens := make(hclwrite.Tokens, 0, len(tokens))
			for i := 0; i < len(tokens); i++ {
				token := tokens[i]

				// Check if this is "flags" attribute with a quoted string value
				if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == "flags" {
					// Look ahead for = and then a quoted string
					if i+1 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenEqual {
						// Find the value token (skip whitespace)
						valueIdx := i + 2
						for valueIdx < len(tokens) && tokens[valueIdx].Type == hclsyntax.TokenNewline {
							valueIdx++
						}

						if valueIdx < len(tokens) && tokens[valueIdx].Type == hclsyntax.TokenQuotedLit {
							// This is flags = "number" - convert to flags = number
							quotedValue := string(tokens[valueIdx].Bytes)
							// Remove quotes and check if it's a number
							unquoted := strings.Trim(quotedValue, `"`)
							if _, err := strconv.Atoi(unquoted); err == nil {
								// It's a number in quotes - add tokens without quotes
								newTokens = append(newTokens, token)       // flags
								newTokens = append(newTokens, tokens[i+1]) // =
								// Add the number without quotes
								numberToken := &hclwrite.Token{
									Type:  hclsyntax.TokenNumberLit,
									Bytes: []byte(unquoted),
								}
								newTokens = append(newTokens, numberToken)
								// Skip to after the quoted value
								i = valueIdx
								continue
							}
						}
					}
				}

				// Check if this is "content" identifier inside data - rename to "value"
				if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == "content" {
					// Check if this is an attribute name (followed by =)
					if i+1 < len(tokens) && (tokens[i+1].Type == hclsyntax.TokenEqual ||
						(string(tokens[i+1].Bytes) == " " && i+2 < len(tokens) && tokens[i+2].Type == hclsyntax.TokenEqual)) {
						// Replace "content" with "value"
						valueToken := &hclwrite.Token{
							Type:  hclsyntax.TokenIdent,
							Bytes: []byte("value"),
						}
						newTokens = append(newTokens, valueToken)
					} else {
						// Not an attribute name, keep as is
						newTokens = append(newTokens, token)
					}
				} else {
					// Keep all other tokens as-is, including flags with numeric values
					newTokens = append(newTokens, token)
				}
			}

			// Set the modified expression back
			block.Body().SetAttributeRaw("data", newTokens)
		}
	}
	return nil
}

// transformDNSRecordStateJSON transforms DNS record state entries from v4 to v5

func transformDNSRecordStateJSON(result string, path string, instance gjson.Result) string {
	// Ensure meta field exists as a JSON string (not null)
	meta := instance.Get("attributes.meta")
	if !meta.Exists() || meta.Type == gjson.Null {
		result, _ = sjson.Set(result, path+".attributes.meta", "{}")
	}

	// Ensure settings field exists with proper structure
	settings := instance.Get("attributes.settings")
	if !settings.Exists() || settings.Type == gjson.Null {
		settingsObj := map[string]interface{}{
			"flatten_cname": nil,
			"ipv4_only":     nil,
			"ipv6_only":     nil,
		}
		result, _ = sjson.Set(result, path+".attributes.settings", settingsObj)
	}

	// Ensure proxiable field exists
	proxiable := instance.Get("attributes.proxiable")
	if !proxiable.Exists() {
		result, _ = sjson.Set(result, path+".attributes.proxiable", false)
	}

	// Ensure timestamp fields exist with default values if missing
	// These are computed fields that should always exist in v5
	createdOn := instance.Get("attributes.created_on")
	if !createdOn.Exists() {
		// If created_on doesn't exist, set it to a default RFC3339 timestamp
		result, _ = sjson.Set(result, path+".attributes.created_on", "2024-01-01T00:00:00Z")
	}

	modifiedOn := instance.Get("attributes.modified_on")
	if !modifiedOn.Exists() {
		// If modified_on doesn't exist, use created_on or set a default
		if createdOn.Exists() {
			result, _ = sjson.Set(result, path+".attributes.modified_on", createdOn.String())
		} else {
			result, _ = sjson.Set(result, path+".attributes.modified_on", "2024-01-01T00:00:00Z")
		}
	}

	// Handle field renames: value -> content (for simple records)
	value := instance.Get("attributes.value")
	content := instance.Get("attributes.content")
	if value.Exists() && !content.Exists() {
		result, _ = sjson.Set(result, path+".attributes.content", value.String())
		result, _ = sjson.Delete(result, path+".attributes.value")
	} else if value.Exists() && content.Exists() {
		// If both exist, remove value since content takes precedence in v5
		result, _ = sjson.Delete(result, path+".attributes.value")
	}

	// Ensure TTL is present (required in v5)
	ttl := instance.Get("attributes.ttl")
	if !ttl.Exists() {
		// TTL is missing - add default value of 1 (automatic)
		result, _ = sjson.Set(result, path+".attributes.ttl", 1.0)
	}

	// Rename metadata -> meta and ensure it's a JSON string
	metadata := instance.Get("attributes.metadata")
	if metadata.Exists() {
		// Convert metadata to JSON string if it's an object
		if metadata.IsObject() {
			result, _ = sjson.Set(result, path+".attributes.meta", "{}")
		} else {
			result, _ = sjson.Set(result, path+".attributes.meta", metadata.String())
		}
		result, _ = sjson.Delete(result, path+".attributes.metadata")
	}

	// Remove deprecated fields
	if instance.Get("attributes.hostname").Exists() {
		result, _ = sjson.Delete(result, path+".attributes.hostname")
	}
	if instance.Get("attributes.allow_overwrite").Exists() {
		result, _ = sjson.Delete(result, path+".attributes.allow_overwrite")
	}
	if instance.Get("attributes.timeouts").Exists() {
		result, _ = sjson.Delete(result, path+".attributes.timeouts")
	}

	// Get record type to determine how to handle data field
	recordType := instance.Get("attributes.type").String()

	// Simple record types that use 'content' field (no data field)
	simpleTypes := map[string]bool{
		"A": true, "AAAA": true, "CNAME": true, "MX": true,
		"NS": true, "PTR": true, "TXT": true, "OPENPGPKEY": true,
	}

	// If it's a simple record type, ensure data field is null
	if simpleTypes[recordType] {
		// Set data field to null for simple types
		result, _ = sjson.Set(result, path+".attributes.data", nil)
		return result
	}

	// Note: For CAA records in v5, content field stays at root level for display purposes
	// The data field contains the structured data

	// For complex record types that use data field
	// v4: data is an array with one object [{ fields... }]
	// v5: data should be an object { fields... }

	// Get the current data field
	data := instance.Get("attributes.data")
	if !data.Exists() {
		// No data field - create empty object for complex types
		dataObj := make(map[string]interface{})
		// For CAA records, set flags to null
		if recordType == "CAA" {
			dataObj["flags"] = nil
		}

		result, _ = sjson.Set(result, path+".attributes.data", dataObj)
		return result
	}

	// Transform data from array to object if needed
	// Only include fields that were present in v4 state
	dataObj := make(map[string]interface{})

	if data.IsArray() {
		// v4 format - data is an array, need to convert to object
		dataArray := data.Array()
		if len(dataArray) > 0 {
			// Take the first element and copy only its existing fields
			firstElem := dataArray[0]

			// Copy all fields from the array element to the object
			firstElem.ForEach(func(key, value gjson.Result) bool {
				k := key.String()

				// Skip 'name' field - it should not be in data
				if k == "name" {
					return true
				}

				// Skip 'proto' field - not supported in v5
				if k == "proto" {
					return true
				}

				// For CAA records, rename 'content' to 'value' in data
				if recordType == "CAA" && k == "content" {
					dataObj["value"] = value.String()
					return true
				}

				// Special handling for flags field - wrap in correct format
				if k == "flags" {
					switch value.Type {
					case gjson.Number:
						// Preserve as json.Number to maintain exact numeric type
						dataObj[k] = map[string]interface{}{
							"value": json.Number(value.Raw),
							"type":  "number",
						}
					case gjson.String:
						// Convert string to number if possible
						if _, err := strconv.ParseFloat(value.String(), 64); err == nil {
							// It's a valid number, preserve as json.Number
							dataObj[k] = map[string]interface{}{
								"value": json.Number(value.String()),
								"type":  "number",
							}
						} else if value.String() == "" {
							dataObj[k] = nil
						} else {
							dataObj[k] = map[string]interface{}{
								"value": value.String(),
								"type":  "string",
							}
						}
					case gjson.Null:
						dataObj[k] = nil
					}
				} else if recordType == "CAA" && k == "value" {
					// For CAA records, 'value' field stays in data
					dataObj[k] = value.String()
				} else {
					// Handle different value types normally for other fields
					switch value.Type {
					case gjson.Number:
						dataObj[k] = value.Float()
					case gjson.String:
						dataObj[k] = value.String()
					case gjson.True:
						dataObj[k] = true
					case gjson.False:
						dataObj[k] = false
					case gjson.Null:
						dataObj[k] = nil
					default:
						// For arrays/objects, use the raw value
						dataObj[k] = value.Value()
					}
				}
				return true
			})
		}

		// For CAA records, if flags field doesn't exist in dataObj, set it to null
		if recordType == "CAA" {
			if _, hasFlags := dataObj["flags"]; !hasFlags {
				dataObj["flags"] = nil
			}
		}

		// Set the data back as an object (not array) with all fields
		result, _ = sjson.Set(result, path+".attributes.data", dataObj)
	} else if data.IsObject() {
		// Already an object - only include existing fields
		dataObj = make(map[string]interface{})

		// Copy existing values from the current data object
		data.ForEach(func(key, value gjson.Result) bool {
			k := key.String()

			// Skip 'name' field - it should not be in data
			if k == "name" {
				return true
			}

			// Skip 'proto' field - not supported in v5
			if k == "proto" {
				return true
			}

			// Special handling for flags - ensure it's wrapped
			if k == "flags" {
				// Check if already wrapped
				if value.IsObject() && value.Get("value").Exists() {
					// Already wrapped, keep as is
					dataObj[k] = value.Value()
				} else {
					// Need to wrap it
					switch value.Type {
					case gjson.Number:
						// Preserve as json.Number to maintain exact numeric type
						dataObj[k] = map[string]interface{}{
							"value": json.Number(value.Raw),
							"type":  "number",
						}
					case gjson.String:
						// Convert string to number if possible
						if _, err := strconv.ParseFloat(value.String(), 64); err == nil {
							// It's a valid number, preserve as json.Number
							dataObj[k] = map[string]interface{}{
								"value": json.Number(value.String()),
								"type":  "number",
							}
						} else if value.String() == "" {
							dataObj[k] = nil
						} else {
							// Keep as string
							dataObj[k] = map[string]interface{}{
								"value": value.String(),
								"type":  "string",
							}
						}

					case gjson.Null:
						dataObj[k] = nil
					}
				}
			} else {
				// Handle other fields normally
				switch value.Type {
				case gjson.Number:
					dataObj[k] = value.Float()
				case gjson.String:
					dataObj[k] = value.String()
				case gjson.True:
					dataObj[k] = true
				case gjson.False:
					dataObj[k] = false
				case gjson.Null:
					dataObj[k] = nil
				default:
					dataObj[k] = value.Value()
				}
			}
			return true
		})

		// For CAA records, also check if we need to rename content to value
		if recordType == "CAA" {
			if content := data.Get("content"); content.Exists() {
				dataObj["value"] = content.String()
				// Remove content from dataObj if it exists
				delete(dataObj, "content")
			}
		}

		if _, hasFlags := dataObj["flags"]; !hasFlags {
			dataObj["flags"] = nil
		}

		// Set the data back with all fields
		result, _ = sjson.Set(result, path+".attributes.data", dataObj)
	}

	// For SRV and URI records, ensure priority is at root level if it exists in data
	// MX records have priority at root level in both v4 and v5
	if recordType == "SRV" || recordType == "URI" {
		// Check if priority exists in dataObj
		if priority, ok := dataObj["priority"]; ok && priority != nil {
			// Set priority at root level as well
			result, _ = sjson.Set(result, path+".attributes.priority", priority)
		}
	}

	return result
}
