resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"

  build_config = {
    build_caching     = true
    build_command     = "npm run build"
    destination_dir   = "dist"
    root_dir          = "/frontend"
    web_analytics_tag = "my-tag"
  }
}
