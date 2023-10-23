package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareAPIShieldSchema_Create(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd, rnd2 := generateRandomResourceName(), generateRandomResourceName()
	resourceID := "cloudflare_api_shield_schema." + rnd
	resourceID2 := "cloudflare_api_shield_schema." + rnd2
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAPIShieldSchemasAreDeleted,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShieldSchema(rnd, zoneID, "myschema", testAPIShieldFixtureSchema),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "name", "myschema"),
					resource.TestCheckResourceAttr(resourceID, "kind", "openapi_v3"),
					// validation_enabled is not explicitly defined in resource, but should be false
					resource.TestCheckResourceAttr(resourceID, "validation_enabled", "false"),
					resource.TestCheckResourceAttrWith(resourceID, "source", func(value string) error {
						// remove trailing whitespace from template
						value = strings.TrimSpace(value)
						if value != testAPIShieldFixtureSchema {
							return fmt.Errorf("expected source to be: %v but got: %v", testAPIShieldFixtureSchema, value)
						}
						return nil
					}),
				),
			},
			// check new resource with different ID (resourceID2) with optional parameter "validation_enabled" set to true
			{
				Config: testAccCloudflareAPIShieldSchemaValidationEnabled(rnd2, zoneID, "myschema", testAPIShieldFixtureSchema, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID2, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID2, "name", "myschema"),
					resource.TestCheckResourceAttr(resourceID2, "kind", "openapi_v3"),
					resource.TestCheckResourceAttr(resourceID2, "validation_enabled", "true"),
					resource.TestCheckResourceAttrWith(resourceID2, "source", func(value string) error {
						// remove trailing whitespace from template
						value = strings.TrimSpace(value)
						if value != testAPIShieldFixtureSchema {
							return fmt.Errorf("expected source to be: %v but got: %v", testAPIShieldFixtureSchema, value)
						}
						return nil
					}),
				),
			},
		},
	})
}

func TestAccCloudflareAPIShieldSchema_CreateForceNew(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	resourceID := "cloudflare_api_shield_schema." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	var previousSchemaID string
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAPIShieldSchemasAreDeleted,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShieldSchema(rnd, zoneID, "myschema", testAPIShieldFixtureSchema),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceID, "id", func(value string) error {
						previousSchemaID = value
						return nil
					}),
				),
			},
			{
				// changing the name should force a new schema
				Config: testAccCloudflareAPIShieldSchema(rnd, zoneID, "myschema2", testAPIShieldFixtureSchema),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, "name", "myschema2"),
					resource.TestCheckResourceAttrWith(resourceID, "id", func(value string) error {
						if value == previousSchemaID {
							return fmt.Errorf("expected schema ID to have changed")
						}
						previousSchemaID = value
						return nil
					}),
				),
			},
			{
				// changing the source should force a new schema
				Config: testAccCloudflareAPIShieldSchema(rnd, zoneID, "myschema2", testAPIShieldFixtureSchema2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceID, "source", func(value string) error {
						// remove trailing whitespace from template
						value = strings.TrimSpace(value)
						if value != testAPIShieldFixtureSchema2 {
							return fmt.Errorf("expected source to be: %v but got: %v", testAPIShieldFixtureSchema2, value)
						}
						return nil
					}),
					resource.TestCheckResourceAttrWith(resourceID, "id", func(value string) error {
						if value == previousSchemaID {
							return fmt.Errorf("expected schema ID to have changed")
						}
						previousSchemaID = value
						return nil
					}),
				),
			},
		},
	})
}

func TestAccCloudflareAPIShieldSchema_Update(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	resourceID := "cloudflare_api_shield_schema." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	var schemaID string
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAPIShieldSchemasAreDeleted,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShieldSchemaValidationEnabled(rnd, zoneID, "myschema", testAPIShieldFixtureSchema, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, "validation_enabled", "false"),
					resource.TestCheckResourceAttrWith(resourceID, "id", func(value string) error {
						schemaID = value
						return nil
					}),
				),
			},
			{
				// changing the validation_enabled status to "true" should update the existing schema
				Config: testAccCloudflareAPIShieldSchemaValidationEnabled(rnd, zoneID, "myschema", testAPIShieldFixtureSchema, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, "validation_enabled", "true"),
					resource.TestCheckResourceAttrWith(resourceID, "id", func(value string) error {
						if value != schemaID {
							return fmt.Errorf("expected schema ID to have remained the same")
						}
						return nil
					}),
				),
			},
		},
	})
}

func testAccCheckCloudflareAPIShieldSchemasAreDeleted(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_api_shield_schema" {
			continue
		}

		_, err := client.GetAPIShieldSchema(
			context.Background(),
			cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]),
			cloudflare.GetAPIShieldSchemaParams{
				SchemaID: rs.Primary.Attributes["id"],
			},
		)
		if err == nil {
			return fmt.Errorf("schema still exists")
		}

		var notFoundError *cloudflare.NotFoundError
		if !errors.As(err, &notFoundError) {
			return fmt.Errorf("expected not found error but got: %w", err)
		}
	}

	return nil
}

func testAccCloudflareAPIShieldSchema(resourceName, zone string, name, source string) string {
	return fmt.Sprintf(`
	resource "cloudflare_api_shield_schema" "%[1]s" {
		zone_id = "%[2]s"
		name = "%[3]s"
		source = <<-EOT
%[4]s
EOT
	}
`, resourceName, zone, name, source)
}

func testAccCloudflareAPIShieldSchemaValidationEnabled(resourceName, zone string, name, source string, validationEnabled bool) string {
	return fmt.Sprintf(`
	resource "cloudflare_api_shield_schema" "%[1]s" {
		zone_id = "%[2]s"
		name = "%[3]s"
		validation_enabled = %[4]v
		source = <<-EOT
%[5]s
EOT
	}
`, resourceName, zone, name, validationEnabled, source)
}

const testAPIShieldFixtureSchema = `{
	"openapi": "3.0.2",
	"servers": [
		{
			"url": "https://example.com"
		}
	],
	"paths": {
		"/example/path": {}
	}
}`

const testAPIShieldFixtureSchema2 = `{
	"openapi": "3.0.2",
	"servers": [
		{
			"url": "https://developers.example.com"
		}
	],
	"paths": {
		"/example/path": {}
	}
}`
