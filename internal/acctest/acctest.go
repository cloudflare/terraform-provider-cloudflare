package acctest

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/option"
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
		t.Fatal("valid credentials are required for this acceptance test.")
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

// Test helper method checking the Workspace One environment variables are present.
func TestAccPreCheck_WorkspaceOne(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_WORKSPACE_ONE_CLIENT_ID"); v == "" {
		t.Fatal("CLOUDFLARE_WORKSPACE_ONE_CLIENT_ID must be set for this acceptance test.")
	}

	if v := os.Getenv("CLOUDFLARE_WORKSPACE_ONE_CLIENT_SECRET"); v == "" {
		t.Fatal("CLOUDFLARE_WORKSPACE_ONE_CLIENT_SECRET must be set for this acceptance test.")
	}

	if v := os.Getenv("CLOUDFLARE_WORKSPACE_ONE_API_URL"); v == "" {
		t.Fatal("CLOUDFLARE_WORKSPACE_ONE_API_URL must be set for this acceptance test.")
	}

	if v := os.Getenv("CLOUDFLARE_WORKSPACE_ONE_AUTH_URL"); v == "" {
		t.Fatal("CLOUDFLARE_WORKSPACE_ONE_AUTH_URL must be set for this acceptance test.")
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
		for attr, key := range rs.Primary.Attributes {
			fmt.Println(strings.Join([]string{name, attr}, "."), "=", key)
		}
	}
	fmt.Println()

	return nil
}

type debugNonEmptyRefreshPlan struct{}

func (pc debugNonEmptyRefreshPlan) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	fmt.Println("\n---------\n\nRESOURCE DRIFT:")
	for _, d := range req.Plan.ResourceDrift {
		bytes, _ := json.MarshalIndent(d, "  ", "  ")
		fmt.Printf("%s\n\n", string(bytes))
	}
	fmt.Println("---------")
}

var pc plancheck.PlanCheck = debugNonEmptyRefreshPlan{}
var LogResourceDrift = []plancheck.PlanCheck{pc}

type debugNonEmptyPlan struct{}

func (pc debugNonEmptyPlan) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	fmt.Println("\n---------\n\nRESOURCE Changes:")
	for _, d := range req.Plan.ResourceChanges {
		bytes, _ := json.MarshalIndent(d, "  ", "  ")
		fmt.Printf("%s\n\n", string(bytes))
	}
	fmt.Println("---------")
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
			if isSetNestedAttributeField(key) {
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

			// If we get here, it's a disallowed change
			resp.Error = fmt.Errorf("expected empty plan except for falsey-to-null changes, but %s.%s has change from %v to %v",
				rc.Address, key, beforeValue, afterValue)
			return
		}
	}
}

// isSetNestedAttributeField checks if a field name corresponds to a SetNestedAttribute
// These fields should be compared as sets (order-independent) rather than arrays
func isSetNestedAttributeField(fieldName string) bool {
	setFields := []string{
		"include", "exclude", "require", // zero_trust_access_policy
		"approval_groups", // zero_trust_access_policy (after migration)
		// Add other SetNestedAttribute fields as needed
	}

	for _, setField := range setFields {
		if fieldName == setField {
			return true
		}
	}
	return false
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

var ExpectEmptyPlanExceptFalseyToNull = expectEmptyPlanExceptFalseyToNull{}

// / PtrTo is a small helper to get a pointer to a particular value
func PtrTo[T any](v T) *T {
	return &v
}

// WriteOutConfig writes the config to tmpDir
func WriteOutConfig(t *testing.T, v4Config string, tmpDir string) {
	t.Helper()

	// Write the v4 config to tmpDir/test_migration.tf
	testConfigPath := filepath.Join(tmpDir, "test_migration.tf")
	t.Logf("Writing v4 config to: %s", testConfigPath)

	err := os.WriteFile(testConfigPath, []byte(v4Config), 0644)
	if err != nil {
		t.Fatalf("Failed to write test config file: %v", err)
	}
	t.Logf("Successfully wrote v4 config (%d bytes)", len(v4Config))

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
	t.Logf("Migrate path: %s", migratePath)

	// specify patternsDir so we use local patterns instead of the ones from Github
	patternsDir := filepath.Join(projectRoot, ".grit")

	// Find state file in tmpDir
	entries, err := os.ReadDir(tmpDir)
	var stateDir string
	if err != nil {
		t.Logf("Failed to read test directory: %v", err)
	} else {
		for _, entry := range entries {
			if entry.IsDir() {
				inner_entries, _ := os.ReadDir(filepath.Join(tmpDir, entry.Name()))
				for _, inner_entry := range inner_entries {
					if inner_entry.Name() == "terraform.tfstate" {
						stateDir = filepath.Join(tmpDir, entry.Name())
					}
				}
			}

		}
	}

	// Run the migration command on tmpDir (for config) and terraform.tfstate (for state)
	t.Logf("StateDir: %s", stateDir)
	state, err := os.ReadFile(filepath.Join(stateDir, "terraform.tfstate"))
	if err != nil {
		t.Fatalf("Failed to read state file: %v", err)
	}
	t.Logf("State is: %s", string(state))
	cmd := exec.Command("go", "run", "-C", migratePath, ".", "-config", tmpDir, "-patterns-dir", patternsDir, "-state", filepath.Join(stateDir, "terraform.tfstate"))
	cmd.Dir = tmpDir
	// Capture output for debugging
	output, err := cmd.CombinedOutput()

	t.Logf("Migration output:\n%s", string(output))

	if err != nil {
		t.Fatalf("Migration command failed: %v", err)
	}
	newState, err := os.ReadFile(filepath.Join(stateDir, "terraform.tfstate"))
	if err != nil {
		t.Fatalf("Failed to read state file: %v", err)
	}
	t.Logf("New State is: %s", string(newState))
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
				t.Logf("Running migration command for version: %s", exactVersion)
				RunMigrationCommand(t, v4Config, tmpDir)
			} else {
				t.Logf("Skipping migration command for version: %s", exactVersion)
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
