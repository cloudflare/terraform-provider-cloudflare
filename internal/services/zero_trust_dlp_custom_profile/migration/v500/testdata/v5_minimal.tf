resource "cloudflare_zero_trust_dlp_custom_profile" "%[1]s" {
  account_id          = "%[2]s"
  name                = "minimal-%[1]s"
  allowed_match_count = 1

  entries = [{
    name    = "Simple %[1]s"
    enabled = true
    pattern = {
      regex = "test[0-9]{3,5}"
    }
  }]
}
