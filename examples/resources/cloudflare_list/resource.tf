# IP list
resource "cloudflare_list" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "example_list"
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
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "example_list"
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
        include_subdomains    = "enabled"
        subpath_matching      = "enabled"
        status_code           = 301
        preserve_query_string = "enabled"
        preserve_path_suffix  = "disabled"
      }
    }
    comment = "two"
  }
}

# ASN list
resource "cloudflare_list" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "example_list"
  description = "example ASNs for a list"
  kind        = "asn"

  item {
    value {
      asn = 677
    }
    comment = "one"
  }

  item {
    value {
     asn = 989
    }
    comment = "two"
  }
}


# Hostname list
resource "cloudflare_list" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "example_list"
  description = "example hostnames for a list"
  kind        = "hostname"

  item {
    value {
      hostname {
        url_hostname = "example.com"
      }
    }
    comment = "one"
  }

  item {
    value {
      hostname {
        url_hostname = "*.example.com"
      }
    }
    comment = "two"
  }
}
