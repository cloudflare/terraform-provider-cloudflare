---
name: acceptance-tests
description: Write and debug acceptance tests for Cloudflare's public Terraform provider. Use when contributing `resource_test.go`, `data_source_test.go`, `testdata/*.tf`, import-state checks, sweepers, or diagnosing `TF_ACC=1` failures. Do not use for ordinary Terraform HCL authoring.
---

# Cloudflare Terraform - Acceptance Tests

Acceptance tests in this provider interact with real Cloudflare resources. This skill is for maintainers, product engineers, and external contributors changing the provider itself. It is not for ordinary Terraform users writing HCL that consumes the provider.

These tests are hand-maintained alongside generated provider code, and correctness depends on repo-specific helpers, environment variables, import ID formats, test naming, and cleanup conventions.

## Core Workflow

1. Identify the service directory: `internal/services/<service>`.
2. Inspect existing files before editing: `resource_test.go`, `data_source_test.go`, `testdata/*.tf`, schema/model files, and nearby services with similar scope.
3. If tests do not exist, run `scripts/generate-acctest -dry-run <service>` first, then scaffold with `scripts/generate-acctest <service>` only when the paths are correct.
4. Write valid Terraform config in `testdata/*.tf`; generated scaffold configs are placeholders and will fail until completed.
5. Use `acctest.LoadTestCase(...)` to load `testdata` configs from Go tests.
6. Add resource checks for required IDs, user-configured attributes, API-normalized attributes, and computed fields that prove the read path works.
7. Add import-state verification when the resource supports import.
8. Add or update a sweeper only when tests create remote resources that can be orphaned.
9. Run targeted compile/unit checks first. Run `TF_ACC=1` tests only when credentials and the required account/zone fixtures are available.

## Codegen Boundary

Before editing provider implementation files, check the file header. If a file says it is generated, treat it as disposable output from the provider's code generation pipeline. Do not make durable fixes there unless the user explicitly asks for a temporary local patch.

Acceptance-test files are hand-maintained and safe to edit: `resource_test.go`, `data_source_test.go`, and `testdata/*.tf`. Files under `migration/**` are safe only when the task is explicitly about migration tests.

Generated-file headers may change as the provider's code generation system evolves. Detect generated files by the header, not by a specific generator name.

## Example Services To Inspect

Before inventing a new pattern, inspect existing tests with the same shape:

- `internal/services/turnstile_widget/resource_test.go` for straightforward account-scoped resource tests with update and import verification.
- `internal/services/vulnerability_scanner_credential/resource_test.go` for nested resources and composite `ImportStateIdFunc` import IDs.
- `internal/services/logpush_dataset_field/data_source_test.go` for data source tests using `ConfigStateChecks`, `statecheck`, `knownvalue`, and `tfjsonpath`.
- `internal/services/dns_record/resource_test.go` for extensive lifecycle checks, sweepers, API existence checks, and zone/domain fixtures.
- `internal/services/waiting_room/resource_test.go` for explicit idempotency/no-op plan checks.

Use these as references, not templates to copy blindly.

Do not fabricate Cloudflare SDK methods; use pseudocode until real client calls are inspected.

## Test Configs

Prefer `acctest.LoadTestCase` for resource and data source configs:

```go
func testAccExampleConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("basic.tf", rnd, accountID)
}
```

The corresponding `testdata/basic.tf` uses `fmt.Sprintf` placeholders:

```hcl
resource "cloudflare_example" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
}
```

Use inline `fmt.Sprintf` only for very small one-off configs. Use `//go:embed` mostly for migration tests with multiple v4/v5 fixture files.

## Naming And Cleanup

All remote resources created by acceptance tests should use `utils.GenerateRandomResourceName()` unless the API requires a different format. This produces names with the shared test prefix used by sweepers.

Sweepers must filter with `utils.ShouldSweepResource(name)` before deleting anything:

```go
if !utils.ShouldSweepResource(remote.Name) {
	continue
}
```

Never write a sweeper that deletes all resources in an account or zone by default. The only broad-delete path should be the existing explicit danger-mode behavior in `utils.ShouldSweepResource`.

