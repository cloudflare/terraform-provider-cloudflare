package acctest

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
)

var (
	// Integration test account ID.
	TestAccCloudflareAccountID string = "f037e56e89293a057740de681ac9abbe"

	// Integration test account zone ID.
	TestAccCloudflareZoneID string = "0da42c8d2132a9ddaf714f9e7c920711"
	// Integration test account zone name.
	TestAccCloudflareZoneName string = "terraform.cfapi.net"

	// Integration test account alternate zone ID.
	TestAccCloudflareAltZoneID string = "b72110c08e3382597095c29ba7e661ea"
	// Integration test account alternate zone name.
	TestAccCloudflareAltZoneName string = "terraform2.cfapi.net"
)

var TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"cloudflare": func() (tfprotov6.ProviderServer, error) {
		return providerserver.NewProtocol6(internal.NewProvider("dev")())(), nil
	},
}

func TestAccPreCheck(t *testing.T) {
	TestAccPreCheck_Credentials(t)
}

func TestAccPreCheck_Credentials(t *testing.T) {
	apiKey := os.Getenv(consts.APIKeyEnvVarKey)
	apiToken := os.Getenv(consts.APITokenEnvVarKey)
	userServiceKey := os.Getenv(consts.APIUserServiceKeyEnvVarKey)

	if apiToken == "" && apiKey == "" && userServiceKey == "" {
		t.Fatalf(
			"valid credentials are required for this acceptance test: one of %s, %s, or %s must be set",
			consts.APIKeyEnvVarKey,
			consts.APITokenEnvVarKey,
			consts.APIUserServiceKeyEnvVarKey,
		)
	}
}

// Test helper method checking `CLOUDFLARE_ZONE_ID` is present.
func TestAccPreCheck_ZoneID(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_ZONE_ID"); v == "" {
		t.Fatal("CLOUDFLARE_ZONE_ID must be set for this acceptance test.")
	}
}

// Test helper method checking `CLOUDFLARE_ACCOUNT_ID` is present.
func TestAccPreCheck_AccountID(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_ACCOUNT_ID"); v == "" {
		t.Fatal("CLOUDFLARE_ACCOUNT_ID must be set for this acceptance test.")
	}
}

// Test helper method checking `CLOUDFLARE_DOMAIN` is present.
func TestAccPreCheck_Domain(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_DOMAIN"); v == "" {
		t.Fatal("CLOUDFLARE_DOMAIN must be set for acceptance tests. The domain is used to create and destroy record against.")
	}
}

// Test helper method checking `CLOUDFLARE_ALT_DOMAIN` is present.
func TestAccPreCheck_AlternateDomain(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_ALT_DOMAIN"); v == "" {
		t.Fatal("CLOUDFLARE_ALT_DOMAIN must be set for this acceptance test.")
	}
}

// Test helper method checking `CLOUDFLARE_ALT_ZONE_ID` is present.
func TestAccPreCheck_AlternateZoneID(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_ALT_ZONE_ID"); v == "" {
		t.Fatal("CLOUDFLARE_ALT_ZONE_ID must be set for this acceptance test.")
	}
}

// Test helper method checking `CLOUDFLARE_EMAIL` is present.
func TestAccPreCheck_Email(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_EMAIL"); v == "" {
		t.Fatal("CLOUDFLARE_EMAIL must be set for acceptance tests")
	}
}

// Test helper method checking `CLOUDFLARE_API_KEY` is present.
func TestAccPreCheck_APIKey(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_API_KEY"); v == "" {
		t.Fatal("CLOUDFLARE_API_KEY must be set for acceptance tests")
	}
}

// Test helper method checking `CLOUDFLARE_API_TOKEN` is present.
func TestAccPreCheck_APIToken(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_API_TOKEN"); v == "" {
		t.Fatal("CLOUDFLARE_API_TOKEN must be set for acceptance tests")
	}
}

// Test helper method checking `CLOUDFLARE_LOGPUSH_OWNERSHIP_TOKEN` is present.
func TestAccPreCheck_LogpushToken(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_LOGPUSH_OWNERSHIP_TOKEN"); v == "" {
		t.Fatal("CLOUDFLARE_LOGPUSH_OWNERSHIP_TOKEN must be set for this acceptance test.")
	}
}

// Test helper method checking the CrowdStrike environment variables are present.
func TestAccPreCheck_CrowdStrike(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CLIENT_ID"); v == "" {
		t.Skip("Skipping acceptance test as CLOUDFLARE_CROWDSTRIKE_CLIENT_ID is not set")
	}

	if v := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CLIENT_SECRET"); v == "" {
		t.Skip("Skipping acceptance test as CLOUDFLARE_CROWDSTRIKE_CLIENT_SECRET is not set")
	}

	if v := os.Getenv("CLOUDFLARE_CROWDSTRIKE_API_URL"); v == "" {
		t.Skip("Skipping acceptance test as CLOUDFLARE_CROWDSTRIKE_API_URL is not set")
	}

	if v := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CUSTOMER_ID"); v == "" {
		t.Skip("Skipping acceptance test as CLOUDFLARE_CROWDSTRIKE_CUSTOMER_ID is not set")
	}
}

// Test helper method checking the required environment variables for Cloudflare Pages
// are present.
func TestAccPreCheck_Pages(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_PAGES_OWNER"); v == "" {
		t.Fatal("CLOUDFLARE_PAGES_OWNER must be set for this acceptance test.")
	}

	if v := os.Getenv("CLOUDFLARE_PAGES_REPO"); v == "" {
		t.Fatal("CLOUDFLARE_PAGES_REPO must be set for this acceptance test.")
	}
}

// Test helper method checking `CLOUDFLARE_BYO_IP_PREFIX_ID` is present.
func TestAccPreCheck_BYOIPPrefix(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_BYO_IP_PREFIX_ID"); v == "" {
		t.Skip("Skipping acceptance test as CLOUDFLARE_BYO_IP_PREFIX_ID is not set")
	}
}

