resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"

  build_config = {
    build_command   = "npm run build"
    destination_dir = "public"
    root_dir        = "/"
  }
}
