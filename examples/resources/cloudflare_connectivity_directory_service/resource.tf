resource "cloudflare_connectivity_directory_service" "example_connectivity_directory_service" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  host = {
    ipv4 = "10.0.0.1"
    network = {
      tunnel_id = "0191dce4-9ab4-7fce-b660-8e5dec5172da"
    }
  }
  name = "web-app"
  type = "http"
  http_port = 8080
  https_port = 8443
  tls_settings = {
    cert_verification_mode = "verify_full"
    certificate_authority = "705db030-8247-489b-80dd-0d6256e5950d"
    client_certificate = "a493fbb4-34bc-4218-a63c-61a9323ac725"
  }
}
