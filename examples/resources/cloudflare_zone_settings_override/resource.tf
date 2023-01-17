resource "cloudflare_zone_settings_override" "test" {
  zone_id = d41d8cd98f00b204e9800998ecf8427e
  settings {
    brotli                   = "on"
    challenge_ttl            = 2700
    security_level           = "high"
    opportunistic_encryption = "on"
    automatic_https_rewrites = "on"
    mirage                   = "on"
    waf                      = "on"
    minify {
      css  = "on"
      js   = "off"
      html = "off"
    }
    security_header {
      enabled = true
    }
  }
}
