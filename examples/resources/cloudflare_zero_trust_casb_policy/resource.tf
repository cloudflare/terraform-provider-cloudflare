resource "cloudflare_zero_trust_casb_policy" "example_zero_trust_casb_policy" {
  account_id = "46148281d8a93d002ef242d8b0d5f9f6"
  actions = {
    remediation_types = [{
      remediation_type_id = "5a7d9e2f-1b3c-4d5e-8f6a-7b8c9d0e1f2a"
    }]
    webhook_configs = [{
      webhook_config_id = "3f7b8c9d-6e5a-4f3b-9c2d-1e0a8b7c6d5e"
    }]
  }
  applies_to_all_integrations = false
  display_name = "Auto-remediate public files"
  enabled = true
  finding_type_id = "5a7d9e2f-1b3c-4d5e-8f6a-7b8c9d0e1f2a"
  description = "Automatically remove public access from files when detected"
  integration_ids = ["497f6eca-6276-4993-bfeb-53cbbbba6f08"]
}
