# IPv4 host with tunnel network
resource "cloudflare_connectivity_directory_service" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "my-internal-app"
  type       = "http"

  host {
    ipv4 = "192.168.1.100"

    network {
      tunnel_id = "238dccd1-149b-463d-8228-560ab83a54fd"
    }
  }

  http_port  = 80
  https_port = 443
}

# Hostname host with resolver network
resource "cloudflare_connectivity_directory_service" "hostname_example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "my-hostname-app"
  type       = "http"

  host {
    hostname = "internal.example.com"

    resolver_network {
      tunnel_id    = "238dccd1-149b-463d-8228-560ab83a54fd"
      resolver_ips = ["10.0.0.1", "10.0.0.2"]
    }
  }
}
