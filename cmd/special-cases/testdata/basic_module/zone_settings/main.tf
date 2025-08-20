resource "cloudflare_zone_settings_override" "zone_settings" {
  zone_id = var.zone_id

  settings {
    always_use_https         = var.always_use_https
    automatic_https_rewrites = var.automatic_https_rewrites
    browser_check            = var.browser_check
    email_obfuscation        = var.email_obfuscation
    ipv6                     = var.ipv6
    min_tls_version          = var.min_tls_version
    security_level           = var.security_level
    ssl                      = var.ssl
    websockets               = var.websockets

    nel {
      enabled = var.enable_network_error_logging
    }

    security_header {
      enabled            = var.security_header_enabled
      include_subdomains = var.security_header_include_subdomains
      max_age            = var.security_header_max_age
      nosniff            = var.security_header_nosniff
      preload            = var.security_header_preload
    }
  }
}
