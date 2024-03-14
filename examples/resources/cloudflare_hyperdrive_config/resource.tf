resource "cloudflare_hyperdrive_config" "no_defaults" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "my-hyperdrive-config"
  origin = {
    database = "postgres"
    password = "my-password"
    host     = "my-database.example.com"
    port     = 5432
    scheme   = "postgres"
    user     = "my-user"
  }
}
