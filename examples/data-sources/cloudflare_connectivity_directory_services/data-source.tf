data "cloudflare_connectivity_directory_services" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  type       = "http"
}

output "all_services" {
  value = data.cloudflare_connectivity_directory_services.example.services
}
