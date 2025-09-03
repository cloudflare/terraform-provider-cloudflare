package zero_trust_dlp_entry_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

const (
	EnvTfAcc = "TF_ACC"
)

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_dlp_entry", &resource.Sweeper{
		Name: "cloudflare_zero_trust_dlp_entry",
		F:    testSweepCloudflareZeroTrustDLPEntry,
	})
}

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// setupDLPTest handles common test setup and TF_ACC environment check
func setupDLPTest(t *testing.T, testName string, allowedMatchCount int64) (string, string, func()) {
	if os.Getenv(EnvTfAcc) == "" {
		t.Skip(fmt.Sprintf(
			"Acceptance tests skipped unless env '%s' set",
			EnvTfAcc))
		return "", "", func() {}
	}

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	// Create a DLP profile for testing
	profileID := createTestDLPProfile(t, rnd, testName, allowedMatchCount)
	
	// Return cleanup function
	cleanup := func() {
		deleteTestDLPProfile(t, profileID, accountID)
	}
	
	return rnd, profileID, cleanup
}

// createTestDLPProfile creates a DLP profile for testing
func createTestDLPProfile(t *testing.T, rnd, testName string, allowedMatchCount int64) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	client := cloudflare.NewClient()
	
	profileResp, err := client.ZeroTrust.DLP.Profiles.Custom.New(context.Background(), zero_trust.DLPProfileCustomNewParams{
		AccountID: cloudflare.F(accountID),
		Name:      cloudflare.F(fmt.Sprintf("test-%s-profile-%s", testName, rnd)),
		AllowedMatchCount: cloudflare.F(allowedMatchCount),
	})
	if err != nil {
		t.Fatalf("Failed to create test DLP profile: %v", err)
	}
	
	return profileResp.ID
}

// deleteTestDLPProfile deletes a DLP profile after testing
func deleteTestDLPProfile(t *testing.T, profileID, accountID string) {
	client := cloudflare.NewClient()
	_, err := client.ZeroTrust.DLP.Profiles.Custom.Delete(context.Background(), profileID, zero_trust.DLPProfileCustomDeleteParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		t.Logf("Failed to delete test DLP profile %s: %v", profileID, err)
	}
}

func testSweepCloudflareZeroTrustDLPEntry(r string) error {
	client := cloudflare.NewClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	entries, err := client.ZeroTrust.DLP.Entries.List(context.Background(), zero_trust.DLPEntryListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		return err
	}

	for _, entry := range entries.Result {
		if entry.Type == "custom" {
			_, err := client.ZeroTrust.DLP.Entries.Delete(context.Background(), entry.ID, zero_trust.DLPEntryDeleteParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				continue
			}
		}
	}

	return nil
}

func TestAccCloudflareZeroTrustDLPEntry_Basic(t *testing.T) {
	rnd, profileID, cleanup := setupDLPTest(t, "basic", 5)
	defer cleanup()
	
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_entry.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDLPEntryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZeroTrustDLPEntryConfigBasic(rnd, accountID, profileID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("test-dlp-entry-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"regex":      knownvalue.StringExact("[0-9]{4}[[:space:]]?-?[0-9]{4}[[:space:]]?-?[0-9]{4}[[:space:]]?-?[0-9]{4}"),
						"validation": knownvalue.StringExact("luhn"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("profile_id"), knownvalue.StringExact(profileID)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"type", "case_sensitive", "secret", "confidence", "variant", "word_list", "created_at", "updated_at"},
			},
			{
				Config: testAccCloudflareZeroTrustDLPEntryConfigUpdated(rnd, accountID, profileID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("test-dlp-entry-%s-updated", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"regex": knownvalue.StringExact("[0-9]{3}[[:space:]]?-?[0-9]{2}[[:space:]]?-?[0-9]{4}"),
					})),
				},
			},
		},
	})
}

func TestAccCloudflareZeroTrustDLPEntry_Minimal(t *testing.T) {
	rnd, profileID, cleanup := setupDLPTest(t, "minimal", 1)
	defer cleanup()
	
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_entry.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDLPEntryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZeroTrustDLPEntryConfigMinimal(rnd, accountID, profileID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("minimal-test-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("profile_id"), knownvalue.StringExact(profileID)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"type", "case_sensitive", "secret", "confidence", "variant", "word_list", "created_at", "updated_at"},
			},
		},
	})
}

func testAccCheckCloudflareZeroTrustDLPEntryDestroy(s *terraform.State) error {
	client := cloudflare.NewClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_dlp_entry" {
			continue
		}

		_, err := client.ZeroTrust.DLP.Entries.Get(context.Background(), rs.Primary.ID, zero_trust.DLPEntryGetParams{
			AccountID: cloudflare.F(accountID),
		})
		if err == nil {
			return fmt.Errorf("DLP Entry still exists")
		}
	}

	return nil
}

