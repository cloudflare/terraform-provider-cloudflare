package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/netip"
	"sort"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/rules"
)

// CFListStateV4 represents the complete Terraform state file structure for v4
type CFListStateV4 struct {
	Version          int                    `json:"version"`
	TerraformVersion string                 `json:"terraform_version"`
	Serial           int                    `json:"serial"`
	Lineage          string                 `json:"lineage"`
	Outputs          map[string]interface{} `json:"outputs"`
	Resources        []ResourceV4           `json:"resources"`
	CheckResults     json.RawMessage        `json:"check_results"`
}

// ResourceV4 represents a single resource in the state
type ResourceV4 struct {
	Mode      string       `json:"mode"`
	Type      string       `json:"type"`
	Name      string       `json:"name"`
	Provider  string       `json:"provider"`
	Instances []InstanceV4 `json:"instances"`
}

// InstanceV4 represents an instance of a resource
type InstanceV4 struct {
	SchemaVersion       int                        `json:"schema_version"`
	Attributes          CloudflareListAttributesV4 `json:"attributes"`
	SensitiveAttributes []interface{}              `json:"sensitive_attributes"`
	Private             json.RawMessage            `json:"private,omitempty"`
	Dependencies        []string                   `json:"dependencies,omitempty"`
}

// CloudflareListAttributesV4 represents the attributes of a cloudflare_list resource in v4
type CloudflareListAttributesV4 struct {
	AccountID   string       `json:"account_id"`
	Description string       `json:"description,omitempty"`
	ID          string       `json:"id"`
	Item        []ListItemV4 `json:"item,omitempty"`
	Kind        string       `json:"kind"`
	Name        string       `json:"name"`

	// Additional fields that may be present
	CreatedOn             string `json:"created_on,omitempty"`
	ModifiedOn            string `json:"modified_on,omitempty"`
	NumItems              int    `json:"num_items,omitempty"`
	NumReferencingFilters int    `json:"num_referencing_filters,omitempty"`
}

// ListItemV4 represents an item in a cloudflare_list in v4
type ListItemV4 struct {
	Comment string            `json:"comment,omitempty"`
	Value   []ListItemValueV4 `json:"value"`
}

// ListItemValueV4 represents the value of a list item in v4
// Note: In v4, value is an array with a single element
type ListItemValueV4 struct {
	// For ASN lists
	ASN interface{} `json:"asn,omitempty"` // Can be int or string

	// For IP lists
	IP interface{} `json:"ip,omitempty"` // Can be string or null

	// For hostname lists - v4 uses array format
	Hostname []HostnameValueV4 `json:"hostname,omitempty"`

	// For redirect lists - v4 uses array format
	Redirect []RedirectValueV4 `json:"redirect,omitempty"`
}

// HostnameValueV4 represents a hostname value in v4
type HostnameValueV4 struct {
	URLHostname string `json:"url_hostname"`
	// Only applies to wildcard hostnames (e.g., *.example.com). When true (default),
	// only subdomains are blocked. When false, both the root domain and subdomains are
	// blocked.
	ExcludeExactHostname bool `json:"exclude_exact_hostname"`
}

// RedirectValueV4 represents a redirect value in v4
type RedirectValueV4 struct {
	SourceURL           string      `json:"source_url"`
	TargetURL           string      `json:"target_url"`
	IncludeSubdomains   string      `json:"include_subdomains,omitempty"`    // string ("enabled"/"disabled") in v4
	SubpathMatching     string      `json:"subpath_matching,omitempty"`      // string ("enabled"/"disabled") in v4
	PreserveQueryString string      `json:"preserve_query_string,omitempty"` // string ("enabled"/"disabled") in v4
	PreservePathSuffix  string      `json:"preserve_path_suffix,omitempty"`  // string ("enabled"/"disabled") in v4
	StatusCode          interface{} `json:"status_code,omitempty"`           // Can be int or string
}

// Helper methods for working with the v4 state

// ParseCFListStateV4 parses JSON bytes into a CFListStateV4 struct
func ParseCFListStateV4(data []byte) (*CFListStateV4, error) {
	var state CFListStateV4
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, err
	}
	return &state, nil
}

