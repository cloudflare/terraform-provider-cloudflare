resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  builds {
    enabled        = true
    build_command  = "npm run build"
    deploy_command = "npm run deploy"
  }
}
