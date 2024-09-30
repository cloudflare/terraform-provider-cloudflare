data "cloudflare_infrastructure_access_targets" "example" {
  account_id        = "f037e56e89293a057740de681ac9abbe"
  hostname_contains = "example"
  ipv4              = "210.26.29.230"
}

# output the list of targets the data source contains
output "targets" {
  value = data.cloudflare_infrastructure_access_targets.example.targets
}
