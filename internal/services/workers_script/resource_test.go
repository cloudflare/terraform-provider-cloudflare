package workers_script_test

import (
	"context"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

const (
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

func TestAccCloudflareWorkerScript_ServiceWorker(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_workers_script." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerScriptConfigServiceWorkerInitial(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("content"), knownvalue.StringExact(scriptContent1)),
				},
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigServiceWorkerUpdate(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("content"), knownvalue.StringExact(scriptContent2)),
				},
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigServiceWorkerUpdateBinding(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("content"), knownvalue.StringExact(scriptContent2)),
				},
			},
			{
				PreConfig: func() {
					client := acctest.SharedClient()
					result, err := client.Workers.Scripts.Settings.Edit(context.Background(), rnd, workers.ScriptSettingEditParams{AccountID: cloudflare.F(accountID), ScriptSetting: workers.ScriptSettingParam{Logpush: cloudflare.Bool(true)}})
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
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_workers_script." + rnd
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
				Config: testAccCheckCloudflareWorkerScriptUploadModule(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
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
					result, err := client.Workers.Scripts.Settings.Edit(context.Background(), rnd, workers.ScriptSettingEditParams{AccountID: cloudflare.F(accountID), ScriptSetting: workers.ScriptSettingParam{Logpush: cloudflare.Bool(true)}})
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
	name := "cloudflare_workers_script." + rnd
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
				Config: testAccWorkersScriptConfigWithContentFile(rnd, accountID, contentFile),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("content_file"), knownvalue.StringExact(contentFile)),
				},
			},
			{
				PreConfig: func() {
					writeContentFile(t, fmt.Sprintf("%s // v2", moduleContent))
				},
				Config: testAccWorkersScriptConfigWithContentFile(rnd, accountID, contentFile),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
			},
			{
				Config: testAccWorkersScriptConfigWithContentFile(rnd, accountID, contentFile),
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
						rnd,
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
				Config:      testAccWorkersScriptConfigWithInvalidContentSHA256(rnd, accountID, contentFile),
				ExpectError: regexp.MustCompile(`SHA-256 Hash Mismatch`),
			},
		},
	})
}

func TestAccCloudflareWorkerScript_PythonWorker(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_workers_script." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("python_worker.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
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

func TestAccCloudflareWorkerScript_ModuleWithDurableObject(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_workers_script." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("module_with_durable_object.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("main_module"), knownvalue.StringExact("worker.js")),
				},
			},
			{
				ResourceName:            name,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"bindings.0.namespace_id", "has_modules", "main_module", "startup_time_ms"},
			},
		},
	})
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
	return acctest.LoadTestCase("module.tf", rnd, moduleContent, accountID, compatibilityDate, strings.Join(compatibilityFlags, `","`))
}

func testAccWorkersScriptConfigWithContentFile(rnd, accountID, contentFile string) string {
	return acctest.LoadTestCase("module_with_content_file.tf", rnd, accountID, contentFile)
}

func testAccWorkersScriptConfigWithInvalidContentSHA256(rnd, accountID, contentFile string) string {
	return acctest.LoadTestCase("module_with_invalid_content_sha256.tf", rnd, accountID, contentFile)
}
