resource "cloudflare_managed_transforms" "%[1]s" {
  zone_id = "%[2]s"
  managed_request_headers = []
  
  managed_response_headers = [{
    id = "add_security_headers"
    enabled = true
  }]
}
