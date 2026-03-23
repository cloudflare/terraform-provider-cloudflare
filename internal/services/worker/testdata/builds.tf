resource "cloudflare_worker" {
  account_id = "%[2]s"
  name       = "%[1]s"

  builds = {
    enabled        = true
    branch         = "main"
    build_command  = "npm run build"
    deploy_command = "npx wrangler deploy"
  }
}
