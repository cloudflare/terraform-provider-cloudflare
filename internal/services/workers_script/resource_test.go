package workers_script_test

import (
	"context"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

const (
	resourcePrefix    = "tfacctest-"
	scriptContent1    = `addEventListener('fetch', event => {event.respondWith(new Response('test 1'))});`
	scriptContent2    = `addEventListener('fetch', event => {event.respondWith(new Response('test 2'))});`
	moduleContent     = `import {DurableObject} from 'cloudflare:workers'; export class MyDurableObject extends DurableObject {}; export default { fetch() { return new Response('Hello world'); }, };`
	encodedWasm       = "AGFzbQEAAAAGgYCAgAAA" // wat source: `(module)`, so literally just an empty wasm module
	compatibilityDate = "2023-03-19"
	d1DatabaseID      = "ce8b95dc-b376-4ff8-9b9e-1801ed6d745d"
)

var (
	compatibilityFlags = []string{"nodejs_compat", "web_socket_compression"}
)

func init() {
	resource.AddTestSweepers("cloudflare_workers_script", &resource.Sweeper{
		Name: "cloudflare_workers_script",
		F:    testSweepCloudflareWorkerScripts,
	})
}

func testSweepCloudflareWorkerScripts(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		tflog.Info(ctx, "Skipping worker scripts sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	list, err := client.Workers.Scripts.List(ctx, workers.ScriptListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list worker scripts: %s", err))
		return fmt.Errorf("failed to list worker scripts: %w", err)
	}

	if len(list.Result) == 0 {
		tflog.Info(ctx, "No worker scripts to sweep")
		return nil
	}

	for _, script := range list.Result {
		// Use standard filtering helper to only delete test worker scripts
		if !utils.ShouldSweepResource(script.ID) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting worker script: %s (account: %s)", script.ID, accountID))
		_, err := client.Workers.Scripts.Delete(ctx, script.ID, workers.ScriptDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete worker script %s: %s", script.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted worker script: %s", script.ID))
	}

	return nil
}

func TestAccCloudflareWorkerScript_ServiceWorker(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	resourceName := resourcePrefix + rnd
	name := "cloudflare_workers_script." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerScriptConfigServiceWorkerInitial(resourceName, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(resourceName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("content"), knownvalue.StringExact(scriptContent1)),
				},
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigServiceWorkerUpdate(resourceName, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(resourceName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("content"), knownvalue.StringExact(scriptContent2)),
				},
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigServiceWorkerUpdateBinding(resourceName, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(resourceName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("content"), knownvalue.StringExact(scriptContent2)),
				},
			},
			{
				PreConfig: func() {
					client := acctest.SharedClient()
					result, err := client.Workers.Scripts.Settings.Edit(context.Background(), resourceName, workers.ScriptSettingEditParams{AccountID: cloudflare.F(accountID), ScriptSetting: workers.ScriptSettingParam{Logpush: cloudflare.Bool(true)}})
					if err != nil {
						t.Errorf("Error updating script settings out-of-band to test drift detection: %s", err)
					}
					if result == nil {
						t.Error("Could not update script settings out-of-band to test drift detection.")
					}
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
				RefreshPlanChecks: resource.RefreshPlanChecks{
					PostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
			},
			{
				ResourceName:            name,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"startup_time_ms"},
			},
		},
	})
}

func TestAccCloudflareWorkerScript_ModuleUpload(t *testing.T) {
	t.Skip("issue: API behavior change has caused this test to start failing")
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	resourceName := resourcePrefix + rnd
	name := "cloudflare_workers_script." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		// CheckDestroy:             testAccCheckCloudflareWorkerScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerScriptUploadModule(resourceName, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(resourceName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("content"), knownvalue.StringExact(moduleContent)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("compatibility_date"), knownvalue.StringExact(compatibilityDate)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("compatibility_flags"), knownvalue.SetSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("compatibility_flags").AtSliceIndex(0), knownvalue.StringExact(compatibilityFlags[0])),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("logpush"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("placement").AtMapKey("mode"), knownvalue.StringExact("smart")),
				},
			},
			{
				PreConfig: func() {
					client := acctest.SharedClient()
					result, err := client.Workers.Scripts.Settings.Edit(context.Background(), resourceName, workers.ScriptSettingEditParams{AccountID: cloudflare.F(accountID), ScriptSetting: workers.ScriptSettingParam{Logpush: cloudflare.Bool(true)}})
					if err != nil {
						t.Errorf("Error updating script settings out-of-band to test drift detection: %s", err)
					}
					if result == nil {
						t.Error("Could not update script settings out-of-band to test drift detection.")
					}
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
				RefreshPlanChecks: resource.RefreshPlanChecks{
					PostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
			},
			{
				ResourceName:            name,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"bindings.2", "bindings.#", "main_module", "startup_time_ms"},
			},
		},
	})
}