// ToJSON converts the state to JSON bytes
func (s *CFListStateV4) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}

// ToJSONIndented converts the state to indented JSON bytes
func (s *CFListStateV4) ToJSONIndented() ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}

// FindResource finds all resources by type and name
func (s *CFListStateV4) FindResource(resourceType, name string) []ResourceV4 {
	var matches []ResourceV4
	for _, r := range s.Resources {
		if r.Type == resourceType && r.Name == name {
			matches = append(matches, r)
		}
	}
	return matches
}

// GetCloudflareListResources returns all cloudflare_list resources
func (s *CFListStateV4) GetCloudflareListResources() []ResourceV4 {
	var lists []ResourceV4
	for _, r := range s.Resources {
		if r.Type == "cloudflare_list" {
			lists = append(lists, r)
		}
	}
	return lists
}

// ExtractItems extracts all items from a cloudflare_list for migration to v5
func (s *CFListStateV4) ExtractItems(listResourceName string) []ListItemV4 {
	resources := s.FindResource("cloudflare_list", listResourceName)
	if len(resources) == 0 || len(resources[0].Instances) == 0 {
		return nil
	}
	return resources[0].Instances[0].Attributes.Item
}

// RemoveItems removes item field from cloudflare_list attributes
// This is used during v4 to v5 migration
func (s *CFListStateV4) RemoveItems() {
	for i := range s.Resources {
		if s.Resources[i].Type == "cloudflare_list" {
			for j := range s.Resources[i].Instances {
				s.Resources[i].Instances[j].Attributes.Item = nil
			}
		}
	}
}

// GetItemValue extracts the actual value from a ListItemV4 based on list kind
func (item *ListItemV4) GetItemValue(kind string) interface{} {
	if len(item.Value) == 0 {
		return nil
	}

	v := item.Value[0]
	switch kind {
	case "asn":
		return v.ASN
	case "ip":
		return v.IP
	case "hostname":
		if len(v.Hostname) > 0 {
			return v.Hostname[0]
		}
		return nil
	case "redirect":
		if len(v.Redirect) > 0 {
			return v.Redirect[0]
		}
		return nil
	default:
		return nil
	}
}

// ConvertToV5State creates a v5 state structure from v4
// This is a helper for migration - it creates the structure but doesn't populate item IDs
func (s *CFListStateV4) ConvertToV5State() *CFListStateV5 {
	v5State := &CFListStateV5{
		Version:          s.Version,
		TerraformVersion: s.TerraformVersion,
		Serial:           s.Serial,
		Lineage:          s.Lineage,
		Outputs:          s.Outputs,
		Resources:        make([]ResourceV5, 0),
		CheckResults:     s.CheckResults,
	}

	// Copy cloudflare_list resources without items
	for _, resource := range s.Resources {
		if resource.Type == "cloudflare_list" {
			v5Resource := ResourceV5{
				Mode:      resource.Mode,
				Type:      resource.Type,
				Name:      resource.Name,
				Provider:  resource.Provider,
				Instances: make([]InstanceV5, len(resource.Instances)),
			}

			for i, instance := range resource.Instances {
				// Create a copy of attributes without items
				v5Attrs := CloudflareListAttributes{
					AccountID:             instance.Attributes.AccountID,
					CreatedOn:             instance.Attributes.CreatedOn,
					Description:           instance.Attributes.Description,
					ID:                    instance.Attributes.ID,
					Kind:                  instance.Attributes.Kind,
					ModifiedOn:            instance.Attributes.ModifiedOn,
					Name:                  instance.Attributes.Name,
					NumItems:              instance.Attributes.NumItems,
					NumReferencingFilters: instance.Attributes.NumReferencingFilters,
					// Item field is intentionally not copied
				}

				v5Instance := InstanceV5{
					SchemaVersion:       instance.SchemaVersion,
					SensitiveAttributes: instance.SensitiveAttributes,
					Private:             instance.Private,
					Dependencies:        instance.Dependencies,
				}

				// Set the attributes
				attrBytes, _ := json.Marshal(v5Attrs)
				v5Instance.Attributes = attrBytes

				v5Resource.Instances[i] = v5Instance
			}

			v5State.Resources = append(v5State.Resources, v5Resource)
		}
	}

	return v5State
}

