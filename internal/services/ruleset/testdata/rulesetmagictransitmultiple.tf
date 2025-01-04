
resource "cloudflare_ruleset" "%[1]s" {
  account_id  = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s magic transit ruleset description"
  kind        = "root"
  phase       = "magic_transit"

  rules = [{
    action      = "block"
    expression  = "udp.dstport in { 32768..65535 }"
    description = "Block UDP Ephemeral Ports"
    enabled     = true
    },
    {
      action = "skip"
      action_parameters = {
        ruleset = "current"
      }
      expression  = "tcp.dstport in { 32768..65535 }"
      description = "Allow TCP Ephemeral Ports"
      enabled     = true
      logging = {
        enabled = false
      }
  }]

}