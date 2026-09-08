
resource "cloudflare_turnstile_widget" "%[1]s" {
  account_id = "%[2]s"
  name       = "my-tf-widget-order-%[1]s"
  mode       = "managed"

  # Domains deliberately listed in non-alphabetical order. The API
  # canonically returns them sorted alphabetically, so without the
  # response-reordering fix this config produces a perpetual diff
  # (GitHub #7028).
  domains = [
    "zebra.example.com",
    "alpha.example.com",
    "middle.example.com",
  ]
}
