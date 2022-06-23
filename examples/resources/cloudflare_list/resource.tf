# IP list
resource "cloudflare_list" "example" {
  account_id  = "919f297a62fdfb28844177128ed4d331"
  name        = "example list"
  description = "example IPs for a list"
  kind        = "ip"

  item {
    value {
      ip = "192.0.2.0"
    }
    comment = "one"
  }

  item {
    value {
      ip = "192.0.2.1"
    }
    comment = "two"
  }
}

# Redirect list
resource "cloudflare_list" "example" {
  account_id  = "919f297a62fdfb28844177128ed4d331"
  name        = "example list"
  description = "example redirects for a list"
  kind        = "redirect"

  item {
    value {
      redirect {
        source_url = "example.com/blog"
        target_url = "https://blog.example.com"
      }
    }
    comment = "one"
  }

  item {
    value {
      redirect {
        source_url            = "example.com/foo"
        target_url            = "https://foo.example.com"
        include_subdomains    = true
        subpath_matching      = true
        status_code           = 301
        preserve_query_string = true
        preserve_path_suffix  = false
      }
    }
    comment = "two"
  }
}
