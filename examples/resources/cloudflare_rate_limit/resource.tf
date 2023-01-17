resource "cloudflare_rate_limit" "example" {
  zone_id   = "0da42c8d2132a9ddaf714f9e7c920711"
  threshold = 2000
  period    = 2
  match {
    request {
      url_pattern = "${var.cloudflare_zone}/*"
      schemes     = ["HTTP", "HTTPS"]
      methods     = ["GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"]
    }
    response {
      statuses       = [200, 201, 202, 301, 429]
      origin_traffic = false
      headers = [
        {
          name  = "Host"
          op    = "eq"
          value = "localhost"
        },
        {
          name  = "X-Example"
          op    = "ne"
          value = "my-example"
        }
      ]
    }
  }
  action {
    mode    = "simulate"
    timeout = 43200
    response {
      content_type = "text/plain"
      body         = "custom response body"
    }
  }
  correlate {
    by = "nat"
  }
  disabled            = false
  description         = "example rate limit for a zone"
  bypass_url_patterns = ["example.com/bypass1", "example.com/bypass2"]
}
