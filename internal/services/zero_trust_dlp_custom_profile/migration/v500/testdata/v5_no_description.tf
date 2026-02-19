resource "cloudflare_zero_trust_dlp_custom_profile" "%[1]s" {
  account_id          = "%[2]s"
  name                = "no-desc-%[1]s"
  allowed_match_count = 0

  entries = [{
    name    = "Test %[1]s"
    enabled = false
    pattern = {
      regex = "test[0-9]{3}"
    }
  }]
}
