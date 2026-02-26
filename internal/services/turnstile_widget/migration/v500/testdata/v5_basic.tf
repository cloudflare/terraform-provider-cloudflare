resource "cloudflare_turnstile_widget" "%s" {
  account_id = "%s"
  name       = "%s"
  domains    = ["example.com"]
  mode       = "managed"
}
