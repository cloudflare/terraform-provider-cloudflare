# Challenge requests coming from known Tor exit nodes.
resource "cloudflare_access_rule" "tor_exit_nodes" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  notes   = "Requests coming from known Tor exit nodes"
  mode    = "challenge"
  configuration {
    target = "country"
    value  = "T1"
  }
}

# Allowlist requests coming from Antarctica, but only for single zone.
resource "cloudflare_access_rule" "antarctica" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  notes   = "Requests coming from Antarctica"
  mode    = "whitelist"
  configuration {
    target = "country"
    value  = "AQ"
  }
}

# Allowlist office's network IP ranges on all account zones (or other lists of
# resources).
variable "my_office" {
  type    = list(string)
  default = ["192.0.2.0/24", "198.51.100.0/24", "2001:db8::/56"]
}

resource "cloudflare_access_rule" "office_network" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  count      = length(var.my_office)
  notes      = "Requests coming from office network"
  mode       = "whitelist"
  configuration {
    target = "ip_range"
    value  = element(var.my_office, count.index)
  }
}
