resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[3]s"
  kind        = "ip"
  description = "%[4]s"
  
  item {
    value {
      ip = "192.0.2.1"
    }
    comment = "Test IP 1"
  }
  
  item {
    value {
      ip = "192.0.2.2"
    }
    comment = "Test IP 2"
  }
}