func TestAcc_WorkerScriptWithContentFile(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	resourceName := resourcePrefix + rnd
	name := "cloudflare_workers_script." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()
	contentFile := path.Join(tmpDir, "worker.mjs")

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

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					writeContentFile(t, moduleContent)
				},
				Config: testAccWorkersScriptConfigWithContentFile(resourceName, accountID, contentFile),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(resourceName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("content_file"), knownvalue.StringExact(contentFile)),
				},
			},
			{
				PreConfig: func() {
					writeContentFile(t, fmt.Sprintf("%s // v2", moduleContent))
				},
				Config: testAccWorkersScriptConfigWithContentFile(resourceName, accountID, contentFile),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
			},
			{
				Config: testAccWorkersScriptConfigWithContentFile(resourceName, accountID, contentFile),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				PreConfig: func() {
					// revert remote state back to the original module content
					client := acctest.SharedClient()
					boundary := "--form-data-boundary-tkqpb9sps99x33zg"
					body := []byte(fmt.Sprintf(`--%s
Content-Disposition: form-data; name="files"; filename="worker.mjs"
Content-Type: application/javascript+module

%s
--%s
Content-Disposition: form-data; name="metadata"; filename="metadata.json"
Content-Type: application/json

{"main_module": "worker.mjs"}
--%s--
`,
						boundary, moduleContent, boundary, boundary,
					))
					result, err := client.Workers.Scripts.Update(context.Background(),
						resourceName,
						workers.ScriptUpdateParams{AccountID: cloudflare.F(accountID)},
						option.WithRequestBody("multipart/form-data;boundary="+boundary, body),
					)
					if err != nil {
						t.Errorf("Error updating script content out-of-band to test drift detection: %s", err)
					}
					if result == nil {
						t.Error("Could not update script content out-of-band to test drift detection.")
					}
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
				RefreshPlanChecks: resource.RefreshPlanChecks{
					PostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
			},
			{
				ResourceName:            name,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"content", "content_file", "content_sha256", "main_module", "startup_time_ms"},
			},
		},
	})
}

func TestAcc_WorkerScriptWithInvalidContentSHA256(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	resourceName := resourcePrefix + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()
	contentFile := path.Join(tmpDir, "worker.mjs")

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

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					writeContentFile(t, moduleContent)
				},
				Config:      testAccWorkersScriptConfigWithInvalidContentSHA256(resourceName, accountID, contentFile),
				ExpectError: regexp.MustCompile(`SHA-256 Hash Mismatch`),
			},
		},
	})
}

func TestAccCloudflareWorkerScript_PythonWorker(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	resourceName := resourcePrefix + rnd
	name := "cloudflare_workers_script." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("python_worker.tf", resourceName, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(resourceName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("main_module"), knownvalue.StringExact("index.py")),
				},
			},
			{
				ResourceName:            name,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"content_type", "has_modules", "main_module", "startup_time_ms"},
			},
		},
	})
}

