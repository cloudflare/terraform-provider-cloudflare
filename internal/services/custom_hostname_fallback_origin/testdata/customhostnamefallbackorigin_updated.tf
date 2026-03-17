# Only the fallback origin resource - DNS records are managed outside Terraform
# to avoid destruction order issues with the async delete API.
resource "cloudflare_custom_hostname_fallback_origin" "%[2]s" {
  zone_id = "%[1]s"
  origin  = "fallback-origin-updated.%[3]s.%[4]s"
}
