package worker_version_test

import (
	"context"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_worker_version", &resource.Sweeper{
		Name: "cloudflare_worker_version",
		F:    testSweepCloudflareWorkerVersion,
	})
}

func testSweepCloudflareWorkerVersion(r string) error {
	ctx := context.Background()
	// Worker Version is a worker script-level resource.
	// When worker scripts are swept, versions are cleaned up automatically.
	// No sweeping required.
	tflog.Info(ctx, "Worker Version doesn't require sweeping (worker script resource)")
	return nil
}

func TestAccCloudflareWorkerVersion_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	workerName := "cloudflare_worker." + rnd
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

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					writeContentFile(t, `export default {fetch() {return new Response()}}`)
				},
				Config: testAccCloudflareWorkerVersionConfig(rnd, accountID, contentFile),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.CompareValuePairs(workerName, tfjsonpath.New("id"), resourceName, tfjsonpath.New("worker_id"), compare.ValuesSame()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("main_module"), knownvalue.StringExact("index.js")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modules"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"name":           knownvalue.StringExact("index.js"),
							"content_file":   knownvalue.StringExact(contentFile),
							"content_base64": knownvalue.Null(),
							"content_type":   knownvalue.StringExact("application/javascript+module"),
							"content_sha256": knownvalue.StringExact("e06650aadafc1df60cbf34d68dab2bb20b20d175c9310ed0006169f1a266ef08"),
						}),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("compatibility_date"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("assets"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("migrations"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("placement"), knownvalue.Null()),
					// Verify computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("usage_model"), knownvalue.StringExact("standard")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("compatibility_flags"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("annotations"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"workers_message":      knownvalue.Null(),
						"workers_tag":          knownvalue.Null(),
						"workers_triggered_by": knownvalue.NotNull(),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("limits"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("number"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("source"), knownvalue.NotNull()),
				},
			},
			{
				PreConfig: func() {
					// Update the content file
					writeContentFile(t, `export default {fetch() {return new Response("Hello World!")}}`)
				},
				Config: testAccCloudflareWorkerVersionConfigUpdate(rnd, accountID, contentFile),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.CompareValuePairs(workerName, tfjsonpath.New("id"), resourceName, tfjsonpath.New("worker_id"), compare.ValuesSame()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("main_module"), knownvalue.StringExact("index.js")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modules"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"name":           knownvalue.StringExact("index.js"),
							"content_file":   knownvalue.StringExact(contentFile),
							"content_base64": knownvalue.Null(),
							"content_type":   knownvalue.StringExact("application/javascript+module"),
							"content_sha256": knownvalue.StringExact("abba0df0e36536eb43b5f543dfd4ce55afc9059fa6a400ccaed8002dbcbedb7b"),
						}),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("compatibility_date"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("assets"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("migrations"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"name": knownvalue.StringExact("NODE_ENV"),
							"type": knownvalue.StringExact("plain_text"),
							"text": knownvalue.StringExact("production"),
						}),
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"name": knownvalue.StringExact("VERSION_METADATA"),
							"type": knownvalue.StringExact("version_metadata"),
						}),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("placement"), knownvalue.Null()),
					// Verify computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("usage_model"), knownvalue.StringExact("standard")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("compatibility_flags"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("annotations"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"workers_message":      knownvalue.Null(),
						"workers_tag":          knownvalue.Null(),
						"workers_triggered_by": knownvalue.NotNull(),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("limits"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("number"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("source"), knownvalue.NotNull()),
				},
			},
			{
				PreConfig: func() {
					writeContentFile(t, `export default {fetch() {return new Response("Hello World!")}}`)
				},
				Config: testAccCloudflareWorkerVersionConfigBindingOrder(rnd, accountID, contentFile),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings"), knownvalue.SetExact([]knownvalue.Check{
						// namespace_id should be set for KV bindings, while id
						// should be null.
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"name":         knownvalue.StringExact("KV1"),
							"type":         knownvalue.StringExact("kv_namespace"),
							"id":           knownvalue.Null(),
							"namespace_id": knownvalue.NotNull(),
						}),
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"name":         knownvalue.StringExact("KV2"),
							"type":         knownvalue.StringExact("kv_namespace"),
							"id":           knownvalue.Null(),
							"namespace_id": knownvalue.NotNull(),
						}),
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"name":         knownvalue.StringExact("KV3"),
							"type":         knownvalue.StringExact("kv_namespace"),
							"id":           knownvalue.Null(),
							"namespace_id": knownvalue.NotNull(),
						}),
						// id should be set for D1 bindings, while namespace_id
						// should be null.
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"name":         knownvalue.StringExact("DB1"),
							"type":         knownvalue.StringExact("d1"),
							"id":           knownvalue.NotNull(),
							"namespace_id": knownvalue.Null(),
						}),
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"name":         knownvalue.StringExact("DB2"),
							"type":         knownvalue.StringExact("d1"),
							"id":           knownvalue.NotNull(),
							"namespace_id": knownvalue.Null(),
						}),
					})),
				},
			},
			{
				ResourceName:            resourceName,
				ImportStateIdFunc:       testAccCloudflareWorkerVersionImportStateIdFunc(resourceName, accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modules.0.content_file", "modules.0.content_base64", "bindings"}, // content_file not stored in API, content_base64 populated on import; binding order is different
			},
		},
	})
}

