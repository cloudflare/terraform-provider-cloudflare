# Use DNS servers 192.0.2.0 or 192.0.2.1 for example.com
resource "cloudflare_fallback_domain" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  domains {
    suffix      = "example.com"
    description = "Example domain"
    dns_server  = ["192.0.2.0", "192.0.2.1"]
  }
}

# Explicitly adding example.com to the default entries.
resource "cloudflare_fallback_domain" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  dynamic "domains" {
    for_each = toset(["intranet", "internal", "private", "localdomain", "domain", "lan", "home", "host", "corp", "local", "localhost", "home.arpa", "invalid", "test"])
    content {
      suffix = domains.value
    }
  }

  domains {
    suffix      = "example.com"
    description = "Example domain"
    dns_server  = ["192.0.2.0", "192.0.2.1"]
  }
}

# Create a device policy
resource "cloudflare_device_settings_policy" "developer_warp_policy" {
  account_id    = "f037e56e89293a057740de681ac9abbe"
  name          = "Developers"
  precedence    = 10
  match         = "any(identity.groups.name[*] in {\"Developers\"})"
  switch_locked = true
}

# Use DNS servers 192.0.2.0 or 192.0.2.1 for example.com for a particular device policy
resource "cloudflare_fallback_domain" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  policy_id  = cloudflare_device_settings_policy.developer_warp_policy.id
  domains {
    suffix      = "example.com"
    description = "Example domain"
    dns_server  = ["192.0.2.0", "192.0.2.1"]
  }
}
