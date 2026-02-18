resource "cloudflare_dlp_profile" "%[1]s" {
  account_id          = "%[2]s"
  name                = "no-desc-%[1]s"
  type                = "custom"
  allowed_match_count = 0

  entry {
    name    = "Test %[1]s"
    enabled = false
    pattern {
      regex = "test[0-9]{3}"
    }
  }
}
