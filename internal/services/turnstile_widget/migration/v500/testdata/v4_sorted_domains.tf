resource "cloudflare_turnstile_widget" "%s" {
  account_id = "%s"
  name       = "%s"
  domains    = toset(["zzz.example.com", "aaa.example.com", "mmm.example.com"])
  mode       = "managed"
}
