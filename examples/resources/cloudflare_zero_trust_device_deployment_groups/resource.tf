resource "cloudflare_zero_trust_device_deployment_groups" "example_zero_trust_device_deployment_groups" {
  account_id = "account_id"
  name = "Engineering Ring 0"
  version_config = [{
    target_environment = "windows"
    version = "2026.6.234.0"
  }]
  policy_ids = ["string"]
}
