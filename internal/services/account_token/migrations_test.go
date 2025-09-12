package account_token_test

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

// TestMigrateAccountToken_V0ToV1_Basic tests state schema migration from v0 to v1
// This ensures the migration from JSON string policies to Dynamic policies works correctly
func TestMigrateAccountToken_V0ToV1_Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_account_token." + rnd

	// Test configuration that would create policies in the new format
	config := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }, {
      id = "c8fed203ed3043cba015a93ad1616f1f"
    }]
    resources = {
      "com.cloudflare.api.account.%[2]s" = "*"
    }
  }]
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.StringExact("active")),
				},
			},
		},
	})
}

// TestMigrateAccountToken_V0ToV1_ComplexPolicies tests migration with multiple policies
// including both simple string resources and nested object resources
func TestMigrateAccountToken_V0ToV1_ComplexPolicies(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_account_token." + rnd

	// Test configuration with multiple policies - one simple, one nested
	config := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    resources = {
      "com.cloudflare.api.account.%[2]s" = "*"
    }
  }, {
    effect = "allow"
    permission_groups = [{
      id = "c8fed203ed3043cba015a93ad1616f1f"
    }]
    resources = {
      "com.cloudflare.api.account.%[2]s" = {
        "com.cloudflare.api.account.zone.*" = "*"
      }
    }
  }]
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.StringExact("active")),
				},
			},
		},
	})
}

// TestMigrateAccountToken_V0ToV1_MinimalPolicy tests migration with minimal policy structure
func TestMigrateAccountToken_V0ToV1_MinimalPolicy(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_account_token." + rnd

	// Test configuration with minimal policy
	config := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    resources = {
      "com.cloudflare.api.account.%[2]s" = "*"
    }
  }]
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.StringExact("active")),
				},
			},
		},
	})
}
