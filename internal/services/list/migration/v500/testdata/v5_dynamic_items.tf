variable "ip_list_%[1]s" {
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

  items = [for item in var.ip_list_%[1]s : {
    ip      = item.ip
    comment = item.comment
  }]
}
