# Enable Content Scanning
resource "cloudflare_content_scanning" "example" {
    zone_id = "399c6f4950c01a5a141b99ff7fbcbd8b"
    enabled = true
}