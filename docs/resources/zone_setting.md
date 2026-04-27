---
page_title: "cloudflare_zone_setting Resource - Cloudflare"
subcategory: ""
description: |-
  Accepted Permissions
  Zone Settings ReadZone Settings Write
---

# cloudflare_zone_setting (Resource)

Accepted Permissions

- `Zone Settings Read`
- `Zone Settings Write`

-> If using the `ssl_recommender` zone setting, use the `enabled` attribute instead of `value`.

## Available Setting IDs

The following table lists all valid `setting_id` values that can be used with this resource. Some settings may require a specific Cloudflare plan level. For more details, see the [API documentation](https://developers.cloudflare.com/api/resources/zones/subresources/settings/methods/list/).

| Setting ID | Value Type | Description |
|---|---|---|
| `0rtt` | `"on"` / `"off"` | 0-RTT (Zero Round Trip Time) session resumption |
| `advanced_ddos` | `"on"` / `"off"` | Advanced DDoS protection (Enterprise) |
| `aegis` | Special | Aegis protection |
| `always_online` | `"on"` / `"off"` | Always Online - serve stale content when origin is unreachable |
| `always_use_https` | `"on"` / `"off"` | Redirect all HTTP requests to HTTPS |
| `automatic_https_rewrites` | `"on"` / `"off"` | Rewrite HTTP links to HTTPS in page content |
| `automatic_platform_optimization` | Object | Automatic Platform Optimization for WordPress |
| `brotli` | `"on"` / `"off"` | Brotli compression |
| `browser_cache_ttl` | Number | Browser cache TTL in seconds |
| `browser_check` | `"on"` / `"off"` | Browser Integrity Check |
| `cache_level` | `"bypass"`, `"basic"`, `"simplified"`, `"aggressive"` | Cache level |
| `challenge_ttl` | Number | Challenge TTL in seconds (300-2592000) |
| `china_network_enabled` | Special | China Network access (Enterprise) |
| `ciphers` | Array of strings | Allowed cipher suites |
| `cname_flattening` | Special | CNAME Flattening |
| `content_converter` | Special | Content converter |
| `development_mode` | `"on"` / `"off"` | Development Mode (temporarily bypass cache) |
| `early_hints` | `"on"` / `"off"` | Early Hints |
| `edge_cache_ttl` | Number | Edge cache TTL in seconds |
| `email_obfuscation` | `"on"` / `"off"` | Email address obfuscation |
| `h2_prioritization` | `"on"`, `"off"`, `"custom"` | HTTP/2 Edge Prioritization |
| `hotlink_protection` | `"on"` / `"off"` | Hotlink protection |
| `http2` | `"on"` / `"off"` | HTTP/2 |
| `http3` | `"on"` / `"off"` | HTTP/3 (with QUIC) |
| `image_resizing` | `"on"`, `"off"`, `"open"` | Image Resizing |
| `ip_geolocation` | `"on"` / `"off"` | IP Geolocation header |
| `ipv6` | `"on"` / `"off"` | IPv6 compatibility |
| `max_upload` | Number | Max upload size in MB (plan-dependent) |
| `min_tls_version` | `"1.0"`, `"1.1"`, `"1.2"`, `"1.3"` | Minimum TLS version |
| `mirage` | `"on"` / `"off"` | Mirage image optimization |
| `nel` | Object (`{enabled: bool}`) | Network Error Logging |
| `opportunistic_encryption` | `"on"` / `"off"` | Opportunistic Encryption |
| `opportunistic_onion` | `"on"` / `"off"` | Onion Routing |
| `orange_to_orange` | `"on"` / `"off"` | Orange to Orange (O2O) |
| `origin_error_page_pass_thru` | `"on"` / `"off"` | Pass-through origin error pages (Enterprise) |
| `origin_h2_max_streams` | Number | Origin HTTP/2 max concurrent streams |
| `origin_max_http_version` | String | Maximum HTTP version used for origin connections |
| `polish` | `"off"`, `"lossless"`, `"lossy"` | Image optimization (Polish) |
| `prefetch_preload` | `"on"` / `"off"` | Prefetch preload |
| `privacy_pass` | Special | Privacy Pass support |
| `proxy_read_timeout` | Number | Proxy read timeout in seconds (Enterprise) |
| `pseudo_ipv4` | `"off"`, `"add_header"`, `"overwrite_header"` | Pseudo IPv4 |
| `redirects_for_ai_training` | Special | Redirects for AI training |
| `replace_insecure_js` | Special | Replace insecure JavaScript |
| `response_buffering` | `"on"` / `"off"` | Response buffering (Enterprise) |
| `rocket_loader` | `"on"` / `"off"` | Rocket Loader |
| `security_header` | Object | HTTP Strict Transport Security (HSTS) settings |
| `security_level` | `"off"`, `"essentially_off"`, `"low"`, `"medium"`, `"high"`, `"under_attack"` | Security Level |
| `server_side_exclude` | `"on"` / `"off"` | Server Side Excludes |
| `sha1_support` | Special | SHA-1 certificate support |
| `sort_query_string_for_cache` | `"on"` / `"off"` | Sort query string for cache |
| `ssl` | `"off"`, `"flexible"`, `"full"`, `"strict"` | SSL/TLS encryption mode |
| `ssl_recommender` | Uses `enabled` attr | SSL/TLS Recommender (use `enabled` instead of `value`) |
| `tls_1_2_only` | Special | TLS 1.2 only |
| `tls_1_3` | `"on"`, `"off"`, `"zrt"` | TLS 1.3 |
| `tls_client_auth` | `"on"` / `"off"` | TLS Client Auth |
| `transformations` | Special | URL Transformations |
| `transformations_allowed_origins` | Special | Transformations allowed origins |
| `true_client_ip_header` | `"on"` / `"off"` | True-Client-IP header (Enterprise) |
| `waf` | `"on"` / `"off"` | Web Application Firewall (legacy) |
| `webp` | `"on"` / `"off"` | WebP image format |
| `websockets` | `"on"` / `"off"` | WebSockets |

## Example Usage

```terraform
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
```

### Additional Examples

#### String Value with Choices
```terraform
# Minimum TLS Version
resource "cloudflare_zone_setting" "min_tls" {
  zone_id    = "023e105f4ecef8ad9ca31a8372d0c353"
  setting_id = "min_tls_version"
  value      = "1.2"  # Options: "1.0", "1.1", "1.2", "1.3"
}

# SSL/TLS Mode
resource "cloudflare_zone_setting" "ssl" {
  zone_id    = "023e105f4ecef8ad9ca31a8372d0c353"
  setting_id = "ssl"
  value      = "strict"  # Options: "off", "flexible", "full", "strict"
}

# Security Level
resource "cloudflare_zone_setting" "security_level" {
  zone_id    = "023e105f4ecef8ad9ca31a8372d0c353"
  setting_id = "security_level"
  value      = "medium"  # Options: "off", "essentially_off", "low", "medium", "high", "under_attack"
}

# Cache Level
resource "cloudflare_zone_setting" "cache_level" {
  zone_id    = "023e105f4ecef8ad9ca31a8372d0c353"
  setting_id = "cache_level"
  value      = "aggressive"  # Options: "bypass", "basic", "simplified", "aggressive"
}
```

#### Numeric Values
```terraform
# Browser Cache TTL
resource "cloudflare_zone_setting" "browser_cache_ttl" {
  zone_id    = "023e105f4ecef8ad9ca31a8372d0c353"
  setting_id = "browser_cache_ttl"
  value      = 14400  # Seconds (4 hours). Common values: 30, 60, 120, 300, 1200, 1800, 3600, 7200, 10800, 14400, 18000, 28800, 43200, 57600, 72000, 86400, 172800, 259200, 345600, 432000, 691200, 1382400, 2073600, 2678400, 5356800, 16070400, 31536000
}

# Challenge TTL
resource "cloudflare_zone_setting" "challenge_ttl" {
  zone_id    = "023e105f4ecef8ad9ca31a8372d0c353"
  setting_id = "challenge_ttl"
  value      = 1800  # Seconds (30 minutes). Range: 300-2592000
}

# Max Upload Size
resource "cloudflare_zone_setting" "max_upload" {
  zone_id    = "023e105f4ecef8ad9ca31a8372d0c353"
  setting_id = "max_upload"
  value      = 100  # MB. Range: 1-5000 (depending on plan)
}
```

#### Special Cases
```terraform
# 0-RTT (Zero Round Trip Time)
resource "cloudflare_zone_setting" "zero_rtt" {
  zone_id    = "023e105f4ecef8ad9ca31a8372d0c353"
  setting_id = "0rtt"
  value      = "on"
}

# Network Error Logging (NEL)
resource "cloudflare_zone_setting" "nel" {
  zone_id    = "023e105f4ecef8ad9ca31a8372d0c353"
  setting_id = "nel"
  value = {
    enabled = true
  }
}
```

### Common Configuration Sets

#### Security Hardening Configuration
```terraform
# Enable HTTPS everywhere
resource "cloudflare_zone_setting" "always_use_https" {
  zone_id    = var.zone_id
  setting_id = "always_use_https"
  value      = "on"
}

# Automatic HTTPS Rewrites
resource "cloudflare_zone_setting" "automatic_https_rewrites" {
  zone_id    = var.zone_id
  setting_id = "automatic_https_rewrites"
  value      = "on"
}

# Minimum TLS 1.2
resource "cloudflare_zone_setting" "min_tls_version" {
  zone_id    = var.zone_id
  setting_id = "min_tls_version"
  value      = "1.2"
}

# Enable TLS 1.3
resource "cloudflare_zone_setting" "tls_1_3" {
  zone_id    = var.zone_id
  setting_id = "tls_1_3"
  value      = "on"
}

# Strict SSL
resource "cloudflare_zone_setting" "ssl" {
  zone_id    = var.zone_id
  setting_id = "ssl"
  value      = "strict"
}
```

#### Performance Optimization Configuration
```terraform
# Enable HTTP/3
resource "cloudflare_zone_setting" "http3" {
  zone_id    = var.zone_id
  setting_id = "http3"
  value      = "on"
}

# Enable Brotli Compression
resource "cloudflare_zone_setting" "brotli" {
  zone_id    = var.zone_id
  setting_id = "brotli"
  value      = "on"
}

# Early Hints
resource "cloudflare_zone_setting" "early_hints" {
  zone_id    = var.zone_id
  setting_id = "early_hints"
  value      = "on"
}

# Aggressive Caching
resource "cloudflare_zone_setting" "cache_level" {
  zone_id    = var.zone_id
  setting_id = "cache_level"
  value      = "aggressive"
}

# Browser Cache TTL
resource "cloudflare_zone_setting" "browser_cache" {
  zone_id    = var.zone_id
  setting_id = "browser_cache_ttl"
  value      = 14400  # 4 hours
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `setting_id` (String) Setting name
- `value` (Dynamic) Current value of the zone setting.

### Optional

- `enabled` (Boolean) ssl-recommender enrollment setting.
- `zone_id` (String) Identifier

### Read-Only

- `editable` (Boolean) Whether or not this setting can be modified for this zone (based on your Cloudflare plan level).
- `id` (String) Setting name
- `modified_on` (String) last time this setting was modified.
- `time_remaining` (Number) Value of the zone setting.
Notes: The interval (in seconds) from when development mode expires (positive integer) or last expired (negative integer) for the domain. If development mode has never been enabled, this value is false.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_zone_setting.example '<zone_id>/<setting_id>'
```

