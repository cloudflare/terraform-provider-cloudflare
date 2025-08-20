resource "cloudflare_zone" "example_com" {
  account_id = var.account_id
  paused     = false
  plan       = "enterprise"
  type       = "full"
  zone       = "example.com"
}

module "zone_settings" {
  source         = "../zone_settings"
  zone_id        = cloudflare_zone.example_com.id
  security_level = "high"
  ssl            = "origin_pull"
}