// Test helper method checking all required Hyperdrive configurations are present.
func TestAccPreCheck_Hyperdrive(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_NAME"); v == "" {
		t.Fatal("CLOUDFLARE_HYPERDRIVE_DATABASE_NAME must be set for this acceptance test")
	}

	if v := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_HOSTNAME"); v == "" {
		t.Fatal("CLOUDFLARE_HYPERDRIVE_DATABASE_HOSTNAME must be set for this acceptance test")
	}

	if v := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_PORT"); v == "" {
		t.Fatal("CLOUDFLARE_HYPERDRIVE_DATABASE_PORT must be set for this acceptance test")
	}

	if v := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_USER"); v == "" {
		t.Fatal("CLOUDFLARE_HYPERDRIVE_DATABASE_USER must be set for this acceptance test")
	}

	if v := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_PASSWORD"); v == "" {
		t.Fatal("CLOUDFLARE_HYPERDRIVE_DATABASE_PASSWORD must be set for this acceptance test")
	}
}

func TestAccPreCheck_HyperdriveWithAccess(t *testing.T) {
	TestAccPreCheck_Hyperdrive(t)

	if v := os.Getenv("CLOUDFLARE_HYPERDRIVE_ACCESS_CLIENT_ID"); v == "" {
		t.Fatal("CLOUDFLARE_HYPERDRIVE_ACCESS_CLIENT_ID must be set for this acceptance test")
	}

	if v := os.Getenv("CLOUDFLARE_HYPERDRIVE_ACCESS_CLIENT_SECRET"); v == "" {
		t.Fatal("CLOUDFLARE_HYPERDRIVE_ACCESS_CLIENT_SECRET must be set for this acceptance test")
	}
}

// Test helper method checking `CLOUDFLARE_INTERNAL_ZONE_ID` is present.
func TestAccPreCheck_InternalZoneID(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_INTERNAL_ZONE_ID"); v == "" {
		t.Skip("Skipping acceptance test as CLOUDFLARE_INTERNAL_ZONE_ID is not set")
	}
}

// TestAccSkipForDefaultZone is used for skipping over tests that are not run by
// default on usual acceptance test suite account.
func TestAccSkipForDefaultZone(t *testing.T, reason string) {
	if os.Getenv("CLOUDFLARE_ZONE_ID") == TestAccCloudflareZoneID {
		t.Skipf("Skipping acceptance test for default zone (%s). %s", TestAccCloudflareZoneID, reason)
	}
}

// TestAccSkipForDefaultAccount is used for skipping over tests that are not run by
// default on usual acceptance test suite account.
func TestAccSkipForDefaultAccount(t *testing.T, reason string) {
	if os.Getenv("CLOUDFLARE_ACCOUNT_ID") == TestAccCloudflareAccountID {
		t.Skipf("Skipping acceptance test for default account (%s). %s", TestAccCloudflareAccountID, reason)
	}
}

// SharedV1Client returns a common Cloudflare V1 client setup needed for the
// sweeper functions.
func SharedV1Client() (*cfv1.API, error) {
	client, err := cfv1.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_EMAIL"))

	if err != nil {
		return client, err
	}

	return client, nil
}

// SharedClient returns a common Cloudflare V2 client setup needed for the
// sweeper functions.
func SharedClient() *cloudflare.Client {
	return cloudflare.NewClient(
		option.WithAPIKey(os.Getenv("CLOUDFLARE_API_KEY")),
		option.WithAPIEmail(os.Getenv("CLOUDFLARE_EMAIL")),
	)
}

// LoadTestCase takes a filename and variadic parameters to build test case output.
//
// Example: If you have a "basic" test case that for `r2_bucket` resource, inside
// of the `testdata` directory, you need to have a `basic.tf` file.
//
//	  $ tree internal/services/r2_bucket/
//	  ├── ...
//	  ├── schema.go
//	  └── testdata
//		  └── basic.tf
//
// The invocation would be `LoadTestCase("basic.tf", rnd, acountID)`.
func LoadTestCase(filename string, parameters ...interface{}) string {
	pwd, err := os.Getwd()
	if err != nil {
		return ""
	}

	fullPath := filepath.Join(pwd, "testdata", filename)
	f, err := os.ReadFile(fullPath)
	if err != nil {
		return ""
	}

	return fmt.Sprintf(string(f), parameters...)
}

// DumpState returns the state representation of the resource for inspection
// inside of a `resource.TestCase`.
//
// Example:
//
//	resource.Test(t, resource.TestCase{
//		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
//		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
//		Steps: []resource.TestStep{
//			{
//				Config: myConfig(),
//				Check: resource.ComposeTestCheckFunc(
//					acctest.DumpState,
//					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
//				),
//			},
//		},
//	})
//
// Returns the resource, identifier and attributes all keyed to the value.
//
//	cloudflare_foo.bar.id = 4dd549e6d9b22cdac402b33af89d57ab
//	cloudflare_foo.bar.other_thing = true
//	cloudflare_foo.bar.zone_id = 0da42c8d2132a9ddaf714f9e7c920711
//	cloudflare_foo.bar.% = 12
func DumpState(s *terraform.State) error {
	fmt.Println()
	for name, rs := range s.RootModule().Resources {
		// sort the keys
		keys := make([]string, 0, len(rs.Primary.Attributes))
		for attr := range rs.Primary.Attributes {
			keys = append(keys, attr)
		}
		sort.Strings(keys)
		for _, attr := range keys {
			key := rs.Primary.Attributes[attr]
			fmt.Println(strings.Join([]string{name, attr}, "."), "=", key)
		}
	}
	fmt.Println()

	return nil
}

type debugNonEmptyRefreshPlan struct{}

func (pc debugNonEmptyRefreshPlan) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	if os.Getenv("TF_LOG") == "DEBUG" {
		fmt.Println("\n---------\n\nRESOURCE DRIFT:")
		for _, d := range req.Plan.ResourceDrift {
			bytes, _ := json.MarshalIndent(d, "  ", "  ")
			fmt.Printf("%s\n\n", string(bytes))
		}
		fmt.Println("---------")
	}
}