// Helper function to parse time strings
func parseTimeV4(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

// CFListStateV5 represents the complete Terraform state file structure for v5
type CFListStateV5 struct {
	Version          int                    `json:"version"`
	TerraformVersion string                 `json:"terraform_version"`
	Serial           int                    `json:"serial"`
	Lineage          string                 `json:"lineage"`
	Outputs          map[string]interface{} `json:"outputs"`
	Resources        []ResourceV5           `json:"resources"`
	CheckResults     json.RawMessage        `json:"check_results"`
}

// ResourceV5 represents a single resource in the state
type ResourceV5 struct {
	Mode      string       `json:"mode"`
	Type      string       `json:"type"`
	Name      string       `json:"name"`
	Provider  string       `json:"provider"`
	Instances []InstanceV5 `json:"instances"`
}

// InstanceV5 represents an instance of a resource
type InstanceV5 struct {
	SchemaVersion       int             `json:"schema_version"`
	Attributes          json.RawMessage `json:"attributes"`
	SensitiveAttributes []interface{}   `json:"sensitive_attributes"`
	Private             json.RawMessage `json:"private,omitempty"`
	Dependencies        []string        `json:"dependencies,omitempty"`
}

// CloudflareListAttributes represents the attributes of a cloudflare_list resource
type CloudflareListAttributes struct {
	AccountID             string `json:"account_id"`
	CreatedOn             string `json:"created_on"`
	Description           string `json:"description"`
	ID                    string `json:"id"`
	Kind                  string `json:"kind"`
	ModifiedOn            string `json:"modified_on"`
	Name                  string `json:"name"`
	NumItems              int    `json:"num_items"`
	NumReferencingFilters int    `json:"num_referencing_filters"`

	// Item fields that existed in v4 but are removed in v5
	// These are only used when reading v4 state for migration
	Item  []ListItemV5  `json:"item,omitempty"`
	Items []interface{} `json:"items,omitempty"`
}

// ListItemV5 represents an item in a cloudflare_list (for v4 lists with nested items)
type ListItemV5 struct {
	Value   ListItemValueV5 `json:"value"`
	Comment string          `json:"comment,omitempty"`
}

// ListItemValueV5 represents the value of a list item
// This is a flexible structure that can hold different types based on list kind
type ListItemValueV5 struct {
	// For IP lists
	IP string `json:"ip,omitempty"`

	// For ASN lists
	ASN interface{} `json:"asn,omitempty"` // Can be string or int

	// For hostname lists
	Hostname *HostnameValue `json:"hostname,omitempty"`

	// For redirect lists
	Redirect *RedirectValue `json:"redirect,omitempty"`
}

// HostnameValue represents a hostname value
type HostnameValue struct {
	URLHostname string `json:"url_hostname,required"`
	// Only applies to wildcard hostnames (e.g., *.example.com). When true (default),
	// only subdomains are blocked. When false, both the root domain and subdomains are
	// blocked.
	ExcludeExactHostname bool `json:"exclude_exact_hostname"`
}

// RedirectValue represents a redirect value
type RedirectValue struct {
	SourceURL           string      `json:"source_url"`
	TargetURL           string      `json:"target_url"`
	IncludeSubdomains   interface{} `json:"include_subdomains,omitempty"`    // bool in v4, string ("enabled"/"disabled") in v5
	SubpathMatching     interface{} `json:"subpath_matching,omitempty"`      // bool in v4, string ("enabled"/"disabled") in v5
	PreserveQueryString interface{} `json:"preserve_query_string,omitempty"` // bool in v4, string ("enabled"/"disabled") in v5
	PreservePathSuffix  interface{} `json:"preserve_path_suffix,omitempty"`  // bool in v4, string ("enabled"/"disabled") in v5
	StatusCode          interface{} `json:"status_code,omitempty"`           // Can be int or string
}

// CloudflareListItemAttributes represents the attributes of a cloudflare_list_item resource (v5)
type CloudflareListItemAttributes struct {
	AccountID   string         `json:"account_id"`
	ASN         *int64         `json:"asn"`     // Pointer to handle null
	Comment     *string        `json:"comment"` // Pointer to handle null
	CreatedOn   string         `json:"created_on"`
	Hostname    *HostnameValue `json:"hostname"` // Pointer to handle null
	ID          string         `json:"id"`
	IP          *string        `json:"ip"` // Pointer to handle null
	ListID      string         `json:"list_id"`
	ModifiedOn  string         `json:"modified_on"`
	OperationID *string        `json:"operation_id"` // Pointer to handle null
	Redirect    *RedirectValue `json:"redirect"`     // Pointer to handle null or object
}

// Helper methods for working with the state

// ParseCFListStateV5 parses JSON bytes into a CFListStateV5 struct
func ParseCFListStateV5(data []byte) (*CFListStateV5, error) {
	var state CFListStateV5
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, err
	}
	return &state, nil
}

