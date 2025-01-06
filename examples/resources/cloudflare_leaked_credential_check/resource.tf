# Enable the Leaked Credentials Check detection
resource "cloudflare_leaked_credential_check" "example" {
    zone_id = "399c6f4950c01a5a141b99ff7fbcbd8b"
    enabled = true
}