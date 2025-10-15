resource "cloudflare_content_scanning" "%[2]s" {
  zone_id = "%[1]s"
  value   = "enabled"
}

output "modified" {
  value = cloudflare_content_scanning.%[2]s.modified
}
