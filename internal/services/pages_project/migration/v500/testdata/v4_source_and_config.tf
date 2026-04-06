resource "cloudflare_pages_project" "%s" {
  account_id        = "%s"
  name              = "%s"
  production_branch = "main"

  source {
    type = "github"
    config {
      owner                         = "cloudflare"
      repo_name                     = "%s"
      production_branch             = "main"
      production_deployment_enabled = true
      pr_comments_enabled           = true
    }
  }
}