## PreCheck Selection

Always include credential precheck:

```go
PreCheck: func() { acctest.TestAccPreCheck(t) },
```

Add scope-specific prechecks based on the config. Common helpers:

| Helper | Required environment variable |
| --- | --- |
| `acctest.TestAccPreCheck_AccountID(t)` | `CLOUDFLARE_ACCOUNT_ID` |
| `acctest.TestAccPreCheck_ZoneID(t)` | `CLOUDFLARE_ZONE_ID` |
| `acctest.TestAccPreCheck_Domain(t)` | `CLOUDFLARE_DOMAIN` |
| `acctest.TestAccPreCheck_AlternateZoneID(t)` | `CLOUDFLARE_ALT_ZONE_ID` |
| `acctest.TestAccPreCheck_AlternateDomain(t)` | `CLOUDFLARE_ALT_DOMAIN` |
| `acctest.TestAccPreCheck_Email(t)` | `CLOUDFLARE_EMAIL` |
| `acctest.TestAccPreCheck_APIToken(t)` | `CLOUDFLARE_API_TOKEN` |
| `acctest.TestAccPreCheck_APIKey(t)` | `CLOUDFLARE_API_KEY` |

Check `internal/acctest/acctest.go` for specialized prechecks before adding new environment handling.

## Resource Test Structure

Use protocol v6 provider factories:

```go
resource.Test(t, resource.TestCase{
	PreCheck: func() {
		acctest.TestAccPreCheck(t)
		acctest.TestAccPreCheck_AccountID(t)
	},
	ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
	Steps: []resource.TestStep{
		{
			Config: testAccExampleConfig(rnd, accountID),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
				resource.TestCheckResourceAttr(resourceName, "name", rnd),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
			),
		},
	},
})
```

Use `resource.ParallelTest` only when the resource and backing API are safe for parallel creation, update, and deletion in the shared test account or zone. Singleton resources, per-zone settings, ordered resources, and APIs with eventual consistency often need sequential tests.

Test names are not perfectly uniform across the repo. Some use `TestAccCloudflare<Service>_...`; others use shorter names such as `TestAcc<Service>_...`. Follow the existing service naming style when a test file already exists, and use `^TestAcc` as the broad targeted pattern.

## Data Source Tests

Data source tests usually verify that config inputs and read results are present in state. For simple data sources, `ConfigStateChecks` with `statecheck.ExpectKnownValue` can be clearer than `resource.ComposeTestCheckFunc`:

```go
resource.Test(t, resource.TestCase{
	PreCheck:                 func() { acctest.TestAccPreCheck(t) },
	ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
	Steps: []resource.TestStep{
		{
			Config: testAccExampleDataSourceConfig(rnd, accountID),
			ConfigStateChecks: []statecheck.StateCheck{
				statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
			},
		},
	},
})
```

Use data source tests to cover lookup behavior that resource tests do not prove: account vs zone inputs, list filters, or lookup by name.

## Attribute Checks

Check attributes that prove the full lifecycle works:

- Required scope fields such as `account_id` or `zone_id`.
- User-provided fields from config.
- Important computed fields with `resource.TestCheckResourceAttrSet`.
- API-normalized fields when normalization is expected and stable.
- Collection sizes with `field.#` when order is not the behavior under test.

Prefer constants for shared schema keys:

- `consts.AccountIDSchemaKey`
- `consts.ZoneIDSchemaKey`

Avoid asserting unstable timestamps, API-generated ordering, secrets, or fields known to vary between create/read/import unless the test is specifically about that behavior.

## Import-State Checks

Add import verification when the resource docs and implementation support import:

```go
{
	ResourceName:        resourceName,
	ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
	ImportState:         true,
	ImportStateVerify:   true,
}
```

Choose the prefix based on resource scope. Account-scoped resources usually use `accountID/`; zone-scoped resources usually use `zoneID/`. Confirm against existing tests for the same resource family or the resource import docs.

Use `ImportStateVerifyIgnore` for fields that cannot round-trip through import:

- Write-only or sensitive values.
- Values intentionally omitted by API reads.
- API-computed timestamps.
- Deprecated compatibility fields.
- Fields normalized differently by create and read.

