# Enable security headers using Managed Meaders
resource "cloudflare_managed_headers" "add_security_headers_example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"

  managed_request_headers {
    id      = "add_true_client_ip_headers"
    enabled = true
  }

  managed_response_headers {
    id      = "remove_x-powered-by_header"
    enabled = true
  }
}
