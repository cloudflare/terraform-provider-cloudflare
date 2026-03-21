resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"

  source = {
    type = "github"
    config = {
      branch                         = "main"
      build_command                  = "npm run build"
      deploy_command                 = "npm run deploy"
      owner                          = "my-org"
      owner_id                       = "123456"
      production_deployments_enabled = true
      repo_id                        = "789012"
      repo_name                      = "my-worker-repo"
      path_includes                  = ["src/**"]
      path_excludes                  = ["test/**"]
      preview_branch_includes        = ["feature/**"]
      preview_branch_excludes        = ["experimental/**"]
    }
  }
}
