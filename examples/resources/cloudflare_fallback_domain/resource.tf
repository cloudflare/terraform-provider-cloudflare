# Use DNS servers 192.0.2.0 or 192.0.2.1 for example.com
resource "cloudflare_fallback_domain" "example" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  domains {
    suffix      = "example.com"
    description = "Example domain"
    dns_server  = ["192.0.2.0", "192.0.2.1"]
  }
}

# Explicitly adding example.com to the default entries.
resource "cloudflare_fallback_domain" "example" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
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
resource "cloudflare_device_policy" "developer_warp_policy" {
  account_id    = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name          = "Developers"
  precedence    = 10
  match         = "any(identity.groups.name[*] in {\"Developers\"})"
  switch_locked = true
}

# Use DNS servers 192.0.2.0 or 192.0.2.1 for example.com for a particular device policy
resource "cloudflare_fallback_domain" "example" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  policy_id  = cloudflare_device_policy.developer_warp_policy.id
  domains {
    suffix      = "example.com"
    description = "Example domain"
    dns_server  = ["192.0.2.0", "192.0.2.1"]
  }
}