func TestAccCloudflareWorkerVersion_ContentBase64(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	workerName := "cloudflare_worker." + rnd
	resourceName := "cloudflare_worker_version." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWorkerVersionConfigContentBase64(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.CompareValuePairs(workerName, tfjsonpath.New("id"), resourceName, tfjsonpath.New("worker_id"), compare.ValuesSame()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("main_module"), knownvalue.StringExact("index.js")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modules"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"name":           knownvalue.StringExact("index.js"),
							"content_file":   knownvalue.Null(),
							"content_base64": knownvalue.StringExact("ZXhwb3J0IGRlZmF1bHQge2FzeW5jIGZldGNoKCkge3JldHVybiBuZXcgUmVzcG9uc2UoJ0hlbGxvIGZyb20gYmFzZTY0IScpfX0="),
							"content_type":   knownvalue.StringExact("application/javascript+module"),
							"content_sha256": knownvalue.StringExact("d0e1cf792981b29449943830ca7119f5e9c839d0419baf161083747f785d887f"),
						}),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("usage_model"), knownvalue.StringExact("standard")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("number"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("source"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccCloudflareWorkerVersionConfigContentBase64Updated(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modules"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"name":           knownvalue.StringExact("index.js"),
							"content_file":   knownvalue.Null(),
							"content_base64": knownvalue.StringExact("ZXhwb3J0IGRlZmF1bHQge2FzeW5jIGZldGNoKCkge3JldHVybiBuZXcgUmVzcG9uc2UoJ1VwZGF0ZWQgZnJvbSBiYXNlNjQhJyl9fQ=="),
							"content_type":   knownvalue.StringExact("application/javascript+module"),
							"content_sha256": knownvalue.StringExact("c028232f838abb823bee8f749cade1a3b6b1a5ced950cfa902cfb804c7f68e85"),
						}),
					})),
				},
			},
			{
				ResourceName:            resourceName,
				ImportStateIdFunc:       testAccCloudflareWorkerVersionImportStateIdFunc(resourceName, accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"bindings"},
			},
		},
	})
}

func TestAccCloudflareWorkerVersion_WithAssets(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_worker_version." + rnd
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
				Config: testAccCloudflareWorkerVersionConfigWithAssets(rnd, accountID, assetsDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("assets").AtMapKey("directory"), knownvalue.StringExact(assetsDir)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("assets").AtMapKey("asset_manifest_sha256"), knownvalue.StringExact("b098d2898ca7ae5677c7291d97323e7894137515043f3e560f3bd155870eea9e")),
				},
			},
			{
				PreConfig: func() {
					writeAssetFile(t, "v2")
				},
				Config: testAccCloudflareWorkerVersionConfigWithAssets(rnd, accountID, assetsDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("assets").AtMapKey("directory"), knownvalue.StringExact(assetsDir)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("assets").AtMapKey("asset_manifest_sha256"), knownvalue.StringExact("46f07eb8a3fa881af81ce2b6b3fc1627edccf115526aa5c631308c45c75d2fb1")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
			},
			{
				Config: testAccCloudflareWorkerVersionConfigWithAssets(rnd, accountID, assetsDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				ResourceName:      resourceName,
				ImportStateIdFunc: testAccCloudflareWorkerVersionImportStateIdFunc(resourceName, accountID),
				ImportState:       true,
				ImportStateVerify: true,
				// FIXME: handle refreshing assets.config
				ImportStateVerifyIgnore: []string{"assets.%", "assets.directory", "assets.asset_manifest_sha256", "assets.config", "startup_time_ms"},
			},
		},
	})
}

func testAccCloudflareWorkerVersionConfig(rnd, accountID, contentFile string) string {
	return acctest.LoadTestCase("basic.tf", rnd, accountID, contentFile)
}

