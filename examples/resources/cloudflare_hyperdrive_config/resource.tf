# Config using all default values
resource "cloudflare_hyperdrive_config" "no_defaults" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "my-hyperdrive-config"
  password   = "my-password"
  origin     = {
    database = "postgres"
    host     = "my-database.example.com"
  }
}

# Config not using any default values
resource "cloudflare_hyperdrive_config" "no_defaults" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "my-hyperdrive-config"
  password   = "my-password"
  origin     = {
    database = "postgres"
    host     = "my-database.example.com"
    port     = 5432
    scheme   = "postgres"
    user     = "my-user"
  }
  caching    = {
    disabled               = false
    max_age                = 60
    stale_while_revalidate = 30
  }
}