var pc plancheck.PlanCheck = debugNonEmptyRefreshPlan{}
var LogResourceDrift = []plancheck.PlanCheck{pc}

type debugNonEmptyPlan struct{}

func (pc debugNonEmptyPlan) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	if os.Getenv("TF_LOG") == "DEBUG" {
		fmt.Println("\n---------\n\nRESOURCE Changes:")
		for _, d := range req.Plan.ResourceChanges {
			bytes, _ := json.MarshalIndent(d, "  ", "  ")
			fmt.Printf("%s\n\n", string(bytes))
		}
		fmt.Println("---------")
	}
}

var DebugNonEmptyPlan = debugNonEmptyPlan{}

// ExpectEmptyPlanExceptFalseyToNull is a plan check that expects an empty plan,
// except for changes from falsey values (false, empty string, empty arrays) to null.
// This is useful for migration tests where v4 had defaults that v5 doesn't have.
type expectEmptyPlanExceptFalseyToNull struct{}

func (e expectEmptyPlanExceptFalseyToNull) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	for _, rc := range req.Plan.ResourceChanges {
		if rc.Change.Actions[0] == "no-op" || rc.Change.Actions[0] == "read" {
			continue
		}

		// Check if this is an update action
		if rc.Change.Actions[0] != "update" {
			resp.Error = fmt.Errorf("expected empty plan, but %s has planned action(s): %v", rc.Address, rc.Change.Actions)
			return
		}

		// For updates, check each attribute change
		beforeMap, beforeOk := rc.Change.Before.(map[string]interface{})
		afterMap, afterOk := rc.Change.After.(map[string]interface{})

		if !beforeOk || !afterOk {
			resp.Error = fmt.Errorf("expected empty plan, but %s has non-map changes", rc.Address)
			return
		}

		// Check each attribute that's different
		for key, afterValue := range afterMap {
			beforeValue, _ := beforeMap[key]

			// Skip if values are the same
			if reflect.DeepEqual(beforeValue, afterValue) {
				continue
			}

			// Special handling for SetNestedAttribute fields (like include, exclude, require)
			if isSetNestedAttributeField(rc.Address, key) {
				if areSetNestedAttributesEquivalent(beforeValue, afterValue) {
					continue // Sets are equivalent despite ordering differences
				}
			}

			// Special handling for map attributes that may contain json.Number types or DynamicAttribute fields
			// This applies to all resources, not just dns_record, to handle Terraform state type variations
			if beforeMap, ok := beforeValue.(map[string]interface{}); ok {
				if afterMap, ok := afterValue.(map[string]interface{}); ok {
					if areMapsSemanticallySame(beforeMap, afterMap) {
						continue // Maps are semantically equivalent
					}
				}
			}

			// Allow changes from falsey to null
			if afterValue == nil {
				if isFalseyValue(beforeValue) {
					continue // This change is allowed
				}
			}

			// Allow session_duration changes from nil to a default value (API sets defaults)
			if key == "session_duration" && beforeValue == nil && afterValue != nil {
				continue // This change is allowed - API sets default session duration
			}

			// If we get here, it's a disallowed change
			resp.Error = fmt.Errorf("expected empty plan except for falsey-to-null changes, but %s.%s has change from %v to %v",
				rc.Address, key, beforeValue, afterValue)
			return
		}
	}
}

// isSetNestedAttributeField checks if a field name corresponds to a SetNestedAttribute
// for a specific resource type. These fields should be compared as sets (order-independent) rather than arrays
func isSetNestedAttributeField(resourceAddress, fieldName string) bool {
	// Extract resource type from address (e.g., "cloudflare_zero_trust_access_policy.example" -> "cloudflare_zero_trust_access_policy")
	resourceType := extractResourceTypeFromAddress(resourceAddress)

	// Map of resource types to their SetNestedAttribute field names
	// SetNestedAttribute fields are compared as sets (order-independent) rather than arrays
	// This prevents ambiguity when multiple resources have fields with the same name
	setFieldsByResource := map[string][]string{
		"cloudflare_zero_trust_access_policy": {"include", "exclude", "require", "approval_groups"},
		"cloudflare_access_policy":            {"include", "exclude", "require", "approval_group"}, // v4 resource name
		"cloudflare_ruleset":                  {"rules"},
		"cloudflare_load_balancer_pool":       {"origins"},
		"cloudflare_workers_script":           {"bindings"}, // v4->v5 migration transforms separate binding blocks to unified bindings list
		// Add other resources with SetNestedAttribute fields as needed
	}

	fields, exists := setFieldsByResource[resourceType]
	if !exists {
		return false
	}

	for _, setField := range fields {
		if fieldName == setField {
			return true
		}
	}
	return false
}

// extractResourceTypeFromAddress extracts the resource type from a resource address
// e.g., "cloudflare_zero_trust_access_policy.example" -> "cloudflare_zero_trust_access_policy"
func extractResourceTypeFromAddress(address string) string {
	parts := strings.Split(address, ".")
	if len(parts) < 2 {
		return address // fallback to full address if parsing fails
	}
	return parts[0]
}

