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

const EnvTfAcc = "TF_ACC"

// dlpTestSetup performs common test setup including TF_ACC check and client/profile creation
type dlpTestSetup struct {
	client    *cloudflare.Client
	accountID string
	profileID string
	cleanupFn func()
}

// setupDLPTest handles common test setup and returns setup struct with cleanup function
func setupDLPTest(t *testing.T, profileName string) *dlpTestSetup {
	// Check TF_ACC environment variable first
	if os.Getenv(EnvTfAcc) == "" {
		t.Skip(fmt.Sprintf(
			"Acceptance tests skipped unless env '%s' set",
			EnvTfAcc))
		return nil
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Fatal("CLOUDFLARE_ACCOUNT_ID must be set for acceptance tests")
	}

	// Create Cloudflare client
	client := cloudflare.NewClient()

	// Create test DLP profile
	profileResp, err := client.ZeroTrust.DLP.Profiles.Custom.New(context.Background(), zero_trust.DLPProfileCustomNewParams{
		AccountID:         cloudflare.F(accountID),
		Name:              cloudflare.F(profileName),
		AllowedMatchCount: cloudflare.F(int64(5)),
	})
	if err != nil {
		t.Fatalf("Failed to create test DLP profile: %v", err)
	}

	setup := &dlpTestSetup{
		client:    client,
		accountID: accountID,
		profileID: profileResp.ID,
		cleanupFn: func() {
			client.ZeroTrust.DLP.Profiles.Custom.Delete(context.Background(), profileResp.ID, zero_trust.DLPProfileCustomDeleteParams{
				AccountID: cloudflare.F(accountID),
			})
		},
	}

	// Register cleanup function
	t.Cleanup(setup.cleanupFn)

	return setup
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_dlp_entry", &resource.Sweeper{
		Name: "cloudflare_zero_trust_dlp_entry",
		F:    testSweepCloudflareZeroTrustDLPEntry,
	})
}

func TestMain(m *testing.M) {
	resource.TestMain(m)
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
		if entry.Type == "custom" && utils.ShouldSweepResource(entry.Name) {
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
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_entry.%s", rnd)

	setup := setupDLPTest(t, fmt.Sprintf("test-basic-profile-%s", rnd))
	if setup == nil {
		return // Test was skipped
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDLPEntryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZeroTrustDLPEntryConfigBasic(rnd, setup.accountID, setup.profileID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(setup.accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("test-dlp-entry-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"regex":      knownvalue.StringExact("[0-9]{4}[[:space:]]?-?[0-9]{4}[[:space:]]?-?[0-9]{4}[[:space:]]?-?[0-9]{4}"),
						"validation": knownvalue.StringExact("luhn"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("profile_id"), knownvalue.StringExact(setup.profileID)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", setup.accountID),
				ImportStateVerifyIgnore: []string{"type", "case_sensitive", "secret", "confidence", "variant", "word_list", "created_at", "updated_at"},
			},
			{
				Config: testAccCloudflareZeroTrustDLPEntryConfigUpdated(rnd, setup.accountID, setup.profileID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(setup.accountID)),
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
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_entry.%s", rnd)

	setup := setupDLPTest(t, fmt.Sprintf("test-profile-%s", rnd))
	if setup == nil {
		return // Test was skipped
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDLPEntryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZeroTrustDLPEntryConfigMinimal(rnd, setup.accountID, setup.profileID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(setup.accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("minimal-test-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("profile_id"), knownvalue.StringExact(setup.profileID)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", setup.accountID),
				ImportStateVerifyIgnore: []string{"type", "case_sensitive", "secret", "confidence", "variant", "word_list", "created_at", "updated_at"},
			},
		},
	})
}

func TestAccCloudflareZeroTrustDLPEntry_PatternValidations(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_entry.%s", rnd)

	setup := setupDLPTest(t, fmt.Sprintf("test-validation-profile-%s", rnd))
	if setup == nil {
		return // Test was skipped
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDLPEntryDestroy,
		Steps: []resource.TestStep{
			// Test credit_card validation
			{
				Config: testAccCloudflareZeroTrustDLPEntryConfigPatternValidations(rnd, setup.accountID, setup.profileID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(setup.accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("test-dlp-entry-validation-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"regex":      knownvalue.StringExact("4[0-9]{3}-[0-9]{4}-[0-9]{4}-[0-9]{4}"),
						"validation": knownvalue.StringExact("luhn"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("profile_id"), knownvalue.StringExact(setup.profileID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("case_sensitive"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("secret"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			// Test pattern without validation
			{
				Config: testAccCloudflareZeroTrustDLPEntryConfigNoValidation(rnd, setup.accountID, setup.profileID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("test-dlp-entry-no-validation-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"regex": knownvalue.StringExact("\\b[a-z]{1,10}@[a-z]{1,10}\\.[a-z]{2,4}\\b"),
					})),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", setup.accountID),
				ImportStateVerifyIgnore: []string{"type", "case_sensitive", "secret", "confidence", "variant", "word_list", "created_at", "updated_at"},
			},
		},
	})
}

func TestAccCloudflareZeroTrustDLPEntry_Comprehensive(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_entry.%s", rnd)

	setup := setupDLPTest(t, fmt.Sprintf("test-comprehensive-profile-%s", rnd))
	if setup == nil {
		return // Test was skipped
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDLPEntryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZeroTrustDLPEntryConfigComprehensive(rnd, setup.accountID, setup.profileID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(setup.accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("comprehensive-dlp-entry-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"regex": knownvalue.StringExact("[0-9]{3}-[0-9]{2}-[0-9]{4}"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("profile_id"), knownvalue.StringExact(setup.profileID)),
					// Verify computed attributes are present
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("case_sensitive"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("secret"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("updated_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("confidence"), knownvalue.Null()),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", setup.accountID),
				ImportStateVerifyIgnore: []string{"type", "case_sensitive", "secret", "confidence", "variant", "word_list", "created_at", "updated_at"},
			},
		},
	})
}

func TestAccCloudflareZeroTrustDLPEntry_ToggleEnabled(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_entry.%s", rnd)

	setup := setupDLPTest(t, fmt.Sprintf("test-toggle-profile-%s", rnd))
	if setup == nil {
		return // Test was skipped
	}

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
				Config: testAccCloudflareZeroTrustDLPEntryConfigToggleEnabled(rnd, setup.accountID, setup.profileID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("toggle-enabled-%s", rnd))),
				},
			},
			// Toggle to enabled=true
			{
				Config: testAccCloudflareZeroTrustDLPEntryConfigMinimal(rnd, setup.accountID, setup.profileID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("minimal-test-%s", rnd))),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", setup.accountID),
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
