resource "cloudflare_zone_subscription" "%[1]s" {
  zone_id = "%[2]s"

  # Keep enterprise plan but test that we can manage the resource
  rate_plan = {
    id = "enterprise"
  }

  # For enterprise plans, frequency might not be changeable
  # so we're just verifying the resource can be managed
}