package worker_version_test

import (
	"os"
	"path"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareWorkerVersionDataSource_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	workerName := "cloudflare_worker." + rnd
	resourceName := "cloudflare_worker_version." + rnd
	dataSourceName := "data.cloudflare_worker_version." + rnd

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
				Config: testAccWorkerVersionDataSourceConfig(rnd, accountID, contentFile),
				ConfigStateChecks: []statecheck.StateCheck{
					// Check the resource was created properly
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),

					// Check the data source fetches the resource correctly
					statecheck.CompareValuePairs(dataSourceName, tfjsonpath.New("version_id"), resourceName, tfjsonpath.New("id"), compare.ValuesSame()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.CompareValuePairs(dataSourceName, tfjsonpath.New("worker_id"), workerName, tfjsonpath.New("id"), compare.ValuesSame()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("main_module"), knownvalue.StringExact("index.js")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("modules"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"name":           knownvalue.StringExact("index.js"),
							"content_type":   knownvalue.StringExact("application/javascript+module"),
							"content_base64": knownvalue.StringExact("ZXhwb3J0IGRlZmF1bHQge2ZldGNoKCkge3JldHVybiBuZXcgUmVzcG9uc2UoKX19"),
						}),
					})),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("compatibility_date"), knownvalue.Null()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("assets"), knownvalue.Null()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("migrations"), knownvalue.Null()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("placement"), knownvalue.Null()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("usage_model"), knownvalue.StringExact("standard")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("compatibility_flags"), knownvalue.Null()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("annotations"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"workers_message":      knownvalue.Null(),
						"workers_tag":          knownvalue.Null(),
						"workers_triggered_by": knownvalue.NotNull(),
					})),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("bindings"), knownvalue.Null()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("limits"), knownvalue.Null()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("number"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("source"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func testAccWorkerVersionDataSourceConfig(rnd, accountID, contentFile string) string {
	return acctest.LoadTestCase("datasource_basic.tf", rnd, accountID, contentFile)
}
