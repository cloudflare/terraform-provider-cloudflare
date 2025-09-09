variable "ip_list" {
  default = [
    { ip = "10.0.0.1", comment = "Dynamic IP 1" },
    { ip = "10.0.0.2", comment = "Dynamic IP 2" },
  ]
}

resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[3]s"
  kind        = "ip"
  description = "%[4]s"
  
  dynamic "item" {
    for_each = var.ip_list
    content {
      value {
        ip = item.value.ip
      }
      comment = item.value.comment
    }
  }
}