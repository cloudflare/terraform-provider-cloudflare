resource "cloudflare_dlp_profile" "%[1]s" {
  account_id          = "%[2]s"
  name                = "AWS Keys"
  type                = "predefined"
  allowed_match_count = 3

  entry {
    id      = "%[3]s"
    name    = "AWS Access Key ID"
    enabled = true
  }
}
