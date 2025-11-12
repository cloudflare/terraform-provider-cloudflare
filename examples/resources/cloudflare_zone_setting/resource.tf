# Basic on/off setting
resource "cloudflare_zone_setting" "always_online" {
  zone_id    = "023e105f4ecef8ad9ca31a8372d0c353"
  setting_id = "always_online"
  value      = "on"
}

# String value with specific choices
resource "cloudflare_zone_setting" "min_tls_version" {
  zone_id    = "023e105f4ecef8ad9ca31a8372d0c353"
  setting_id = "min_tls_version"
  value      = "1.2"
}

# Numeric value
resource "cloudflare_zone_setting" "browser_cache_ttl" {
  zone_id    = "023e105f4ecef8ad9ca31a8372d0c353"
  setting_id = "browser_cache_ttl"
  value      = 14400  # 4 hours in seconds
}

# Array/List value
resource "cloudflare_zone_setting" "ciphers" {
  zone_id    = "023e105f4ecef8ad9ca31a8372d0c353"
  setting_id = "ciphers"
  value = [
    "ECDHE-ECDSA-AES128-GCM-SHA256",
    "ECDHE-ECDSA-CHACHA20-POLY1305"
  ]
}

# Nested object value
resource "cloudflare_zone_setting" "security_header" {
  zone_id    = "023e105f4ecef8ad9ca31a8372d0c353"
  setting_id = "security_header"
  value = {
    strict_transport_security = {
      enabled            = true
      include_subdomains = true
      max_age            = 86400
      nosniff            = true
      preload            = false
    }
  }
}

# Special case: ssl_recommender uses 'enabled' instead of 'value'
resource "cloudflare_zone_setting" "ssl_recommender" {
  zone_id    = "023e105f4ecef8ad9ca31a8372d0c353"
  setting_id = "ssl_recommender"
  enabled    = true
}