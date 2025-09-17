
resource "cloudflare_rate_limit" "%[1]s" {
  zone_id = "%[2]s"
  threshold = 2000
  period = 10
  match = [{
    request = [{
      url_pattern = "%[3]s/tfacc-full-%[1]s"
      schemes = ["HTTP", "HTTPS"]
      methods = ["GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"]
    }]
    response = {
      statuses = [200, 201, 202, 301, 429]
      origin_traffic = false
      headers = [
        {
          name  = "Test"
          op    = "ne"
          value = "test"
	    },
        {
          name  = "Host"
          op    = "eq"
          value = "localhost"
	    }
      ]
    }]
  }
  action = [{
    mode = "simulate"
    timeout = 43200
    response = {
      content_type = "text/plain"
      body = "my response body"
    }
  }]
  correlate = [{
	  by = "nat"
  }]
  disabled = true
  description = "my fully specified rate limit for a zone"
  bypass_url_patterns = ["%[3]s/bypass1","%[3]s/bypass2"]
}
