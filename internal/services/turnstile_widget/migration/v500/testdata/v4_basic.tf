resource "cloudflare_turnstile_widget" "%s" {
  account_id = "%s"
  name       = "%s"
  domains    = toset(["example.com"])
  mode       = "managed"
}
