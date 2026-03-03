data "cloudflare_connectivity_directory_service" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  service_id = "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}

output "service_name" {
  value = data.cloudflare_connectivity_directory_service.example.name
}