// ToJSON converts the state to JSON bytes
func (s *CFListStateV5) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}

// ToJSONIndented converts the state to indented JSON bytes
func (s *CFListStateV5) ToJSONIndented() ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}

// FindResource finds all resources by type and name
func (s *CFListStateV5) FindResource(resourceType, name string) []ResourceV5 {
	var matches []ResourceV5
	for _, r := range s.Resources {
		if r.Type == resourceType && r.Name == name {
			matches = append(matches, r)
		}
	}
	return matches
}

// GetCloudflareListResources returns all cloudflare_list resources
func (s *CFListStateV5) GetCloudflareListResources() []ResourceV5 {
	var lists []ResourceV5
	for _, r := range s.Resources {
		if r.Type == "cloudflare_list" {
			lists = append(lists, r)
		}
	}
	return lists
}

// GetListAttributes parses the raw attributes as CloudflareListAttributes
func (i *InstanceV5) GetListAttributes() (*CloudflareListAttributes, error) {
	var attrs CloudflareListAttributes
	if err := json.Unmarshal(i.Attributes, &attrs); err != nil {
		return nil, err
	}
	return &attrs, nil
}

// GetListItemAttributes parses the raw attributes as CloudflareListItemAttributes
func (i *InstanceV5) GetListItemAttributes() (*CloudflareListItemAttributes, error) {
	var attrs CloudflareListItemAttributes
	if err := json.Unmarshal(i.Attributes, &attrs); err != nil {
		return nil, err
	}
	return &attrs, nil
}

// SetListAttributes sets the attributes from a CloudflareListAttributes struct
func (i *InstanceV5) SetListAttributes(attrs *CloudflareListAttributes) error {
	data, err := json.Marshal(attrs)
	if err != nil {
		return err
	}
	i.Attributes = data
	return nil
}

// SetListItemAttributes sets the attributes from a CloudflareListItemAttributes struct
func (i *InstanceV5) SetListItemAttributes(attrs *CloudflareListItemAttributes) error {
	data, err := json.Marshal(attrs)
	if err != nil {
		return err
	}
	i.Attributes = data
	return nil
}

