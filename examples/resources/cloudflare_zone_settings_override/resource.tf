resource "cloudflare_zone_settings_override" "test" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  settings = [{
    brotli                   = "on"
    challenge_ttl            = 2700
    security_level           = "high"
    opportunistic_encryption = "on"
    automatic_https_rewrites = "on"
    mirage                   = "on"
    waf                      = "on"
    minify = [{
      css  = "on"
      js   = "off"
      html = "off"
    }]
    security_header = [{
      enabled = true
    }]
  }]
}