func TestAcc_WorkerScriptWithAssets(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	resourceName := resourcePrefix + rnd
	name := "cloudflare_workers_script." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	assetsDir := t.TempDir()
	assetFile := path.Join(assetsDir, "index.html")

	writeAssetFile := func(t *testing.T, content string) {
		err := os.WriteFile(assetFile, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Error creating temp file at path %s: %s", assetFile, err.Error())
		}
	}

	cleanup := func(t *testing.T) {
		err := os.Remove(assetFile)
		if err != nil {
			t.Logf("Error removing temp file at path %s: %s", assetFile, err.Error())
		}
	}

	defer cleanup(t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					writeAssetFile(t, "v1")
				},
				Config: testAccWorkersScriptConfigWithAssets(resourceName, accountID, assetsDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("directory"), knownvalue.StringExact(assetsDir)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("asset_manifest_sha256"), knownvalue.StringExact("b098d2898ca7ae5677c7291d97323e7894137515043f3e560f3bd155870eea9e")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("has_assets"), knownvalue.Bool(true)),
				},
			},
			{
				PreConfig: func() {
					writeAssetFile(t, "v2")
				},
				Config: testAccWorkersScriptConfigWithAssets(resourceName, accountID, assetsDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("directory"), knownvalue.StringExact(assetsDir)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("asset_manifest_sha256"), knownvalue.StringExact("46f07eb8a3fa881af81ce2b6b3fc1627edccf115526aa5c631308c45c75d2fb1")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("has_assets"), knownvalue.Bool(true)),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
			},
			{
				Config: testAccWorkersScriptConfigWithAssets(resourceName, accountID, assetsDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				ResourceName:            name,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"assets.%", "assets.directory", "assets.asset_manifest_sha256", "startup_time_ms"},
			},
		},
	})
}

func TestAccCloudflareWorkerScript_ModuleWithDurableObject(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	resourceName := resourcePrefix + rnd
	name := "cloudflare_workers_script." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("module_with_durable_object.tf", resourceName, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(resourceName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("main_module"), knownvalue.StringExact("worker.js")),
				},
			},
			{
				ResourceName:            name,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"bindings.0.namespace_id", "has_modules", "main_module", "migrations", "startup_time_ms"},
			},
		},
	})
}

func TestAccCloudflareWorkerScript_AssetsConfigRunWorkerFirst(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_workers_script." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	contentDir := t.TempDir()
	contentFile := path.Join(contentDir, "index.js")
	writeContentFile := func(t *testing.T) {
		err := os.WriteFile(contentFile, []byte(`export default { fetch() { return new Response('Hello world'); } };`), 0644)
		if err != nil {
			t.Fatalf("Error creating temp file at path %s: %s", contentFile, err.Error())
		}
	}

	assetsDir := t.TempDir()
	assetFile := path.Join(assetsDir, "index.html")
	writeAssetFile := func(t *testing.T) {
		err := os.WriteFile(assetFile, []byte("Hello world"), 0644)
		if err != nil {
			t.Fatalf("Error creating temp file at path %s: %s", assetFile, err.Error())
		}
	}

	cleanup := func(t *testing.T) {
		for _, file := range []string{assetFile, contentFile} {
			err := os.Remove(file)
			if err != nil {
				t.Logf("Error removing temp file at path %s: %s", file, err.Error())
			}
		}
	}
	defer cleanup(t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					writeContentFile(t)
					writeAssetFile(t)
				},
				Config: testAccCheckCloudflareWorkerScriptConfigWithAssetsWithRunWorkerFirst(rnd, accountID, contentFile, assetsDir, `false`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("config").AtMapKey("run_worker_first"), knownvalue.Bool(false)),
				},
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigWithAssetsWithRunWorkerFirst(rnd, accountID, contentFile, assetsDir, `["/api/*"]`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("config").AtMapKey("run_worker_first"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.StringExact("/api/*"),
					})),
				},
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigWithAssetsWithRunWorkerFirst(rnd, accountID, contentFile, assetsDir, `true`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("config").AtMapKey("run_worker_first"), knownvalue.Bool(true)),
				},
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigWithAssetsWithRunWorkerFirst(rnd, accountID, contentFile, assetsDir, `["/api/*", "!/api/health"]`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("config").AtMapKey("run_worker_first"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.StringExact("/api/*"),
						knownvalue.StringExact("!/api/health"),
					})),
				},
			},
		},
	})
}

