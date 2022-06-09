# Enable security headers using Managed Meaders
resource "cloudflare_managed_headers" "add_security_headers_example" {
    zone_id     = "cb029e245cfdd66dc8d2e570d5dd3322"
    managed_request_headers {
        id = "add_true_client_ip_headers"
        enabled = true
    }

    managed_request_headers {
        id = "add_visitor_location_headers"
        enabled = false
    }

    managed_response_headers {
        id = "add_security_headers"
        enabled = false
    }

    managed_response_headers {
        id = "remove_x-powered-by_header"
        enabled = true
    }
}
