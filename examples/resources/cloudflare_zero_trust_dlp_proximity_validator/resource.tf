resource "cloudflare_zero_trust_dlp_proximity_validator" "example_zero_trust_dlp_proximity_validator" {
  account_id = "account_id"
  name = "HIPAA context words"
  words = [{
    weight = "strong"
    word = "social security number"
  }]
}
