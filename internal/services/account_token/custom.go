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
		
		// Handle permission_groups - sort for consistent ordering
		if pgVal, ok := attrs["permission_groups"].(basetypes.ListValue); ok && !pgVal.IsNull() {
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

// extractAndSortPermissionGroups extracts permission groups and sorts them by ID
func extractAndSortPermissionGroups(pgList basetypes.ListValue) []map[string]interface{} {
	type pgItem struct {
		id   string
		data map[string]interface{}
	}
	
	items := make([]pgItem, 0, len(pgList.Elements()))
	for _, elem := range pgList.Elements() {
		if obj, ok := elem.(basetypes.ObjectValue); ok {
			attrs := obj.Attributes()
			item := pgItem{data: make(map[string]interface{})}
			
			if idVal, ok := attrs["id"].(basetypes.StringValue); ok && !idVal.IsNull() {
				item.id = idVal.ValueString()
				item.data["id"] = item.id
			}
			
			if nameVal, ok := attrs["name"].(basetypes.StringValue); ok && !nameVal.IsNull() {
				item.data["name"] = nameVal.ValueString()
			}
			
			items = append(items, item)
		}
	}
	
	// Sort by ID for consistent ordering
	sort.Slice(items, func(i, j int) bool {
		return items[i].id < items[j].id
	})
	
	result := make([]map[string]interface{}, len(items))
	for i, item := range items {
		result[i] = item.data
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