Do not add broad ignore lists without explaining the API behavior in the test or by following an existing pattern in the service.

For a named service failure, inspect that service's current tests and schema before recommending field-specific `ImportStateVerifyIgnore`; do not guess likely ignore fields from general API intuition.

For composite import IDs that need values from multiple resources, use `ImportStateIdFunc` instead of trying to precompute the ID before apply:

```go
{
	ResourceName:            resourceName,
	ImportState:             true,
	ImportStateVerify:       true,
	ImportStateIdFunc: func(s *terraform.State) (string, error) {
		parent, ok := s.RootModule().Resources[parentResourceName]
		if !ok {
			return "", fmt.Errorf("parent resource %q not found in state", parentResourceName)
		}

		child, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("child resource %q not found in state", resourceName)
		}

		return fmt.Sprintf("%s/%s/%s", accountID, parent.Primary.ID, child.Primary.ID), nil
	},
}
```

This pattern is common for nested Cloudflare APIs where the Terraform resource ID alone is not sufficient to import the object.

Only add `ImportStateVerifyIgnore` to this step when you can name the field-specific API reason. For example, API token and credential values are write-only and are not returned by API reads, so `value` may need to be ignored. Do not copy ignore lists from unrelated resources.

## Plan And Idempotency Checks

Add plan checks when the behavior under test is specifically about update actions, replacement avoidance, or no-drift idempotency. Do not add plan checks as a blanket workaround for unstable tests.

Useful patterns:

- `plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate)` to prove an update happens rather than replacement.
- `plancheck.ExpectEmptyPlan()` after reapplying identical config to prove idempotency.
- `PlanOnly: true` plus `ExpectNonEmptyPlan: false` for an explicit no-op plan step.
- `plancheck.ExpectKnownValue(...)` when asserting a planned value before apply is more useful than checking final state only.

When adding `ExpectNonEmptyPlan: true` or `plancheck.ExpectNonEmptyPlan()`, document why the non-empty plan is expected. Non-empty plans are usually a smell unless the test is explicitly covering replacement, API behavior, or a known provider limitation.

For post-apply drift, first capture the exact planned diff. Use a repeated-config step with `plancheck.ExpectEmptyPlan()` or a `PlanOnly: true` / `ExpectNonEmptyPlan: false` step to prove idempotency. Use `PostApplyPostRefresh` checks only when the drift is specifically after refresh.

## CheckDestroy

Use `CheckDestroy` when the API has a reliable get/list endpoint that proves deletion. Iterate Terraform state resources, skip other types, and query the API using IDs from state.

Do not add `CheckDestroy` for singleton resources or settings that are reset rather than deleted. For these, either omit `CheckDestroy` or assert reset behavior in a follow-up read/plan step if the existing service patterns do that.

## Sweepers

Add `TestMain` when registering sweepers:

```go
func TestMain(m *testing.M) {
	resource.TestMain(m)
}
```

Register sweepers in `init()`:

```go
func init() {
	resource.AddTestSweepers("cloudflare_example", &resource.Sweeper{
		Name: "cloudflare_example",
		F:    testSweepCloudflareExample,
	})
}
```

Sweeper rules:

- Return early if required scope env vars are missing.
- List remote resources using the provider's Cloudflare client helpers.
- Delete only resources passing `utils.ShouldSweepResource`.
- Log failures and continue when deleting one stale resource should not prevent the rest from being swept.
- Run `scripts/lint-sweepers` after adding or changing sweepers.

## Credentials And Environment

Acceptance tests require `TF_ACC=1` plus Cloudflare credentials. The base credential precheck accepts supported credential forms, but many tests need account/zone/domain variables.

Check `contributing/environment-variable-dictionary.md` and `internal/acctest/acctest.go` for current environment variables.

Common variables:

```sh
TF_ACC=1
CLOUDFLARE_API_TOKEN=...
CLOUDFLARE_API_KEY=...
CLOUDFLARE_EMAIL=...
CLOUDFLARE_ACCOUNT_ID=...
CLOUDFLARE_ZONE_ID=...
CLOUDFLARE_DOMAIN=...
```