// areSetNestedAttributesEquivalent compares two sets of nested attributes
// Returns true if they contain the same elements (ignoring order)
func areSetNestedAttributesEquivalent(before, after interface{}) bool {
	beforeSlice, beforeOk := before.([]interface{})
	afterSlice, afterOk := after.([]interface{})

	// If either is not a slice, fall back to regular comparison
	if !beforeOk || !afterOk {
		return reflect.DeepEqual(before, after)
	}

	// Different lengths means different sets
	if len(beforeSlice) != len(afterSlice) {
		return false
	}

	// For each element in before, find a matching element in after
	for _, beforeElement := range beforeSlice {
		found := false
		for _, afterElement := range afterSlice {
			if areNestedAttributeElementsEquivalent(beforeElement, afterElement) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

// areNestedAttributeElementsEquivalent compares two nested attribute elements
// with special handling for falsey-to-null transitions
func areNestedAttributeElementsEquivalent(before, after interface{}) bool {
	beforeMap, beforeOk := before.(map[string]interface{})
	afterMap, afterOk := after.(map[string]interface{})

	// If either is not a map, use regular comparison
	if !beforeOk || !afterOk {
		return reflect.DeepEqual(before, after)
	}

	// Get all unique keys from both maps
	allKeys := make(map[string]bool)
	for key := range beforeMap {
		allKeys[key] = true
	}
	for key := range afterMap {
		allKeys[key] = true
	}

	// Compare each key's value
	for key := range allKeys {
		beforeValue, beforeHasKey := beforeMap[key]
		afterValue, afterHasKey := afterMap[key]

		// If key missing from one side, treat as nil
		if !beforeHasKey {
			beforeValue = nil
		}
		if !afterHasKey {
			afterValue = nil
		}

		// Skip if values are identical
		if reflect.DeepEqual(beforeValue, afterValue) {
			continue
		}

		// Allow falsey-to-null transitions
		if afterValue == nil && isFalseyValue(beforeValue) {
			continue
		}
		if beforeValue == nil && isFalseyValue(afterValue) {
			continue
		}

		// Values are different and not a valid transition
		return false
	}

	return true
}

// isFalseyValue returns true if the value is considered "falsey"
// (false, empty string, empty slice/array, zero value, or map with all nil values)
func isFalseyValue(v interface{}) bool {
	if v == nil {
		return true
	}

	switch val := v.(type) {
	case bool:
		return !val
	case string:
		return val == ""
	case []interface{}:
		return len(val) == 0
	case map[string]interface{}:
		if len(val) == 0 {
			return true
		}
		// Check if all values in the map are nil or falsey
		for _, mapValue := range val {
			if !isFalseyValue(mapValue) {
				return false
			}
		}
		return true
	case int, int64, float64:
		// Handle numeric zero values
		return reflect.ValueOf(val).IsZero()
	default:
		// For other types, check if it's the zero value
		rv := reflect.ValueOf(val)
		if rv.IsValid() {
			return rv.IsZero()
		}
		return true
	}
}

// areMapsSemanticallySame compares two map values for semantic equality.
// This handles the case where reflect.DeepEqual fails on maps that are semantically identical
// but represented differently in memory (e.g., different object instances with same content).
func areMapsSemanticallySame(before, after interface{}) bool {
	beforeMap, beforeOk := before.(map[string]interface{})
	afterMap, afterOk := after.(map[string]interface{})

	if !beforeOk || !afterOk {
		// If either isn't a map, fall back to reflect.DeepEqual
		return reflect.DeepEqual(before, after)
	}

	// Check if maps have the same keys
	if len(beforeMap) != len(afterMap) {
		return false
	}

	// Check each key-value pair
	for key, beforeVal := range beforeMap {
		afterVal, exists := afterMap[key]
		if !exists {
			return false
		}

		// For nested maps (like DynamicAttribute structures), recurse
		if beforeMapNested, ok := beforeVal.(map[string]interface{}); ok {
			if afterMapNested, ok := afterVal.(map[string]interface{}); ok {
				if !areMapsSemanticallySame(beforeMapNested, afterMapNested) {
					return false
				}
				continue
			}
		}

		// For nil values, both must be nil
		if beforeVal == nil && afterVal == nil {
			continue
		}
		if beforeVal == nil || afterVal == nil {
			return false
		}

		// Handle json.Number type specially - convert to string for comparison
		// This handles cases where state has json.Number but plan has string or vice versa
		beforeStr, beforeIsJSONNumber := beforeVal.(json.Number)
		afterStr, afterIsJSONNumber := afterVal.(json.Number)

		if beforeIsJSONNumber || afterIsJSONNumber {
			// At least one is json.Number - compare string representations
			var beforeAsString, afterAsString string
			if beforeIsJSONNumber {
				beforeAsString = string(beforeStr)
			} else if str, ok := beforeVal.(string); ok {
				beforeAsString = str
			} else {
				beforeAsString = fmt.Sprintf("%v", beforeVal)
			}

			if afterIsJSONNumber {
				afterAsString = string(afterStr)
			} else if str, ok := afterVal.(string); ok {
				afterAsString = str
			} else {
				afterAsString = fmt.Sprintf("%v", afterVal)
			}

			if beforeAsString != afterAsString {
				return false
			}
			continue // Values are semantically equal
		}

		// Compare types first
		beforeType := reflect.TypeOf(beforeVal)
		afterType := reflect.TypeOf(afterVal)
		if beforeType != afterType {
			return false
		}

		// For other values, use reflect.DeepEqual
		if !reflect.DeepEqual(beforeVal, afterVal) {
			return false
		}
	}

	return true
}

var ExpectEmptyPlanExceptFalseyToNull = expectEmptyPlanExceptFalseyToNull{}

// isAllowedRuleSettingsChange checks if a rule_settings change is allowed
// for Gateway Policy resources. Allows nil-to-empty-collection changes for
// add_headers and override_ips fields which the API populates automatically.
func isAllowedRuleSettingsChange(before, after interface{}) bool {
	beforeMap, beforeOk := before.(map[string]interface{})
	afterMap, afterOk := after.(map[string]interface{})

	if !beforeOk || !afterOk {
		return false
	}

	// Get all keys from both maps
	allKeys := make(map[string]bool)
	for key := range beforeMap {
		allKeys[key] = true
	}
	for key := range afterMap {
		allKeys[key] = true
	}

	// Check each field
	for key := range allKeys {
		beforeVal, beforeExists := beforeMap[key]
		afterVal, afterExists := afterMap[key]

		// Skip if values are the same
		if reflect.DeepEqual(beforeVal, afterVal) {
			continue
		}

		// Allow nil fields in before to be removed in after (cleaned up nil fields)
		if beforeExists && !afterExists && beforeVal == nil {
			continue
		}

		// Allow fields being added if they were missing before (beforeExists == false)
		if !beforeExists && afterExists {
			// Allow adding add_headers as empty map
			if key == "add_headers" && isEmptyMap(afterVal) {
				continue
			}
			// Allow adding override_ips as empty slice
			if key == "override_ips" && isEmptySlice(afterVal) {
				continue
			}
		}

		// Allow nil -> map{} for add_headers
		if key == "add_headers" {
			if beforeVal == nil && isEmptyMap(afterVal) {
				continue
			}
		}

		// Allow nil -> [] for override_ips
		if key == "override_ips" {
			if (beforeVal == nil || !beforeExists) && isEmptySlice(afterVal) {
				continue
			}
		}

		// Allow fields removed from v5 schema to be removed or change
		removedFields := []string{"allow_child_bypass", "insecure_disable_dnssec_validation",
			"ignore_cname_category_matches", "resolve_dns_through_cloudflare", "block_page",
			"override_host", "ip_indicator_feeds"}
		isRemovedField := false
		for _, removedField := range removedFields {
			if key == removedField {
				isRemovedField = true
				break
			}
		}
		if isRemovedField {
			continue // Removed field changes are allowed
		}

		// If we get here and the field is different, this change is not allowed
		return false
	}

	return true
}

// isEmptyMap checks if a value is an empty map
func isEmptyMap(v interface{}) bool {
	m, ok := v.(map[string]interface{})
	return ok && len(m) == 0
}

// isEmptySlice checks if a value is an empty slice
func isEmptySlice(v interface{}) bool {
	s, ok := v.([]interface{})
	return ok && len(s) == 0
}

// ExpectEmptyPlanExceptGatewayPolicyAPIChanges is a plan check specifically for
// cloudflare_zero_trust_gateway_policy resources. It expects an empty plan except for:
// - Falsey-to-null changes (like the base checker)
// - Precedence changes (API auto-calculates with random offset)
// - rule_settings changes (API populates empty collections, removes deprecated fields)
type expectEmptyPlanExceptGatewayPolicyAPIChanges struct{}

func (e expectEmptyPlanExceptGatewayPolicyAPIChanges) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	for _, rc := range req.Plan.ResourceChanges {
		if rc.Change.Actions[0] == "no-op" || rc.Change.Actions[0] == "read" {
			continue
		}

		// Check if this is an update action
		if rc.Change.Actions[0] != "update" {
			resp.Error = fmt.Errorf("expected empty plan, but %s has planned action(s): %v", rc.Address, rc.Change.Actions)
			return
		}

		// For updates, check each attribute change
		beforeMap, beforeOk := rc.Change.Before.(map[string]interface{})
		afterMap, afterOk := rc.Change.After.(map[string]interface{})

		if !beforeOk || !afterOk {
			resp.Error = fmt.Errorf("expected empty plan, but %s has non-map changes", rc.Address)
			return
		}

		// Check each attribute that's different
		for key, afterValue := range afterMap {
			beforeValue, _ := beforeMap[key]

			// Skip if values are the same
			if reflect.DeepEqual(beforeValue, afterValue) {
				continue
			}

			// Special handling for SetNestedAttribute fields (like include, exclude, require)
			if isSetNestedAttributeField(rc.Address, key) {
				if areSetNestedAttributesEquivalent(beforeValue, afterValue) {
					continue // Sets are equivalent despite ordering differences
				}
			}

			// Allow changes from falsey to null
			if afterValue == nil {
				if isFalseyValue(beforeValue) {
					continue // This change is allowed
				}
			}

			// Allow session_duration changes from nil to a default value (API sets defaults)
			if key == "session_duration" && beforeValue == nil && afterValue != nil {
				continue // This change is allowed - API sets default session duration
			}

			// Gateway Policy specific: Allow precedence changes (API auto-calculates with random offset)
			if strings.Contains(rc.Address, "cloudflare_zero_trust_gateway_policy") && key == "precedence" {
				continue // This change is allowed - API modifies precedence values
			}

			// Gateway Policy specific: Allow Computed field changes (v5 provider schema issues)
			// These fields are marked as Computed in v5 schema and show as (known after apply) during refresh
			if strings.Contains(rc.Address, "cloudflare_zero_trust_gateway_policy") {
				computedFields := []string{"created_at", "updated_at", "version", "sharable",
					"deleted_at", "expiration", "read_only", "schedule", "source_account", "warning_status"}
				isComputedField := false
				for _, computedField := range computedFields {
					if key == computedField {
						isComputedField = true
						break
					}
				}
				if isComputedField {
					continue // This change is allowed - v5 provider Computed field
				}
			}

			// Gateway Policy specific: Allow rule_settings changes (API normalization)
			if strings.Contains(rc.Address, "cloudflare_zero_trust_gateway_policy") && key == "rule_settings" {
				if isAllowedRuleSettingsChange(beforeValue, afterValue) {
					continue // This change is allowed
				}
			}

			// If we get here, it's a disallowed change
			resp.Error = fmt.Errorf("expected empty plan except for Gateway Policy API changes, but %s.%s has change from %v to %v",
				rc.Address, key, beforeValue, afterValue)
			return
		}
	}
}

var ExpectEmptyPlanExceptGatewayPolicyAPIChanges = expectEmptyPlanExceptGatewayPolicyAPIChanges{}

// debugLogf logs a message only when TF_LOG=DEBUG is set
func debugLogf(t *testing.T, format string, args ...interface{}) {
	t.Helper()
	if strings.ToLower(os.Getenv("TF_LOG")) == "debug" {
		t.Logf(format, args...)
	}
}

// / PtrTo is a small helper to get a pointer to a particular value
func PtrTo[T any](v T) *T {
	return &v
}

// WriteOutConfig writes the config to tmpDir
func WriteOutConfig(t *testing.T, v4Config string, tmpDir string) {
	t.Helper()

	// Write the v4 config to tmpDir/test_migration.tf
	testConfigPath := filepath.Join(tmpDir, "test_migration.tf")
	debugLogf(t, "Writing v4 config to: %s", testConfigPath)

	err := os.WriteFile(testConfigPath, []byte(v4Config), 0644)
	if err != nil {
		t.Fatalf("Failed to write test config file: %v", err)
	}
	debugLogf(t, "Successfully wrote v4 config (%d bytes)", len(v4Config))

}

// RunMigrationV2Command runs the new tf-migrate binary to transform config and state
// NOTE: assumes config and state are already in tmpDir
func RunMigrationV2Command(t *testing.T, v4Config string, tmpDir string, sourceVersion string, targetVersion string) {
	t.Helper()

	// Get the migration binary path from environment variable
	migratorPath := os.Getenv("TF_MIGRATE_BINARY_PATH")
	if migratorPath == "" {
		// Fall back to default location relative to project root
		cwd, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get current working directory: %v", err)
		}
		projectRoot := filepath.Join(cwd, "..", "..", "..")
		migratorPath = filepath.Join(projectRoot, "tf-migrate", "tf-migrate")
	}

	// Check if the binary exists
	if _, err := os.Stat(migratorPath); os.IsNotExist(err) {
		t.Fatalf("tf-migrate binary not found at %s. Please set TF_MIGRATE_BINARY_PATH or ensure the binary is built.", migratorPath)
	}

	// Find state file in tmpDir
	// First check if state file exists directly in tmpDir (from v4 import)
	var stateFilePath string
	directStateFile := filepath.Join(tmpDir, "terraform.tfstate")
	if _, err := os.Stat(directStateFile); err == nil {
		stateFilePath = directStateFile
	} else {
		// Look for state file in subdirectories (from test framework)
		entries, err := os.ReadDir(tmpDir)
		if err != nil {
			t.Logf("Failed to read test directory: %v", err)
		} else {
			for _, entry := range entries {
				if entry.IsDir() {
					inner_entries, _ := os.ReadDir(filepath.Join(tmpDir, entry.Name()))
					for _, inner_entry := range inner_entries {
						if inner_entry.Name() == "terraform.tfstate" {
							stateFilePath = filepath.Join(tmpDir, entry.Name(), "terraform.tfstate")
							break
						}
					}
				}
				if stateFilePath != "" {
					break
				}
			}
		}
	}

	// Build the command
	args := []string{
		"migrate",
		"--config-dir", tmpDir,
		"--source-version", sourceVersion,
		"--target-version", targetVersion,
	}

	// Add state file argument if found
	if stateFilePath != "" {
		args = append(args, "--state-file", stateFilePath)
	}

	// Add debug logging if TF_LOG is set
	if strings.ToLower(os.Getenv("TF_LOG")) == "debug" {
		args = append(args, "--log-level", "debug")
	}

	// Run the migration command
	cmd := exec.Command(migratorPath, args...)
	cmd.Dir = tmpDir

	// Capture output for debugging
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("tf-migrate command failed: %v\nMigration output:\n%s", err, string(output))
	}
}