func TestAccCloudflareWorkerScript_AssetsConfigRunWorkerFirstMigration(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_workers_script." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	contentDir := t.TempDir()
	contentFile := path.Join(contentDir, "index.js")
	writeContentFile := func(t *testing.T) {
		err := os.WriteFile(contentFile, []byte(`export default { fetch() { return new Response('Hello world'); } };`), 0644)
		if err != nil {
			t.Fatalf("Error creating temp file at path %s: %s", contentFile, err.Error())
		}
	}

	assetsDir := t.TempDir()
	assetFile := path.Join(assetsDir, "index.html")
	writeAssetFile := func(t *testing.T) {
		err := os.WriteFile(assetFile, []byte("Hello world"), 0644)
		if err != nil {
			t.Fatalf("Error creating temp file at path %s: %s", assetFile, err.Error())
		}
	}

	cleanup := func(t *testing.T) {
		for _, file := range []string{assetFile, contentFile} {
			err := os.Remove(file)
			if err != nil {
				t.Logf("Error removing temp file at path %s: %s", file, err.Error())
			}
		}
	}
	defer cleanup(t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					writeContentFile(t)
					writeAssetFile(t)
				},
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.10.1",
					},
				},
				Config: testAccCheckCloudflareWorkerScriptConfigWithAssetsWithRunWorkerFirst(rnd, accountID, contentFile, assetsDir, `false`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("config").AtMapKey("run_worker_first"), knownvalue.Bool(false)),
				},
			},
			{
				Config:                   testAccCheckCloudflareWorkerScriptConfigWithAssetsWithRunWorkerFirst(rnd, accountID, contentFile, assetsDir, `["/api/*"]`),
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("config").AtMapKey("run_worker_first"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.StringExact("/api/*"),
					})),
				},
			},
		},
	})
}

// TestAccCloudflareWorkerScript_UnmanagedSecretNoDrift verifies that unmanaged secrets
// (secrets that exist on the Worker but are not defined in Terraform config) do not cause drift.
// This is a regression test for https://github.com/cloudflare/terraform-provider-cloudflare/issues/5892
func TestAccCloudflareWorkerScript_UnmanagedSecretNoDrift(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	resourceName := resourcePrefix + rnd
	name := "cloudflare_workers_script." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccWorkersScriptConfigWithManagedSecret(resourceName, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(resourceName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("bindings"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("MANAGED_SECRET")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("secret_text")),
				},
			},
			{
				PreConfig: func() {
					// Add an unmanaged secret via the API (out-of-band)
					client := acctest.SharedClient()
					boundary := "--form-data-boundary-unmanaged-secret"
					body := []byte(fmt.Sprintf(`--%s
Content-Disposition: form-data; name="files"; filename="worker.mjs"
Content-Type: application/javascript+module

export default { fetch() { return new Response('Hello world'); } };
--%s
Content-Disposition: form-data; name="metadata"; filename="metadata.json"
Content-Type: application/json

{"main_module": "worker.mjs", "bindings": [{"type": "secret_text", "name": "MANAGED_SECRET", "text": "managed-secret-value"}, {"type": "secret_text", "name": "UNMANAGED_SECRET", "text": "unmanaged-secret-value"}]}
--%s--
`,
						boundary, boundary, boundary,
					))
					result, err := client.Workers.Scripts.Update(context.Background(),
						resourceName,
						workers.ScriptUpdateParams{AccountID: cloudflare.F(accountID)},
						option.WithRequestBody("multipart/form-data;boundary="+boundary, body),
					)
					if err != nil {
						t.Errorf("Error adding unmanaged secret out-of-band: %s", err)
					}
					if result == nil {
						t.Error("Could not add unmanaged secret out-of-band.")
					}
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
				RefreshPlanChecks: resource.RefreshPlanChecks{
					PostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				// Verify state still only contains the managed secret and plan is empty
				Config: testAccWorkersScriptConfigWithManagedSecret(resourceName, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("bindings"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("MANAGED_SECRET")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("secret_text")),
				},
			},
		},
	})
}