// RemoveListItems removes item and items fields from cloudflare_list attributes
// This is used to clean up v5 state after migration from v4
func (s *CFListStateV5) RemoveListItems() error {
	for i := range s.Resources {
		if s.Resources[i].Type == "cloudflare_list" {
			for j := range s.Resources[i].Instances {
				// Parse the attributes
				attrs, err := s.Resources[i].Instances[j].GetListAttributes()
				if err != nil {
					return err
				}
				// Clear the item fields
				attrs.Item = nil
				attrs.Items = nil
				// Set back the modified attributes
				if err := s.Resources[i].Instances[j].SetListAttributes(attrs); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// GetCloudflareListItemResources returns all cloudflare_list_item resources
func (s *CFListStateV5) GetCloudflareListItemResources() []ResourceV5 {
	var items []ResourceV5
	for _, r := range s.Resources {
		if r.Type == "cloudflare_list_item" {
			items = append(items, r)
		}
	}
	return items
}

// APIListItem represents a simplified version of a list item from the API
type APIListItem struct {
	ID         string
	CreatedOn  string
	ModifiedOn string
}

// MigrateCloudflareListToV5 migrates cloudflare_list resources from v4 to v5 state format
// In v4, items are part of the list resource. In v5, they become separate cloudflare_list_item resources.
func MigrateCloudflareListToV5(v4JSON []byte) (*CFListStateV5, error) {
	c := cloudflare.NewClient()
	// Step 1: Parse v4 state
	v4State, err := ParseCFListStateV4(v4JSON)
	if err != nil {
		return nil, fmt.Errorf("failed to parse v4 state: %w", err)
	}

	// Step 2: Process all cloudflare_list resources
	allListStateResources := v4State.GetCloudflareListResources()
	if len(allListStateResources) == 0 {
		//TODO:: handle empty state
		return v4State.ConvertToV5State(), nil
	}

	// Step 3: Convert to v5 state (removes items from lists)
	v5State := v4State.ConvertToV5State()

	// Step 4: Process each list and create list_item resources
	for _, listResource := range allListStateResources {
		listName := listResource.Name

		// Extract items from this specific list
		if len(listResource.Instances) == 0 {
			continue
		}

		listAttrs := listResource.Instances[0].Attributes
		items := listAttrs.Item

		// Update the list's num_items field in v5 state
		if err := updateListNumItems(v5State, listResource.Name, len(items)); err != nil {
			fmt.Printf("  Warning: Failed to update num_items for list %s: %v\n", listName, err)
		}

		if len(items) == 0 {
			continue
		}

		cloudflareListItems := make([]rules.ListItemListResponse, 0)
		cloudflareListItemsResponse, err := c.Rules.Lists.Items.List(context.Background(), listAttrs.ID, rules.ListItemListParams{
			AccountID: cloudflare.F(listAttrs.AccountID),
		})
		if err != nil {
			log.Printf("  Warning: Failed to fetch items for list %s: %v\n", listName, err)
			return nil, err
		}
		cloudflareListItems = append(cloudflareListItems, cloudflareListItemsResponse.Result...)
		nextPage, err := cloudflareListItemsResponse.GetNextPage()
		if err != nil {
			log.Printf("  Warning: Failed to fetch next page %s: %v\n", listName, err)
			return nil, err
		}
		for nextPage != nil {
			cloudflareListItems = append(cloudflareListItems, nextPage.Result...)
			nextPage, err = nextPage.GetNextPage()
			if err != nil {
				log.Printf("  Warning: Failed to fetch next page %s: %v\n", listName, err)
				return nil, err
			}
		}
		// Fetch actual item IDs from the Cloudflare API
		// Note: For testing purposes, we'll use mock data
		apiItemMap := createAPIItemMap(cloudflareListItems, listAttrs.Kind, listAttrs.ID)

		// Sort items by their values for deterministic ordering
		// This ensures config and state migrations use the same order
		sort.Slice(items, func(i, j int) bool {
			keyI := getItemSortKeyFromState(items[i], listAttrs.Kind)
			keyJ := getItemSortKeyFromState(items[j], listAttrs.Kind)
			return keyI < keyJ
		})

		// Create cloudflare_list_item resources for this list
		for i, item := range items {
			itemResource := ResourceV5{
				Mode:     "managed",
				Type:     "cloudflare_list_item",
				Name:     fmt.Sprintf("%s_item_%d", listName, i),
				Provider: listResource.Provider,
				Instances: []InstanceV5{
					{
						SchemaVersion:       0,
						SensitiveAttributes: []interface{}{},
					},
				},
			}

			// Create the item attributes
			itemAttrs := CloudflareListItemAttributes{
				AccountID: listAttrs.AccountID,
				ListID:    listAttrs.ID,
			}

			// Set comment if not empty
			if item.Comment != "" {
				itemAttrs.Comment = &item.Comment
			}

			// Try to find matching API item to get the real ID and timestamps
			itemKey := getItemKeyFromState(item, listAttrs.Kind, listAttrs.ID, apiItemMap)
			if apiItem, found := apiItemMap[itemKey]; found {
				itemAttrs.ID = apiItem.ID
				itemAttrs.CreatedOn = apiItem.CreatedOn
				itemAttrs.ModifiedOn = apiItem.ModifiedOn
			} else {
				// Use placeholder if no match found

				itemAttrs.ID = fmt.Sprintf("item_id_%d_placeholder", i)
				fmt.Printf("  Warning: No API match found for item %d (key: %s)item %+v kind %s \n", i, itemKey, item, listAttrs.Kind)
			}

			// Set the value based on list kind
			switch listAttrs.Kind {
			case "asn":
				if len(item.Value) > 0 && item.Value[0].ASN != nil {
					// Convert ASN to int64
					switch v := item.Value[0].ASN.(type) {
					case float64:
						asn := int64(v)
						itemAttrs.ASN = &asn
					case int:
						asn := int64(v)
						itemAttrs.ASN = &asn
					case int64:
						itemAttrs.ASN = &v
					}
				}
			case "ip":
				if len(item.Value) > 0 && item.Value[0].IP != nil {
					ipStr := fmt.Sprintf("%v", item.Value[0].IP)
					itemAttrs.IP = &ipStr
				}
			case "hostname":
				if len(item.Value) > 0 && len(item.Value[0].Hostname) > 0 {
					// Convert v4 hostname to v5 format
					v4Hostname := item.Value[0].Hostname[0]
					itemAttrs.Hostname = &HostnameValue{
						URLHostname:          v4Hostname.URLHostname,
						ExcludeExactHostname: false, // Default value, adjust as needed
					}
				}
			case "redirect":
				if len(item.Value) > 0 && len(item.Value[0].Redirect) > 0 {
					redirect := item.Value[0].Redirect[0]
					// Convert v4 redirect to v5 format (strings in v4 to booleans in v5)
					v5Redirect := &RedirectValue{
						SourceURL:  redirect.SourceURL,
						TargetURL:  redirect.TargetURL,
						StatusCode: redirect.StatusCode,
					}

					// Convert string ("enabled"/"disabled") to boolean
					if redirect.IncludeSubdomains == "enabled" {
						v5Redirect.IncludeSubdomains = true
					} else {
						v5Redirect.IncludeSubdomains = false
					}

					if redirect.SubpathMatching == "enabled" {
						v5Redirect.SubpathMatching = true
					} else {
						v5Redirect.SubpathMatching = false
					}

					if redirect.PreserveQueryString == "enabled" {
						v5Redirect.PreserveQueryString = true
					} else {
						v5Redirect.PreserveQueryString = false
					}

					if redirect.PreservePathSuffix == "enabled" {
						v5Redirect.PreservePathSuffix = true
					} else {
						v5Redirect.PreservePathSuffix = false
					}

					itemAttrs.Redirect = v5Redirect
				}
			}

			// Set the attributes on the instance
			attrBytes, _ := json.Marshal(itemAttrs)
			itemResource.Instances[0].Attributes = attrBytes

			v5State.Resources = append(v5State.Resources, itemResource)
		}
	}
	return v5State, nil
}

// updateListNumItems updates the num_items field for a cloudflare_list resource in v5 state
func updateListNumItems(v5State *CFListStateV5, listName string, numItems int) error {
	resources := v5State.FindResource("cloudflare_list", listName)
	if len(resources) == 0 {
		return fmt.Errorf("list resource %s not found in v5 state", listName)
	}

	for _, resource := range resources {
		for i := range resource.Instances {
			attrs, err := resource.Instances[i].GetListAttributes()
			if err != nil {
				return fmt.Errorf("failed to get list attributes: %w", err)
			}
			attrs.NumItems = numItems
			if err := resource.Instances[i].SetListAttributes(attrs); err != nil {
				return fmt.Errorf("failed to set list attributes: %w", err)
			}
		}
	}
	return nil
}

// createAPIItemMap creates a mock map of API items for testing
// In production, this would be replaced with actual API calls
func createAPIItemMap(response []rules.ListItemListResponse, listKind, listID string) map[string]APIListItem {
	apiItemMap := make(map[string]APIListItem)

	for _, item := range response {
		key := getItemKey(item, listKind, listID)
		if key != "" {
			apiItemMap[key] = APIListItem{
				ID:         item.ID,
				CreatedOn:  item.CreatedOn,
				ModifiedOn: item.ModifiedOn,
			}
		}
	}

	return apiItemMap
}

// getItemKey generates a unique key for a v4 item based on its value
func getItemKey(item rules.ListItemListResponse, listKind, listID string) string {
	switch listKind {
	case "ip":
		if item.IP != "" {
			return fmt.Sprintf("%s:ip:%v", listID, item.IP)
		}
	case "asn":
		if item.ASN >= 0 {
			return fmt.Sprintf("%s:asn:%d", listID, item.ASN)
		}
	case "hostname":
		if len(item.Hostname.URLHostname) > 0 {
			return fmt.Sprintf("%s:hostname:%s", listID, item.Hostname.URLHostname)
		}
	case "redirect":
		return fmt.Sprintf("%s:redirect:%s->%s", listID, item.Redirect.SourceURL, item.Redirect.TargetURL)
	}
	return ""
}

// getItemSortKeyFromState extracts a sort key from a state item for deterministic ordering
func getItemSortKeyFromState(item ListItemV4, listKind string) string {
	if len(item.Value) == 0 {
		// If no value, use comment as fallback
		if item.Comment != "" {
			return "zzz_" + item.Comment
		}
		return "zzz_unknown"
	}

	v := item.Value[0]
	switch listKind {
	case "ip":
		if v.IP != nil {
			return fmt.Sprintf("%v", v.IP)
		}
	case "asn":
		if v.ASN != nil {
			// Pad ASN for numeric sorting
			var asnStr string
			switch asn := v.ASN.(type) {
			case float64:
				asnStr = fmt.Sprintf("%020d", int64(asn))
			case int:
				asnStr = fmt.Sprintf("%020d", asn)
			case int64:
				asnStr = fmt.Sprintf("%020d", asn)
			case string:
				// Remove quotes and pad
				asnStr = strings.Trim(asn, "\"")
				asnStr = fmt.Sprintf("%020s", asnStr)
			default:
				asnStr = fmt.Sprintf("%020v", asn)
			}
			return asnStr
		}
	case "hostname":
		if len(v.Hostname) > 0 {
			return v.Hostname[0].URLHostname
		}
	case "redirect":
		if len(v.Redirect) > 0 {
			return v.Redirect[0].SourceURL
		}
	}

	// Fallback to comment
	if item.Comment != "" {
		return "zzz_" + item.Comment
	}
	return "zzz_unknown"
}

func getItemKeyFromState(item ListItemV4, listKind string, listID string, apiItemMap map[string]APIListItem) string {
	if len(item.Value) == 0 {
		return ""
	}
	for _, v := range item.Value {
		switch listKind {
		case "ip":
			if v.IP != nil {
				ip := fmt.Sprintf("%v", v.IP)
				var err error
				if strings.Contains(ip, "/") {
					ip, err = hostFromCIDR(fmt.Sprintf("%v", ip))
					if err != nil {
						continue
					}
				}
				key := fmt.Sprintf("%s:ip:%v", listID, ip)
				if _, ok := apiItemMap[key]; ok {
					return key
				}
			}
		case "asn":
			key := fmt.Sprintf("%s:asn:%d", listID, int64(v.ASN.(float64)))
			if _, ok := apiItemMap[key]; ok {
				return key
			}
		case "hostname":
			key := fmt.Sprintf("%s:hostname:%s", listID, v.Hostname[0].URLHostname)
			if _, ok := apiItemMap[key]; ok {
				return key
			}
		case "redirect":
			if len(v.Redirect) > 0 {
				r := v.Redirect[0]
				key := fmt.Sprintf("%s:redirect:%s->%s", listID, r.SourceURL, r.TargetURL)
				if _, ok := apiItemMap[key]; ok {
					return key
				}
			}
		default:
			// Unknown list kind
		}
	}

	return ""
}

func hostFromCIDR(s string) (string, error) {
	p, err := netip.ParsePrefix(s)
	if err != nil {
		return "", err
	}
	a := p.Addr()
	want := 32
	if !a.Is4() {
		want = 128
	}
	if p.Bits() != want {
		return s, nil
	}
	return a.String(), nil
}
