package zone_tags_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// Example test for Option B - separate zone_tags resource (not runnable, just for demonstration)
func TestAccCloudflareZoneTags_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: `
resource "cloudflare_zone" "example" {
  account = { id = "023e105f4ecef8ad9ca31a8372d0c353" }
  name = "example.com"
  type = "full"
}

resource "cloudflare_zone_tags" "example_tags" {
  zone_id = cloudflare_zone.example.id
  tags = {
    env = "production"
    CostCenter = "4300-systems-engineering"
  }
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_zone_tags.example_tags", "tags.env", "production"),
					resource.TestCheckResourceAttr("cloudflare_zone_tags.example_tags", "tags.CostCenter", "4300-systems-engineering"),
				),
			},
			{
				Config: `
resource "cloudflare_zone" "example" {
  account = { id = "023e105f4ecef8ad9ca31a8372d0c353" }
  name = "example.com"
  type = "full"
}

resource "cloudflare_zone_tags" "example_tags" {
  zone_id = cloudflare_zone.example.id
  tags = {
    env = "staging"
    CostCenter = "4300-systems-engineering"
    Project = "terraform-poc"
  }
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_zone_tags.example_tags", "tags.env", "staging"),
					resource.TestCheckResourceAttr("cloudflare_zone_tags.example_tags", "tags.CostCenter", "4300-systems-engineering"),
					resource.TestCheckResourceAttr("cloudflare_zone_tags.example_tags", "tags.Project", "terraform-poc"),
				),
			},
		},
	})
}

// Test for import functionality
func TestAccCloudflareZoneTags_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: `
resource "cloudflare_zone" "example" {
  account = { id = "023e105f4ecef8ad9ca31a8372d0c353" }
  name = "example.com"
  type = "full"
}

resource "cloudflare_zone_tags" "example_tags" {
  zone_id = cloudflare_zone.example.id
  tags = {
    env = "production"
    CostCenter = "4300-systems-engineering"
  }
}`,
			},
			{
				ResourceName:      "cloudflare_zone_tags.example_tags",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources["cloudflare_zone.example"]
					if !ok {
						return "", fmt.Errorf("not found: cloudflare_zone.example")
					}
					return rs.Primary.ID, nil
				},
			},
		},
	})
}