resource "cloudflare_turnstile_widget" "%s" {
  account_id = "%s"
  name       = "%s"
  domains    = ["aaa.example.com", "mmm.example.com", "zzz.example.com"]
  mode       = "managed"
}
