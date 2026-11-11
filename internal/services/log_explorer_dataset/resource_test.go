package log_explorer_dataset_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/logs"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	testresource "github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestMain(m *testing.M) {
	testresource.TestMain(m)
}

func init() {
	testresource.AddTestSweepers("cloudflare_log_explorer_dataset", &testresource.Sweeper{
		Name: "cloudflare_log_explorer_dataset",
		F:    testSweepCloudflareLogExplorerDataset,
	})
}

// testSweepCloudflareLogExplorerDataset removes leftover test datasets that a
// failed acceptance run could not destroy because deletion_protection was still
// enabled. Unlike name-prefixed resources, a dataset has no free-form identifier
// to filter test resources from real ones, so this only sweeps the fixed dataset
// types this suite creates (dns_logs for the test zone, magic_bgp_logs for the
// test account).
func testSweepCloudflareLogExplorerDataset(_ string) error {
	ctx := context.Background()
	client := testAccClient()

	if zoneID := os.Getenv("CLOUDFLARE_ZONE_ID"); zoneID != "" {
		if err := testSweepLogExplorerDatasetsForScope(ctx, client, zoneID, "", "dns_logs"); err != nil {
			return err
		}
	} else {
		tflog.Info(ctx, "Skipping log explorer dataset sweep for zone: CLOUDFLARE_ZONE_ID not set")
	}

	if accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID"); accountID != "" {
		if err := testSweepLogExplorerDatasetsForScope(ctx, client, "", accountID, "magic_bgp_logs"); err != nil {
			return err
		}
	} else {
		tflog.Info(ctx, "Skipping log explorer dataset sweep for account: CLOUDFLARE_ACCOUNT_ID not set")
	}

	return nil
}

func testSweepLogExplorerDatasetsForScope(
	ctx context.Context,
	client *cloudflare.Client,
	zoneID, accountID, datasetType string,
) error {
	params := logs.LogExplorerDatasetListParams{}
	if accountID != "" {
		params.AccountID = cloudflare.F(accountID)
	} else {
		params.ZoneID = cloudflare.F(zoneID)
	}

	page, err := client.Logs.LogExplorer.Datasets.List(ctx, params)
	if err != nil {
		return fmt.Errorf("failed to list log explorer datasets: %w", err)
	}

	for page != nil {
		for _, dataset := range page.Result {
			if dataset.Dataset != datasetType {
				continue
			}
			tflog.Info(ctx, fmt.Sprintf("Sweeping log explorer dataset: %s", dataset.DatasetID))
			if err := testSweepDeleteLogExplorerDataset(ctx, client, zoneID, accountID, dataset.DatasetID); err != nil {
				tflog.Error(ctx, fmt.Sprintf("failed to sweep log explorer dataset %s: %s", dataset.DatasetID, err))
			}
		}
		page, err = page.GetNextPage()
		if err != nil {
			return fmt.Errorf("failed to list log explorer datasets: %w", err)
		}
	}

	return nil
}

func testSweepDeleteLogExplorerDataset(
	ctx context.Context,
	client *cloudflare.Client,
	zoneID, accountID, datasetID string,
) error {
	updateParams := logs.LogExplorerDatasetUpdateParams{
		UpdateRequest: logs.UpdateRequestParam{
			Enabled:            cloudflare.F(true),
			DeletionProtection: cloudflare.F(false),
		},
	}
	deleteParams := logs.LogExplorerDatasetDeleteParams{}
	if accountID != "" {
		updateParams.AccountID = cloudflare.F(accountID)
		deleteParams.AccountID = cloudflare.F(accountID)
	} else {
		updateParams.ZoneID = cloudflare.F(zoneID)
		deleteParams.ZoneID = cloudflare.F(zoneID)
	}

	if _, err := client.Logs.LogExplorer.Datasets.Update(ctx, datasetID, updateParams); err != nil {
		return fmt.Errorf("failed to disable deletion protection: %w", err)
	}
	if _, err := client.Logs.LogExplorer.Datasets.Delete(ctx, datasetID, deleteParams); err != nil {
		return fmt.Errorf("failed to delete dataset: %w", err)
	}
	return nil
}

