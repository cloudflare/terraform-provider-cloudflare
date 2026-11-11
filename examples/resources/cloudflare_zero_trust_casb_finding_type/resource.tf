resource "cloudflare_zero_trust_casb_finding_type" "example_zero_trust_casb_finding_type" {
  account_id = "46148281d8a93d002ef242d8b0d5f9f6"
  applies_to_all_integrations = true
  asset_class = "GoogleWorkspaceFile"
  asset_vendor = "GOOGLE_WORKSPACE"
  description = "Detects files shared with too many users"
  expression = <<EOT
  default result := "pass"

  result := "fail" if input.asset.publicLinkDisabled == false

  EOT
  name = "Excessive File Sharing"
  severity = "High"
  integration_ids = ["550e8400-e29b-41d4-a716-446655440000"]
}
