resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Response Compression Ruleset %[2]s"
  phase   = "http_response_compression"
  kind    = "zone"

  rules {
    expression = "true"
    action     = "compress_response"
    description = "Enable gzip compression"

    action_parameters {
      algorithms {
        name = "gzip"
      }
    }
  }

  rules {
    expression = "http.host eq \"example.com\""
    action     = "compress_response"
    description = "Enable multiple compression algorithms"

    action_parameters {
      algorithms {
        name = "gzip"
      }

      algorithms {
        name = "brotli"
      }
    }
  }

  rules {
    expression = "http.request.uri.path contains \"/api\""
    action     = "compress_response"
    description = "Compression with specific algorithms"

    action_parameters {
      algorithms {
        name = "brotli"
      }

      algorithms {
        name = "zstd"
      }

      algorithms {
        name = "gzip"
      }
    }
  }
}
