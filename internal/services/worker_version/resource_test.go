package worker_version_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareWorkerVersion_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_worker." + rnd
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
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
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
