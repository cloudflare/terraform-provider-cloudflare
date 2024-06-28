data "cloudflare_api_token_permission_groups" "all" {}

# Get zone level DNS read permission ID.
output "dns_read_permission_id" {
  value = data.cloudflare_api_token_permission_groups.all.zone["DNS Read"] // 82e64a83756745bbbb1c9c2701bf816b
}

# Get account level "Load Balancing: Monitors and Pools Read" permission ID.
output "account_lb_monitors_and_read_id" {
  value = data.cloudflare_api_token_permission_groups.all.account["Load Balancing: Monitors and Pools Read"] // 9d24387c6e8544e2bc4024a03991339f
}

# Get user level "Memberships Read" permission ID.
output "user_memberships_read_id" {
  value = data.cloudflare_api_token_permission_groups.all.user["Memberships Read"] // 3518d0f75557482e952c6762d3e64903
}
