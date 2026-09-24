resource "cloudflare_zone_lockdown" "%[1]s" {
    zone_id = "%[2]s"
    urls    = ["%[3]s", "%[4]s", "%[5]s"]
    configurations = [
      {
        target = "ip"
        value  = "198.51.100.4"
      }
    ]
}
