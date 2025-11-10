resource "cloudflare_connectivity_directory_service" "example_connectivity_directory_service" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  host = {
    hostname = "api.example.com"
    resolver_network = {
      tunnel_id = "0191dce4-9ab4-7fce-b660-8e5dec5172da"
      resolver_ips = ["string"]
    }
  }
  name = "web-server"
  type = "http"
  http_port = 8080
  https_port = 8443
}
