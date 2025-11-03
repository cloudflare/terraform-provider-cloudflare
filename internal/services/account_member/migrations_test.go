package account_member_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMain(m *testing.M) {
	// Write IMMEDIATELY to confirm binary execution - before any other calls
	os.Stderr.WriteString("\n===========================================\n")
	os.Stderr.WriteString("=== TestMain: ENTRY POINT ===\n")
	os.Stderr.WriteString("=== Test binary started successfully ===\n")
	os.Stderr.WriteString("===========================================\n")

	fmt.Println("TestMain: About to call resource.TestMain")
	fmt.Println("TestMain: This call may hang if provider setup blocks")

	// This call is suspected to hang
	resource.TestMain(m)

	// If we see this, resource.TestMain completed
	os.Stderr.WriteString("=== TestMain: resource.TestMain returned successfully ===\n")
	fmt.Println("TestMain: Completed")
}

// TestMigrateAccountMember_Basic tests basic migration from v4 to v5
// This test verifies that:
// - email_address is renamed to email
// - role_ids is renamed to roles
func TestMigrateAccountMember_Basic(t *testing.T) {
	t.Logf("=== TEST START: TestMigrateAccountMember_Basic ===")
	t.Logf("=== This log confirms test code is running and producing output ===")

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	email := fmt.Sprintf("%s@example.com", rnd)
	resourceName := "cloudflare_account_member." + rnd
	tmpDir := t.TempDir()

	t.Logf("=== Test setup complete. About to call resource.Test() which will download v4 provider ===")
	t.Logf("=== Resource name: %s, TempDir: %s ===", resourceName, tmpDir)

	// V4 config using old attribute names
	v4Config := fmt.Sprintf(`
resource "cloudflare_account_member" "%[1]s" {
  account_id    = "%[2]s"
  email_address = "%[3]s"
  role_ids      = ["05784afa30c1afe1440e79d9351c7430"]
}`, rnd, accountID, email)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			t.Logf("=== PreCheck: Running pre-flight checks ===")
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			t.Logf("=== PreCheck: Checks passed. Provider download will happen next ===")
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider (this will download v4.52.1 if not cached)
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				// Verify email_address was renamed to email
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
				// Verify role_ids was renamed to roles
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("roles"), knownvalue.SetExact([]knownvalue.Check{
					knownvalue.StringExact("05784afa30c1afe1440e79d9351c7430"),
				})),
			}),
		},
	})
}

// TestMigrateAccountMember_MultipleRoles tests migration with multiple roles
func TestMigrateAccountMember_MultipleRoles(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	email := fmt.Sprintf("%s@example.com", rnd)
	resourceName := "cloudflare_account_member." + rnd
	tmpDir := t.TempDir()

	// V4 config with multiple role_ids
	v4Config := fmt.Sprintf(`
resource "cloudflare_account_member" "%[1]s" {
  account_id    = "%[2]s"
  email_address = "%[3]s"
  role_ids      = [
    "05784afa30c1afe1440e79d9351c7430",
    "3536bcfad5faccb999b47003c79917fb"
  ]
}`, rnd, accountID, email)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("roles"), knownvalue.SetSizeExact(2)),
			}),
		},
	})
}

// TestMigrateAccountMember_WithStatus tests migration with explicit status
func TestMigrateAccountMember_WithStatus(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	email := fmt.Sprintf("%s@example.com", rnd)
	resourceName := "cloudflare_account_member." + rnd
	tmpDir := t.TempDir()

	// V4 config with status
	v4Config := fmt.Sprintf(`
resource "cloudflare_account_member" "%[1]s" {
  account_id    = "%[2]s"
  email_address = "%[3]s"
  role_ids      = ["05784afa30c1afe1440e79d9351c7430"]
  status        = "accepted"
}`, rnd, accountID, email)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("roles"), knownvalue.SetExact([]knownvalue.Check{
					knownvalue.StringExact("05784afa30c1afe1440e79d9351c7430"),
				})),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.StringExact("accepted")),
			}),
		},
	})
}