// RunMigrationCommand runs the migration script to transform config and state
// NOTE: assumes config and state are already in tmpDir
func RunMigrationCommand(t *testing.T, v4Config string, tmpDir string) {
	t.Helper()

	// Get the current working directory to find the migration binary
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}

	// Build the path to the migration binary
	// The test runs from internal/services/zone, so we need to go up to the root
	projectRoot := filepath.Join(cwd, "..", "..", "..")
	migratePath := filepath.Join(projectRoot, "cmd", "migrate")
	debugLogf(t, "Migrate path: %s", migratePath)

	// Path to the transformations directories
	transformerDir := filepath.Join(projectRoot, "cmd", "migrate", "transformations", "config")

	debugLogf(t, "Using YAML transformations from: %s", transformerDir)

	// Find state file in tmpDir
	// First check if state file exists directly in tmpDir (from v4 import)
	var stateFilePath string
	directStateFile := filepath.Join(tmpDir, "terraform.tfstate")
	if _, err := os.Stat(directStateFile); err == nil {
		stateFilePath = directStateFile
	} else {
		// Look for state file in subdirectories (from test framework)
		entries, err := os.ReadDir(tmpDir)
		if err != nil {
			t.Logf("Failed to read test directory: %v", err)
		} else {
			for _, entry := range entries {
				if entry.IsDir() {
					inner_entries, _ := os.ReadDir(filepath.Join(tmpDir, entry.Name()))
					for _, inner_entry := range inner_entries {
						if inner_entry.Name() == "terraform.tfstate" {
							stateFilePath = filepath.Join(tmpDir, entry.Name(), "terraform.tfstate")
							break
						}
					}
				}
				if stateFilePath != "" {
					break
				}
			}
		}
	}

	// Run the migration command on tmpDir (for config) and terraform.tfstate (for state)
	debugLogf(t, "State file path: %s", stateFilePath)
	state, err := os.ReadFile(stateFilePath)
	if err != nil {
		t.Fatalf("Failed to read state file: %v", err)
	}
	debugLogf(t, "Config is: %s", string(state))
	debugLogf(t, "State is: %s", string(state))

	var cmd *exec.Cmd
	// Use the new Go-based YAML transformations
	debugLogf(t, "Running migration with YAML transformations")
	cmd = exec.Command("go", "run", "-C", migratePath, ".",
		"-config", tmpDir,
		"-state", stateFilePath,
		"-grit=false",                      // Disable Grit transformations
		"-transformer=true",                // Enable YAML transformations
		"-transformer-dir", transformerDir) // Use local YAML configs
	cmd.Dir = tmpDir
	// Capture output for debugging
	output, err := cmd.CombinedOutput()

	debugLogf(t, "Migration output:\n%s", string(output))

	if err != nil {
		t.Fatalf("Migration command failed: %v\nMigration output:\n%s", err, string(output))
	}
	newState, err := os.ReadFile(stateFilePath)
	if err != nil {
		t.Fatalf("Failed to read state file: %v", err)
	}
	debugLogf(t, "New State is: %s", string(newState))
}

