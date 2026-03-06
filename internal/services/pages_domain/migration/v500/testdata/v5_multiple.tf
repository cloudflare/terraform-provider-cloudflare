resource "cloudflare_pages_project" "%[1]s_%[2]s" {
  account_id        = "%[3]s"
  name              = "%[4]s"
  production_branch = "main"
}

resource "cloudflare_pages_domain" "%[1]s_domain1" {
  account_id   = "%[3]s"
  project_name = cloudflare_pages_project.%[1]s_%[2]s.name
  name         = "%[5]s"
}

resource "cloudflare_pages_domain" "%[1]s_domain2" {
  account_id   = "%[3]s"
  project_name = cloudflare_pages_project.%[1]s_%[2]s.name
  name         = "%[6]s"
}
