# Challenge requests coming from known Tor exit nodes.
resource "cloudflare_access_rule" "tor_exit_nodes" {
  notes = "Requests coming from known Tor exit nodes"
  mode  = "challenge"
  configuration {
    target = "country"
    value  = "T1"
  }
}

# Whitelist (sic!) requests coming from Antarctica, but only for single zone.
resource "cloudflare_access_rule" "antarctica" {
  notes = "Requests coming from Antarctica"
  mode  = "whitelist"
  configuration {
    target = "country"
    value  = "AQ"
  }
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
}

# Whitelist office's network IP ranges on all account zones (or other lists of resources).
# Resulting Terraform state will be a list of resources.
provider "cloudflare" {
  # ... other provider configuration
  account_id = "d41d8cd98f00b204e9800998ecf8427e"
}
variable "my_office" {
  type    = list(string)
  default = ["192.0.2.0/24", "198.51.100.0/24", "2001:db8::/56"]
}
resource "cloudflare_access_rule" "office_network" {
  count = length(var.my_office)
  notes = "Requests coming from office network"
  mode  = "whitelist"
  configuration {
    target = "ip_range"
    value  = element(var.my_office, count.index)
  }
}
