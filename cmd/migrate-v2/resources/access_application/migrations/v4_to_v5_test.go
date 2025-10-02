package migrations

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/utils/testhelpers"
	"github.com/stretchr/testify/require"
)

func TestAccessApplicationPoliciesTransformation(t *testing.T) {
	migration, err := NewV4ToV5Migration()
	require.NoError(t, err)

	runner := testhelpers.NewMigrationTestRunner(migration, "cloudflare_zero_trust_access_application")

	tests := []testhelpers.MigrationTestCase{
		{
			Name: "transform policies from list of strings to list of objects",
			Input: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  
  policies = [
    cloudflare_zero_trust_access_policy.allow.id,
    cloudflare_zero_trust_access_policy.deny.id
  ]
}`,
			Expected: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"

  policies = [{ id = cloudflare_zero_trust_access_policy.allow.id }, { id = cloudflare_zero_trust_access_policy.deny.id }]
  type     = "self_hosted"
}`,
		},
		{
			Name: "transform policies with literal IDs",
			Input: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  
  policies = ["policy-id-1", "policy-id-2"]
}`,
			Expected: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"

  policies = [{ id = "policy-id-1" }, { id = "policy-id-2" }]
  type     = "self_hosted"
}`,
		},
		{
			Name: "mixed references and literals",
			Input: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  
  policies = [
    cloudflare_zero_trust_access_policy.allow.id,
    "literal-policy-id",
    cloudflare_zero_trust_access_policy.deny.id
  ]
}`,
			Expected: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"

  policies = [{ id = cloudflare_zero_trust_access_policy.allow.id }, { id = "literal-policy-id" }, { id = cloudflare_zero_trust_access_policy.deny.id }]
  type     = "self_hosted"
}`,
		},
		{
			Name: "handle old resource name references",
			Input: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  
  policies = [
    cloudflare_access_policy.old_style.id
  ]
}`,
			Expected: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"

  policies = [{ id = cloudflare_zero_trust_access_policy.old_style.id }]
  type     = "self_hosted"
}`,
		},
		{
			Name: "already transformed policies should not change",
			Input: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  
  policies = [
    { id = cloudflare_zero_trust_access_policy.allow.id },
    { id = "literal-policy-id" }
  ]
}`,
			Expected: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"

  policies = [{ id = cloudflare_zero_trust_access_policy.allow.id }, { id = "literal-policy-id" }]
  type     = "self_hosted"
}`,
		},
		{
			Name: "empty policies list",
			Input: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  
  policies = []
}`,
			Expected: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"

  policies = []
  type     = "self_hosted"
}`,
		},
	}

	runner.RunConfigTests(t, tests)
}

func TestAccessApplicationDestinationsBlocksToAttribute(t *testing.T) {
	migration, err := NewV4ToV5Migration()
	require.NoError(t, err)

	runner := testhelpers.NewMigrationTestRunner(migration, "cloudflare_zero_trust_access_application")

	tests := []testhelpers.MigrationTestCase{
		{
			Name: "single destinations block conversion",
			Input: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "self_hosted"
  
  destinations {
    uri = "https://example.com"
  }
}`,
			Expected: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "self_hosted"

  destinations = [
    {
      uri = "https://example.com"
    }
  ]
}`,
		},
		{
			Name: "multiple destinations blocks conversion",
			Input: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "self_hosted"
  
  destinations {
    uri = "https://example.com"
  }
  
  destinations {
    uri = "tcp://db.example.com:5432"
  }
}`,
			Expected: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "self_hosted"

  destinations = [
    {
      uri = "https://example.com"
    },
    {
      uri = "tcp://db.example.com:5432"
    }
  ]
}`,
		},
		{
			Name: "destinations with nested uri block",
			Input: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "self_hosted"
  
  destinations {
    uri {
      path = "/admin"
    }
  }
}`,
			Expected: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "self_hosted"

  destinations = [
    {
      uri = {
        path = "/admin"
      }
    }
  ]
}`,
		},
		{
			Name: "no destinations blocks",
			Input: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "self_hosted"
}`,
			Expected: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "self_hosted"
}`,
		},
	}

	runner.RunConfigTests(t, tests)
}

