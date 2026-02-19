resource "cloudflare_access_identity_provider" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "azureAD"
  config {
    client_id     = "test"
    client_secret = "test"
    directory_id  = "directory"
  }
}