func testAccCloudflareZeroTrustDLPEntryConfigBasic(rnd, accountID, profileID string) string {
	return acctest.LoadTestCase("dlpentry_basic.tf", rnd, accountID, profileID)
}

func testAccCloudflareZeroTrustDLPEntryConfigUpdated(rnd, accountID, profileID string) string {
	return acctest.LoadTestCase("dlpentry_updated.tf", rnd, accountID, profileID)
}

func testAccCloudflareZeroTrustDLPEntryConfigMinimal(rnd, accountID, profileID string) string {
	return acctest.LoadTestCase("dlpentry_minimal.tf", rnd, accountID, profileID)
}

func TestAccCloudflareZeroTrustDLPEntry_PatternValidations(t *testing.T) {
	rnd, profileID, cleanup := setupDLPTest(t, "validation", 3)
	defer cleanup()
	
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_entry.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDLPEntryDestroy,
		Steps: []resource.TestStep{
			// Test luhn validation
			{
				Config: testAccCloudflareZeroTrustDLPEntryConfigPatternValidations(rnd, accountID, profileID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("test-dlp-entry-validation-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"regex":      knownvalue.StringExact("[0-9]{4}[[:space:]]?-?[0-9]{4}[[:space:]]?-?[0-9]{4}[[:space:]]?-?[0-9]{4}"),
						"validation": knownvalue.StringExact("luhn"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("profile_id"), knownvalue.StringExact(profileID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			// Test pattern without validation
			{
				Config: testAccCloudflareZeroTrustDLPEntryConfigNoValidation(rnd, accountID, profileID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("test-dlp-entry-no-validation-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"regex": knownvalue.StringExact("\\b[A-Z0-9]{1,20}@[A-Z0-9]{1,10}\\.[A-Z]{2,4}\\b"),
					})),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"type", "case_sensitive", "secret", "confidence", "variant", "word_list", "created_at", "updated_at"},
			},
		},
	})
}

func TestAccCloudflareZeroTrustDLPEntry_Comprehensive(t *testing.T) {
	rnd, profileID, cleanup := setupDLPTest(t, "comprehensive", 2)
	defer cleanup()
	
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_entry.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDLPEntryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZeroTrustDLPEntryConfigComprehensive(rnd, accountID, profileID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("comprehensive-dlp-entry-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"regex":      knownvalue.StringExact("[0-9]{16}"),
						"validation": knownvalue.StringExact("luhn"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("profile_id"), knownvalue.StringExact(profileID)),
					// Verify computed attributes are present
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"type", "case_sensitive", "secret", "confidence", "variant", "word_list", "created_at", "updated_at"},
			},
		},
	})
}

func TestAccCloudflareZeroTrustDLPEntry_ToggleEnabled(t *testing.T) {
	rnd, profileID, cleanup := setupDLPTest(t, "toggle", 1)
	defer cleanup()
	
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_entry.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDLPEntryDestroy,
		Steps: []resource.TestStep{
			// Start with enabled=false
			{
				Config: testAccCloudflareZeroTrustDLPEntryConfigToggleEnabled(rnd, accountID, profileID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("toggle-enabled-%s", rnd))),
				},
			},
			// Toggle to enabled=true
			{
				Config: testAccCloudflareZeroTrustDLPEntryConfigMinimal(rnd, accountID, profileID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("minimal-test-%s", rnd))),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"type", "case_sensitive", "secret", "confidence", "variant", "word_list", "created_at", "updated_at"},
			},
		},
	})
}

func testAccCloudflareZeroTrustDLPEntryConfigPatternValidations(rnd, accountID, profileID string) string {
	return acctest.LoadTestCase("dlpentry_pattern_validations.tf", rnd, accountID, profileID)
}

func testAccCloudflareZeroTrustDLPEntryConfigNoValidation(rnd, accountID, profileID string) string {
	return acctest.LoadTestCase("dlpentry_no_validation.tf", rnd, accountID, profileID)
}

func testAccCloudflareZeroTrustDLPEntryConfigComprehensive(rnd, accountID, profileID string) string {
	return acctest.LoadTestCase("dlpentry_comprehensive.tf", rnd, accountID, profileID)
}

func testAccCloudflareZeroTrustDLPEntryConfigToggleEnabled(rnd, accountID, profileID string) string {
	return acctest.LoadTestCase("dlpentry_toggle_enabled.tf", rnd, accountID, profileID)
}