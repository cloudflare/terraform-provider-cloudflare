resource "cloudflare_managed_headers" "%[1]s" {
  zone_id = "%[2]s"

  managed_response_headers {
    id      = "add_security_headers"
    enabled = true
  }
}
