resource "cloudflare_managed_transforms" "%[1]s" {
  zone_id = "%[2]s"
  managed_request_headers = [{
    id = "remove_visitor_ip_headers"
    enabled = true
  }]
  
  managed_response_headers = []
}