// MigrationTestStepWithPlan creates test steps for migrations that need plan processing after migration
// This handles resources that can't use state upgraders and need plan/refresh to correct state
// Returns multiple steps: migration step, plan step to process changes, then validation step
func MigrationTestStepWithPlan(t *testing.T, v4Config string, tmpDir string, exactVersion string, stateChecks []statecheck.StateCheck) []resource.TestStep {
	// First step: run migration
	migrationStep := MigrationTestStep(t, v4Config, tmpDir, exactVersion, nil) // No state checks yet

	// Second step: run plan to process import blocks and state corrections
	planStep := resource.TestStep{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		ConfigDirectory:          config.StaticDirectory(tmpDir),
		PlanOnly:                 true, // Just run plan to process imports/corrections
	}

	// Third step: verify final plan is clean and state is correct
	validationStep := resource.TestStep{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		ConfigDirectory:          config.StaticDirectory(tmpDir),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				DebugNonEmptyPlan,
				ExpectEmptyPlanExceptFalseyToNull, // Should be clean after processing
			},
		},
		ConfigStateChecks: stateChecks,
	}

	return []resource.TestStep{migrationStep, planStep, validationStep}
}

// MigrationV2TestStepWithPlan creates multiple test steps for v2 migration with plan processing
// This is similar to MigrationTestStepWithPlan but uses the v2 migration command with explicit version parameters
func MigrationV2TestStepWithPlan(t *testing.T, v4Config string, tmpDir string, exactVersion string, sourceVersion string, targetVersion string, stateChecks []statecheck.StateCheck) []resource.TestStep {
	// First step: run migration
	migrationStep := MigrationV2TestStep(t, v4Config, tmpDir, exactVersion, sourceVersion, targetVersion, nil) // No state checks yet

	// Second step: run plan to process import blocks and state corrections
	planStep := resource.TestStep{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		ConfigDirectory:          config.StaticDirectory(tmpDir),
		PlanOnly:                 true, // Just run plan to process imports/corrections
	}

	// Third step: verify final plan is clean and state is correct
	validationStep := resource.TestStep{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		ConfigDirectory:          config.StaticDirectory(tmpDir),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				DebugNonEmptyPlan,
				ExpectEmptyPlanExceptFalseyToNull, // Should be clean after processing
			},
		},
		ConfigStateChecks: stateChecks,
	}

	return []resource.TestStep{migrationStep, planStep, validationStep}
}

