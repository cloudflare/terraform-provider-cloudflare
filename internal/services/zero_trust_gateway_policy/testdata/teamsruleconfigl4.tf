resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"
  description = "desc"
  enabled = true
  precedence = 12308
  action = "l4_override"
  filters = ["l4"]
  traffic = "net.dst.ip in {10.0.0.0/8} and net.dst.port in {80 443 8080 53} and not(net.dst.ip in {10.217.0.0/16})"
  device_posture = "any(device_posture.checks.passed[*] == \"51fe39d9-d584-48f5-9eed-36cd14ada791\")"
  rule_settings = {
    l4override ={ "ip": "1.1.1.1", port: 80}
  }
}
