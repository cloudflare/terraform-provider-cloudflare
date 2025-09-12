package account_token

import (
	"context"
	"encoding/json"
	"sort"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func (m AccountTokenModel) marshalCustom() (data []byte, err error) {
	if data, err = apijson.MarshalRoot(m); err != nil {
		return
	}
	return m.marshalPolicies(data)
}

func (m AccountTokenModel) marshalCustomForUpdate(state AccountTokenModel) (data []byte, err error) {
	if data, err = apijson.MarshalForUpdate(m, state); err != nil {
		return
	}
	return m.marshalPolicies(data)
}

func (m AccountTokenModel) marshalPolicies(b []byte) ([]byte, error) {
	var base map[string]json.RawMessage
	if err := json.Unmarshal(b, &base); err != nil {
		return nil, err
	}
	if _, ok := base["policies"]; !ok {
		return b, nil
	}

	// Convert policies Dynamic to proper JSON structure
	if !m.Policies.IsNull() && !m.Policies.IsUnknown() {
		policiesJSON, err := convertDynamicPoliciesToJSON(context.Background(), m.Policies)
		if err == nil && policiesJSON != nil {
			policiesBytes, err := json.Marshal(policiesJSON)
			if err == nil {
				base["policies"] = policiesBytes
				return json.Marshal(base)
			}
		}
	}

	// Fallback to original approach if conversion fails
	type onlyPolicies struct {
		Policies types.Dynamic `json:"policies,required"`
	}
	pb, err := apijson.MarshalRoot(onlyPolicies{Policies: m.Policies})
	if err != nil {
		return nil, err
	}
	var pm map[string]json.RawMessage
	if err := json.Unmarshal(pb, &pm); err != nil {
		return nil, err
	}
	if raw, ok := pm["policies"]; ok {
		base["policies"] = raw
	}
	return json.Marshal(base)
}

// convertDynamicPoliciesToJSON converts Dynamic policies to proper JSON handling resource polymorphism
func convertDynamicPoliciesToJSON(ctx context.Context, policies types.Dynamic) ([]interface{}, error) {
	dynValue, _ := policies.ToDynamicValue(ctx)
	underlying := dynValue.UnderlyingValue()
	
	listVal, ok := underlying.(basetypes.ListValue)
	if !ok {
		return nil, nil
	}

	result := make([]interface{}, 0, len(listVal.Elements()))
	for _, elem := range listVal.Elements() {
		objVal, ok := elem.(basetypes.ObjectValue)
		if !ok {
			continue
		}

		attrs := objVal.Attributes()
		policy := make(map[string]interface{})
		
		// Handle effect
		if effect, ok := attrs["effect"].(basetypes.StringValue); ok && !effect.IsNull() {
			policy["effect"] = effect.ValueString()
		}
		
		// Handle permission_groups - sort for consistent ordering (handles both lists and sets)
		if pgVal, ok := attrs["permission_groups"]; ok && !pgVal.IsNull() && !pgVal.IsUnknown() {
			pgs := extractAndSortPermissionGroups(pgVal)
			if len(pgs) > 0 {
				policy["permission_groups"] = pgs
			}
		}
		
		// Handle resources - this is the polymorphic field
		if resVal, ok := attrs["resources"]; ok && !resVal.IsNull() {
			if resources := extractResources(resVal); resources != nil {
				policy["resources"] = resources
			}
		}
		
		// Handle ID if present
		if idVal, ok := attrs["id"].(basetypes.StringValue); ok && !idVal.IsNull() {
			policy["id"] = idVal.ValueString()
		}
		
		result = append(result, policy)
	}
	
	return result, nil
}

// extractAndSortPermissionGroups extracts permission groups and always treats them as a set
// by sorting them consistently by ID, regardless of input type
func extractAndSortPermissionGroups(pgVal attr.Value) []map[string]interface{} {
	var elements []attr.Value
	
	// Extract elements from any collection type
	switch v := pgVal.(type) {
	case basetypes.ListValue:
		if v.IsNull() {
			return nil
		}
		elements = v.Elements()
	case basetypes.SetValue:
		if v.IsNull() {
			return nil
		}
		elements = v.Elements()
	case basetypes.TupleValue:
		if v.IsNull() {
			return nil
		}
		elements = v.Elements()
	default:
		return nil
	}
	
	// Use a map to deduplicate by ID (set behavior)
	uniqueItems := make(map[string]map[string]interface{})
	
	for _, elem := range elements {
		if obj, ok := elem.(basetypes.ObjectValue); ok {
			attrs := obj.Attributes()
			item := make(map[string]interface{})
			
			var id string
			if idVal, ok := attrs["id"].(basetypes.StringValue); ok && !idVal.IsNull() {
				id = idVal.ValueString()
				item["id"] = id
			}
			
			if nameVal, ok := attrs["name"].(basetypes.StringValue); ok && !nameVal.IsNull() {
				item["name"] = nameVal.ValueString()
			}
			
			// Only keep unique IDs (set behavior)
			if id != "" {
				uniqueItems[id] = item
			}
		}
	}
	
	// Extract IDs and sort them for consistent ordering
	ids := make([]string, 0, len(uniqueItems))
	for id := range uniqueItems {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	
	// Build result in sorted order
	result := make([]map[string]interface{}, 0, len(ids))
	for _, id := range ids {
		result = append(result, uniqueItems[id])
	}
	
	return result
}

// extractResources handles the polymorphic resources field
func extractResources(resVal attr.Value) map[string]interface{} {
	// Handle MapValue (most common case)
	if mapVal, ok := resVal.(basetypes.MapValue); ok && !mapVal.IsNull() {
		result := make(map[string]interface{})
		for key, val := range mapVal.Elements() {
			// Simple string value
			if strVal, ok := val.(basetypes.StringValue); ok && !strVal.IsNull() {
				result[key] = strVal.ValueString()
			} else if dynVal, ok := val.(basetypes.DynamicValue); ok && !dynVal.IsNull() {
				// Nested dynamic value (could be string or map)
				underlying := dynVal.UnderlyingValue()
				if strVal, ok := underlying.(basetypes.StringValue); ok && !strVal.IsNull() {
					result[key] = strVal.ValueString()
				} else if nestedMap, ok := underlying.(basetypes.MapValue); ok && !nestedMap.IsNull() {
					// Nested map case
					nested := make(map[string]interface{})
					for nKey, nVal := range nestedMap.Elements() {
						if nStrVal, ok := nVal.(basetypes.StringValue); ok && !nStrVal.IsNull() {
							nested[nKey] = nStrVal.ValueString()
						}
					}
					result[key] = nested
				}
			}
		}
		return result
	}
	
	// Handle DynamicValue wrapping a map
	if dynVal, ok := resVal.(basetypes.DynamicValue); ok && !dynVal.IsNull() {
		return extractResources(dynVal.UnderlyingValue())
	}
	
	return nil
}
