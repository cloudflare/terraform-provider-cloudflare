variable "account_id" {
  default = "%[3]s"
}

variable "project_name" {
  default = "%[4]s"
}

variable "domain_name" {
  default = "%[5]s"
}

resource "cloudflare_pages_project" "%[1]s_%[2]s" {
  account_id        = var.account_id
  name              = var.project_name
  production_branch = "main"
}

resource "cloudflare_pages_domain" "%[1]s" {
  account_id   = var.account_id
  project_name = cloudflare_pages_project.%[1]s_%[2]s.name
  name         = var.domain_name
}