func testAccCloudflareWorkerVersionConfigUpdate(rnd, accountID, contentFile string) string {
	return acctest.LoadTestCase("basic_update.tf", rnd, accountID, contentFile)
}

func testAccCloudflareWorkerVersionConfigWithAssets(rnd, accountID, assetsDir string) string {
	return acctest.LoadTestCase("assets.tf", rnd, accountID, assetsDir)
}

func testAccCloudflareWorkerVersionConfigBindingOrder(rnd, accountID, contentFile string) string {
	return acctest.LoadTestCase("basic_binding_order.tf", rnd, accountID, contentFile)
}

func TestAccCloudflareWorkerVersion_AssetsConfigRunWorkerFirst(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_worker_version." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	tmpDir := t.TempDir()
	assetsDir := path.Join(tmpDir, "assets")
	err := os.Mkdir(assetsDir, 0755)
	if err != nil {
		t.Fatalf("Error creating assets directory: %s", err.Error())
	}

	assetFile := path.Join(assetsDir, "index.html")
	workerFile := path.Join(tmpDir, "worker.js")

	writeFiles := func(t *testing.T) {
		err := os.WriteFile(assetFile, []byte("Hello world"), 0644)
		if err != nil {
			t.Fatalf("Error creating asset file at path %s: %s", assetFile, err.Error())
		}
		err = os.WriteFile(workerFile, []byte("export default { fetch() { return new Response('Hello from worker'); } };"), 0644)
		if err != nil {
			t.Fatalf("Error creating worker file at path %s: %s", workerFile, err.Error())
		}
	}

	cleanup := func(t *testing.T) {
		for _, file := range []string{assetFile, workerFile} {
			err := os.Remove(file)
			if err != nil {
				t.Logf("Error removing temp file at path %s: %s", file, err.Error())
			}
		}
		err := os.Remove(assetsDir)
		if err != nil {
			t.Logf("Error removing assets directory: %s", err.Error())
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
					writeFiles(t)
				},
				Config: testAccCloudflareWorkerVersionConfigWithAssetsWithRunWorkerFirst(rnd, accountID, assetsDir, workerFile, `false`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("config").AtMapKey("run_worker_first"), knownvalue.Bool(false)),
				},
			},
			{
				Config: testAccCloudflareWorkerVersionConfigWithAssetsWithRunWorkerFirst(rnd, accountID, assetsDir, workerFile, `["/api/*"]`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("config").AtMapKey("run_worker_first"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.StringExact("/api/*"),
					})),
				},
			},
			{
				Config: testAccCloudflareWorkerVersionConfigWithAssetsWithRunWorkerFirst(rnd, accountID, assetsDir, workerFile, `true`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("config").AtMapKey("run_worker_first"), knownvalue.Bool(true)),
				},
			},
			{
				Config: testAccCloudflareWorkerVersionConfigWithAssetsWithRunWorkerFirst(rnd, accountID, assetsDir, workerFile, `["/api/*", "!/api/health"]`),
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

func testAccCloudflareWorkerVersionConfigContentBase64(rnd, accountID string) string {
	return acctest.LoadTestCase("content_base64.tf", rnd, accountID)
}

func testAccCloudflareWorkerVersionConfigContentBase64Updated(rnd, accountID string) string {
	return acctest.LoadTestCase("content_base64_update.tf", rnd, accountID)
}

func testAccCloudflareWorkerVersionImportStateIdFunc(resourceName, accountID string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rnd := resourceName[len("cloudflare_worker_version."):]
		workerResourceName := "cloudflare_worker." + rnd

		worker, ok := s.RootModule().Resources[workerResourceName]
		if !ok {
			return "", fmt.Errorf("worker resource not found: %s", workerResourceName)
		}

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("worker_version resource not found: %s", resourceName)
		}

		return fmt.Sprintf("%s/%s/%s", accountID, worker.Primary.ID, rs.Primary.ID), nil
	}
}

func testAccCloudflareWorkerVersionConfigWithAssetsWithRunWorkerFirst(rnd, accountID, assetsDir, workerFile, runWorkerFirst string) string {
	return acctest.LoadTestCase("assets_with_run_worker_first.tf", rnd, accountID, assetsDir, workerFile, runWorkerFirst)
}

func testAccCloudflareWorkerVersionConfigSensitiveBindings(rnd, accountID, contentFile string) string {
	return acctest.LoadTestCase("sensitive_bindings.tf", rnd, accountID, contentFile)
}