// MigrationTestStep creates a test step that runs the migration command and validates with v5 provider
func MigrationTestStep(t *testing.T, v4Config string, tmpDir string, exactVersion string, stateChecks []statecheck.StateCheck) resource.TestStep {
	// Choose the appropriate plan check based on the version
	var planChecks []plancheck.PlanCheck
	if strings.HasPrefix(exactVersion, "4.") {
		// When upgrading from v4, allow falsey-to-null changes due to removed defaults
		planChecks = []plancheck.PlanCheck{
			DebugNonEmptyPlan,
			ExpectEmptyPlanExceptFalseyToNull,
		}
	} else {
		// When upgrading from v5, expect a completely empty plan
		planChecks = []plancheck.PlanCheck{
			DebugNonEmptyPlan,
			plancheck.ExpectEmptyPlan(),
		}
	}

	return resource.TestStep{
		PreConfig: func() {
			WriteOutConfig(t, v4Config, tmpDir)
			// we only run the migration command if the version is 4.x.x, because users will not expect to run it within v5 versions.
			if strings.HasPrefix(exactVersion, "4.") {
				debugLogf(t, "Running migration command for version: %s", exactVersion)
				RunMigrationCommand(t, v4Config, tmpDir)
			} else {
				debugLogf(t, "Skipping migration command for version: %s", exactVersion)
			}
		},
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		ConfigDirectory:          config.StaticDirectory(tmpDir),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: planChecks,
		},
		ConfigStateChecks: stateChecks,
	}
}

// MigrationV2TestStep creates a test step that runs the migration command and validates with v5 provider
// Parameters:
//   - t: testing context
//   - v4Config: the configuration to migrate
//   - tmpDir: temporary directory for the test
//   - exactVersion: the exact version of the provider used to create the state (e.g., "4.52.1")
//   - sourceVersion: the source version for migration (e.g., "v4")
//   - targetVersion: the target version for migration (e.g., "v5")
//   - stateChecks: state validation checks to run after migration
func MigrationV2TestStep(t *testing.T, v4Config string, tmpDir string, exactVersion string, sourceVersion string, targetVersion string, stateChecks []statecheck.StateCheck) resource.TestStep {
	// Choose the appropriate plan check based on the source version
	var planChecks []plancheck.PlanCheck
	if sourceVersion == "v4" {
		// When upgrading from v4, allow falsey-to-null changes due to removed defaults
		planChecks = []plancheck.PlanCheck{
			DebugNonEmptyPlan,
			ExpectEmptyPlanExceptFalseyToNull,
		}
	} else {
		// When upgrading from other versions, expect a completely empty plan
		planChecks = []plancheck.PlanCheck{
			DebugNonEmptyPlan,
			plancheck.ExpectEmptyPlan(),
		}
	}

	return resource.TestStep{
		PreConfig: func() {
			WriteOutConfig(t, v4Config, tmpDir)
			debugLogf(t, "Running migration command for version: %s (%s -> %s)", exactVersion, sourceVersion, targetVersion)
			RunMigrationV2Command(t, v4Config, tmpDir, sourceVersion, targetVersion)
		},
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		ConfigDirectory:          config.StaticDirectory(tmpDir),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: planChecks,
		},
		ConfigStateChecks: stateChecks,
	}
}

