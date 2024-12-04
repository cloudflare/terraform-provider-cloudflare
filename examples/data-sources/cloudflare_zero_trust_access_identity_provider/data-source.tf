data "cloudflare_zero_trust_access_identity_provider" "example_zero_trust_access_identity_provider" {
  identity_provider_id = "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"
  account_id = "account_id"
  zone_id = "zone_id"
}
