data "cloudflare_zero_trust_casb_finding_types" "example_zero_trust_casb_finding_types" {
  account_id = "46148281d8a93d002ef242d8b0d5f9f6"
  direction = "asc"
  max_updated_at = "2019-12-27T18:11:19.117Z"
  min_updated_at = "2019-12-27T18:11:19.117Z"
  order = "created"
  origins = ["system"]
  quarantined = true
  search = "search"
  severities = ["high"]
  types = ["Content"]
  vendors = ["GOOGLE_WORKSPACE"]
}
