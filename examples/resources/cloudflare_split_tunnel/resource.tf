# Excluding *.example.com from WARP routes
resource "cloudflare_split_tunnel" "example_split_tunnel_exclude" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  mode       = "exclude"
  tunnels {
    host        = "*.example.com"
    description = "example domain"
  }
}

# Including *.example.com in WARP routes
resource "cloudflare_split_tunnel" "example_split_tunnel_include" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  mode       = "include"
  tunnels {
    host        = "*.example.com"
    description = "example domain"
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

# Excluding *.example.com from WARP routes for a particular device policy
resource "cloudflare_split_tunnel" "example_device_policy_split_tunnel_exclude" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  policy_id  = cloudflare_device_policy.developer_warp_policy.id
  mode       = "exclude"
  tunnels {
    host        = "*.example.com"
    description = "example domain"
  }
}

# Including *.example.com in WARP routes for a particular device policy
resource "cloudflare_split_tunnel" "example_split_tunnel_include" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  policy_id  = cloudflare_device_policy.developer_warp_policy.id
  mode       = "include"
  tunnels {
    host        = "*.example.com"
    description = "example domain"
  }
}