// TestAccCloudflareWorkerScript_NoSecretsInConfigWithUnmanagedSecrets verifies that workers with
// unmanaged secrets do not cause drift when config has no secrets at all.
// This is a common scenario where users manage secrets outside of Terraform (e.g., via wrangler or dashboard).
func TestAccCloudflareWorkerScript_NoSecretsInConfigWithUnmanagedSecrets(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	resourceName := resourcePrefix + rnd
	name := "cloudflare_workers_script." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccWorkersScriptConfigNoSecrets(resourceName, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(resourceName)),
				},
			},
			{
				PreConfig: func() {
					// Add secrets via the API (out-of-band) - simulating wrangler secret put
					client := acctest.SharedClient()
					boundary := "--form-data-boundary-unmanaged-secrets"
					body := []byte(fmt.Sprintf(`--%s
Content-Disposition: form-data; name="files"; filename="worker.mjs"
Content-Type: application/javascript+module

export default { fetch() { return new Response('Hello world'); } };
--%s
Content-Disposition: form-data; name="metadata"; filename="metadata.json"
Content-Type: application/json

{"main_module": "worker.mjs", "bindings": [{"type": "secret_text", "name": "API_KEY", "text": "secret-api-key"}, {"type": "secret_text", "name": "DB_PASSWORD", "text": "secret-db-password"}]}
--%s--
`,
						boundary, boundary, boundary,
					))
					result, err := client.Workers.Scripts.Update(context.Background(),
						resourceName,
						workers.ScriptUpdateParams{AccountID: cloudflare.F(accountID)},
						option.WithRequestBody("multipart/form-data;boundary="+boundary, body),
					)
					if err != nil {
						t.Errorf("Error adding secrets out-of-band: %s", err)
					}
					if result == nil {
						t.Error("Could not add secrets out-of-band.")
					}
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
				RefreshPlanChecks: resource.RefreshPlanChecks{
					PostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				// Verify we can still apply the same config without issues
				Config: testAccWorkersScriptConfigNoSecrets(resourceName, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func testAccWorkersScriptConfigWithManagedSecret(rnd, accountID string) string {
	return acctest.LoadTestCase("module_with_unmanaged_secret.tf", rnd, accountID)
}

func testAccWorkersScriptConfigNoSecrets(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_workers_script" "%[1]s" {
  account_id = "%[2]s"
  script_name = "%[1]s"
  content = "export default { fetch() { return new Response('Hello world'); } };"
  main_module = "worker.mjs"
}
`, rnd, accountID)
}

func testAccCheckCloudflareWorkerScriptConfigServiceWorkerInitial(rnd, accountID string) string {
	return acctest.LoadTestCase("service_worker_initial.tf", rnd, scriptContent1, accountID)
}

func testAccCheckCloudflareWorkerScriptConfigServiceWorkerUpdate(rnd, accountID string) string {
	return acctest.LoadTestCase("service_worker_update.tf", rnd, scriptContent2, accountID)
}

func testAccCheckCloudflareWorkerScriptConfigServiceWorkerUpdateBinding(rnd, accountID string) string {
	return acctest.LoadTestCase("service_worker_update_binding.tf", rnd, scriptContent2, encodedWasm, accountID)
}

func testAccCheckCloudflareWorkerScriptUploadModule(rnd, accountID string) string {
	// Use non-migration template to support Terraform < 1.11
	return acctest.LoadTestCase("module_no_migrations.tf", rnd, moduleContent, accountID, compatibilityDate, strings.Join(compatibilityFlags, `","`))
}

func testAccWorkersScriptConfigWithContentFile(rnd, accountID, contentFile string) string {
	return acctest.LoadTestCase("module_with_content_file.tf", rnd, accountID, contentFile)
}

func testAccWorkersScriptConfigWithInvalidContentSHA256(rnd, accountID, contentFile string) string {
	return acctest.LoadTestCase("module_with_invalid_content_sha256.tf", rnd, accountID, contentFile)
}

func testAccWorkersScriptConfigWithAssets(rnd, accountID, assetsDir string) string {
	return acctest.LoadTestCase("module_with_assets.tf", rnd, accountID, assetsDir)
}

func testAccCheckCloudflareWorkerScriptConfigWithAssetsWithRunWorkerFirst(rnd, accountID, contentFile, assetsDir, runWorkerFirst string) string {
	return acctest.LoadTestCase("module_with_assets_with_run_worker_first.tf", rnd, accountID, contentFile, assetsDir, runWorkerFirst)
}
