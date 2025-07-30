resource "cloudflare_snippet_rules" "example_snippet_rules" {
  zone_id = "9f1839b6152d298aca64c4e906b6d074"
  rules = [{
    expression   = "ip.src eq 1.1.1.1"
    snippet_name = "my_snippet"
    description  = "Execute my_snippet when IP address is 1.1.1.1."
    enabled      = true
  }]
}
