package integration

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/utils/testhelpers"
)

// TestAccessApplicationMigration tests the migration of cloudflare_zero_trust_access_application
// from v4 to v5 by comparing input against expected output.
//
// To add new test cases:
// 1. Add a .tf file to testdata/v4/
// 2. Add the expected .tf file to testdata/v5/ with the same name
// 3. Optionally add .tfstate files for state migration testing
//
// The test framework handles everything else automatically.
func TestAccessApplicationMigration(t *testing.T) {
	testhelpers.RunIntegrationTests(t, testhelpers.IntegrationTestOptions{
		FromVersion: "v4",
		ToVersion:   "v5",
	})
}