# Excluding *.example.com from WARP routes
resource "cloudflare_split_tunnel" "example_split_tunnel_exclude" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  mode       = "exclude"
  tunnels {
    host        = "*.example.com"
    description = "example domain"
  }
}

# Including *.example.com in WARP routes
resource "cloudflare_split_tunnel" "example_split_tunnel_include" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  mode       = "include"
  tunnels {
    host        = "*.example.com"
    description = "example domain"
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

# Excluding *.example.com from WARP routes for a particular device policy
resource "cloudflare_split_tunnel" "example_device_settings_policy_split_tunnel_exclude" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  policy_id  = cloudflare_device_settings_policy.developer_warp_policy.id
  mode       = "exclude"
  tunnels {
    host        = "*.example.com"
    description = "example domain"
  }
}

# Including *.example.com in WARP routes for a particular device policy
resource "cloudflare_split_tunnel" "example_split_tunnel_include" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  policy_id  = cloudflare_device_policy.developer_warp_policy.id
  mode       = "include"
  tunnels {
    host        = "*.example.com"
    description = "example domain"
  }
}
