package snippet_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/snippets"
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
	resource.AddTestSweepers("cloudflare_snippet", &resource.Sweeper{
		Name: "cloudflare_snippet",
		F:    testSweepCloudflareSnippets,
	})
}

func testSweepCloudflareSnippets(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	if zoneID == "" {
		tflog.Info(ctx, "Skipping snippets sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	// List all snippets in the zone
	list, err := client.Snippets.List(ctx, snippets.SnippetListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list snippets: %s", err))
		return fmt.Errorf("failed to list snippets: %w", err)
	}

	hasSnippets := false
	// Delete all test snippets in the zone
	for list != nil {
		for _, snippet := range list.Result {
			hasSnippets = true
			// Use standard filtering helper
			if !utils.ShouldSweepResource(snippet.SnippetName) {
				continue
			}

			tflog.Info(ctx, fmt.Sprintf("Deleting snippet: %s (zone: %s)", snippet.SnippetName, zoneID))
			_, err := client.Snippets.Delete(ctx, snippet.SnippetName, snippets.SnippetDeleteParams{
				ZoneID: cloudflare.F(zoneID),
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete snippet %s: %s", snippet.SnippetName, err))
				continue
			}
			tflog.Info(ctx, fmt.Sprintf("Deleted snippet: %s", snippet.SnippetName))
		}

		list, err = list.GetNextPage()
		if err != nil {
			break
		}
	}

	if !hasSnippets {
		tflog.Info(ctx, "No snippets to sweep")
	}

	return nil
}

func TestAccCloudflareSnippet_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_snippet." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareSnippetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareSnippetConfig(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"name": knownvalue.StringExact("main.js"),
							"content": knownvalue.StringExact(`export default {
  async fetch(request) {
    // Get the current timestamp
    const timestamp = Date.now();
    // Convert the timestamp to hexadecimal format
    const hexTimestamp = timestamp.toString(16);
    // Clone the request and add the custom header
    const modifiedRequest = new Request(request, {
        headers: new Headers(request.headers)
    });
    modifiedRequest.headers.set("X-Hex-Timestamp", hexTimestamp);
    // Pass the modified request to the origin
    const response = await fetch(modifiedRequest);
    return response;
  },
}
`),
						}),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"main_module": knownvalue.StringExact("main.js"),
					})),
					// Verify computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccCloudflareSnippetConfigUpdate(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"name": knownvalue.StringExact("main.js"),
							"content": knownvalue.StringExact(`export default {
  async fetch(request) {
    return new Response('Hello, World!');
  }
}
`),
						}),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"main_module": knownvalue.StringExact("main.js"),
					})),
					// Verify computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func TestAccCloudflareSnippet_Import(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_snippet." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareSnippetDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create the snippet
			{
				Config: testAccCloudflareSnippetConfig(rnd, zoneID),
			},
			// Step 2: Import and verify state matches config.
			// This covers the nil-pointer dereference fix: ImportState
			// leaves Metadata nil, then Read must handle that without
			// crashing when it populates Metadata.MainModule from the
			// Content.Get response header.
			{
				Config:            testAccCloudflareSnippetConfig(rnd, zoneID),
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateId:     zoneID + "/" + rnd,
				ImportStateVerify: true,
			},
			// Step 3: Plan after import produces no diff, confirming
			// that Read fully populated files and metadata from the
			// Content.Get endpoint.
			{
				Config: testAccCloudflareSnippetConfig(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
		},
	})
}

func TestAccCloudflareSnippet_ImportFieldValidation(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareSnippetDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create with the basic config
			{
				Config: testAccCloudflareSnippetConfig(rnd, zoneID),
			},
			// Step 2: Import with field-level validation
			{
				ResourceName: "cloudflare_snippet." + rnd,
				ImportState:  true,
				ImportStateId: zoneID + "/" + rnd,
				ImportStateCheck: func(states []*terraform.InstanceState) error {
					if len(states) != 1 {
						return fmt.Errorf("expected 1 instance state, got %d", len(states))
					}
					s := states[0]

					if v := s.Attributes["zone_id"]; v != zoneID {
						return fmt.Errorf("expected zone_id %q, got %q", zoneID, v)
					}
					if v := s.Attributes["snippet_name"]; v != rnd {
						return fmt.Errorf("expected snippet_name %q, got %q", rnd, v)
					}
					if v := s.Attributes["metadata.main_module"]; v != "main.js" {
						return fmt.Errorf("expected metadata.main_module %q, got %q", "main.js", v)
					}
					if v := s.Attributes["files.#"]; v != "1" {
						return fmt.Errorf("expected 1 file, got %s", v)
					}
					if v := s.Attributes["files.0.name"]; v != "main.js" {
						return fmt.Errorf("expected files.0.name %q, got %q", "main.js", v)
					}
					if v := s.Attributes["files.0.content"]; v == "" {
						return fmt.Errorf("expected files.0.content to be non-empty")
					}
					if v := s.Attributes["created_on"]; v == "" {
						return fmt.Errorf("expected created_on to be set")
					}
					if v := s.Attributes["modified_on"]; v == "" {
						return fmt.Errorf("expected modified_on to be set")
					}
					return nil
				},
			},
		},
	})
}

func TestAccCloudflareSnippet_ImportAfterUpdate(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_snippet." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareSnippetDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create
			{
				Config: testAccCloudflareSnippetConfig(rnd, zoneID),
			},
			// Step 2: Update content
			{
				Config: testAccCloudflareSnippetConfigUpdate(rnd, zoneID),
			},
			// Step 3: Import the updated resource
			{
				Config:            testAccCloudflareSnippetConfigUpdate(rnd, zoneID),
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateId:     zoneID + "/" + rnd,
				ImportStateVerify: true,
			},
			// Step 4: Plan after import is clean
			{
				Config: testAccCloudflareSnippetConfigUpdate(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
		},
	})
}

func testAccCheckCloudflareSnippetDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_snippet" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		snippetName := rs.Primary.Attributes["snippet_name"]

		_, err := client.Snippets.Get(context.Background(), snippetName, snippets.SnippetGetParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err == nil {
			return fmt.Errorf("snippet still exists")
		}
	}

	return nil
}

func testAccCloudflareSnippetConfig(rnd, zoneID string) string {
	return acctest.LoadTestCase("basic.tf", rnd, zoneID)
}

func testAccCloudflareSnippetConfigUpdate(rnd, zoneID string) string {
	return acctest.LoadTestCase("basic_update.tf", rnd, zoneID)
}

func TestAccUpgradeSnippet_FromPublishedV5(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	config := testAccCloudflareSnippetConfig(rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.16.0",
					},
				},
				Config: config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}
