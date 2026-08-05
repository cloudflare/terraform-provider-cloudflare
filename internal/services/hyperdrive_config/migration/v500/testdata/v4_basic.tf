resource "cloudflare_hyperdrive_config" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-hyperdrive-%[1]s"

  origin = {
    database = "%[3]s"
    host     = "%[4]s"
    port     = %[5]s
    scheme   = "postgres"
    user     = "%[6]s"
    password = "%[7]s"
  }
}