func TestAccessApplicationCombinedMigrations(t *testing.T) {
	migration, err := NewV4ToV5Migration()
	require.NoError(t, err)

	runner := testhelpers.NewMigrationTestRunner(migration, "cloudflare_zero_trust_access_application")

	tests := []testhelpers.MigrationTestCase{
		{
			Name: "combined domain_type removal and destinations conversion",
			Input: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id  = "abc123"
  name        = "Test App"
  type        = "warp"
  domain_type = "public"
  
  destinations {
    uri = "https://example.com"
  }
  
  policies = ["policy-id-1", "policy-id-2"]
}`,
			Expected: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  policies = [{ id = "policy-id-1" }, { id = "policy-id-2" }]
  destinations = [
    {
      uri = "https://example.com"
    }
  ]
}`,
		},
		{
			Name: "all transformations together with allowed_idps",
			Input: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id   = "abc123"
  name         = "Test App"
  type         = "warp"
  domain_type  = "public"
  allowed_idps = toset(["idp-1", "idp-2"])
  
  destinations {
    uri = "https://example.com"
  }
  
  destinations {
    uri = "tcp://db.example.com:5432"
  }
  
  policies = [
    cloudflare_zero_trust_access_policy.allow.id,
    "literal-policy-id"
  ]
}`,
			Expected: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id   = "abc123"
  name         = "Test App"
  type         = "warp"
  allowed_idps = ["idp-1", "idp-2"]
  
  policies = [{ id = cloudflare_zero_trust_access_policy.allow.id }, { id = "literal-policy-id" }]
  destinations = [
    {
      uri = "https://example.com"
    },
    {
      uri = "tcp://db.example.com:5432"
    }
  ]
}`,
		},
	}

	runner.RunConfigTests(t, tests)
}

func TestAccessApplicationSkipAppLauncherLoginPageRemoval(t *testing.T) {
	migration, err := NewV4ToV5Migration()
	require.NoError(t, err)

	runner := testhelpers.NewMigrationTestRunner(migration, "cloudflare_zero_trust_access_application")

	tests := []testhelpers.MigrationTestCase{
		{
			Name: "remove skip_app_launcher_login_page when type is not app_launcher",
			Input: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id                   = "abc123"
  name                         = "Test App"
  domain                       = "test.example.com"
  type                         = "self_hosted"
  skip_app_launcher_login_page = false
}`,
			Expected: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  type       = "self_hosted"
}`,
		},
		{
			Name: "preserve skip_app_launcher_login_page when type is app_launcher",
			Input: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id                   = "abc123"
  name                         = "Test App"
  type                         = "app_launcher"
  skip_app_launcher_login_page = true
}`,
			Expected: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id                   = "abc123"
  name                         = "Test App"
  type                         = "app_launcher"
  skip_app_launcher_login_page = true
}`,
		},
	}

	runner.RunConfigTests(t, tests)
}

func TestAccessApplicationStateMigration(t *testing.T) {
	migration, err := NewV4ToV5Migration()
	require.NoError(t, err)

	runner := testhelpers.NewMigrationTestRunner(migration, "cloudflare_zero_trust_access_application")

	tests := []testhelpers.StateTestCase{
		{
			Name: "remove domain_type from state",
			Input: map[string]interface{}{
				"id":          "app-123",
				"account_id":  "acc-123",
				"name":        "Test App",
				"domain_type": "public",
				"type":        "self_hosted",
			},
			Expected: map[string]interface{}{
				"id":             "app-123",
				"account_id":     "acc-123",
				"name":           "Test App",
				"type":           "self_hosted",
				"schema_version": 1,
			},
		},
		{
			Name: "add default type if missing",
			Input: map[string]interface{}{
				"id":         "app-123",
				"account_id": "acc-123",
				"name":       "Test App",
			},
			Expected: map[string]interface{}{
				"id":             "app-123",
				"account_id":     "acc-123",
				"name":           "Test App",
				"type":           "self_hosted",
				"schema_version": 1,
			},
		},
	}

	runner.RunStateTests(t, tests)
}