package worker_version_test

import (
	"context"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/worker_version"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// TestMigrateWorkerVersionFromV5_14 tests migration from v5.14.0 to latest.
// See: https://github.com/cloudflare/terraform-provider-cloudflare/issues/6567
func TestMigrateWorkerVersionFromV5_14(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_worker_version." + rnd
	tmpDir := t.TempDir()
	contentFile := path.Join(tmpDir, "index.js")

	writeContentFile := func(t *testing.T, content string) {
		err := os.WriteFile(contentFile, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Error creating temp file at path %s: %s", contentFile, err.Error())
		}
	}

	cleanup := func(t *testing.T) {
		err := os.Remove(contentFile)
		if err != nil {
			t.Logf("Error removing temp file at path %s: %s", contentFile, err.Error())
		}
	}

	defer cleanup(t)

	v5_14Config := fmt.Sprintf(`
resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"
}

resource "cloudflare_worker_version" "%[1]s" {
  account_id = "%[2]s"
  worker_id = cloudflare_worker.%[1]s.id
  modules = [
    {
      name         = "index.js"
      content_file = "%[3]s"
      content_type = "application/javascript+module"
    }
  ]
  main_module = "index.js"
}`, rnd, accountID, contentFile)

	latestConfig := fmt.Sprintf(`
resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"
}

resource "cloudflare_worker_version" "%[1]s" {
  account_id = "%[2]s"
  worker_id = cloudflare_worker.%[1]s.id
  modules = [
    {
      name         = "index.js"
      content_file = "%[3]s"
      content_type = "application/javascript+module"
    }
  ]
  main_module = "index.js"
}`, rnd, accountID, contentFile)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					writeContentFile(t, `export default {fetch() {return new Response()}}`)
				},
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.14.0",
					},
				},
				Config: v5_14Config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("main_module"), knownvalue.StringExact("index.js")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("number"), knownvalue.NotNull()),
				},
			},
		},
	})
}

// TestMigrateWorkerVersionFromV5_14_WithBindings tests migration with bindings.
func TestMigrateWorkerVersionFromV5_14_WithBindings(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_worker_version." + rnd
	tmpDir := t.TempDir()
	contentFile := path.Join(tmpDir, "index.js")

	writeContentFile := func(t *testing.T, content string) {
		err := os.WriteFile(contentFile, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Error creating temp file at path %s: %s", contentFile, err.Error())
		}
	}

	cleanup := func(t *testing.T) {
		err := os.Remove(contentFile)
		if err != nil {
			t.Logf("Error removing temp file at path %s: %s", contentFile, err.Error())
		}
	}

	defer cleanup(t)

	v5_14Config := fmt.Sprintf(`
resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"
}

resource "cloudflare_worker_version" "%[1]s" {
  account_id = "%[2]s"
  worker_id = cloudflare_worker.%[1]s.id
  modules = [
    {
      name         = "index.js"
      content_file = "%[3]s"
      content_type = "application/javascript+module"
    }
  ]
  main_module = "index.js"
  
  bindings = [
    {
      name = "PLATFORM_DOMAIN"
      text = "example.com"
      type = "plain_text"
    }
  ]
}`, rnd, accountID, contentFile)

	latestConfig := fmt.Sprintf(`
resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"
}

resource "cloudflare_worker_version" "%[1]s" {
  account_id = "%[2]s"
  worker_id = cloudflare_worker.%[1]s.id
  modules = [
    {
      name         = "index.js"
      content_file = "%[3]s"
      content_type = "application/javascript+module"
    }
  ]
  main_module = "index.js"
  
  bindings = [
    {
      name = "PLATFORM_DOMAIN"
      text = "example.com"
      type = "plain_text"
    }
  ]
}`, rnd, accountID, contentFile)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					writeContentFile(t, `export default {fetch() {return new Response()}}`)
				},
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.14.0",
					},
				},
				Config: v5_14Config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("main_module"), knownvalue.StringExact("index.js")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"name": knownvalue.StringExact("PLATFORM_DOMAIN"),
							"type": knownvalue.StringExact("plain_text"),
							"text": knownvalue.StringExact("example.com"),
						}),
					})),
				},
			},
		},
	})
}

// TestMigrateWorkerVersionFromV0_WithStartupTimeMs tests state upgrade with startup_time_ms.
// See: https://github.com/cloudflare/terraform-provider-cloudflare/issues/6567
func TestMigrateWorkerVersionFromV0_WithStartupTimeMs(t *testing.T) {
	ctx := context.Background()

	schemaV0 := worker_version.ResourceSchemaV0ForTest(ctx)

	// Build a state value that includes startup_time_ms
	stateValue := tftypes.NewValue(schemaV0.Type().TerraformType(ctx), map[string]tftypes.Value{
		"id":                  tftypes.NewValue(tftypes.String, "test-version-id"),
		"account_id":          tftypes.NewValue(tftypes.String, "test-account-id"),
		"worker_id":           tftypes.NewValue(tftypes.String, "test-worker-id"),
		"compatibility_date":  tftypes.NewValue(tftypes.String, nil),
		"main_module":         tftypes.NewValue(tftypes.String, "index.js"),
		"migrations":          tftypes.NewValue(schemaV0.Attributes["migrations"].GetType().TerraformType(ctx), nil),
		"modules":             tftypes.NewValue(schemaV0.Attributes["modules"].GetType().TerraformType(ctx), nil),
		"placement":           tftypes.NewValue(schemaV0.Attributes["placement"].GetType().TerraformType(ctx), nil),
		"usage_model":         tftypes.NewValue(tftypes.String, "standard"),
		"compatibility_flags": tftypes.NewValue(schemaV0.Attributes["compatibility_flags"].GetType().TerraformType(ctx), nil),
		"annotations":         tftypes.NewValue(schemaV0.Attributes["annotations"].GetType().TerraformType(ctx), nil),
		"assets":              tftypes.NewValue(schemaV0.Attributes["assets"].GetType().TerraformType(ctx), nil),
		"bindings":            tftypes.NewValue(schemaV0.Attributes["bindings"].GetType().TerraformType(ctx), nil),
		"limits":              tftypes.NewValue(schemaV0.Attributes["limits"].GetType().TerraformType(ctx), nil),
		"created_on":          tftypes.NewValue(tftypes.String, "2024-01-01T00:00:00Z"),
		"number":              tftypes.NewValue(tftypes.Number, 1),
		"source":              tftypes.NewValue(tftypes.String, "api"),
		"main_script_base64":  tftypes.NewValue(tftypes.String, nil),
		"startup_time_ms":     tftypes.NewValue(tftypes.Number, 100),
	})

	state := tfsdk.State{
		Schema: schemaV0,
		Raw:    stateValue,
	}

	// Without the fix, this fails with "Object defines fields not found in struct: startup_time_ms"
	var model worker_version.ResourceModelV0ForTest
	diags := state.Get(ctx, &model)

	if diags.HasError() {
		for _, d := range diags.Errors() {
			t.Errorf("Error: %s - %s", d.Summary(), d.Detail())
		}
		t.Fatalf("Failed to read state into resourceModelV0")
	}

	if model.ID.ValueString() != "test-version-id" {
		t.Errorf("Expected ID to be 'test-version-id', got '%s'", model.ID.ValueString())
	}
	if model.MainModule.ValueString() != "index.js" {
		t.Errorf("Expected MainModule to be 'index.js', got '%s'", model.MainModule.ValueString())
	}
	if model.StartupTimeMs.ValueInt64() != 100 {
		t.Errorf("Expected StartupTimeMs to be 100, got %d", model.StartupTimeMs.ValueInt64())
	}
}

