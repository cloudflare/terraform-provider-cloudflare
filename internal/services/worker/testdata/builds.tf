resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"

  builds = {
    enabled                            = true
    branch                             = "main"
    build_command                      = "npm run build"
    deploy_command                     = "npm run deploy"
    non_production_deployments_enabled = false
    provider_type                      = "github"
    provider_account_name              = "my-org"
    provider_account_id                = "123456"
    repo_id                            = "789012"
    repo_name                          = "my-worker-repo"
    branch_includes                    = ["feature/**"]
    branch_excludes                    = ["experimental/**"]
  }
}
