resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  builds {
    provider_type         = "github"
    provider_account_name = "my-org"
    repo_id               = "123456"
    repo_name             = "my-worker-repo"
  }
}
