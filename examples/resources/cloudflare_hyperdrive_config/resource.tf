resource "cloudflare_hyperdrive_config" "example_hyperdrive_config" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = "example-hyperdrive"
  origin = {
    database = "postgres"
    host = "database.example.com"
    password = "password"
    port = 5432
    scheme = "postgres"
    user = "postgres"
  }
  caching = {
    disabled = true
    max_age = 0
    stale_while_revalidate = 0
  }
  integration = {

  }
  mtls = {
    ca_certificate_id = "00000000-0000-0000-0000-0000000000"
    mtls_certificate_id = "00000000-0000-0000-0000-0000000000"
    sslmode = "verify-full"
  }
  origin_connection_limit = 60
}