// TestAccCloudflareWorkerVersion_SensitiveBindingsImport tests that importing
// with sensitive bindings (plain_text/secret_text) doesn't force replacement.
func TestAccCloudflareWorkerVersion_SensitiveBindingsImport(t *testing.T) {
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

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					writeContentFile(t, `export default {fetch() {return new Response()}}`)
				},
				Config: testAccCloudflareWorkerVersionConfigSensitiveBindings(rnd, accountID, contentFile),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("annotations"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"workers_message": knownvalue.StringExact("Test import with annotations"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"name": knownvalue.StringExact("PLAIN_TEXT_VAR"),
							"type": knownvalue.StringExact("plain_text"),
							"text": knownvalue.StringExact("plain-text-value"),
						}),
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"name": knownvalue.StringExact("SECRET_VAR"),
							"type": knownvalue.StringExact("secret_text"),
							"text": knownvalue.StringExact("secret-value"),
						}),
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"name": knownvalue.StringExact("VERSION_METADATA"),
							"type": knownvalue.StringExact("version_metadata"),
						}),
					})),
				},
			},
			// Re-apply same config - expect no changes
			{
				Config: testAccCloudflareWorkerVersionConfigSensitiveBindings(rnd, accountID, contentFile),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Import (API won't return sensitive text values)
			{
				ResourceName:            resourceName,
				ImportStateIdFunc:       testAccCloudflareWorkerVersionImportStateIdFunc(resourceName, accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modules.0.content_file", "modules.0.content_base64", "bindings"},
			},
			// After import, applying config should NOT force replacement
			{
				Config: testAccCloudflareWorkerVersionConfigSensitiveBindings(rnd, accountID, contentFile),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"name": knownvalue.StringExact("PLAIN_TEXT_VAR"),
							"type": knownvalue.StringExact("plain_text"),
							"text": knownvalue.StringExact("plain-text-value"),
						}),
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"name": knownvalue.StringExact("SECRET_VAR"),
							"type": knownvalue.StringExact("secret_text"),
							"text": knownvalue.StringExact("secret-value"),
						}),
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"name": knownvalue.StringExact("VERSION_METADATA"),
							"type": knownvalue.StringExact("version_metadata"),
						}),
					})),
				},
			},
		},
	})
}

func TestAccCloudflareWorkerVersion_RunWorkerFirstUpgrade(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_worker_version." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	tmpDir := t.TempDir()
	assetsDir := path.Join(tmpDir, "assets")
	err := os.Mkdir(assetsDir, 0755)
	if err != nil {
		t.Fatalf("Error creating assets directory: %s", err.Error())
	}

	assetFile := path.Join(assetsDir, "index.html")
	workerFile := path.Join(tmpDir, "worker.js")

	writeFiles := func(t *testing.T) {
		err := os.WriteFile(assetFile, []byte("Hello world"), 0644)
		if err != nil {
			t.Fatalf("Error creating asset file at path %s: %s", assetFile, err.Error())
		}
		err = os.WriteFile(workerFile, []byte("export default { fetch() { return new Response('Hello from worker'); } };"), 0644)
		if err != nil {
			t.Fatalf("Error creating worker file at path %s: %s", workerFile, err.Error())
		}
	}

	cleanup := func(t *testing.T) {
		for _, file := range []string{assetFile, workerFile} {
			err := os.Remove(file)
			if err != nil {
				t.Logf("Error removing temp file at path %s: %s", file, err.Error())
			}
		}
		err := os.Remove(assetsDir)
		if err != nil {
			t.Logf("Error removing assets directory: %s", err.Error())
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
					writeFiles(t)
				},
				Config: testAccCloudflareWorkerVersionConfigWithAssetsWithRunWorkerFirst(rnd, accountID, assetsDir, workerFile, `["/api/*"]`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("config").AtMapKey("run_worker_first"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.StringExact("/api/*"),
					})),
				},
			},
			{
				Config: testAccCloudflareWorkerVersionConfigWithAssetsWithRunWorkerFirst(rnd, accountID, assetsDir, workerFile, `true`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("config").AtMapKey("run_worker_first"), knownvalue.Bool(true)),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
			{
				Config: testAccCloudflareWorkerVersionConfigWithAssetsWithRunWorkerFirst(rnd, accountID, assetsDir, workerFile, `["/admin/*", "!/admin/public"]`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("config").AtMapKey("run_worker_first"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.StringExact("/admin/*"),
						knownvalue.StringExact("!/admin/public"),
					})),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
			{
				Config: testAccCloudflareWorkerVersionConfigWithAssetsWithRunWorkerFirst(rnd, accountID, assetsDir, workerFile, `false`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("assets").AtMapKey("config").AtMapKey("run_worker_first"), knownvalue.Bool(false)),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
		},
	})
}
