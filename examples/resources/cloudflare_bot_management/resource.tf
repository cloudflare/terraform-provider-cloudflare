# Bot Management
resource "cloudflare_bot_management" "sbfm_biz_example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"

  enable_js = true

  sbfm_definitely_automated = "block"
  sbfm_likely_automated = "managed_challenge"
  sbfm_verified_bots = "allow"
  sbfm_static_resource_protection = false
  optimize_wordpress = true
}