// MigrationV2TestStepForGatewayPolicy creates a test step for cloudflare_zero_trust_gateway_policy migration
// that uses a custom plan checker to handle Gateway Policy API normalization behaviors:
// - Precedence changes (API auto-calculates with random offset 1-100)
// - rule_settings changes (API populates empty collections, removes deprecated fields)
//
// Parameters:
//   - expectNonEmptyPlan: Set to true for tests with rule_settings that have v5 provider schema issues
func MigrationV2TestStepForGatewayPolicy(t *testing.T, v4Config string, tmpDir string, exactVersion string, sourceVersion string, targetVersion string, expectNonEmptyPlan bool, stateChecks []statecheck.StateCheck) resource.TestStep {
	return resource.TestStep{
		PreConfig: func() {
			WriteOutConfig(t, v4Config, tmpDir)
			debugLogf(t, "Running migration command for Gateway Policy: %s (%s -> %s)", exactVersion, sourceVersion, targetVersion)
			RunMigrationV2Command(t, v4Config, tmpDir, sourceVersion, targetVersion)
		},
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		ConfigDirectory:          config.StaticDirectory(tmpDir),
		ExpectNonEmptyPlan:       expectNonEmptyPlan,
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				DebugNonEmptyPlan,
				ExpectEmptyPlanExceptGatewayPolicyAPIChanges,
			},
			// Note: PostApplyPostRefresh checks are intentionally omitted to allow
			// the ExpectNonEmptyPlan field to control refresh plan expectations.
		},
		ConfigStateChecks: stateChecks,
	}
}

// MigrationV2TestStepWithStateNormalization creates test steps for migrations where the v5 provider's
// schema causes state normalization issues. This is needed when:
// - The v5 provider returns all fields from the API (including nil/empty ones)
// - The migrated state has only populated fields
// - Terraform needs a plan cycle to normalize the state (remove nil fields)
//
// This helper expects an empty plan in the migration step, then runs a plan-only step
// to normalize the state, before validating with a clean plan check.
//
// Use this when `ExpectEmptyPlanExceptFalseyToNull` is too restrictive because the state
// changes involve removing entire nil fields rather than just falsey-to-null conversions.
//
// Example use case: cloudflare_zero_trust_access_group where selector fields are removed
// from state during normalization.
func MigrationV2TestStepWithStateNormalization(t *testing.T, v4Config string, tmpDir string, exactVersion string, sourceVersion string, targetVersion string, stateChecks []statecheck.StateCheck) []resource.TestStep {
	// Step 1: Run migration
	migrationStep := resource.TestStep{
		PreConfig: func() {
			WriteOutConfig(t, v4Config, tmpDir)
			debugLogf(t, "Running migration command for version: %s (%s -> %s)", exactVersion, sourceVersion, targetVersion)
			RunMigrationV2Command(t, v4Config, tmpDir, sourceVersion, targetVersion)
		},
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		ConfigDirectory:          config.StaticDirectory(tmpDir),
	}

	// Step 2: Run plan-only to normalize state (removes nil/empty fields)
	planStep := resource.TestStep{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		ConfigDirectory:          config.StaticDirectory(tmpDir),
		PlanOnly:                 true,
	}

	// Step 3: Verify final plan is clean and state is correct after normalization
	validationStep := resource.TestStep{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		ConfigDirectory:          config.StaticDirectory(tmpDir),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				DebugNonEmptyPlan,
				ExpectEmptyPlanExceptFalseyToNull, // Should be clean after normalization
			},
		},
		ConfigStateChecks: stateChecks,
	}

	return []resource.TestStep{migrationStep, planStep, validationStep}
}

// MigrationV2TestStepAllowCreate allows non-empty plans when a v4 resource needs to be split into multiple v5 resources,
//
// Example: cloudflare_argo with both smart_routing and tiered_caching becomes two separate resources
// Returns two steps: one for migration+apply, one for verification
func MigrationV2TestStepAllowCreate(t *testing.T, v4Config string, tmpDir string, exactVersion string, sourceVersion string, targetVersion string, stateChecks []statecheck.StateCheck) []resource.TestStep {
	return []resource.TestStep{
		{
			// Step 1: Run migration and apply any creates
			PreConfig: func() {
				WriteOutConfig(t, v4Config, tmpDir)
				debugLogf(t, "Running migration command for version: %s (%s -> %s)", exactVersion, sourceVersion, targetVersion)
				RunMigrationV2Command(t, v4Config, tmpDir, sourceVersion, targetVersion)
			},
			ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
			ConfigDirectory:          config.StaticDirectory(tmpDir),
		},
		{
			// Step 2: Verify final state and expect empty plan
			ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
			ConfigDirectory:          config.StaticDirectory(tmpDir),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					DebugNonEmptyPlan,
					ExpectEmptyPlanExceptFalseyToNull,
				},
			},
			ConfigStateChecks: stateChecks,
		},
	}
}

// ImportResourceWithV4Provider imports a resource using the v4 provider before migration testing.
// This is used for import-only resources that cannot be created via Terraform.
func ImportResourceWithV4Provider(t *testing.T, v4Config string, tmpDir string, providerVersion string, resourceName string, importID string) {
	t.Helper()

	// Write the config file
	configPath := filepath.Join(tmpDir, "test_migration.tf")
	if err := os.WriteFile(configPath, []byte(v4Config), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Run terraform init with the specific provider version
	initCmd := exec.Command("terraform", "init")
	initCmd.Dir = tmpDir
	initCmd.Env = append(os.Environ(),
		"TF_CLI_CONFIG_FILE=/dev/null", // Prevent local terraform config from interfering
	)

	initOutput, err := initCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("terraform init failed: %v\nOutput:\n%s", err, string(initOutput))
	}

	// Run terraform import
	importCmd := exec.Command("terraform", "import", resourceName, importID)
	importCmd.Dir = tmpDir
	importCmd.Env = append(os.Environ(),
		"TF_CLI_CONFIG_FILE=/dev/null",
	)

	importOutput, err := importCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("terraform import failed: %v\nOutput:\n%s", err, string(importOutput))
	}

	t.Logf("Successfully imported %s with ID %s using v4 provider", resourceName, importID)
}
