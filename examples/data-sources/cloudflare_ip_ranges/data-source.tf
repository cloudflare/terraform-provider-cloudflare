data "cloudflare_ip_ranges" "cloudflare" {}

resource "example_firewall_resource" "example" {
  name    = "from-cloudflare"
  network = "default"

  source_ranges = data.cloudflare_ip_ranges.cloudflare.ipv4_cidr_blocks

  allow {
    ports    = "443"
    protocol = "tcp"
  }
}
