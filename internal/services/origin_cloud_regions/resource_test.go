package origin_cloud_regions_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	resource.AddTestSweepers("cloudflare_origin_cloud_regions", &resource.Sweeper{
		Name: "cloudflare_origin_cloud_regions",
		F:    testSweepCloudflareOriginCloudRegions,
	})
}

func testSweepCloudflareOriginCloudRegions(r string) error {
	ctx := context.Background()
	// origin_cloud_regions is a zone-level collection managed per-entry.
	// Acceptance tests clean up via destroy; no sweep required.
	tflog.Info(ctx, "origin_cloud_regions does not require sweeping")
	return nil
}

func testOriginCloudRegionsConfig(rnd, zoneID string) string {
	return acctest.LoadTestCase("basic.tf", rnd, zoneID)
}

func testOriginCloudRegionsConfigUpdated(rnd, zoneID string) string {
	return acctest.LoadTestCase("updated.tf", rnd, zoneID)
}

func testOriginCloudRegionsConfigVendorChanged(rnd, zoneID string) string {
	return acctest.LoadTestCase("vendor_changed.tf", rnd, zoneID)
}

func testOriginCloudRegionsConfigEntryRemoved(rnd, zoneID string) string {
	return acctest.LoadTestCase("entry_removed.tf", rnd, zoneID)
}

func testAccCheckOriginCloudRegionsDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_origin_cloud_regions" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]

		res := new(http.Response)
		err := client.Get(
			context.Background(),
			fmt.Sprintf("/zones/%s/cache/origin_public_cloud_region", zoneID),
			nil,
			&res,
		)
		if err != nil {
			continue
		}

		body, _ := io.ReadAll(res.Body)

		var envelope struct {
			Result struct {
				Value []interface{} `json:"value"`
			} `json:"result"`
		}
		if err := json.Unmarshal(body, &envelope); err != nil {
			continue
		}

		if len(envelope.Result.Value) > 0 {
			return fmt.Errorf("origin cloud region mappings still exist for zone %s", zoneID)
		}
	}

	return nil
}

func TestAccCloudflareOriginCloudRegions_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_origin_cloud_regions." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOriginCloudRegionsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testOriginCloudRegionsConfig(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mappings"), knownvalue.SetSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mappings"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"origin_ip": knownvalue.StringExact("192.0.2.1"),
							"vendor":    knownvalue.StringExact("aws"),
							"region":    knownvalue.StringExact("us-east-1"),
						}),
					})),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareOriginCloudRegions_AddEntry(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_origin_cloud_regions." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOriginCloudRegionsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testOriginCloudRegionsConfig(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mappings"), knownvalue.SetSizeExact(1)),
				},
			},
			{
				Config: testOriginCloudRegionsConfigUpdated(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mappings"), knownvalue.SetSizeExact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mappings"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"origin_ip": knownvalue.StringExact("192.0.2.1"),
							"vendor":    knownvalue.StringExact("aws"),
							"region":    knownvalue.StringExact("us-west-2"),
						}),
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"origin_ip": knownvalue.StringExact("2001:db8::1"),
							"vendor":    knownvalue.StringExact("gcp"),
							"region":    knownvalue.StringExact("us-central1"),
						}),
					})),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareOriginCloudRegions_ChangeVendor(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_origin_cloud_regions." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOriginCloudRegionsDestroy,
		Steps: []resource.TestStep{
			{
				// Start: 192.0.2.1 → aws/us-east-1
				Config: testOriginCloudRegionsConfig(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mappings"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"origin_ip": knownvalue.StringExact("192.0.2.1"),
							"vendor":    knownvalue.StringExact("aws"),
							"region":    knownvalue.StringExact("us-east-1"),
						}),
					})),
				},
			},
			{
				// Change: 192.0.2.1 → gcp/us-central1
				Config: testOriginCloudRegionsConfigVendorChanged(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mappings"), knownvalue.SetSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mappings"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"origin_ip": knownvalue.StringExact("192.0.2.1"),
							"vendor":    knownvalue.StringExact("gcp"),
							"region":    knownvalue.StringExact("us-central1"),
						}),
					})),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareOriginCloudRegions_DeleteEntry(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_origin_cloud_regions." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOriginCloudRegionsDestroy,
		Steps: []resource.TestStep{
			{
				// Start: two entries
				Config: testOriginCloudRegionsConfigUpdated(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mappings"), knownvalue.SetSizeExact(2)),
				},
			},
			{
				// Remove 2001:db8::1, leaving only 192.0.2.1 → aws/us-east-1
				Config: testOriginCloudRegionsConfigEntryRemoved(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mappings"), knownvalue.SetSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mappings"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"origin_ip": knownvalue.StringExact("192.0.2.1"),
							"vendor":    knownvalue.StringExact("aws"),
							"region":    knownvalue.StringExact("us-east-1"),
						}),
					})),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
