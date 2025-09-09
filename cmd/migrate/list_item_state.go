package main

import (
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// migrateListItemsInState performs cross-resource state migration for cloudflare_list_item
// This merges all cloudflare_list_item resources into their parent cloudflare_list resources
// and then removes the list_item resources from state
func migrateListItemsInState(jsonStr string) string {
	result := jsonStr

	// Step 1: Collect all cloudflare_list and cloudflare_list_item resources
	lists := make(map[string]int)      // list_id -> resource index
	listItems := []ListItemStateInfo{} // All list items to migrate

	resources := gjson.Get(jsonStr, "resources")
	if resources.Exists() && resources.IsArray() {
		for i, resource := range resources.Array() {
			resourceType := resource.Get("type").String()

			if resourceType == "cloudflare_list" {
				// Store list resource index by its ID
				listID := resource.Get("instances.0.attributes.id").String()
				if listID != "" {
					lists[listID] = i
				}
			} else if resourceType == "cloudflare_list_item" {
				// Collect list item information
				item := extractListItemStateInfo(resource, i)
				if item != nil {
					listItems = append(listItems, *item)
				}
			}
		}
	}

	// Step 2: Group list items by their parent list and merge them
	for listID, listIndex := range lists {
		var itemsForList []ListItemStateInfo
		for _, item := range listItems {
			if item.ListID == listID {
				itemsForList = append(itemsForList, item)
			}
		}

		if len(itemsForList) > 0 {
			// Merge items into the parent list
			listPath := fmt.Sprintf("resources.%d", listIndex)
			result = mergeItemsIntoListState(result, listPath, itemsForList)
		}
	}

	// Step 3: Remove all cloudflare_list_item resources from state
	result = removeListItemResourcesFromState(result)

	return result
}

// ListItemStateInfo holds state information for a cloudflare_list_item
type ListItemStateInfo struct {
	ListID       string
	AccountID    string
	ResourcePath string // Path to remove this resource
	ItemData     map[string]interface{}
}

// extractListItemStateInfo extracts information from a cloudflare_list_item resource
func extractListItemStateInfo(resource gjson.Result, index int) *ListItemStateInfo {
	attrs := resource.Get("instances.0.attributes")
	if !attrs.Exists() {
		return nil
	}

	listID := attrs.Get("list_id").String()
	if listID == "" {
		return nil
	}

	info := &ListItemStateInfo{
		ListID:       listID,
		AccountID:    attrs.Get("account_id").String(),
		ResourcePath: fmt.Sprintf("resources.%d", index),
		ItemData:     make(map[string]interface{}),
	}

	// Extract item data based on what fields are present
	// Handle IP items
	if ip := attrs.Get("ip"); ip.Exists() && ip.String() != "" {
		info.ItemData["ip"] = ip.String()
	}

	// Handle ASN items
	if asn := attrs.Get("asn"); asn.Exists() && asn.Type == gjson.Number {
		info.ItemData["asn"] = asn.Int()
	}

	// Handle hostname items - transform from array to object
	if hostname := attrs.Get("hostname"); hostname.Exists() {
		if hostname.IsArray() && len(hostname.Array()) > 0 {
			// v4 format: hostname is an array
			info.ItemData["hostname"] = hostname.Array()[0].Value()
		} else if hostname.IsObject() {
			// Already in v5 format
			info.ItemData["hostname"] = hostname.Value()
		}
	}

	// Handle redirect items - transform from array to object and convert booleans
	if redirect := attrs.Get("redirect"); redirect.Exists() {
		var redirectObj map[string]interface{}
		
		if redirect.IsArray() && len(redirect.Array()) > 0 {
			// v4 format: redirect is an array
			redirectData := redirect.Array()[0]
			redirectObj = make(map[string]interface{})
			
			// Copy fields
			if sourceURL := redirectData.Get("source_url"); sourceURL.Exists() {
				redirectObj["source_url"] = sourceURL.String()
			}
			if targetURL := redirectData.Get("target_url"); targetURL.Exists() {
				redirectObj["target_url"] = targetURL.String()
			}
			if statusCode := redirectData.Get("status_code"); statusCode.Exists() {
				redirectObj["status_code"] = statusCode.Int()
			}
			
			// Transform string booleans to actual booleans
			transformBoolField := func(fieldName string) {
				if field := redirectData.Get(fieldName); field.Exists() {
					if field.String() == "enabled" {
						redirectObj[fieldName] = true
					} else if field.String() == "disabled" {
						redirectObj[fieldName] = false
					} else if field.Type == gjson.True || field.Type == gjson.False {
						redirectObj[fieldName] = field.Bool()
					}
				}
			}
			
			transformBoolField("include_subdomains")
			transformBoolField("subpath_matching")
			transformBoolField("preserve_query_string")
			transformBoolField("preserve_path_suffix")
		} else if redirect.IsObject() {
			// Already in object format, but might need boolean conversion
			redirectObj = redirect.Value().(map[string]interface{})
		}
		
		if redirectObj != nil {
			info.ItemData["redirect"] = redirectObj
		}
	}

	// Handle comment
	if comment := attrs.Get("comment"); comment.Exists() && comment.String() != "" {
		info.ItemData["comment"] = comment.String()
	}

	return info
}

// mergeItemsIntoListState merges list items into a cloudflare_list resource in state
func mergeItemsIntoListState(jsonStr string, listResourcePath string, items []ListItemStateInfo) string {
	result := jsonStr

	// Get existing items array if present
	itemsPath := listResourcePath + ".instances.0.attributes.items"
	existingItems := gjson.Get(jsonStr, itemsPath)

	var allItems []interface{}

	// Parse existing items
	if existingItems.Exists() && existingItems.IsArray() {
		for _, item := range existingItems.Array() {
			allItems = append(allItems, item.Value())
		}
	}

	// Add new items from list_item resources
	for _, item := range items {
		allItems = append(allItems, item.ItemData)
	}

	// Update state with merged items and num_items count
	if len(allItems) > 0 {
		result, _ = sjson.Set(result, itemsPath, allItems)
		// Set num_items to match the count of items
		numItemsPath := listResourcePath + ".instances.0.attributes.num_items"
		result, _ = sjson.Set(result, numItemsPath, float64(len(allItems)))
	} else {
		// Set num_items to 0 for empty lists
		numItemsPath := listResourcePath + ".instances.0.attributes.num_items"
		result, _ = sjson.Set(result, numItemsPath, float64(0))
	}

	return result
}

// removeListItemResourcesFromState removes all cloudflare_list_item resources from state
func removeListItemResourcesFromState(jsonStr string) string {
	result := jsonStr

	// Get all resources
	resources := gjson.Get(jsonStr, "resources")
	if !resources.Exists() || !resources.IsArray() {
		return result
	}

	// Build new resources array without list_item resources
	var newResources []interface{}
	for _, resource := range resources.Array() {
		resourceType := resource.Get("type").String()
		if resourceType != "cloudflare_list_item" {
			newResources = append(newResources, resource.Value())
		}
	}

	// Update state with filtered resources
	result, _ = sjson.Set(result, "resources", newResources)

	return result
}