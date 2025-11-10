resource "cloudflare_connectivity_directory_service" "example_connectivity_directory_service" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  host = {
    hostname = "hostname"
    ipv4 = "ipv4"
    ipv6 = "ipv6"
    network = {

    }
    resolver_network = {

    }
  }
  name = "name"
  type = "http"
  http_port = 1
  https_port = 1
}
