resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[3]s"
  kind        = "asn"
  description = "%[4]s"
  
  item {
    value {
      asn = 12345
    }
    comment = "Test ASN 1"
  }
  
  item {
    value {
      asn = 67890
    }
  }
}