Some older API paths require API key and email instead of API token. Follow existing tests in the same service family before changing credential behavior.

If a test requires optional external fixtures that most contributors will not have, prefer `t.Skip` with a clear missing-variable message. Use `t.Fatal` for standard required fixtures that the acceptance test cannot meaningfully run without, such as account or zone IDs after the matching precheck has been selected.

## Verification

Start with checks that do not require real API access:

```sh
go test ./internal/services/<service> -run '^TestAcc' -count=0
```

Run targeted acceptance tests only when credentials are configured:

```sh
TF_ACC=1 go test ./internal/services/<service> -run '^TestAcc.*Basic$' -v -count=1 -timeout 30m
```

Repo script equivalents:

```sh
./scripts/test ./internal/services/<service>
TF_ACC=1 ./scripts/test ./internal/services/<service> -run '^TestAcc' -count=1
```

Run formatting after edits:

```sh
./scripts/format
```

If credentials are unavailable, run targeted compile checks and report that `TF_ACC=1` verification was not performed.

For dangling resources, prefer targeted sweep commands:

```sh
go test ./internal/services/<service> -v -sweep=all -sweep-run='cloudflare_<service>' -timeout 10m
./scripts/sweep --resource <service> --dry-run
```

Only use dangerous sweep modes with explicit user approval.

## Debugging Failures

Schema/config errors:

- Re-read the service docs or schema.
- Verify `testdata/*.tf` uses v5 resource names and current attribute names.
- Remove read-only fields from config.
- Confirm `fmt.Sprintf` placeholder ordering in `acctest.LoadTestCase` calls.

Import failures:

- Confirm import ID format from docs or neighboring tests.
- Add `ImportStateVerifyIgnore` only for fields that cannot round-trip.
- Compare post-import state with temporary `acctest.DumpState` in `resource.ComposeTestCheckFunc(...)` if needed, then remove it before finalizing unless the user wants debug output kept.

Non-empty plans after apply:

- Check whether the API normalizes empty strings, empty collections, ordering, casing, or defaults.
- Prefer fixing config or provider behavior over weakening the test. If provider behavior must change, respect the codegen boundary instead of patching generated read logic directly.
- Use plan checks only when an existing service pattern shows the diff is expected.
- Do not use `ImportStateVerifyIgnore` for ordinary apply drift; it only affects import verification.

Destroy failures:

- Determine whether the resource is a singleton/resettable setting rather than deletable.
- Check eventual consistency and API delete behavior.
- Add or fix sweepers if resources can be orphaned.

## Migration Tests Are Related But Separate

Do not turn a normal acceptance-test task into a migration-test rewrite unless the user explicitly asks about v4 to v5 migration, state upgraders, `tf-migrate`, or files under `migration/v500`.

When the task is explicitly migration-related, inspect existing migration tests and helpers in `internal/acctest/acctest.go` before writing new patterns. Migration tests use specialized helpers such as `MigrationV2TestStep`, `MigrationV2TestStepWithPlan`, `MigrationV2TestStepWithStateNormalization`, and resource-specific helpers.

## Common Pitfalls

- Editing generated provider files instead of hand-maintained tests or codegen inputs.
- Leaving generated scaffold placeholders in `resource_test.go` or `testdata/*.tf`.
- Using fixed resource names that sweepers will not clean up.
- Adding a sweeper that does not call `utils.ShouldSweepResource`.
- Forgetting `TestMain` when adding a sweeper.
- Using `account_id` when a resource is zone-scoped, or `zone_id` when it is account-scoped.
- Guessing import ID formats instead of checking docs or nearby tests.
- Asserting API-generated timestamps or unstable collection order.
- Running broad `TF_ACC=1 ./scripts/test` instead of a targeted service test while iterating.
- Mutating global Terraform CLI config such as `~/.terraformrc` without explicit user approval.

## Evaluating This Skill

When iterating on this skill, use realistic prompts and compare agent output with and without the skill. Start with the eval prompts in `evals/evals.json`, then add cases from real provider PRs when agents miss conventions or produce unsafe tests.