func testAccCheckLogExplorerDatasetDestroy(s *terraform.State) error {
	client := testAccClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_log_explorer_dataset" {
			continue
		}

		params := logs.LogExplorerDatasetGetParams{}
		if accountID := rs.Primary.Attributes["account_id"]; accountID != "" {
			params.AccountID = cloudflare.F(accountID)
		} else {
			params.ZoneID = cloudflare.F(rs.Primary.Attributes["zone_id"])
		}

		_, err := client.Logs.LogExplorer.Datasets.Get(context.Background(), rs.Primary.ID, params)
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			if apierr.StatusCode != 404 {
				return fmt.Errorf("log explorer dataset %s still exists", rs.Primary.ID)
			}
			continue
		}
		if err != nil {
			return err
		}
		return fmt.Errorf("log explorer dataset %s still exists", rs.Primary.ID)
	}

	return nil
}

func TestAccCloudflareLogExplorerDataset_Lifecycle(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("set TF_ACC=1 to run acceptance tests")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	const resourceName = "cloudflare_log_explorer_dataset.test"
	const dataSourceName = "data.cloudflare_log_explorer_datasets.test"
	const partialFieldName = "queryname"
	const configuredFilter = `{"where":{"and":[{"key":"QueryName","operator":"eq","value":"example.com"}]}}`
	var datasetID string
	var defaultEnabledFields []string
	testAccCleanupDataset(t, zoneID, &datasetID)

	testresource.Test(t, testresource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogExplorerDatasetDestroy,
		Steps: []testresource.TestStep{
			{
				Config: testAccLogExplorerDatasetConfig(t, "basic.tf", zoneID, true,
					testAccLogExplorerDatasetFieldsConfig(false, "QueryName", "timestamp")),
				Check: testresource.ComposeTestCheckFunc(
					testAccCaptureDatasetID(resourceName, &datasetID),
					testresource.TestCheckResourceAttr(resourceName, "enabled", "false"),
					testAccCheckFieldEnabled(resourceName, "QueryName", true),
					testAccCheckAPIEnabledFields(zoneID, &datasetID, "queryname", "timestamp"),
				),
			},
			{
				Config: testAccLogExplorerDatasetConfig(t, "basic.tf", zoneID, true,
					testAccLogExplorerDatasetFieldsConfig(false, "QueryName", "timestamp")),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				Config: testAccLogExplorerDatasetConfig(t, "basic.tf", zoneID, true, ""),
				Check: testresource.ComposeTestCheckFunc(
					testAccCaptureDatasetID(resourceName, &datasetID),
					testAccCaptureAPIEnabledFields(zoneID, &datasetID, &defaultEnabledFields),
					testresource.TestCheckResourceAttr(resourceName, "zone_id", zoneID),
					testresource.TestCheckResourceAttr(resourceName, "dataset", "dns_logs"),
					testresource.TestCheckResourceAttr(resourceName, "object_type", "zone"),
					testresource.TestCheckResourceAttr(resourceName, "deletion_protection", "true"),
					testresource.TestCheckResourceAttr(resourceName, "filter", ""),
					testresource.TestCheckNoResourceAttr(resourceName, "fields"),
				),
			},
			{
				Config:             testAccLogExplorerDatasetConfig(t, "basic.tf", "replacement-zone-id", true, ""),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				ConfigPlanChecks: testresource.ConfigPlanChecks{
					PostApplyPreRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
			{
				Config: fmt.Sprintf(`resource "cloudflare_log_explorer_dataset" "test" {
  zone_id             = %q
  dataset             = "http_requests"
  deletion_protection = true
}`, zoneID),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				ConfigPlanChecks: testresource.ConfigPlanChecks{
					PostApplyPreRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
			{
				Config: fmt.Sprintf(`resource "cloudflare_log_explorer_dataset" "test" {
  account_id          = %q
  dataset             = "dns_logs"
  deletion_protection = true
}`, "replacement-account-id"),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				ConfigPlanChecks: testresource.ConfigPlanChecks{
					PostApplyPreRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
			{
				Config:      testAccLogExplorerDatasetConfig(t, "basic.tf", zoneID, true, ""),
				Destroy:     true,
				ExpectError: regexp.MustCompile(`(?s)deletion_protection\s*=\s*false.*terraform\s+apply`),
			},
			{
				PreConfig: func() {
					dataset, err := testAccGetDataset(zoneID, &datasetID)
					if err != nil {
						t.Fatalf("Failed to read the dataset after protected destruction: %v", err)
					}
					if !dataset.DeletionProtection {
						t.Fatal("Deletion protection was disabled during protected destruction")
					}
				},
				Config:             testAccLogExplorerDatasetConfig(t, "basic.tf", zoneID, true, ""),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				Config: testAccLogExplorerDatasetConfig(t, "basic.tf", zoneID, false, ""),
				Check:  testresource.TestCheckResourceAttr(resourceName, "deletion_protection", "false"),
			},
			{
				Config: testAccLogExplorerDatasetConfig(
					t,
					"basic_with_filter.tf",
					zoneID,
					false,
					testAccLogExplorerDatasetFieldsConfig(false, partialFieldName, "timestamp"),
				),
				Check: testresource.ComposeTestCheckFunc(
					testresource.TestCheckResourceAttr(resourceName, "enabled", "false"),
					testAccCheckFieldCount(resourceName, 2),
					testAccCheckFieldEnabled(resourceName, partialFieldName, true),
					testAccCheckFieldEnabled(resourceName, "timestamp", true),
					testAccCheckAPIEnabledFields(zoneID, &datasetID, partialFieldName, "timestamp"),
					testresource.TestCheckResourceAttr(resourceName, "filter", configuredFilter),
					testAccCheckAPIFilter(zoneID, &datasetID, configuredFilter),
				),
			},
			{
				Config: testAccLogExplorerDatasetConfig(
					t,
					"basic_with_filter.tf",
					zoneID,
					false,
					testAccLogExplorerDatasetFieldsConfig(false, "QueryName", "timestamp"),
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				Config: testAccLogExplorerDatasetConfig(
					t,
					"basic.tf",
					zoneID,
					false,
					testAccLogExplorerDatasetFieldsConfig(true, partialFieldName, "timestamp"),
				),
				Check: testresource.ComposeTestCheckFunc(
					testresource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					testAccCheckFieldCount(resourceName, 2),
					testAccCheckFieldEnabled(resourceName, partialFieldName, true),
					testAccCheckFieldEnabled(resourceName, "timestamp", true),
					testresource.TestCheckResourceAttr(resourceName, "filter", ""),
					testAccCheckAPIFilter(zoneID, &datasetID, ""),
				),
			},
			{
				Config:             testAccLogExplorerDatasetConfig(t, "basic.tf", zoneID, false, testAccLogExplorerDatasetFieldsConfig(true, partialFieldName, "timestamp")),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				Config: testAccLogExplorerDatasetConfig(
					t,
					"basic.tf",
					zoneID,
					false,
					testAccLogExplorerDatasetFieldsConfig(true, partialFieldName),
				),
				ExpectError: regexp.MustCompile(`(?is)the\s+following\s+fields\s+must\s+be\s+enabled\s+for\s+this\s+dataset`),
			},
			{
				Config: testAccLogExplorerDatasetConfig(t, "basic.tf", zoneID, false, ""),
				Check: testresource.ComposeTestCheckFunc(
					testresource.TestCheckResourceAttr(resourceName, "filter", ""),
					testresource.TestCheckNoResourceAttr(resourceName, "fields"),
					testAccCheckAPIEnabledFieldsMatchCaptured(zoneID, &datasetID, &defaultEnabledFields),
				),
			},
			{
				Config:             testAccLogExplorerDatasetConfig(t, "basic.tf", zoneID, false, ""),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       testAccLogExplorerDatasetImportID("zones", zoneID),
				ImportStateVerifyIgnore: []string{"deletion_protection"},
			},
			{
				Config: testAccLogExplorerDatasetConfig(t, "basic.tf", zoneID, false, "") + fmt.Sprintf(`
data "cloudflare_log_explorer_datasets" "test" {
  zone_id   = %q
  max_items = 1000
}
`, zoneID),
				Check: testAccListContainsDataset(resourceName, dataSourceName),
			},
			{
				PreConfig:          testAccDeleteDataset(t, zoneID, &datasetID),
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
				RefreshPlanChecks: testresource.RefreshPlanChecks{
					PostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
				},
			},
		},
	})
}

func TestAccCloudflareLogExplorerDataset_ScopeValidation(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("set TF_ACC=1 to run acceptance tests")
	}

	testresource.Test(t, testresource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []testresource.TestStep{
			{
				Config: `resource "cloudflare_log_explorer_dataset" "test" {
  dataset = "dns_logs"
}`,
				ExpectError: regexp.MustCompile(`(?i)exactly one.*configured`),
			},
			{
				Config: `resource "cloudflare_log_explorer_dataset" "test" {
  dataset    = "dns_logs"
  account_id = "not-used"
  zone_id    = "not-used"
}`,
				ExpectError: regexp.MustCompile(`(?i)exactly one.*configured`),
			},
			{
				Config: `resource "cloudflare_log_explorer_dataset" "test" {
  zone_id = "not-used"
}`,
				ExpectError: regexp.MustCompile(`(?i)missing required argument|required attribute`),
			},
			{
				Config: `resource "cloudflare_log_explorer_dataset" "test" {
  zone_id = "not-used"
  dataset = "dns_logs"
  fields = [{
    name = "timestamp"
  }]
}`,
				ExpectError: regexp.MustCompile(`(?is)enabled.*required`),
			},
		},
	})
}

func TestAccCloudflareLogExplorerDataset_NoDefaultFilter(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("set TF_ACC=1 to run acceptance tests")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	const resourceName = "cloudflare_log_explorer_dataset.test"
	var datasetID string
	testAccCleanupDataset(t, zoneID, &datasetID)

	testresource.Test(t, testresource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogExplorerDatasetDestroy,
		Steps: []testresource.TestStep{
			{
				Config: fmt.Sprintf(`resource "cloudflare_log_explorer_dataset" "test" {
  zone_id             = %q
  dataset             = "http_requests"
  deletion_protection = false
}`, zoneID),
				Check: testresource.ComposeTestCheckFunc(
					testAccCaptureDatasetID(resourceName, &datasetID),
					testresource.TestCheckResourceAttr(resourceName, "filter", ""),
					testAccCheckAPIFilter(zoneID, &datasetID, ""),
				),
			},
			{
				Config: fmt.Sprintf(`resource "cloudflare_log_explorer_dataset" "test" {
  zone_id             = %q
  dataset             = "http_requests"
  deletion_protection = false
  filter              = ""
}`, zoneID),
				Check: testresource.ComposeTestCheckFunc(
					testresource.TestCheckResourceAttr(resourceName, "filter", ""),
					testAccCheckAPIFilter(zoneID, &datasetID, ""),
				),
			},
			{
				Config: fmt.Sprintf(`resource "cloudflare_log_explorer_dataset" "test" {
  zone_id               = %q
  dataset               = "http_requests"
  deletion_protection   = false
  filter_attack_traffic = false
}`, zoneID),
				Check: testresource.ComposeTestCheckFunc(
					testresource.TestCheckResourceAttr(resourceName, "filter_attack_traffic", "false"),
					testresource.TestCheckResourceAttr(resourceName, "filter", ""),
					testAccCheckAPIFilterAttackTraffic(zoneID, &datasetID, false),
				),
			},
			{
				Config: fmt.Sprintf(`resource "cloudflare_log_explorer_dataset" "test" {
  zone_id               = %q
  dataset               = "http_requests"
  deletion_protection   = false
  filter_attack_traffic = true
}`, zoneID),
				Check: testresource.ComposeTestCheckFunc(
					testresource.TestCheckResourceAttr(resourceName, "filter_attack_traffic", "true"),
					testresource.TestCheckResourceAttr(resourceName, "filter", ""),
					testAccCheckAPIFilterAttackTraffic(zoneID, &datasetID, true),
				),
			},
		},
	})
}

func TestAccCloudflareLogExplorerDataset_AccountLifecycle(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("set TF_ACC=1 to run acceptance tests")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	const resourceName = "cloudflare_log_explorer_dataset.test"

	testresource.Test(t, testresource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogExplorerDatasetDestroy,
		Steps: []testresource.TestStep{
			{
				Config: testAccLogExplorerDatasetConfig(t, "account.tf", accountID, true, ""),
				Check: testresource.ComposeTestCheckFunc(
					testresource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					testresource.TestCheckResourceAttr(resourceName, "dataset", "magic_bgp_logs"),
					testresource.TestCheckResourceAttr(resourceName, "object_type", "account"),
				),
			},
			{
				Config: testAccLogExplorerDatasetConfig(t, "account.tf", accountID, false, ""),
				Check:  testresource.TestCheckResourceAttr(resourceName, "deletion_protection", "false"),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       testAccLogExplorerDatasetImportID("accounts", accountID),
				ImportStateVerifyIgnore: []string{"deletion_protection"},
			},
		},
	})
}

func testAccLogExplorerDatasetConfig(
	t *testing.T,
	fixture, scopeID string,
	deletionProtection bool,
	extraConfig string,
) string {
	t.Helper()

	config, err := os.ReadFile("testdata/" + fixture)
	if err != nil {
		t.Fatal(err)
	}
	return fmt.Sprintf(string(config), scopeID, deletionProtection, extraConfig)
}

// testAccLogExplorerDatasetFieldsConfig builds a `fields` block enabling exactly
// fieldNames. dns_logs requires its own timestamp field ("timestamp") to remain
// enabled in any explicit selection, so callers must include it themselves; the
// provider does not add it automatically, since the required field name varies
// per dataset type and is not something the provider can know.
func testAccLogExplorerDatasetFieldsConfig(enabled bool, fieldNames ...string) string {
	var fields strings.Builder
	for _, fieldName := range fieldNames {
		fmt.Fprintf(&fields, "    { name = %q, enabled = true },\n", fieldName)
	}
	return fmt.Sprintf("  enabled = %t\n  fields = [\n%s  ]\n", enabled, fields.String())
}

// testAccCleanupDataset removes a dataset the test framework could not destroy,
// which happens when a step fails while deletion protection is still enabled.
func testAccCleanupDataset(t *testing.T, zoneID string, datasetID *string) {
	t.Helper()

	t.Cleanup(func() {
		if *datasetID == "" {
			return
		}
		ctx := context.Background()
		client := testAccClient()
		if _, err := client.Logs.LogExplorer.Datasets.Get(
			ctx,
			*datasetID,
			logs.LogExplorerDatasetGetParams{ZoneID: cloudflare.F(zoneID)},
		); err != nil {
			return
		}
		if _, err := client.Logs.LogExplorer.Datasets.Update(
			ctx,
			*datasetID,
			logs.LogExplorerDatasetUpdateParams{
				ZoneID: cloudflare.F(zoneID),
				UpdateRequest: logs.UpdateRequestParam{
					Enabled:            cloudflare.F(true),
					DeletionProtection: cloudflare.F(false),
				},
			},
		); err != nil {
			t.Logf("failed to disable deletion protection for %s: %v", *datasetID, err)
			return
		}
		if _, err := client.Logs.LogExplorer.Datasets.Delete(
			ctx,
			*datasetID,
			logs.LogExplorerDatasetDeleteParams{ZoneID: cloudflare.F(zoneID)},
		); err != nil {
			t.Logf("failed to delete leftover dataset %s: %v", *datasetID, err)
		}
	})
}

// testAccCheckAPIEnabledFields asserts the API enables exactly the named fields, so
// a partial Terraform configuration disables everything it does not list.
func testAccCheckAPIEnabledFields(
	zoneID string,
	datasetID *string,
	enabledNames ...string,
) testresource.TestCheckFunc {
	return func(_ *terraform.State) error {
		dataset, err := testAccGetDataset(zoneID, datasetID)
		if err != nil {
			return err
		}
		wantEnabled := make(map[string]struct{}, len(enabledNames))
		for _, name := range enabledNames {
			wantEnabled[name] = struct{}{}
		}
		for _, field := range dataset.Fields {
			_, want := wantEnabled[field.Name]
			if field.Enabled != want {
				return fmt.Errorf(
					"API field %s enabled is %t, want %t",
					field.Name,
					field.Enabled,
					want,
				)
			}
			delete(wantEnabled, field.Name)
		}
		for name := range wantEnabled {
			return fmt.Errorf("API did not return expected field %s", name)
		}
		return nil
	}
}

// testAccCaptureAPIEnabledFields records the API's default field selection before
// Terraform starts managing fields.
func testAccCaptureAPIEnabledFields(
	zoneID string,
	datasetID *string,
	destination *[]string,
) testresource.TestCheckFunc {
	return func(_ *terraform.State) error {
		dataset, err := testAccGetDataset(zoneID, datasetID)
		if err != nil {
			return err
		}
		*destination = testAccAPIEnabledFieldNames(dataset)
		return nil
	}
}

// testAccCheckAPIEnabledFieldsMatchCaptured proves that removing fields from the
// Terraform configuration restores the API's original default field selection.
func testAccCheckAPIEnabledFieldsMatchCaptured(
	zoneID string,
	datasetID *string,
	want *[]string,
) testresource.TestCheckFunc {
	return func(_ *terraform.State) error {
		if *want == nil {
			return fmt.Errorf("API default enabled fields were not captured")
		}
		dataset, err := testAccGetDataset(zoneID, datasetID)
		if err != nil {
			return err
		}
		got := testAccAPIEnabledFieldNames(dataset)
		if !slices.Equal(got, *want) {
			return fmt.Errorf("API enabled fields are %v, want restored defaults %v", got, *want)
		}
		return nil
	}
}

func testAccAPIEnabledFieldNames(dataset *logs.Dataset) []string {
	enabled := make([]string, 0, len(dataset.Fields))
	for _, field := range dataset.Fields {
		if field.Enabled {
			enabled = append(enabled, field.Name)
		}
	}
	slices.Sort(enabled)
	return enabled
}

// testAccCheckAPIFilter verifies the filter stored by the API independently of
// the normalized Terraform state.
func testAccCheckAPIFilter(
	zoneID string,
	datasetID *string,
	want string,
) testresource.TestCheckFunc {
	return func(_ *terraform.State) error {
		dataset, err := testAccGetDataset(zoneID, datasetID)
		if err != nil {
			return err
		}
		if dataset.Filter != want {
			return fmt.Errorf("API filter is %q, want %q", dataset.Filter, want)
		}
		return nil
	}
}

func testAccCheckAPIFilterAttackTraffic(zoneID string, datasetID *string, want bool) testresource.TestCheckFunc {
	return func(_ *terraform.State) error {
		dataset, err := testAccGetDataset(zoneID, datasetID)
		if err != nil {
			return err
		}
		if dataset.FilterAttackTraffic != want {
			return fmt.Errorf("API filter_attack_traffic is %t, want %t", dataset.FilterAttackTraffic, want)
		}
		return nil
	}
}

func testAccGetDataset(zoneID string, datasetID *string) (*logs.Dataset, error) {
	if *datasetID == "" {
		return nil, fmt.Errorf("dataset_id was not captured")
	}
	return testAccClient().Logs.LogExplorer.Datasets.Get(
		context.Background(),
		*datasetID,
		logs.LogExplorerDatasetGetParams{ZoneID: cloudflare.F(zoneID)},
	)
}

func testAccLogExplorerDatasetImportID(scope, scopeID string) testresource.ImportStateIdFunc {
	return func(state *terraform.State) (string, error) {
		resourceState, ok := state.RootModule().Resources["cloudflare_log_explorer_dataset.test"]
		if !ok {
			return "", fmt.Errorf("resource state not found")
		}
		return fmt.Sprintf("%s/%s/%s", scope, scopeID, resourceState.Primary.Attributes["dataset_id"]), nil
	}
}

func testAccDeleteDataset(t *testing.T, zoneID string, datasetID *string) func() {
	return func() {
		_, err := testAccClient().Logs.LogExplorer.Datasets.Delete(
			context.Background(),
			*datasetID,
			logs.LogExplorerDatasetDeleteParams{ZoneID: cloudflare.F(zoneID)},
		)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testAccCaptureDatasetID(resourceName string, datasetID *string) testresource.TestCheckFunc {
	return func(state *terraform.State) error {
		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource state not found: %s", resourceName)
		}
		*datasetID = resourceState.Primary.Attributes["dataset_id"]
		if *datasetID == "" {
			return fmt.Errorf("dataset_id is empty")
		}
		return nil
	}
}

func testAccListContainsDataset(resourceName, dataSourceName string) testresource.TestCheckFunc {
	return func(state *terraform.State) error {
		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource state not found: %s", resourceName)
		}
		datasetID := resourceState.Primary.Attributes["dataset_id"]
		dataSourceState, ok := state.RootModule().Resources[dataSourceName]
		if !ok {
			return fmt.Errorf("data source state not found: %s", dataSourceName)
		}
		count, err := strconv.Atoi(dataSourceState.Primary.Attributes["result.#"])
		if err != nil {
			return err
		}
		for index := 0; index < count; index++ {
			if dataSourceState.Primary.Attributes[fmt.Sprintf("result.%d.dataset_id", index)] == datasetID {
				return nil
			}
		}
		return fmt.Errorf("dataset %s not found in list data source", datasetID)
	}
}

func testAccCheckFieldEnabled(resourceName, fieldName string, enabled bool) testresource.TestCheckFunc {
	return func(state *terraform.State) error {
		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource state not found: %s", resourceName)
		}
		fieldCount, err := strconv.Atoi(resourceState.Primary.Attributes["fields.#"])
		if err != nil {
			return err
		}
		for index := range fieldCount {
			if resourceState.Primary.Attributes[fmt.Sprintf("fields.%d.name", index)] != fieldName {
				continue
			}
			actual := resourceState.Primary.Attributes[fmt.Sprintf("fields.%d.enabled", index)]
			if actual == strconv.FormatBool(enabled) {
				return nil
			}
			return fmt.Errorf("field %s enabled state is %s, want %t", fieldName, actual, enabled)
		}
		return fmt.Errorf("field %s not found in resource state", fieldName)
	}
}

func testAccCheckFieldCount(resourceName string, expected int) testresource.TestCheckFunc {
	return func(state *terraform.State) error {
		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource state not found: %s", resourceName)
		}
		count, err := strconv.Atoi(resourceState.Primary.Attributes["fields.#"])
		if err != nil {
			return err
		}
		if count != expected {
			return fmt.Errorf("got %d fields, want %d", count, expected)
		}
		return nil
	}
}

func testAccClient() *cloudflare.Client {
	opts := []option.RequestOption{option.WithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))}
	if baseURL := os.Getenv("CLOUDFLARE_BASE_URL"); baseURL != "" {
		opts = append(opts, option.WithBaseURL(baseURL))
	}
	return cloudflare.NewClient(opts...)
}
