resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  builds {
    branch          = "main"
    branch_includes = ["develop", "feature/*"]
    branch_excludes = ["test-*", "temp-*"]
  }
}
